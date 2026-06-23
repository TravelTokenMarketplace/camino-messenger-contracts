import { type ReactNode } from "react";

/**
 * Wraps a row action (e.g. Remove) so it stays out of the way until the row is
 * hovered or focused on pointer devices, while remaining always visible on
 * touch/small screens and reachable by keyboard. It also stays visible while a
 * contained transaction is pending (TxButton sets `data-pending`), so the
 * "Confirming…" state isn't hidden when the cursor leaves the row. Place inside
 * an element with the `group` class.
 */
export function RowAction({ children }: { children: ReactNode }) {
  return (
    <span className="shrink-0 opacity-100 transition-opacity md:opacity-0 md:group-hover:opacity-100 md:group-focus-within:opacity-100 md:has-[[data-pending]]:opacity-100">
      {children}
    </span>
  );
}
