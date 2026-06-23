import { Lock } from "lucide-react";
import { Tooltip } from "./Tooltip";

/**
 * Compact, non-intrusive indicator shown in place of an action the connected
 * wallet lacks the role for. Names the action so it is meaningful on its own,
 * and reveals the required role in a tooltip on hover/focus.
 */
export function PermissionHint({ roleName, action }: { roleName: string; action?: string }) {
  const verb = action ? action.toLowerCase() : undefined;
  const label = verb ? `Can't ${verb}` : "No permission";

  return (
    <Tooltip
      content={
        <>
          Requires the <code className="font-mono">{roleName}</code> role on this account{verb ? ` to ${verb}` : ""}.
        </>
      }
    >
      <button
        type="button"
        aria-label={verb ? `You lack the ${roleName} role required to ${verb}` : `You lack the ${roleName} role`}
        className="inline-flex items-center gap-1 rounded border border-amber-300 bg-amber-50 px-2 py-1 text-xs text-amber-700 dark:border-amber-800 dark:bg-amber-950 dark:text-amber-300"
      >
        <Lock className="h-3.5 w-3.5" /> {label}
      </button>
    </Tooltip>
  );
}
