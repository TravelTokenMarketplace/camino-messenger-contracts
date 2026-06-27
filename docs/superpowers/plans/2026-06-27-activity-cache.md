# Activity Feed Persistent Cache Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Persist scanned activity-feed events to `localStorage` so returning visitors render history instantly, re-scan only the gap since their last visit, and accumulate deeper history across sessions — without trusting reorg-prone recent blocks.

**Architecture:** A pure, React-free cache module (`ui/src/lib/activity/cache.ts`) stores per-`(chainId, sourcesKey)` *contiguous scanned segments* of confirmed events, with bigint-safe serialization and an event-count cap. `useActivity` hydrates from it synchronously on mount, wraps its existing backward-paging `getLogs` fetch so confirmed+covered batches short-circuit to the cache (and newly-scanned confirmed ranges are persisted), and auto-pulls a bounded number of batches on return to bridge the gap to the chain tip. Only blocks at or below `tip − confirmations` are ever persisted; the unconfirmed tail is always re-scanned live.

**Tech Stack:** React 18 + TypeScript, viem v2 (`getLogs`, `bigint` block numbers), TanStack Query (`useInfiniteQuery`), Vitest + Testing Library.

## Global Constraints

- **Run all commands from `ui/`.** If `yarn test` resolves to the root Hardhat project, call `./node_modules/.bin/vitest run <path>` directly (see `ui/CLAUDE.md`).
- **Never widen `eth_getLogs` ranges.** Reuse the existing adaptive batch machinery in `useActivity.ts` (`fetchActivityPage`); do not introduce new unbounded log queries.
- **`ActivityEvent.blockNumber` is a `bigint`, and `args` may contain nested `bigint`s** (e.g. `tokenId`). All serialization must round-trip nested bigints.
- **The cache is a pure optimization, never a source of truth.** Any parse failure, version mismatch, corruption, or quota error must degrade silently to a live scan — never throw into the feed.
- Events are ordered newest-first everywhere via `compareEventsDesc` (block desc, then `logIndex` desc) and deduped by `ActivityEvent.id` (`${txHash}#${logIndex}`).

---

## File Structure

- **Create** `ui/src/lib/activity/sort.ts` — shared pure ordering helpers (`compareEventsDesc`, `dedupeById`) extracted so both the hook and the cache module use one implementation (no import cycle).
- **Create** `ui/src/lib/activity/cache.ts` — the persistence layer: `Segment`/`CacheEntry` types, bigint-safe serialize/deserialize, `readCache`/`writeCache`, `mergeSegment`, `capEntry`, and coverage helpers (`cachedHigh`, `isRangeCovered`, `eventsInRange`, `totalEvents`).
- **Create** `ui/src/lib/activity/cache.test.ts` — unit tests for the cache module.
- **Modify** `ui/src/config/activity.ts` — add cache-related constants.
- **Modify** `ui/src/hooks/useActivity.ts` — re-export from `sort.ts`, add `currentBatchSize` + `fetchActivityPageCached`, and wire hydrate / cache-aware queryFn / auto catch-up into the hook.
- **Modify** `ui/src/hooks/useActivity.test.ts` — add tests for `fetchActivityPageCached`.

---

## Task 1: Shared ordering helpers + cache storage & serialization

**Files:**
- Create: `ui/src/lib/activity/sort.ts`
- Create: `ui/src/lib/activity/cache.ts`
- Create: `ui/src/lib/activity/cache.test.ts`
- Modify: `ui/src/config/activity.ts`
- Modify: `ui/src/hooks/useActivity.ts`

**Interfaces:**
- Consumes: `ActivityEvent` from `../lib/activity/types`.
- Produces:
  - `sort.ts`: `compareEventsDesc(a: ActivityEvent, b: ActivityEvent): number`, `dedupeById(events: ActivityEvent[]): ActivityEvent[]`.
  - `cache.ts`: `type Segment = { low: bigint; high: bigint; events: ActivityEvent[] }`, `type CacheEntry = { version: number; segments: Segment[] }`, `serializeEntry(entry: CacheEntry): string`, `deserializeEntry(json: string): CacheEntry | null`, `readCache(chainId: number, sourcesKey: string): CacheEntry | null`, `writeCache(chainId: number, sourcesKey: string, entry: CacheEntry): void`, `totalEvents(entry: CacheEntry): number`.
  - `config/activity.ts`: `ACTIVITY_CACHE_VERSION: number`, `ACTIVITY_CACHE_MAX_EVENTS: number`.

- [ ] **Step 1: Add cache config constants**

In `ui/src/config/activity.ts`, append:

```ts
/** Bump when the persisted cache shape changes; older entries are discarded on read. */
export const ACTIVITY_CACHE_VERSION = 1;

/**
 * Max events kept per cache entry. When exceeded, the oldest events are dropped
 * (see capEntry) to stay well under the ~5MB localStorage origin budget.
 */
export const ACTIVITY_CACHE_MAX_EVENTS = 2000;
```

- [ ] **Step 2: Extract shared ordering helpers into `sort.ts`**

Create `ui/src/lib/activity/sort.ts`:

```ts
import { type ActivityEvent } from "./types";

/** Newest first: higher block, then higher logIndex. */
export function compareEventsDesc(a: ActivityEvent, b: ActivityEvent): number {
  if (a.blockNumber !== b.blockNumber) return a.blockNumber > b.blockNumber ? -1 : 1;
  return b.logIndex - a.logIndex;
}

/** Dedupe by stable id (`${txHash}#${logIndex}`), preserving first-seen order. */
export function dedupeById(events: ActivityEvent[]): ActivityEvent[] {
  const seen = new Set<string>();
  const out: ActivityEvent[] = [];
  for (const e of events) {
    if (seen.has(e.id)) continue;
    seen.add(e.id);
    out.push(e);
  }
  return out;
}
```

- [ ] **Step 3: Re-export the helpers from `useActivity.ts` so existing imports keep working**

In `ui/src/hooks/useActivity.ts`:

1. Add an import near the top (after the existing imports):

```ts
import { compareEventsDesc, dedupeById } from "../lib/activity/sort";
```

2. **Delete** the local `compareEventsDesc` function (the `/** Newest first... */` block) and the local `dedupeById` function.

3. Add a re-export immediately after the new import so the existing test import (`import { compareEventsDesc } from "./useActivity"`) still resolves:

```ts
export { compareEventsDesc, dedupeById };
```

- [ ] **Step 4: Run the existing suite to confirm the refactor is behavior-preserving**

Run: `cd ui && ./node_modules/.bin/vitest run src/hooks/useActivity.test.ts`
Expected: PASS (all existing cases green — `compareEventsDesc` now sourced from `sort.ts`).

- [ ] **Step 5: Write the failing serialization + storage tests**

Create `ui/src/lib/activity/cache.test.ts`:

```ts
import { afterEach, describe, expect, it, vi } from "vitest";
import {
  deserializeEntry,
  readCache,
  serializeEntry,
  totalEvents,
  writeCache,
  type CacheEntry,
} from "./cache";
import { ACTIVITY_CACHE_VERSION } from "../../config/activity";
import { type ActivityEvent } from "./types";

const CHAIN = 84532;
const SK = "bookingToken:0xbeef";

function ev(block: bigint, logIndex: number, tokenId: bigint): ActivityEvent {
  return {
    id: `0x${block}${logIndex}#${logIndex}`,
    source: "bookingToken",
    category: "Bookings",
    contract: "0x000000000000000000000000000000000000beef",
    blockNumber: block,
    logIndex,
    txHash: `0x${block}${logIndex}`,
    eventName: "TokenBought",
    args: { tokenId, buyer: "0x0000000000000000000000000000000000000001" },
    sentence: "bought",
  };
}

function entry(...events: ActivityEvent[]): CacheEntry {
  return { version: ACTIVITY_CACHE_VERSION, segments: [{ low: 1n, high: 100n, events }] };
}

afterEach(() => {
  localStorage.clear();
  vi.restoreAllMocks();
});

describe("serializeEntry / deserializeEntry", () => {
  it("round-trips nested bigints in blockNumber and args", () => {
    const e = entry(ev(100n, 0, 42n));
    const back = deserializeEntry(serializeEntry(e));
    expect(back).not.toBeNull();
    expect(back!.segments[0].events[0].blockNumber).toBe(100n);
    expect(back!.segments[0].high).toBe(100n);
    expect(back!.segments[0].events[0].args.tokenId).toBe(42n);
  });

  it("returns null on malformed JSON", () => {
    expect(deserializeEntry("{not json")).toBeNull();
  });

  it("returns null on a version mismatch", () => {
    const stale = serializeEntry({ ...entry(ev(1n, 0, 1n)), version: ACTIVITY_CACHE_VERSION + 1 });
    expect(deserializeEntry(stale)).toBeNull();
  });
});

describe("readCache / writeCache", () => {
  it("persists and reads back an entry", () => {
    writeCache(CHAIN, SK, entry(ev(100n, 0, 7n)));
    const back = readCache(CHAIN, SK);
    expect(back!.segments[0].events[0].args.tokenId).toBe(7n);
  });

  it("drops and removes a corrupt stored entry", () => {
    localStorage.setItem("cm:activity:84532:bookingToken:0xbeef", "{garbage");
    expect(readCache(CHAIN, SK)).toBeNull();
    expect(localStorage.getItem("cm:activity:84532:bookingToken:0xbeef")).toBeNull();
  });

  it("swallows quota errors instead of throwing", () => {
    vi.spyOn(Storage.prototype, "setItem").mockImplementation(() => {
      throw new DOMException("quota", "QuotaExceededError");
    });
    expect(() => writeCache(CHAIN, SK, entry(ev(100n, 0, 1n)))).not.toThrow();
  });

  it("totalEvents counts across segments", () => {
    expect(totalEvents({ version: ACTIVITY_CACHE_VERSION, segments: [{ low: 1n, high: 9n, events: [ev(1n, 0, 1n), ev(2n, 0, 2n)] }] })).toBe(2);
  });
});
```

- [ ] **Step 6: Run the tests to verify they fail**

Run: `cd ui && ./node_modules/.bin/vitest run src/lib/activity/cache.test.ts`
Expected: FAIL — `cache.ts` does not exist yet (module-resolution error).

- [ ] **Step 7: Implement serialization + storage in `cache.ts`**

Create `ui/src/lib/activity/cache.ts`:

```ts
import { ACTIVITY_CACHE_MAX_EVENTS, ACTIVITY_CACHE_VERSION } from "../../config/activity";
import { type ActivityEvent } from "./types";

/** A fully-scanned, inclusive block range `[low, high]` and the events found in it (newest-first). */
export type Segment = { low: bigint; high: bigint; events: ActivityEvent[] };

/** Persisted per `(chainId, sourcesKey)`. Usually exactly one segment. */
export type CacheEntry = { version: number; segments: Segment[] };

const BIGINT_TAG = "$bigint";

// JSON can't carry bigint, and args may hold nested bigints (tokenId, amounts).
// Tag every bigint as { $bigint: "<decimal>" } so deserialize can restore it
// without knowing which fields were originally bigints.
function replacer(_key: string, value: unknown): unknown {
  return typeof value === "bigint" ? { [BIGINT_TAG]: value.toString() } : value;
}

function reviver(_key: string, value: unknown): unknown {
  if (value && typeof value === "object" && BIGINT_TAG in value && Object.keys(value).length === 1) {
    return BigInt((value as Record<string, string>)[BIGINT_TAG]);
  }
  return value;
}

export function serializeEntry(entry: CacheEntry): string {
  return JSON.stringify(entry, replacer);
}

export function deserializeEntry(json: string): CacheEntry | null {
  try {
    const parsed = JSON.parse(json, reviver) as CacheEntry;
    if (!parsed || parsed.version !== ACTIVITY_CACHE_VERSION || !Array.isArray(parsed.segments)) return null;
    return parsed;
  } catch {
    return null;
  }
}

function keyFor(chainId: number, sourcesKey: string): string {
  return `cm:activity:${chainId}:${sourcesKey}`;
}

export function totalEvents(entry: CacheEntry): number {
  return entry.segments.reduce((n, s) => n + s.events.length, 0);
}

export function readCache(chainId: number, sourcesKey: string): CacheEntry | null {
  if (typeof localStorage === "undefined") return null;
  const key = keyFor(chainId, sourcesKey);
  const raw = localStorage.getItem(key);
  if (raw == null) return null;
  const entry = deserializeEntry(raw);
  if (!entry) {
    localStorage.removeItem(key); // discard corrupt / stale-version data
    return null;
  }
  return entry;
}

export function writeCache(chainId: number, sourcesKey: string, entry: CacheEntry): void {
  if (typeof localStorage === "undefined") return;
  let toStore = capEntry(entry, ACTIVITY_CACHE_MAX_EVENTS);
  for (let attempt = 0; attempt < 3; attempt++) {
    try {
      localStorage.setItem(keyFor(chainId, sourcesKey), serializeEntry(toStore));
      return;
    } catch {
      // Likely QuotaExceededError — halve the kept events and retry; give up silently after.
      const half = Math.floor(totalEvents(toStore) / 2);
      if (half < 1) return;
      toStore = capEntry(toStore, half);
    }
  }
}
```

> Note: `writeCache` calls `capEntry`, added in Task 2. Define a temporary local `capEntry` stub now so this file type-checks, then replace it in Task 2 Step 4:
>
> ```ts
> // TEMP stub — replaced in Task 2.
> function capEntry(entry: CacheEntry, _maxEvents: number): CacheEntry {
>   return entry;
> }
> ```

- [ ] **Step 8: Run the tests to verify they pass**

Run: `cd ui && ./node_modules/.bin/vitest run src/lib/activity/cache.test.ts`
Expected: PASS (all serialization + storage cases green).

- [ ] **Step 9: Commit**

```bash
cd ui && git add src/lib/activity/sort.ts src/lib/activity/cache.ts src/lib/activity/cache.test.ts src/config/activity.ts src/hooks/useActivity.ts
git commit -m "feat(ui): activity cache storage + bigint-safe serialization"
```

---

## Task 2: Segment algebra & coverage helpers

**Files:**
- Modify: `ui/src/lib/activity/cache.ts`
- Modify: `ui/src/lib/activity/cache.test.ts`

**Interfaces:**
- Consumes: `Segment`, `CacheEntry`, `totalEvents` from Task 1; `compareEventsDesc`, `dedupeById` from `./sort`.
- Produces:
  - `mergeSegment(segments: Segment[], seg: Segment): Segment[]` — insert a freshly-scanned range, coalescing overlapping/adjacent segments (`low <= last.high + 1`), events deduped + newest-first.
  - `capEntry(entry: CacheEntry, maxEvents: number): CacheEntry` — drop oldest events first; raise the trimmed segment's `low` to the lowest kept event's block; drop fully-emptied segments.
  - `cachedHigh(entry: CacheEntry | null): bigint | null` — highest covered block, or `null`.
  - `isRangeCovered(segments: Segment[], from: bigint, to: bigint): boolean` — some single segment spans `[from, to]`.
  - `eventsInRange(segments: Segment[], from: bigint, to: bigint): ActivityEvent[]` — events with `from <= blockNumber <= to`.

- [ ] **Step 1: Write the failing segment-algebra tests**

Append to `ui/src/lib/activity/cache.test.ts` (add the new imports to the existing top import from `./cache`: `cachedHigh`, `capEntry`, `eventsInRange`, `isRangeCovered`, `mergeSegment`):

```ts
describe("mergeSegment", () => {
  const seg = (low: bigint, high: bigint, ...events: ActivityEvent[]) => ({ low, high, events });

  it("coalesces an adjacent range and keeps events newest-first", () => {
    const a = seg(1n, 10n, ev(10n, 0, 1n));
    const b = seg(11n, 20n, ev(20n, 0, 2n)); // touches a (10+1 == 11)
    const out = mergeSegment([a], b);
    expect(out).toHaveLength(1);
    expect([out[0].low, out[0].high]).toEqual([1n, 20n]);
    expect(out[0].events.map((e) => e.blockNumber)).toEqual([20n, 10n]);
  });

  it("keeps a disjoint range as a separate segment", () => {
    const out = mergeSegment([seg(1n, 10n, ev(5n, 0, 1n))], seg(100n, 200n, ev(150n, 0, 2n)));
    expect(out).toHaveLength(2);
    expect(out.map((s) => s.low)).toEqual([1n, 100n]);
  });

  it("dedupes overlapping events by id", () => {
    const shared = ev(10n, 0, 1n);
    const out = mergeSegment([seg(1n, 10n, shared)], seg(5n, 15n, shared, ev(15n, 0, 2n)));
    expect(out).toHaveLength(1);
    expect(out[0].events).toHaveLength(2);
  });
});

describe("capEntry", () => {
  it("drops the oldest events and raises low to the lowest kept block", () => {
    const e: CacheEntry = {
      version: ACTIVITY_CACHE_VERSION,
      segments: [{ low: 1n, high: 30n, events: [ev(30n, 0, 3n), ev(20n, 0, 2n), ev(10n, 0, 1n)] }],
    };
    const capped = capEntry(e, 2);
    expect(capped.segments[0].events.map((x) => x.blockNumber)).toEqual([30n, 20n]);
    expect(capped.segments[0].low).toBe(20n);
    expect(capped.segments[0].high).toBe(30n);
  });

  it("drops a whole segment when all its events are evicted", () => {
    const e: CacheEntry = {
      version: ACTIVITY_CACHE_VERSION,
      segments: [
        { low: 1n, high: 10n, events: [ev(5n, 0, 1n)] },
        { low: 100n, high: 110n, events: [ev(105n, 0, 2n), ev(104n, 0, 3n)] },
      ],
    };
    const capped = capEntry(e, 2);
    expect(capped.segments).toHaveLength(1);
    expect(capped.segments[0].low).toBe(100n);
  });

  it("returns the entry unchanged when under the cap", () => {
    const e = { version: ACTIVITY_CACHE_VERSION, segments: [{ low: 1n, high: 10n, events: [ev(5n, 0, 1n)] }] };
    expect(capEntry(e, 100)).toBe(e);
  });
});

describe("coverage helpers", () => {
  const segs = [{ low: 10n, high: 20n, events: [ev(15n, 0, 1n)] }, { low: 100n, high: 200n, events: [ev(150n, 0, 2n)] }];

  it("cachedHigh returns the highest covered block", () => {
    expect(cachedHigh({ version: ACTIVITY_CACHE_VERSION, segments: segs })).toBe(200n);
    expect(cachedHigh(null)).toBeNull();
    expect(cachedHigh({ version: ACTIVITY_CACHE_VERSION, segments: [] })).toBeNull();
  });

  it("isRangeCovered is true only when one segment spans the whole range", () => {
    expect(isRangeCovered(segs, 12n, 18n)).toBe(true);
    expect(isRangeCovered(segs, 18n, 105n)).toBe(false); // spans the gap
    expect(isRangeCovered(segs, 5n, 15n)).toBe(false); // below segment low
  });

  it("eventsInRange returns events within the bounds", () => {
    expect(eventsInRange(segs, 100n, 200n).map((e) => e.blockNumber)).toEqual([150n]);
    expect(eventsInRange(segs, 0n, 9n)).toEqual([]);
  });
});
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd ui && ./node_modules/.bin/vitest run src/lib/activity/cache.test.ts`
Expected: FAIL — `mergeSegment`, `capEntry`, `cachedHigh`, `isRangeCovered`, `eventsInRange` are not exported.

- [ ] **Step 3: Implement the segment algebra in `cache.ts`**

Add to the import at the top of `ui/src/lib/activity/cache.ts`:

```ts
import { compareEventsDesc, dedupeById } from "./sort";
```

Append these functions:

```ts
function sortDesc(events: ActivityEvent[]): ActivityEvent[] {
  return dedupeById([...events]).sort(compareEventsDesc);
}

/**
 * Insert a freshly-scanned segment, coalescing any segments that overlap or
 * touch it (`low <= high + 1`). Returned segments are sorted ascending by `low`;
 * each segment's events are deduped and newest-first.
 */
export function mergeSegment(segments: Segment[], seg: Segment): Segment[] {
  const all = [...segments, seg].sort((a, b) => (a.low < b.low ? -1 : 1));
  const merged: Segment[] = [];
  for (const s of all) {
    const last = merged[merged.length - 1];
    if (last && s.low <= last.high + 1n) {
      last.low = s.low < last.low ? s.low : last.low;
      last.high = s.high > last.high ? s.high : last.high;
      last.events = sortDesc([...last.events, ...s.events]);
    } else {
      merged.push({ low: s.low, high: s.high, events: sortDesc(s.events) });
    }
  }
  return merged;
}

/**
 * Bound total events by dropping the oldest first. Within a trimmed segment the
 * surviving `low` is raised to the lowest kept event's block — we no longer claim
 * coverage of blocks whose events we discarded. Emptied segments are removed.
 */
export function capEntry(entry: CacheEntry, maxEvents: number): CacheEntry {
  if (totalEvents(entry) <= maxEvents) return entry;
  const ascending = [...entry.segments].sort((a, b) => (a.low < b.low ? -1 : 1));
  let excess = totalEvents(entry) - maxEvents;
  const out: Segment[] = [];
  for (const seg of ascending) {
    if (excess <= 0) {
      out.push(seg);
      continue;
    }
    if (excess >= seg.events.length) {
      excess -= seg.events.length; // whole segment evicted
      continue;
    }
    const kept = seg.events.slice(0, seg.events.length - excess); // events are newest-first; keep the head
    excess = 0;
    out.push({ low: kept[kept.length - 1].blockNumber, high: seg.high, events: kept });
  }
  return { ...entry, segments: out };
}

/** Highest block covered by any segment, or null when empty. */
export function cachedHigh(entry: CacheEntry | null): bigint | null {
  if (!entry || entry.segments.length === 0) return null;
  return entry.segments.reduce((hi, s) => (s.high > hi ? s.high : hi), entry.segments[0].high);
}

/** True when a single segment fully spans `[from, to]`. */
export function isRangeCovered(segments: Segment[], from: bigint, to: bigint): boolean {
  return segments.some((s) => s.low <= from && s.high >= to);
}

/** Events with `from <= blockNumber <= to`, across all segments (unsorted). */
export function eventsInRange(segments: Segment[], from: bigint, to: bigint): ActivityEvent[] {
  const out: ActivityEvent[] = [];
  for (const s of segments) for (const e of s.events) if (e.blockNumber >= from && e.blockNumber <= to) out.push(e);
  return out;
}
```

- [ ] **Step 4: Remove the temporary `capEntry` stub from Task 1**

Delete the `// TEMP stub — replaced in Task 2.` `capEntry` function added in Task 1 Step 7. The real exported `capEntry` above now satisfies `writeCache`'s call.

- [ ] **Step 5: Run the tests to verify they pass**

Run: `cd ui && ./node_modules/.bin/vitest run src/lib/activity/cache.test.ts`
Expected: PASS (serialization, storage, and segment-algebra cases all green).

- [ ] **Step 6: Commit**

```bash
cd ui && git add src/lib/activity/cache.ts src/lib/activity/cache.test.ts
git commit -m "feat(ui): activity cache segment merge, cap, and coverage helpers"
```

---

## Task 3: Wire the cache into `useActivity`

**Files:**
- Modify: `ui/src/config/activity.ts`
- Modify: `ui/src/hooks/useActivity.ts`
- Modify: `ui/src/hooks/useActivity.test.ts`

**Interfaces:**
- Consumes: `readCache`, `writeCache`, `mergeSegment`, `isRangeCovered`, `eventsInRange`, `cachedHigh`, `type Segment`, `type CacheEntry` from `../lib/activity/cache`; `ACTIVITY_CACHE_VERSION`, `ACTIVITY_CONFIRMATIONS`, `ACTIVITY_CATCHUP_MAX_BATCHES` from `../config/activity`; existing `fetchActivityPage`, `batchBlocksFor`, `ActivityPage`, `ActivitySourceInput`.
- Produces:
  - `currentBatchSize(chainId: number): bigint` — the adaptive batch size in effect for a chain (remembered size, else configured default).
  - `interface CachedFetchDeps { readSegments: () => Segment[]; persist: (seg: Segment) => void }`.
  - `fetchActivityPageCached(client, sources, chainId, toBlock, confirmedTip, deps): Promise<ActivityPage>` — serves a confirmed+covered batch from cache (no RPC), otherwise scans via `fetchActivityPage` and persists the confirmed sub-range.

- [ ] **Step 1: Add confirmations + catch-up config constants**

In `ui/src/config/activity.ts`, append (after the constants from Task 1):

```ts
/**
 * Reorg buffer: only blocks at or below `tip − ACTIVITY_CONFIRMATIONS` are ever
 * persisted. The unconfirmed tail above is always re-scanned live, so a reorg
 * shallower than this can't leave a stale cached event.
 */
export const ACTIVITY_CONFIRMATIONS = 64n;

/**
 * Bounded auto catch-up: on return we pull up to this many batches to bridge the
 * gap from the cached high block to the tip before falling back to manual
 * "Load older". Reuses the per-click span (~100k blocks).
 */
export const ACTIVITY_CATCHUP_MAX_BATCHES = ACTIVITY_BATCHES_PER_CLICK;
```

- [ ] **Step 2: Write the failing `fetchActivityPageCached` tests**

Append to `ui/src/hooks/useActivity.test.ts`. First extend the existing import from `./useActivity` to also bring in `fetchActivityPageCached` and `currentBatchSize`. Then add:

```ts
import { type Segment } from "../lib/activity/cache";

describe("fetchActivityPageCached", () => {
  // In-memory cache deps so these stay pure (no localStorage).
  function memCache(initial: Segment[] = []) {
    let segments = initial;
    return {
      deps: { readSegments: () => segments, persist: (seg: Segment) => (segments = [...segments, seg]) },
      get: () => segments,
    };
  }

  it("serves a confirmed, fully-covered batch from cache without calling getLogs", async () => {
    const size = currentBatchSize(CHAIN); // 10000 (no prior failures)
    const from = 100_000n - size + 1n;
    // Minimal ActivityEvent — only id/blockNumber/logIndex matter to this path.
    const cached = { id: "0xcafe#0", blockNumber: 100_000n, logIndex: 0 } as ActivityEvent;
    const { deps } = memCache([{ low: from, high: 100_000n, events: [cached] }]);
    const getLogs = vi.fn();

    const page = await fetchActivityPageCached(client(getLogs), sources, CHAIN, 100_000n, 200_000n, deps);

    expect(getLogs).not.toHaveBeenCalled();
    expect(page.fromBlock).toBe(from);
    expect(page.events.map((e) => e.blockNumber)).toEqual([100_000n]);
  });

  it("scans and persists the confirmed range when not covered", async () => {
    const getLogs = vi.fn().mockResolvedValue([log(1n, 95_000n, 0)]);
    const mem = memCache();

    await fetchActivityPageCached(client(getLogs), sources, CHAIN, 100_000n, 200_000n, mem.deps);

    expect(getLogs).toHaveBeenCalledTimes(1);
    expect(mem.get()).toHaveLength(1);
    expect(mem.get()[0].high).toBe(100_000n); // toBlock <= confirmedTip -> full range persisted
  });

  it("does not persist the unconfirmed tail and filters unconfirmed events", async () => {
    // toBlock 100000 is above confirmedTip 99000: persist only [fromBlock, 99000].
    const getLogs = vi.fn().mockResolvedValue([log(1n, 99_500n, 0), log(2n, 98_000n, 0)]);
    const mem = memCache();

    await fetchActivityPageCached(client(getLogs), sources, CHAIN, 100_000n, 99_000n, mem.deps);

    expect(mem.get()).toHaveLength(1);
    expect(mem.get()[0].high).toBe(99_000n);
    expect(mem.get()[0].events.map((e) => e.blockNumber)).toEqual([98_000n]); // 99_500 (unconfirmed) excluded
  });

  it("persists nothing when the whole batch is unconfirmed", async () => {
    const getLogs = vi.fn().mockResolvedValue([log(1n, 100_000n, 0)]);
    const mem = memCache();

    await fetchActivityPageCached(client(getLogs), sources, CHAIN, 100_000n, 50_000n, mem.deps);

    expect(mem.get()).toHaveLength(0); // fromBlock (90001) > confirmedTip (50000)
  });
});
```

- [ ] **Step 3: Run the tests to verify they fail**

Run: `cd ui && ./node_modules/.bin/vitest run src/hooks/useActivity.test.ts`
Expected: FAIL — `fetchActivityPageCached` and `currentBatchSize` are not exported.

- [ ] **Step 4: Implement `currentBatchSize` + `fetchActivityPageCached`**

In `ui/src/hooks/useActivity.ts`:

1. Extend imports — add the config and cache imports:

```ts
import {
  ACTIVITY_BATCHES_PER_CLICK,
  ACTIVITY_CACHE_VERSION,
  ACTIVITY_CATCHUP_MAX_BATCHES,
  ACTIVITY_CONFIRMATIONS,
  ACTIVITY_MIN_BATCH_BLOCKS,
  batchBlocksFor,
} from "../config/activity";
import {
  cachedHigh,
  eventsInRange,
  isRangeCovered,
  mergeSegment,
  readCache,
  writeCache,
  type CacheEntry,
  type Segment,
} from "../lib/activity/cache";
```

2. Mark `ActivityPage` as exported (change `interface ActivityPage` to `export interface ActivityPage`).

3. Add, after `fetchActivityPage`:

```ts
/** The adaptive batch size currently in effect for a chain (remembered, else configured). */
export function currentBatchSize(chainId: number): bigint {
  return lastWorkingBatch.get(chainId) ?? batchBlocksFor(chainId);
}

export interface CachedFetchDeps {
  readSegments: () => Segment[];
  persist: (seg: Segment) => void;
}

/**
 * Cache-aware page fetch. If the batch window is fully confirmed and already
 * covered by the persisted cache, serve it without an RPC call. Otherwise scan
 * via fetchActivityPage and persist the confirmed sub-range (≤ confirmedTip) and
 * its confirmed events. The unconfirmed tail is never persisted.
 */
export async function fetchActivityPageCached(
  client: Pick<PublicClient, "getLogs">,
  sources: ActivitySourceInput[],
  chainId: number,
  toBlock: bigint,
  confirmedTip: bigint,
  deps: CachedFetchDeps,
): Promise<ActivityPage> {
  const size = currentBatchSize(chainId);
  const windowFrom = toBlock > size - 1n ? toBlock - size + 1n : 0n;

  if (toBlock <= confirmedTip && isRangeCovered(deps.readSegments(), windowFrom, toBlock)) {
    const events = eventsInRange(deps.readSegments(), windowFrom, toBlock).sort(compareEventsDesc);
    return { events, fromBlock: windowFrom };
  }

  const page = await fetchActivityPage(client, sources, chainId, toBlock);

  // Persist only the confirmed slice [page.fromBlock, min(toBlock, confirmedTip)].
  const persistHigh = toBlock < confirmedTip ? toBlock : confirmedTip;
  if (persistHigh >= page.fromBlock) {
    const confirmedEvents = page.events.filter((e) => e.blockNumber <= confirmedTip);
    deps.persist({ low: page.fromBlock, high: persistHigh, events: confirmedEvents });
  }
  return page;
}
```

- [ ] **Step 5: Run the tests to verify they pass**

Run: `cd ui && ./node_modules/.bin/vitest run src/hooks/useActivity.test.ts`
Expected: PASS (the four `fetchActivityPageCached` cases plus all existing cases green).

- [ ] **Step 6: Wire hydrate + cache-aware queryFn + auto catch-up into the hook**

In `ui/src/hooks/useActivity.ts`, update the `useActivity` function body:

1. Add `useEffect` and `useRef` to the React import:

```ts
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
```

2. Inside `useActivity`, after `const sourcesKey = useMemo(...)`, add the synchronous hydrate:

```ts
  // Synchronous hydrate from persisted cache so deep history renders instantly.
  // Re-reads when the (chainId, sources) key changes.
  const hydrated = useMemo<ActivityEvent[]>(() => {
    const entry = readCache(chainId, sourcesKey);
    return entry ? dedupeById(entry.segments.flatMap((s) => s.events)).sort(compareEventsDesc) : [];
  }, [chainId, sourcesKey]);
```

3. Replace the `queryFn` in the `useInfiniteQuery` config with the cache-aware version:

```ts
    queryFn: async ({ pageParam }) => {
      const c = client!;
      const tip = await c.getBlockNumber();
      const toBlock = pageParam ?? tip;
      const confirmedTip = tip > ACTIVITY_CONFIRMATIONS ? tip - ACTIVITY_CONFIRMATIONS : 0n;
      return fetchActivityPageCached(c, sources, chainId, toBlock, confirmedTip, {
        readSegments: () => readCache(chainId, sourcesKey)?.segments ?? [],
        persist: (seg) => {
          const entry: CacheEntry = readCache(chainId, sourcesKey) ?? { version: ACTIVITY_CACHE_VERSION, segments: [] };
          writeCache(chainId, sourcesKey, { ...entry, segments: mergeSegment(entry.segments, seg) });
        },
      });
    },
```

4. Replace the `events` memo so hydrated events merge with the query pages:

```ts
  const events = useMemo(() => {
    const fromQuery = (query.data?.pages ?? []).flatMap((p) => p.events);
    return dedupeById([...fromQuery, ...hydrated]).sort(compareEventsDesc);
  }, [query.data, hydrated]);
```

5. After the existing `loadOlder` `useCallback`, add the bounded auto catch-up. It runs once per `(chainId, sourcesKey)` and only for returning visitors (cache non-empty), bridging the gap by pulling up to `ACTIVITY_CATCHUP_MAX_BATCHES` backward batches (covered batches short-circuit cheaply):

```ts
  // Bounded auto catch-up on return: bridge cached-high -> tip without a click.
  const caughtUpKey = useRef<string | null>(null);
  useEffect(() => {
    const key = `${chainId}:${sourcesKey}`;
    if (hydrated.length === 0) return; // first-ever visit: behave as before (page 0 only)
    if (caughtUpKey.current === key) return; // once per key
    if (!query.isSuccess || query.isFetchingNextPage) return; // wait for page 0
    caughtUpKey.current = key;
    void loadOlderBatches(query.fetchNextPage, ACTIVITY_CATCHUP_MAX_BATCHES);
  }, [hydrated, query.isSuccess, query.isFetchingNextPage, query.fetchNextPage, chainId, sourcesKey]);
```

- [ ] **Step 7: Run the full hook + cache suite**

Run: `cd ui && ./node_modules/.bin/vitest run src/hooks/useActivity.test.ts src/lib/activity/cache.test.ts`
Expected: PASS (all cases green).

- [ ] **Step 8: Type-check and lint the touched files**

Run: `cd ui && yarn build`
Expected: `tsc -b` passes with no type errors and Vite build completes (confirms the hook wiring type-checks end to end).

- [ ] **Step 9: Commit**

```bash
cd ui && git add src/config/activity.ts src/hooks/useActivity.ts src/hooks/useActivity.test.ts
git commit -m "feat(ui): persist activity feed to localStorage with bounded catch-up"
```

---

## Manual verification (after Task 3)

Run `cd ui && yarn dev`, open the Activity page on a chain with deployed contracts, then:

1. Click "Load older" a few times to scan history, then **reload** — the feed should repaint history instantly (hydrated from cache) instead of scanning from scratch.
2. In DevTools → Application → Local Storage, confirm a `cm:activity:<chainId>:<sources>` key exists with segment data.
3. With network throttling / the Network tab open, reload again — confirm far fewer `eth_getLogs` requests fire than on the first load (covered batches short-circuit).
4. Confirm the newest events still appear after reload (live tail re-scanned), and no events are duplicated.
