// Block-range sizing for the Activity feed's eth_getLogs batches. The feed reads
// the chain in bounded windows so free-tier RPCs accept the query (see
// ui/CLAUDE.md). Each "Load older" click fetches one batch further back.

/** Starting batch size per chainId — the known-safe range for most public RPCs. */
export const ACTIVITY_BATCH_BLOCKS: Record<number, bigint> = {
  500: 10_000n, // Camino
  8453: 10_000n, // Base
  84532: 10_000n, // Base Sepolia
};

/** Fallback when a chain has no explicit entry. */
export const DEFAULT_BATCH_BLOCKS = 10_000n;

/**
 * Floor for adaptive shrinking. When a getLogs batch is rejected for range
 * limits we halve and retry down to this size before giving up.
 */
export const ACTIVITY_MIN_BATCH_BLOCKS = 500n;

/**
 * How many batches a single "Load older" click pulls, so the user looks much
 * further back without clicking repeatedly. 10 × 10k blocks ≈ 100k blocks
 * (~2 days on 2s-block chains like Base).
 */
export const ACTIVITY_BATCHES_PER_CLICK = 10;

export function batchBlocksFor(chainId: number): bigint {
  return ACTIVITY_BATCH_BLOCKS[chainId] ?? DEFAULT_BATCH_BLOCKS;
}

/** Bump when the persisted cache shape changes; older entries are discarded on read. */
export const ACTIVITY_CACHE_VERSION = 1;

/**
 * Max events kept per cache entry. When exceeded, the oldest events are dropped
 * (see capEntry) to stay well under the ~5MB localStorage origin budget.
 */
export const ACTIVITY_CACHE_MAX_EVENTS = 2000;
