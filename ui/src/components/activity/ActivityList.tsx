import { useMemo, useState } from "react";
import { type ActivityCategory, type ActivityEvent } from "../../lib/activity/types";
import { CATEGORY_STYLE } from "../../lib/activity/style";
import { ActivityRow } from "./ActivityRow";

interface ActivityListProps {
  events: ActivityEvent[];
  isLoading: boolean;
  error?: Error | null;
  explorerUrl?: string;
  /** Cap the number of rows shown (Dashboard sneak peek). */
  limit?: number;
  /** Render category filter chips (Activity page). */
  showFilters?: boolean;
  /** "Load older" pagination — omit to hide the button (Dashboard). */
  hasNextPage?: boolean;
  isFetchingNextPage?: boolean;
  onLoadOlder?: () => void;
  oldestBlockLoaded?: bigint;
  /** Override the empty-state hint. */
  emptyHint?: string;
}

export function ActivityList({
  events,
  isLoading,
  error,
  explorerUrl,
  limit,
  showFilters = false,
  hasNextPage,
  isFetchingNextPage,
  onLoadOlder,
  oldestBlockLoaded,
  emptyHint = "No activity in the last 10,000 blocks.",
}: ActivityListProps) {
  const [active, setActive] = useState<Set<ActivityCategory>>(new Set());

  const categories = useMemo(() => {
    const set = new Set<ActivityCategory>();
    for (const e of events) set.add(e.category);
    return [...set];
  }, [events]);

  const filtered = active.size ? events.filter((e) => active.has(e.category)) : events;
  const shown = limit != null ? filtered.slice(0, limit) : filtered;

  function toggle(cat: ActivityCategory) {
    setActive((prev) => {
      const next = new Set(prev);
      if (next.has(cat)) next.delete(cat);
      else next.add(cat);
      return next;
    });
  }

  const isEmpty = events.length === 0;

  let body;
  if (error && isEmpty) {
    body = (
      <p className="py-2 text-sm text-amber-600 dark:text-amber-400">
        Couldn&apos;t load activity from this RPC. Try Refresh.
      </p>
    );
  } else if (isLoading && isEmpty) {
    body = <p className="py-2 text-sm text-gray-400 dark:text-gray-500">Loading…</p>;
  } else if (isEmpty) {
    body = <p className="py-2 text-sm text-gray-400 dark:text-gray-500">{emptyHint}</p>;
  } else {
    body = (
      <>
        {showFilters && categories.length > 1 && (
          <div className="mb-3 flex flex-wrap gap-1.5">
            {categories.map((cat) => {
              const on = active.has(cat);
              return (
                <button
                  key={cat}
                  type="button"
                  aria-pressed={on}
                  onClick={() => toggle(cat)}
                  className={`inline-flex items-center gap-1.5 rounded-full border px-2.5 py-0.5 text-xs transition-colors ${
                    on
                      ? "border-indigo-600 bg-indigo-50 text-indigo-700 dark:border-indigo-400 dark:bg-indigo-500/10 dark:text-indigo-300"
                      : "border-gray-300 text-gray-500 hover:bg-gray-50 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-800"
                  }`}
                >
                  <span className={`h-1.5 w-1.5 rounded-full ${CATEGORY_STYLE[cat].dot}`} aria-hidden />
                  {cat}
                </button>
              );
            })}
          </div>
        )}

        <ul className="divide-y divide-gray-100 dark:divide-gray-700/50">
          {shown.map((e) => (
            <ActivityRow key={e.id} event={e} explorerUrl={explorerUrl} />
          ))}
        </ul>
      </>
    );
  }

  return (
    <div>
      {body}

      {onLoadOlder && (
        <div className="mt-3 flex items-center justify-between gap-3 text-xs text-gray-400 dark:text-gray-500">
          <span>{oldestBlockLoaded != null && `Scanned back to block ${oldestBlockLoaded.toString()}`}</span>
          <button
            type="button"
            onClick={onLoadOlder}
            disabled={!hasNextPage || isFetchingNextPage}
            title="Scans several batches further back through history"
            className="rounded-md border border-gray-300 px-2.5 py-1 text-gray-600 transition-colors hover:bg-gray-50 disabled:opacity-50 dark:border-gray-700 dark:text-gray-300 dark:hover:bg-gray-800"
          >
            {isFetchingNextPage ? "Loading…" : hasNextPage ? "Load older" : "No more history"}
          </button>
        </div>
      )}
    </div>
  );
}
