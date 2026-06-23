import { useState } from "react";
import { RefreshCw } from "lucide-react";
import { useQueryClient } from "@tanstack/react-query";
import { Tooltip } from "./Tooltip";

/**
 * Manually re-reads on-chain data by invalidating all queries. A safety net for
 * when a read RPC briefly lags behind a just-mined transaction.
 */
export function RefreshButton({ label = "Refresh" }: { label?: string }) {
  const queryClient = useQueryClient();
  const [spinning, setSpinning] = useState(false);

  async function refresh() {
    setSpinning(true);
    try {
      await queryClient.invalidateQueries();
    } finally {
      setTimeout(() => setSpinning(false), 600);
    }
  }

  return (
    <Tooltip content="Re-read the latest on-chain data">
      <button
        type="button"
        onClick={refresh}
        aria-label={label}
        className="inline-flex items-center gap-1.5 rounded-md border border-gray-300 px-2 py-1 text-xs text-gray-600 transition-colors hover:bg-gray-50 dark:border-gray-700 dark:text-gray-300 dark:hover:bg-gray-800"
      >
        <RefreshCw className={`h-3.5 w-3.5 ${spinning ? "animate-spin" : ""}`} /> {label}
      </button>
    </Tooltip>
  );
}
