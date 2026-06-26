import { type ReactNode } from "react";

/**
 * A board panel — the flat, hairline-ruled surface that frames every section.
 * The title is set as a small uppercase "eyebrow" so panels read like the
 * labelled cells of a departures board rather than generic cards.
 */
export function Card({ title, actions, children }: { title?: string; actions?: ReactNode; children: ReactNode }) {
  return (
    <section className="board rounded-md">
      {(title || actions) && (
        <div className="flex items-center justify-between gap-3 border-b border-tarmac-200/80 px-4 py-2.5 dark:border-tarmac-800">
          {title && <h2 className="eyebrow">{title}</h2>}
          {actions}
        </div>
      )}
      <div className="p-4">{children}</div>
    </section>
  );
}
