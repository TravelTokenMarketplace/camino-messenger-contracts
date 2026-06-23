import { CheckCircle2, ExternalLink, Loader2, X, XCircle } from "lucide-react";
import { APP_CHAINS } from "../config/chains";
import { explorerTxUrl, shortAddress } from "../lib/format";
import { useTx } from "../tx/TxProvider";
import { Tooltip } from "./Tooltip";

const STATE_TEXT = {
  pending: "Waiting for confirmation…",
  confirmed: "Confirmed",
  failed: "Failed",
} as const;

/**
 * Persistent panel listing in-flight and recently finished transactions. Lives
 * outside the action buttons so progress stays visible even when a row action is
 * hidden on mouse-out. Renders nothing when there are no transactions.
 */
export function TxPanel() {
  const { txs, dismiss } = useTx();
  if (txs.length === 0) return null;

  return (
    <aside className="rounded-lg border border-gray-200 bg-white p-4 shadow-sm dark:border-gray-800 dark:bg-gray-800">
      <h2 className="mb-3 text-xs font-semibold uppercase tracking-wide text-gray-400">Transactions</h2>
      <ul className="space-y-3">
        {txs.map((t) => {
          const chain = APP_CHAINS.find((c) => c.id === t.chainId);
          return (
            <li key={t.id} className="flex items-start gap-2 text-sm">
              <span className="mt-0.5 shrink-0">
                {t.state === "pending" && <Loader2 className="h-4 w-4 animate-spin text-indigo-500" />}
                {t.state === "confirmed" && <CheckCircle2 className="h-4 w-4 text-emerald-500" />}
                {t.state === "failed" && <XCircle className="h-4 w-4 text-red-500" />}
              </span>
              <span className="min-w-0 flex-1">
                <span className="block truncate" title={t.label}>{t.label}</span>
                <span className="block text-xs text-gray-400">{STATE_TEXT[t.state]}</span>
                {chain && (
                  <a
                    className="mt-0.5 inline-flex items-center gap-1 text-xs text-indigo-500 hover:underline"
                    href={explorerTxUrl(chain.explorerUrl, t.hash)}
                    target="_blank"
                    rel="noreferrer"
                  >
                    <ExternalLink className="h-3 w-3" /> {shortAddress(t.hash)}
                  </a>
                )}
              </span>
              <Tooltip content="Dismiss from this list (does not affect the transaction)" side="bottom">
                <button
                  type="button"
                  onClick={() => dismiss(t.id)}
                  className="shrink-0 text-gray-400 hover:text-gray-600 dark:hover:text-gray-200"
                  aria-label="Dismiss transaction"
                >
                  <X className="h-3.5 w-3.5" />
                </button>
              </Tooltip>
            </li>
          );
        })}
      </ul>
    </aside>
  );
}
