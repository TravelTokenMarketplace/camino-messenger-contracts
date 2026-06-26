import { useMemo } from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { usePublicClient } from "wagmi";
import { type AbiEvent, type Address, type PublicClient } from "viem";
import { ACTIVITY_MIN_BATCH_BLOCKS, batchBlocksFor } from "../config/activity";
import { toActivityEvent } from "../lib/activity/catalog";
import { type ActivityEvent, type ActivitySource } from "../lib/activity/types";

export interface ActivitySourceInput {
  source: ActivitySource;
  address: Address;
  events: AbiEvent[];
}

interface ActivityPage {
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

/** Newest first: higher block, then higher logIndex. */
export function compareEventsDesc(a: ActivityEvent, b: ActivityEvent): number {
  if (a.blockNumber !== b.blockNumber) return a.blockNumber > b.blockNumber ? -1 : 1;
  return b.logIndex - a.logIndex;
}

function dedupeById(events: ActivityEvent[]): ActivityEvent[] {
  const seen = new Set<string>();
  const out: ActivityEvent[] = [];
  for (const e of events) {
    if (seen.has(e.id)) continue;
    seen.add(e.id);
    out.push(e);
  }
  return out;
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
      if (size <= ACTIVITY_MIN_BATCH_BLOCKS) throw err;
      const halved = size / 2n;
      size = halved < ACTIVITY_MIN_BATCH_BLOCKS ? ACTIVITY_MIN_BATCH_BLOCKS : halved;
    }
  }
}

export interface UseActivityResult {
  events: ActivityEvent[];
  fetchNextPage: () => void;
  hasNextPage: boolean;
  isLoading: boolean;
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

  const query = useInfiniteQuery({
    queryKey: ["activity", chainId, sourcesKey],
    enabled: Boolean(client) && sources.length > 0,
    initialPageParam: undefined as bigint | undefined,
    queryFn: async ({ pageParam }) => {
      const c = client!;
      const toBlock = pageParam ?? (await c.getBlockNumber());
      return fetchActivityPage(c, sources, chainId, toBlock);
    },
    getNextPageParam: (lastPage) => (lastPage.fromBlock > 0n ? lastPage.fromBlock - 1n : undefined),
  });

  const events = useMemo(() => {
    const all = (query.data?.pages ?? []).flatMap((p) => p.events);
    return dedupeById(all).sort(compareEventsDesc);
  }, [query.data]);

  const pages = query.data?.pages ?? [];
  const oldestBlockLoaded = pages.length ? pages[pages.length - 1].fromBlock : undefined;

  return {
    events,
    fetchNextPage: query.fetchNextPage,
    hasNextPage: query.hasNextPage,
    isLoading: query.isLoading,
    isFetchingNextPage: query.isFetchingNextPage,
    error: query.error as Error | null,
    oldestBlockLoaded,
  };
}
