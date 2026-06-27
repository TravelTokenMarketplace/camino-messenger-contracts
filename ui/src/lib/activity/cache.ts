import { ACTIVITY_CACHE_MAX_EVENTS, ACTIVITY_CACHE_VERSION } from "../../config/activity";
import { compareEventsDesc, dedupeById } from "./sort";
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
    // Validate segment shape: a partial/pre-tag write could pass the version
    // check yet carry non-bigint bounds that later throw in mergeSegment's
    // `low <= high + 1n` on the persist path — corruption must degrade to a live
    // scan, never break the feed.
    const shapeOk = parsed.segments.every(
      (s) => s && typeof s.low === "bigint" && typeof s.high === "bigint" && Array.isArray(s.events),
    );
    if (!shapeOk) return null;
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
    if (seg.events.length === 0) {
      out.push(seg);
      continue;
    } // empty coverage carries no evictable events
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
