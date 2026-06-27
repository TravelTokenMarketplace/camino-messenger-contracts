import { afterEach, describe, expect, it, vi } from "vitest";
import {
  cachedHigh,
  capEntry,
  deserializeEntry,
  eventsInRange,
  isRangeCovered,
  mergeSegment,
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

  it("returns null when a segment has non-bigint bounds (would later throw in mergeSegment)", () => {
    // Right version + segments array, but low/high are plain numbers (e.g. an
    // entry written before the $bigint tagging, or a partial write).
    const malformed = JSON.stringify({
      version: ACTIVITY_CACHE_VERSION,
      segments: [{ low: 1, high: 100, events: [] }],
    });
    expect(deserializeEntry(malformed)).toBeNull();
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
    expect(
      totalEvents({
        version: ACTIVITY_CACHE_VERSION,
        segments: [{ low: 1n, high: 9n, events: [ev(1n, 0, 1n), ev(2n, 0, 2n)] }],
      }),
    ).toBe(2);
  });
});

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

  it("preserves empty (no-event) segments while still evicting events from event-bearing ones", () => {
    const e: CacheEntry = {
      version: ACTIVITY_CACHE_VERSION,
      segments: [
        { low: 1n, high: 10n, events: [] }, // empty coverage — must survive eviction
        { low: 20n, high: 30n, events: [ev(25n, 0, 1n), ev(26n, 0, 2n)] },
        { low: 100n, high: 110n, events: [ev(105n, 0, 3n)] },
      ],
    };
    // Cap to 1 event: the two events in the second segment are the oldest; drop the earlier one.
    const capped = capEntry(e, 1);
    // The empty segment must still be present.
    expect(capped.segments.find((s) => s.low === 1n && s.high === 10n)).toBeDefined();
    // The last-segment event (newest) must survive.
    expect(capped.segments.find((s) => s.events.some((ev) => ev.blockNumber === 105n))).toBeDefined();
    // Total events must equal 1.
    expect(capped.segments.reduce((n, s) => n + s.events.length, 0)).toBe(1);
  });
});

describe("coverage helpers", () => {
  const segs = [
    { low: 10n, high: 20n, events: [ev(15n, 0, 1n)] },
    { low: 100n, high: 200n, events: [ev(150n, 0, 2n)] },
  ];

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
