import { useState } from "react";
import { Lock } from "lucide-react";

/**
 * Compact, non-intrusive indicator shown in place of an action the connected
 * wallet lacks the role for. Reveals the required-role message on hover (pointer)
 * or click/focus (touch & keyboard).
 */
export function PermissionHint({ roleName }: { roleName: string }) {
  const [open, setOpen] = useState(false);

  return (
    <span className="relative inline-flex">
      <button
        type="button"
        onClick={() => setOpen((o) => !o)}
        onBlur={() => setOpen(false)}
        aria-label={`You lack the ${roleName} role required for this action`}
        className="peer inline-flex items-center gap-1 rounded border border-amber-300 bg-amber-50 px-2 py-1 text-xs text-amber-700 hover:bg-amber-100 dark:border-amber-800 dark:bg-amber-950 dark:text-amber-300"
      >
        <Lock className="h-3.5 w-3.5" /> No permission
      </button>
      <span
        role="tooltip"
        className={`pointer-events-none absolute left-0 top-full z-10 mt-1 w-max max-w-xs rounded bg-gray-900 px-2 py-1 text-xs text-white shadow-lg peer-hover:block dark:bg-gray-700 ${open ? "block" : "hidden"}`}
      >
        Requires the <code>{roleName}</code> role on this account.
      </span>
    </span>
  );
}
