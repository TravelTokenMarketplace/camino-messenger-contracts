import { type ActivityEvent } from "./types";

/** Newest first: higher block, then higher logIndex. */
export function compareEventsDesc(a: ActivityEvent, b: ActivityEvent): number {
  if (a.blockNumber !== b.blockNumber) return a.blockNumber > b.blockNumber ? -1 : 1;
  return b.logIndex - a.logIndex;
}

/** Dedupe by stable id (`${txHash}#${logIndex}`), preserving first-seen order. */
export function dedupeById(events: ActivityEvent[]): ActivityEvent[] {
  const seen = new Set<string>();
  const out: ActivityEvent[] = [];
  for (const e of events) {
    if (seen.has(e.id)) continue;
    seen.add(e.id);
    out.push(e);
  }
  return out;
}
