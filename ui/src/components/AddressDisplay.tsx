import { useState } from "react";
import { Check, Copy } from "lucide-react";
import { shortAddress } from "../lib/format";
import { Identicon } from "./Identicon";

interface AddressDisplayProps {
  address: string;
  truncate?: boolean;
  className?: string;
  showIcon?: boolean;
}

export function AddressDisplay({ address, truncate = false, className = "", showIcon = true }: AddressDisplayProps) {
  const [copied, setCopied] = useState(false);

  async function copy(e: React.MouseEvent) {
    // Don't let a copy click trigger a parent row's navigation.
    e.stopPropagation();
    e.preventDefault();
    try {
      await navigator.clipboard?.writeText(address);
      setCopied(true);
      setTimeout(() => setCopied(false), 1200);
    } catch {
      // clipboard unavailable; ignore
    }
  }

  return (
    <span className={`inline-flex items-center gap-1.5 font-mono ${className}`}>
      {showIcon && <Identicon address={address} />}
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
