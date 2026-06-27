import { afterEach, describe, expect, it, vi } from "vitest";
import {
  __resetBatchMemory,
  compareEventsDesc,
  currentBatchSize,
  fetchActivityPage,
  fetchActivityPageCached,
  loadOlderBatches,
  type ActivitySourceInput,
} from "./useActivity";
import { type Segment } from "../lib/activity/cache";
import { BOOKING_TOKEN_EVENTS } from "../lib/activity/catalog";
import { type ActivityEvent } from "../lib/activity/types";

const CHAIN = 84532;
const ADDR = "0xbookingtoken000000000000000000000000beef";

const sources: ActivitySourceInput[] = [{ source: "bookingToken", address: ADDR, events: BOOKING_TOKEN_EVENTS }];

// fetchActivityPage only needs a structural `getLogs`; cast fakes to its param type
// rather than satisfy viem's overloaded PublicClient signature.
type ActivityClient = Parameters<typeof fetchActivityPage>[0];
const client = (getLogs: unknown) => ({ getLogs }) as ActivityClient;

function log(tokenId: bigint, blockNumber: bigint, logIndex: number) {
  return {
    eventName: "TokenBought",
    args: { tokenId, buyer: "0x0000000000000000000000000000000000000001" },
    blockNumber,
    logIndex,
    transactionHash: `0x${blockNumber}${logIndex}`,
    address: ADDR,
  };
}

afterEach(() => {
  __resetBatchMemory();
  vi.restoreAllMocks();
});

describe("compareEventsDesc", () => {
  it("orders by block desc then logIndex desc", () => {
    const a = { blockNumber: 10n, logIndex: 1 } as ActivityEvent;
    const b = { blockNumber: 10n, logIndex: 2 } as ActivityEvent;
    const c = { blockNumber: 9n, logIndex: 5 } as ActivityEvent;
    expect([a, c, b].sort(compareEventsDesc).map((e) => e.logIndex)).toEqual([2, 1, 5]);
  });
});

describe("fetchActivityPage", () => {
  it("decodes, merges and sorts logs newest-first", async () => {
    const getLogs = vi.fn().mockResolvedValue([log(1n, 100n, 0), log(2n, 105n, 1), log(3n, 105n, 0)]);
    const page = await fetchActivityPage(client(getLogs), sources, CHAIN, 200n);

    expect(page.fromBlock).toBe(0n); // toBlock (200) < batch (10000) -> floored at 0
    expect(page.events.map((e) => e.blockNumber)).toEqual([105n, 105n, 100n]);
    expect(page.events.map((e) => e.logIndex)).toEqual([1, 0, 0]);
  });

  it("computes fromBlock for a full batch window", async () => {
    const getLogs = vi.fn().mockResolvedValue([]);
    const page = await fetchActivityPage(client(getLogs), sources, CHAIN, 50_000n);
    expect(page.fromBlock).toBe(50_000n - 10_000n + 1n);
    expect(getLogs).toHaveBeenCalledWith(expect.objectContaining({ fromBlock: 40_001n, toBlock: 50_000n }));
  });

  it("halves the range and retries when getLogs rejects, then remembers the size", async () => {
    // Reject any range wider than 2500 blocks; succeed otherwise.
    const getLogs = vi.fn(async ({ fromBlock, toBlock }: { fromBlock: bigint; toBlock: bigint }) => {
      if (toBlock - fromBlock + 1n > 2500n) throw new Error("range too wide");
      return [];
    });

    const page = await fetchActivityPage(client(getLogs), sources, CHAIN, 100_000n);
    // 10000 -> 5000 -> 2500 (first accepted): three attempts.
    expect(getLogs).toHaveBeenCalledTimes(3);
    expect(page.fromBlock).toBe(100_000n - 2500n + 1n);

    // Next page should START at the remembered 2500 size (one call, no re-halving).
    getLogs.mockClear();
    const page2 = await fetchActivityPage(client(getLogs), sources, CHAIN, 90_000n);
    expect(getLogs).toHaveBeenCalledTimes(1);
    expect(page2.fromBlock).toBe(90_000n - 2500n + 1n);
  });

  it("propagates a range error once it fails even at the floor range", async () => {
    const getLogs = vi.fn().mockRejectedValue(new Error("block range too large"));
    await expect(fetchActivityPage(client(getLogs), sources, CHAIN, 100_000n)).rejects.toThrow("range");
    // 10000 -> 5000 -> 2500 -> 1250 -> 625 -> 500 floor: six attempts, then give up.
    expect(getLogs).toHaveBeenCalledTimes(6);
  });

  it("does not halve on a non-range error — it surfaces immediately", async () => {
    const getLogs = vi.fn().mockRejectedValue(new Error("429 Too Many Requests"));
    await expect(fetchActivityPage(client(getLogs), sources, CHAIN, 100_000n)).rejects.toThrow("429");
    expect(getLogs).toHaveBeenCalledTimes(1);
  });
});

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

describe("loadOlderBatches", () => {
  it("pulls up to maxBatches pages in one call", async () => {
    const fetchNextPage = vi.fn().mockResolvedValue({ hasNextPage: true });
    await loadOlderBatches(fetchNextPage, 10);
    expect(fetchNextPage).toHaveBeenCalledTimes(10);
  });

  it("stops early when there is no more history", async () => {
    const fetchNextPage = vi
      .fn()
      .mockResolvedValueOnce({ hasNextPage: true })
      .mockResolvedValueOnce({ hasNextPage: true })
      .mockResolvedValueOnce({ hasNextPage: false });
    await loadOlderBatches(fetchNextPage, 10);
    expect(fetchNextPage).toHaveBeenCalledTimes(3);
  });
});
