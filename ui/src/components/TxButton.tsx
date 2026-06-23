import { useState } from "react";
import { explorerTxUrl } from "../lib/format";

type Status = "idle" | "pending" | "success" | "error";

interface TxButtonProps {
  label: string;
  disabled?: boolean;
  write: () => Promise<`0x${string}`>;
  onConfirmed?: () => void;
  explorerBase?: string;
}

export function TxButton({ label, disabled, write, onConfirmed, explorerBase }: TxButtonProps) {
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

  return (
    <div className="flex flex-col gap-1">
      <button
        type="button"
        disabled={disabled || status === "pending"}
        onClick={handleClick}
        className="rounded bg-indigo-600 px-3 py-1.5 text-white disabled:opacity-50"
      >
        {status === "pending" ? "Confirming…" : status === "success" ? "Confirmed" : label}
      </button>
      {hash && explorerBase && (
        <a className="text-xs text-indigo-500 underline" href={explorerTxUrl(explorerBase, hash)} target="_blank" rel="noreferrer">
          View transaction
        </a>
      )}
      {error && <span className="text-xs text-red-600">{error}</span>}
    </div>
  );
}
