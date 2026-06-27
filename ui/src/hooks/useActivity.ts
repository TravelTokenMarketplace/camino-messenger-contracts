import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { usePublicClient } from "wagmi";
import { type AbiEvent, type Address, type PublicClient } from "viem";
import {
  ACTIVITY_BATCHES_PER_CLICK,
  ACTIVITY_CACHE_VERSION,
  ACTIVITY_CATCHUP_MAX_BATCHES,
  ACTIVITY_CONFIRMATIONS,
  ACTIVITY_MIN_BATCH_BLOCKS,
  batchBlocksFor,
} from "../config/activity";
import { toActivityEvent } from "../lib/activity/catalog";
import {
  eventsInRange,
  isRangeCovered,
  mergeSegment,
  readCache,
  writeCache,
  type CacheEntry,
  type Segment,
} from "../lib/activity/cache";
import { compareEventsDesc, dedupeById } from "../lib/activity/sort";
import { type ActivityEvent, type ActivitySource } from "../lib/activity/types";

export { compareEventsDesc, dedupeById };

export interface ActivitySourceInput {
  source: ActivitySource;
  address: Address;
  events: AbiEvent[];
}

export interface ActivityPage {
  events: ActivityEvent[];
  /** Lowest block actually scanned for this page. */
  fromBlock: bigint;
}

/**
 * Session memory of the largest getLogs range a chain accepted. Once an RPC
 * reveals a smaller real cap (by rejecting a wider range), later batches start
 * from that size instead of re-failing from the configured default each time.
 */
const lastWorkingBatch = new Map<number, bigint>();

/** Test-only: clear adaptive batch-size memory between cases. */
export function __resetBatchMemory() {
  lastWorkingBatch.clear();
}

/**
 * Whether a getLogs failure looks like an RPC block-range / result-size cap
 * (worth retrying with a smaller batch) rather than a transient fault — a
 * timeout, 429, 503 — or a bad request, which should surface immediately instead
 * of triggering a pointless halving spiral.
 */
export function isRangeLimitError(err: unknown): boolean {
  const msg = err instanceof Error ? err.message : String(err);
  // "block range too large", "max block range exceeded", "more than N results",
  // "query exceeds limit" — but NOT "429 Too Many Requests" (transient throttle).
  return /\b(range|results?|limit|exceed(?:s|ed)?|too\s+(?:large|wide|big))\b/i.test(msg);
}

/**
 * Fetch one page (batch) of activity ending at `toBlock`, scanning backwards.
 * Adaptive sizing: if any source's getLogs rejects (range/limit error), the
 * batch is halved and retried for all sources, down to a floor, before the
 * error propagates. The largest size that succeeds is remembered per chain.
 */
export async function fetchActivityPage(
  client: Pick<PublicClient, "getLogs">,
  sources: ActivitySourceInput[],
  chainId: number,
  toBlock: bigint,
): Promise<ActivityPage> {
  let size = lastWorkingBatch.get(chainId) ?? batchBlocksFor(chainId);

  for (;;) {
    const fromBlock = toBlock > size - 1n ? toBlock - size + 1n : 0n;
    try {
      const perSource = await Promise.all(
        sources.map(async (s) => {
          const logs = await client.getLogs({ address: s.address, events: s.events, fromBlock, toBlock });
          return logs.map((log) => toActivityEvent(log as never, s.source)).filter((e): e is ActivityEvent => !!e);
        }),
      );
      lastWorkingBatch.set(chainId, size);
      const events = dedupeById(perSource.flat()).sort(compareEventsDesc);
      return { events, fromBlock };
    } catch (err) {
      // Only a range/size cap is worth retrying smaller; rethrow everything else
      // (timeouts, throttling, provider outages) so it isn't masked as "too wide".
      if (!isRangeLimitError(err) || size <= ACTIVITY_MIN_BATCH_BLOCKS) throw err;
      const halved = size / 2n;
      size = halved < ACTIVITY_MIN_BATCH_BLOCKS ? ACTIVITY_MIN_BATCH_BLOCKS : halved;
    }
  }
}

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

  const segments = deps.readSegments();
  if (toBlock <= confirmedTip && isRangeCovered(segments, windowFrom, toBlock)) {
    const events = eventsInRange(segments, windowFrom, toBlock).sort(compareEventsDesc);
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

/**
 * Pull up to `maxBatches` pages, stopping early once there is no further
 * history. Extracted from the hook so the multi-batch loop is unit-testable.
 */
export async function loadOlderBatches(
  fetchNextPage: () => Promise<{ hasNextPage: boolean }>,
  maxBatches: number = ACTIVITY_BATCHES_PER_CLICK,
): Promise<void> {
  for (let i = 0; i < maxBatches; i++) {
    const res = await fetchNextPage();
    if (!res.hasNextPage) break;
  }
}

export interface UseActivityResult {
  events: ActivityEvent[];
  /** Pull several batches further back in one call (see ACTIVITY_BATCHES_PER_CLICK). */
  loadOlder: () => void;
  hasNextPage: boolean;
  isLoading: boolean;
  /** True while a loadOlder() run (one or more batches) is in flight. */
  isFetchingNextPage: boolean;
  error: Error | null;
  /** Lowest block scanned across all loaded pages. */
  oldestBlockLoaded?: bigint;
}

/**
 * Cursor-paginated activity feed. Page 0 ends at the latest block; each
 * `fetchNextPage` scans the next batch further back. Keyed by chainId + sources
 * so the app's RefreshButton (query invalidation) refetches it. No polling.
 */
export function useActivity({
  sources,
  chainId,
}: {
  sources: ActivitySourceInput[];
  chainId: number;
}): UseActivityResult {
  const client = usePublicClient({ chainId });
  const sourcesKey = useMemo(() => sources.map((s) => `${s.source}:${s.address}`).join(","), [sources]);

  // Synchronous hydrate from persisted cache so deep history renders instantly.
  // Re-reads when the (chainId, sources) key changes.
  const hydrated = useMemo<ActivityEvent[]>(() => {
    const entry = readCache(chainId, sourcesKey);
    return entry ? dedupeById(entry.segments.flatMap((s) => s.events)).sort(compareEventsDesc) : [];
  }, [chainId, sourcesKey]);

  const query = useInfiniteQuery({
    queryKey: ["activity", chainId, sourcesKey],
    enabled: Boolean(client) && sources.length > 0,
    initialPageParam: undefined as bigint | undefined,
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
    getNextPageParam: (lastPage) => (lastPage.fromBlock > 0n ? lastPage.fromBlock - 1n : undefined),
  });

  const events = useMemo(() => {
    const fromQuery = (query.data?.pages ?? []).flatMap((p) => p.events);
    return dedupeById([...fromQuery, ...hydrated]).sort(compareEventsDesc);
  }, [query.data, hydrated]);

  const pages = query.data?.pages ?? [];
  const oldestBlockLoaded = pages.length ? pages[pages.length - 1].fromBlock : undefined;

  // One click pulls several batches so the user spans much more history at once.
  // Sequential so each page's adaptive batch size carries into the next.
  const { fetchNextPage } = query;
  const [isLoadingOlder, setIsLoadingOlder] = useState(false);
  const loadOlder = useCallback(async () => {
    setIsLoadingOlder(true);
    try {
      await loadOlderBatches(fetchNextPage);
    } finally {
      setIsLoadingOlder(false);
    }
  }, [fetchNextPage]);

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

  return {
    events,
    loadOlder,
    hasNextPage: query.hasNextPage,
    isLoading: query.isLoading,
    isFetchingNextPage: isLoadingOlder || query.isFetchingNextPage,
    error: query.error as Error | null,
    oldestBlockLoaded,
  };
}
