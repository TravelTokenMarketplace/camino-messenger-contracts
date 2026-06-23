import { createContext, useCallback, useContext, useState, type ReactNode } from "react";
import type { Hex, TransactionReceipt } from "viem";
import { useQueryClient } from "@tanstack/react-query";
import { useConfig } from "wagmi";
import { getAccount, waitForTransactionReceipt } from "wagmi/actions";

export type TxState = "pending" | "confirmed" | "failed";

export interface TrackedTx {
  id: string;
  label: string;
  hash: Hex;
  chainId?: number;
  state: TxState;
}

interface TrackArgs {
  /** Human-friendly description shown in the transaction panel. */
  label: string;
  /** Submits the transaction and resolves with its hash. */
  write: () => Promise<Hex>;
  /**
   * Called once the transaction is mined successfully (good place to refetch or
   * navigate). Receives the mined receipt. Errors thrown here are swallowed and
   * do NOT mark the (already-confirmed) transaction as failed.
   */
  onConfirmed?: (receipt: TransactionReceipt) => void;
}

interface TxApi {
  txs: TrackedTx[];
  /**
   * Submits a transaction and tracks it to completion. Resolves once the tx is
   * submitted (hash obtained); mining is awaited in the background so the panel
   * keeps showing progress even after the triggering button disappears.
   * Submission errors (e.g. wallet rejection) reject so the caller can show them
   * inline; they do not create a panel entry.
   */
  track: (args: TrackArgs) => Promise<void>;
  dismiss: (id: string) => void;
}

// Default used when no provider is mounted (e.g. unit tests): fire-and-forget
// that still calls write/onConfirmed and surfaces submission errors.
const fallback: TxApi = {
  txs: [],
  track: async ({ write, onConfirmed }) => {
    await write();
    // No real receipt without a provider; tests only assert onConfirmed runs.
    onConfirmed?.({} as TransactionReceipt);
  },
  dismiss: () => {},
};

const TxContext = createContext<TxApi>(fallback);

export function useTx() {
  return useContext(TxContext);
}

let counter = 0;
const nextId = () => `tx-${Date.now()}-${counter++}`;

export function TxProvider({ children }: { children: ReactNode }) {
  const config = useConfig();
  const queryClient = useQueryClient();
  const [txs, setTxs] = useState<TrackedTx[]>([]);

  const update = useCallback((id: string, patch: Partial<TrackedTx>) => {
    setTxs((prev) => prev.map((t) => (t.id === id ? { ...t, ...patch } : t)));
  }, []);

  const dismiss = useCallback((id: string) => {
    setTxs((prev) => prev.filter((t) => t.id !== id));
  }, []);

  const track = useCallback(
    async ({ label, write, onConfirmed }: TrackArgs) => {
      // Submission errors propagate to the caller and create no panel entry.
      const hash = await write();
      const id = nextId();
      const chainId = getAccount(config).chainId;
      setTxs((prev) => [{ id, label, hash, chainId, state: "pending" }, ...prev]);
      // Wait for mining off the critical path so the panel reflects on-chain
      // confirmation even after the triggering row action is hidden.
      void (async () => {
        let receipt: TransactionReceipt;
        try {
          receipt = await waitForTransactionReceipt(config, { hash, chainId });
        } catch {
          update(id, { state: "failed" });
          return;
        }
        if (receipt.status !== "success") {
          update(id, { state: "failed" });
          return;
        }
        // The tx is mined and successful — mark it confirmed before running any
        // side effects so a throwing callback can't flip it back to "failed".
        update(id, { state: "confirmed" });
        try {
          onConfirmed?.(receipt);
        } catch {
          // Side-effect failures (e.g. navigation) must not affect tx state.
        }
        // The mined tx may affect reads anywhere in the app (lists, role badges,
        // balances, count pills), not just the one wired to onConfirmed —
        // invalidate every query so the whole view refreshes.
        void queryClient.invalidateQueries();
      })();
    },
    [config, queryClient, update],
  );

  return <TxContext.Provider value={{ txs, track, dismiss }}>{children}</TxContext.Provider>;
}
