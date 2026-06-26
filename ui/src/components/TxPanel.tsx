import { ExternalLink, X } from "lucide-react";
import { APP_CHAINS } from "../config/chains";
import { explorerTxUrl, shortAddress } from "../lib/format";
import { useTx } from "../tx/TxProvider";
import { Tooltip } from "./Tooltip";

type TxState = "pending" | "confirmed" | "failed";

// Board vocabulary: a transaction is a departure moving through states.
const STATE = {
  pending: { word: "IN TRANSIT", note: "Waiting for confirmation…" },
  confirmed: { word: "CONFIRMED", note: "Settled on-chain" },
  failed: { word: "REVERTED", note: "Transaction failed" },
} as const;

/**
 * Split-flap status chip — the signature element. A pending transaction flips on
 * an amber tile (like a departures board mid-change) and resolves to a teal
 * CONFIRMED, mirroring the truest live moment in the app: a tx clearing.
 */
function StatusFlap({ state }: { state: TxState }) {
  const tone =
    state === "confirmed"
      ? "border-camino-300 bg-camino-50 text-camino-700 dark:border-camino-800 dark:bg-camino-950 dark:text-camino-300"
      : state === "failed"
        ? "border-signal/40 bg-signal/10 text-signal-fg dark:text-signal-dark"
        : "border-departure-300 bg-departure-50 text-departure-700 dark:border-departure-800 dark:bg-departure-900/40 dark:text-departure-300";
  return (
    <span
      className={`inline-flex shrink-0 items-center rounded-[3px] border px-1.5 py-0.5 font-mono text-[0.625rem] font-semibold uppercase tracking-[0.12em] [transform-style:preserve-3d] ${tone} ${
        state === "pending" ? "animate-flap" : ""
      }`}
    >
      {STATE[state].word}
    </span>
  );
}

/**
 * Persistent panel listing in-flight and recently finished transactions. Lives
 * outside the action buttons so progress stays visible even when a row action is
 * hidden on mouse-out. Renders nothing when there are no transactions.
 */
export function TxPanel() {
  const { txs, dismiss } = useTx();
  if (txs.length === 0) return null;

  return (
    <aside className="board rounded-md">
      <div className="flex items-center gap-2 border-b border-tarmac-200/80 px-4 py-2.5 dark:border-tarmac-800">
        <span className="h-1.5 w-1.5 rounded-full bg-camino-500" aria-hidden />
        <h2 className="eyebrow">Transactions</h2>
      </div>
      <ul className="divide-y divide-tarmac-200/60 dark:divide-tarmac-800/80">
        {txs.map((t) => {
          const chain = APP_CHAINS.find((c) => c.id === t.chainId);
          return (
            <li key={t.id} className="flex items-start gap-3 px-4 py-3 text-sm">
              <StatusFlap state={t.state} />
              <span className="min-w-0 flex-1">
                <span className="block truncate font-mono text-xs text-tarmac-800 dark:text-tarmac-100" title={t.label}>
                  {t.label}
                </span>
                <span className="block text-xs text-tarmac-400">{STATE[t.state].note}</span>
                {chain && (
                  <a
                    className="mt-0.5 inline-flex items-center gap-1 font-mono text-xs text-camino-600 hover:underline dark:text-camino-400"
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
                  className="shrink-0 text-tarmac-400 hover:text-tarmac-700 dark:hover:text-tarmac-200"
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
