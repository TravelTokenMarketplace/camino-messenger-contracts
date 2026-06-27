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
