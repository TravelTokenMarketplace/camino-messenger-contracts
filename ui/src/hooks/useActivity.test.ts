import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { renderHook, waitFor } from "@testing-library/react";
import { createElement } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  __resetBatchMemory,
  compareEventsDesc,
  currentBatchSize,
  fetchActivityPage,
  fetchActivityPageCached,
  loadOlderBatches,
  useActivity,
  type ActivitySourceInput,
} from "./useActivity";
import { writeCache, type CacheEntry, type Segment } from "../lib/activity/cache";
import { ACTIVITY_CACHE_VERSION } from "../config/activity";
import { BOOKING_TOKEN_EVENTS } from "../lib/activity/catalog";
import { type ActivityEvent } from "../lib/activity/types";

// ---------------------------------------------------------------------------
// Wagmi mock — scoped to this module so hook tests can control usePublicClient.
// Pure-function tests (fetchActivityPage, fetchActivityPageCached, etc.) don't
// call usePublicClient and are unaffected by this mock.
// ---------------------------------------------------------------------------
const mockGetBlockNumber = vi.hoisted(() => vi.fn());
const mockGetLogs = vi.hoisted(() => vi.fn());

vi.mock("wagmi", () => ({
  usePublicClient: () => ({ getBlockNumber: mockGetBlockNumber, getLogs: mockGetLogs }),
}));

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

// ---------------------------------------------------------------------------
// useActivity hook — wiring tests (hydrate + catch-up).
// Each test gets a fresh QueryClient and clean localStorage.
// ---------------------------------------------------------------------------

/** Minimal valid ActivityEvent for pre-seeding the cache. */
function hookEvent(blockNumber: bigint, logIndex: number, tokenId: bigint): ActivityEvent {
  return {
    id: `${blockNumber}#${logIndex}`,
    source: "bookingToken",
    category: "Bookings",
    contract: ADDR as `0x${string}`,
    blockNumber,
    logIndex,
    txHash: `0xtest${blockNumber}${logIndex}` as `0x${string}`,
    eventName: "TokenBought",
    args: { tokenId, buyer: "0x0000000000000000000000000000000000000001" as `0x${string}` },
    sentence: "bought",
  };
}

/** Cache entry covering [1, 100] with one event at block 100. */
function seedEntry(): CacheEntry {
  return { version: ACTIVITY_CACHE_VERSION, segments: [{ low: 1n, high: 100n, events: [hookEvent(100n, 0, 7n)] }] };
}

/** Fresh QueryClient + Provider wrapper per test so query state never leaks. */
function makeWrapper() {
  const qc = new QueryClient({ defaultOptions: { queries: { retry: false } } });
  return ({ children }: { children: React.ReactNode }) => createElement(QueryClientProvider, { client: qc }, children);
}

// sourcesKey derived the same way the hook does it internally
const SOURCES_KEY = `bookingToken:${ADDR}`;

describe("useActivity hook", () => {
  beforeEach(() => {
    localStorage.clear();
    mockGetBlockNumber.mockReset();
    mockGetLogs.mockReset();
    __resetBatchMemory();
  });

  it("a. hydrate: cached events appear on initial render without waiting for RPC", () => {
    // Pre-seed cache so readCache returns data synchronously in the hook's useMemo.
    writeCache(CHAIN, SOURCES_KEY, seedEntry());

    // Allow the query to start (it won't be awaited for this assertion).
    mockGetBlockNumber.mockResolvedValue(200n);
    mockGetLogs.mockResolvedValue([]);

    const { result } = renderHook(() => useActivity({ sources, chainId: CHAIN }), { wrapper: makeWrapper() });

    // Hydrated events are folded in synchronously on the first render — no await.
    // Guards the useMemo hydrate + the merge into the events array.
    const found = result.current.events.find((e) => e.blockNumber === 100n);
    expect(found).toBeDefined();
    expect((found?.args as { tokenId?: bigint }).tokenId).toBe(7n);
  });

  it("b. catch-up: fires once per key and re-arms when chainId changes", async () => {
    // Use a chain ID distinct from CHAIN to test re-arm.
    const CHAIN_B = 8453;

    // Pre-seed both keys so hydrated.length > 0, which is the catch-up gate.
    writeCache(CHAIN, SOURCES_KEY, seedEntry());
    writeCache(CHAIN_B, SOURCES_KEY, seedEntry());

    // Large enough tip so fromBlock > 0 after page 0, meaning hasNextPage=true
    // and catch-up actually fetches additional pages.
    const TIP = 50_000n;
    mockGetBlockNumber.mockResolvedValue(TIP);
    mockGetLogs.mockResolvedValue([]);

    const wrapper = makeWrapper();
    const { result, rerender } = renderHook(({ chainId }: { chainId: number }) => useActivity({ sources, chainId }), {
      wrapper,
      initialProps: { chainId: CHAIN },
    });

    // Wait for page 0 to succeed; then catch-up fires (hydrated.length > 0 and key unset).
    // Catch-up fetches additional pages → getBlockNumber is called more than once.
    // Guards: caughtUpKey stays unset until page 0 succeeds, then loadOlderBatches runs.
    await waitFor(() => expect(mockGetBlockNumber.mock.calls.length).toBeGreaterThan(1), { timeout: 5000 });

    // Let catch-up drain so isFetchingNextPage settles before we record the count.
    await waitFor(() => expect(result.current.isFetchingNextPage).toBe(false), { timeout: 5000 });
    const countAfterChainA = mockGetBlockNumber.mock.calls.length;

    // Re-render with the same chainId — caughtUpKey is already set, so no new batches.
    // Guards: the once-per-key guard on caughtUpKey ref.
    rerender({ chainId: CHAIN });
    await waitFor(() => expect(result.current.isFetchingNextPage).toBe(false), { timeout: 2000 });
    expect(mockGetBlockNumber.mock.calls.length).toBe(countAfterChainA);

    // Switch to a new chain — fresh key, caughtUpKey must re-arm.
    // Guards: key change clears the once-per-key guard (new key !== caughtUpKey.current).
    rerender({ chainId: CHAIN_B });
    await waitFor(() => expect(mockGetBlockNumber.mock.calls.length).toBeGreaterThan(countAfterChainA + 1), {
      timeout: 5000,
    });
  });

  it("c. first-ever visit (empty cache) does not trigger auto catch-up", async () => {
    // No cache seeded — hydrated will be [] and the catch-up gate returns early.
    mockGetBlockNumber.mockResolvedValue(50_000n);
    mockGetLogs.mockResolvedValue([]);

    const { result } = renderHook(() => useActivity({ sources, chainId: CHAIN }), { wrapper: makeWrapper() });

    // Wait for page 0 to complete and isFetchingNextPage to clear.
    await waitFor(() => expect(result.current.isLoading).toBe(false), { timeout: 3000 });
    await waitFor(() => expect(result.current.isFetchingNextPage).toBe(false), { timeout: 3000 });

    // Exactly one getBlockNumber call (for page 0 only); catch-up must not have fired.
    // Guards: the `if (hydrated.length === 0) return` guard in the catch-up effect.
    expect(mockGetBlockNumber).toHaveBeenCalledTimes(1);
  });
});
