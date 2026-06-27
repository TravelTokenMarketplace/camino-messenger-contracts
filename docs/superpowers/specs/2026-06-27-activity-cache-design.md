# Activity feed persistent cache — design

**Date:** 2026-06-27
**Status:** Approved, ready for planning
**Area:** `ui/` — the ecosystem activity feed

## Problem

The activity feed (`ui/src/hooks/useActivity.ts`) reads contract events from the
chain with `eth_getLogs`, paginating backwards in ~10k-block batches
(`useInfiniteQuery`). Page 0 ends at the latest block; each "Load older" click
scans the next batch further back. All results live only in TanStack Query's
in-memory cache, so every reload re-scans the same block ranges from scratch.

Events on finalized blocks are immutable, so re-scanning them is pure waste. We
want to persist scanned events to the browser so returning visitors:

1. **Load faster** — render cached history instantly instead of waiting for
   getLogs batches.
2. **Make fewer RPC calls** — never re-scan a block range already covered,
   easing load on the app-owned free-tier RPC.
3. **Retain deeper history** — accumulate more history across visits than a
   single session's "Load older" clicks would fetch.
4. **Auto-catch-up on return** — on reopening, scan only the new blocks since
   last visit rather than requiring a manual click.

## Constraints & decisions

- **Reorg safety: confirmation buffer.** Only blocks at or below
  `tip − confirmations` are persisted. The unconfirmed tail above that is always
  re-scanned live each visit and never written to cache.
- **Storage backend: `localStorage` + cap.** Simple key-per-`(chain, sources)`.
  Bound the stored data so we stay well under the ~5MB origin budget.
- **Catch-up: bounded auto, then manual.** Auto-scan the gap up to a cap (reuse
  the existing ~100k-block "Load older" span). If the gap is larger, fill the
  most recent window and let the user click "Load older" to continue.

## Architecture

### 1. Core data model — `ui/src/lib/activity/cache.ts` (pure, no React)

Keyed **per `(chainId, sourcesKey)`** — the same `source:address` key
`useActivity` already builds. Stored value is a set of **contiguous scanned
segments**:

```ts
type Segment = { low: bigint; high: bigint; events: ActivityEvent[] }; // [low,high] fully scanned, inclusive
type CacheEntry = { version: number; segments: Segment[] };            // usually exactly 1 segment
```

**Why segments rather than a single "newest cached block":** the bounded
catch-up cap means a long absence can leave a hole. When the gap from the cached
high block to the tip exceeds the cap, we scan the *recent* window (users want
newest-first) and keep the older cached history as a *separate, older* segment
with a visible gap between them. "Load older" then scans the hole and **merges
adjacent segments** when they meet. The normal case is a single segment; the
segment model just keeps the edge case correct instead of silently dropping
history.

Pure, unit-testable functions:

- `readCache(chainId, sourcesKey): CacheEntry | null`
- `writeCache(chainId, sourcesKey, entry): void`
- `mergeSegment(segments, segment): Segment[]` — insert a freshly-scanned range,
  coalescing overlapping/adjacent segments; events deduped by `id`.
- `capEntry(entry, maxEvents): CacheEntry` — evict oldest events past the bound,
  shrinking the oldest segment's `low` accordingly.
- `serialize` / `deserialize` — bigint-safe JSON (`blockNumber` is a `bigint`).

The cache is a pure optimization and **never a source of truth**: any parse
failure, schema-version mismatch, or corruption discards that entry and falls
back to a live scan.

### 2. Config additions — `ui/src/config/activity.ts`

- `ACTIVITY_CONFIRMATIONS` (e.g. `64n`) — only blocks `≤ tip − confirmations`
  are persisted; the tail above is live-only.
- `ACTIVITY_CATCHUP_MAX_BATCHES` — auto-catch-up cap; reuse the existing
  `ACTIVITY_BATCHES_PER_CLICK` count (~100k-block span).
- `ACTIVITY_CACHE_MAX_EVENTS` (e.g. `2000`) per entry — when exceeded, drop
  oldest events to stay well under the localStorage budget.
- `ACTIVITY_CACHE_VERSION` — schema version for invalidation on shape changes.

### 3. Load flow in `useActivity`

On mount, for the active `(chainId, sources)`:

1. **Hydrate** events synchronously from cache → feed renders instantly.
2. **Catch-up**: compute `confirmedTip = tip − confirmations`. Scan
   `(cachedHigh, confirmedTip]` forward, capped at `ACTIVITY_CATCHUP_MAX_BATCHES`,
   reusing `fetchActivityPage`'s adaptive batch sizing. Merge the scanned range
   into the cache and prepend its events to the view.
3. **Live tail**: always scan `(confirmedTip, tip]` fresh; show but never
   persist these events.
4. **Load older** (unchanged UX): backward batches **short-circuit** to the
   cache for ranges already covered, hitting `eth_getLogs` only for uncovered
   blocks. Newly scanned confirmed ranges are written back via `mergeSegment`.

Existing query invalidation (RefreshButton / tx `onConfirmed`) continues to
work — it re-runs catch-up + live tail, not the full history.

### 4. Edge cases

- **Schema bump / corrupt entry** → discard that entry, re-scan. Cache is never
  authoritative.
- **`QuotaExceededError` on write** → apply `capEntry` more aggressively, then
  skip persisting; the feed still works fully in-memory.
- **Empty scanned ranges** (a window with no matching events) are still recorded
  as covered, so we don't re-scan them.
- **Reorg within the buffer** — by construction we never persist blocks newer
  than `tip − confirmations`, and the live tail is always re-scanned, so a reorg
  shallower than the buffer can't leave a stale cached event.

### 5. Testing

Following the existing pattern (`useActivity.ts` factors logic into pure
exported functions tested in `useActivity.test.ts`):

- `ui/src/lib/activity/cache.test.ts` — bulk of coverage: `mergeSegment`
  coalescing (overlapping, adjacent, disjoint), `capEntry` eviction and `low`
  adjustment, bigint serialize/deserialize round-trip, version-mismatch and
  corruption handling.
- Additions to `useActivity.test.ts` — hydrate-then-catch-up with a mocked
  client: assert cached events render first, only the gap is scanned, the live
  tail is re-scanned and not persisted, and "Load older" short-circuits covered
  ranges.

## Out of scope (YAGNI)

- Caching mutable reads (balances, roles) — only immutable confirmed events.
- IndexedDB — `localStorage` + cap is sufficient for the realistic event volume.
- Cross-tab cache synchronization.
- Tuning per-chain confirmation depths — a single safe default for now.
