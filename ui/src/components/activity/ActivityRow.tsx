import { Activity as ActivityIcon, ExternalLink } from "lucide-react";
import { lookupEntry } from "../../lib/activity/catalog";
import { CATEGORY_STYLE } from "../../lib/activity/style";
import { type ActivityEvent } from "../../lib/activity/types";
import { explorerTxUrl, formatRelativeTime } from "../../lib/format";
import { InlineSentence } from "./InlineSentence";

export function ActivityRow({ event, explorerUrl }: { event: ActivityEvent; explorerUrl?: string }) {
  const Icon = lookupEntry(event.source, event.eventName)?.icon ?? ActivityIcon;
  const when = event.timestamp != null ? formatRelativeTime(event.timestamp) : `block ${event.blockNumber.toString()}`;
  const absolute = event.timestamp != null ? new Date(event.timestamp * 1000).toLocaleString() : undefined;

  return (
    <li className="flex items-center gap-3 py-2">
      <span
        className={`inline-flex h-7 w-7 shrink-0 items-center justify-center rounded-full ${CATEGORY_STYLE[event.category].icon}`}
        title={event.category}
      >
        <Icon className="h-4 w-4" aria-hidden />
      </span>
      <span className="min-w-0 break-words text-sm text-gray-800 dark:text-gray-200">
        <InlineSentence sentence={event.sentence} args={event.args} />
      </span>
      <span className="ml-auto flex shrink-0 items-center gap-2 text-xs text-gray-400 dark:text-gray-500">
        <time title={absolute}>{when}</time>
        {explorerUrl && (
          <a
            href={explorerTxUrl(explorerUrl, event.txHash)}
            target="_blank"
            rel="noreferrer"
            title="View transaction"
            aria-label="View transaction"
            className="transition-colors hover:text-gray-700 dark:hover:text-gray-200"
          >
            <ExternalLink className="h-3.5 w-3.5" />
          </a>
        )}
      </span>
    </li>
  );
}
