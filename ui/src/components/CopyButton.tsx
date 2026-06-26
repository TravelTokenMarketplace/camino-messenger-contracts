import { useState } from "react";
import { Check, Copy } from "lucide-react";
import { Tooltip } from "./Tooltip";

/** Small icon button that copies arbitrary text to the clipboard. */
export function CopyButton({
  value,
  label = "Copy",
  className = "",
}: {
  value: string;
  label?: string;
  className?: string;
}) {
  const [copied, setCopied] = useState(false);

  async function copy(e: React.MouseEvent) {
    e.stopPropagation();
    e.preventDefault();
    try {
      await navigator.clipboard?.writeText(value);
      setCopied(true);
      setTimeout(() => setCopied(false), 1200);
    } catch {
      // clipboard unavailable; ignore
    }
  }

  return (
    <Tooltip content={copied ? "Copied" : label}>
      <button
        type="button"
        onClick={copy}
        aria-label={label}
        className={`shrink-0 text-tarmac-400 transition-colors hover:text-tarmac-700 dark:hover:text-tarmac-200 ${className}`}
      >
        {copied ? <Check className="h-3.5 w-3.5 text-green-600" /> : <Copy className="h-3.5 w-3.5" />}
      </button>
    </Tooltip>
  );
}
