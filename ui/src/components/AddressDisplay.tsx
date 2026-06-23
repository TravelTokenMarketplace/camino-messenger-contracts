import { useState } from "react";
import { Check, Copy } from "lucide-react";
import { shortAddress } from "../lib/format";

interface AddressDisplayProps {
  address: string;
  truncate?: boolean;
  className?: string;
}

export function AddressDisplay({ address, truncate = false, className = "" }: AddressDisplayProps) {
  const [copied, setCopied] = useState(false);

  async function copy() {
    try {
      await navigator.clipboard?.writeText(address);
      setCopied(true);
      setTimeout(() => setCopied(false), 1200);
    } catch {
      // clipboard unavailable; ignore
    }
  }

  return (
    <span className={`inline-flex items-center gap-1 font-mono ${className}`}>
      <span className="break-all">{truncate ? shortAddress(address) : address}</span>
      <button
        type="button"
        onClick={copy}
        title="Copy address"
        aria-label="Copy address"
        className="shrink-0 text-gray-400 transition-colors hover:text-gray-700 dark:hover:text-gray-200"
      >
        {copied ? <Check className="h-3.5 w-3.5 text-green-600" /> : <Copy className="h-3.5 w-3.5" />}
      </button>
    </span>
  );
}
