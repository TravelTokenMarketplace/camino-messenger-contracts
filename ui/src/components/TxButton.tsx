import { type ReactNode, useState } from "react";
import { Loader2 } from "lucide-react";
import type { TransactionReceipt } from "viem";
import { useTx } from "../tx/TxProvider";
import { Tooltip } from "./Tooltip";

interface TxButtonProps {
  label: string;
  disabled?: boolean;
  write: () => Promise<`0x${string}`>;
  onConfirmed?: (receipt: TransactionReceipt) => void;
  icon?: ReactNode;
  variant?: "primary" | "danger";
  /** Rich tooltip; defaults to a note that this sends a wallet transaction. */
  tooltip?: ReactNode;
}

export function TxButton({ label, disabled, write, onConfirmed, icon, variant = "primary", tooltip }: TxButtonProps) {
  const { track } = useTx();
  const [pending, setPending] = useState(false);
  const [error, setError] = useState<string>();

  async function handleClick() {
    setPending(true);
    setError(undefined);
    try {
      // Resolves once submitted; mining and refetch are handled by the panel.
      await track({ label, write, onConfirmed });
    } catch (e) {
      const msg = e instanceof Error ? e.message : String(e);
      setError(msg.split("\n")[0]);
    } finally {
      setPending(false);
    }
  }

  const color =
    variant === "danger" ? "bg-signal text-white hover:bg-signal-fg" : "bg-camino-600 text-white hover:bg-camino-700";

  return (
    <div className="flex flex-col gap-1">
      <Tooltip content={tooltip ?? "Sends a transaction to your wallet to confirm."}>
        <button
          type="button"
          disabled={disabled || pending}
          data-pending={pending || undefined}
          onClick={handleClick}
          className={`inline-flex items-center justify-center gap-1.5 rounded-sm px-3 py-1.5 font-mono text-xs font-medium uppercase tracking-[0.08em] transition-colors disabled:opacity-50 ${
            pending ? "bg-departure-500 text-white" : color
          }`}
        >
          {pending ? <Loader2 className="h-4 w-4 animate-spin" /> : icon}
          <span>{pending ? "Confirming…" : label}</span>
        </button>
      </Tooltip>
      {error && <span className="max-w-xs font-mono text-xs text-signal-fg dark:text-signal-dark">{error}</span>}
    </div>
  );
}
