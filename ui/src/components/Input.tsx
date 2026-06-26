import { forwardRef, type InputHTMLAttributes } from "react";

/** Shared input styling — consistent borders, focus ring, and dark mode. */
export const inputClass =
  "rounded-md border border-tarmac-300 bg-paper-raised px-3 py-1.5 text-sm text-tarmac-900 placeholder-tarmac-400 transition-colors focus:border-camino-500 focus:outline-none focus:ring-1 focus:ring-camino-500 disabled:opacity-50 dark:border-tarmac-600 dark:bg-tarmac-900 dark:text-tarmac-100 dark:placeholder-tarmac-500";

export const Input = forwardRef<HTMLInputElement, InputHTMLAttributes<HTMLInputElement>>(function Input(
  { className = "", ...props },
  ref,
) {
  return <input ref={ref} className={`${inputClass} ${className}`} {...props} />;
});
