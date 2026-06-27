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

// TEMP stub — replaced in Task 2.
function capEntry(entry: CacheEntry, _maxEvents: number): CacheEntry {
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
