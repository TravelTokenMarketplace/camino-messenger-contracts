import { type ReactNode, useState } from "react";
import { ExternalLink, Loader2 } from "lucide-react";
import { explorerTxUrl } from "../lib/format";

type Status = "idle" | "pending" | "success" | "error";

interface TxButtonProps {
  label: string;
  disabled?: boolean;
  write: () => Promise<`0x${string}`>;
  onConfirmed?: () => void;
  explorerBase?: string;
  icon?: ReactNode;
  variant?: "primary" | "danger";
}

export function TxButton({ label, disabled, write, onConfirmed, explorerBase, icon, variant = "primary" }: TxButtonProps) {
  const [status, setStatus] = useState<Status>("idle");
  const [hash, setHash] = useState<string>();
  const [error, setError] = useState<string>();

  async function handleClick() {
    setStatus("pending");
    setError(undefined);
    try {
      const h = await write();
      setHash(h);
      setStatus("success");
      onConfirmed?.();
    } catch (e) {
      setError(e instanceof Error ? e.message : String(e));
      setStatus("error");
    }
  }

  const color =
    variant === "danger"
      ? "bg-red-600 hover:bg-red-700"
      : "bg-indigo-600 hover:bg-indigo-700";

  return (
    <div className="flex flex-col gap-1">
      <button
        type="button"
        disabled={disabled || status === "pending"}
        onClick={handleClick}
        className={`inline-flex items-center justify-center gap-1.5 rounded px-3 py-1.5 text-white transition-colors disabled:opacity-50 ${color}`}
      >
        {status === "pending" ? <Loader2 className="h-4 w-4 animate-spin" /> : icon}
        <span>{status === "pending" ? "Confirming…" : status === "success" ? "Confirmed" : label}</span>
      </button>
      {hash && explorerBase && (
        <a className="inline-flex items-center gap-1 text-xs text-indigo-500 underline" href={explorerTxUrl(explorerBase, hash)} target="_blank" rel="noreferrer">
          <ExternalLink className="h-3 w-3" /> View transaction
        </a>
      )}
      {error && <span className="text-xs text-red-600">{error}</span>}
    </div>
  );
}
