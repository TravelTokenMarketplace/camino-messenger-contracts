import { useQueries } from "@tanstack/react-query";
import { usePublicClient } from "wagmi";

/** Dedupe block numbers, preserving first-seen order. Exported for testing. */
export function dedupeBlockNumbers(blockNumbers: bigint[]): bigint[] {
  const seen = new Set<string>();
  const out: bigint[] = [];
  for (const bn of blockNumbers) {
    const k = bn.toString();
    if (seen.has(k)) continue;
    seen.add(k);
    out.push(bn);
  }
  return out;
}

/**
 * Resolve unix-second timestamps for a set of block numbers. Each block is a
 * separate query cached forever (block timestamps are immutable) and keyed by
 * chain + block, so the Dashboard card, Activity page and account tab never
 * refetch the same block.
 */
export function useBlockTimestamps(chainId: number, blockNumbers: bigint[]): Map<bigint, number> {
  const client = usePublicClient({ chainId });
  const unique = dedupeBlockNumbers(blockNumbers);

  const results = useQueries({
    queries: unique.map((bn) => ({
      queryKey: ["block-timestamp", chainId, bn.toString()],
      enabled: Boolean(client),
      // Timestamps are immutable, so never refetch; but each block is its own
      // query key, so cap retention to keep the cache from growing unbounded as
      // the user pages further back.
      staleTime: Infinity,
      gcTime: 5 * 60 * 1000,
      queryFn: async () => {
        const block = await client!.getBlock({ blockNumber: bn });
        return Number(block.timestamp);
      },
    })),
  });

  const map = new Map<bigint, number>();
  unique.forEach((bn, i) => {
    const ts = results[i]?.data;
    if (typeof ts === "number") map.set(bn, ts);
  });
  return map;
}
