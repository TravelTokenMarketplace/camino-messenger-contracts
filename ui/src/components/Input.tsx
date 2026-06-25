import { forwardRef, type InputHTMLAttributes } from "react";

/** Shared input styling — consistent borders, focus ring, and dark mode. */
export const inputClass =
  "rounded-md border border-gray-300 bg-white px-3 py-1.5 text-sm text-gray-900 placeholder-gray-400 transition-colors focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500 disabled:opacity-50 dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100 dark:placeholder-gray-500";

export const Input = forwardRef<HTMLInputElement, InputHTMLAttributes<HTMLInputElement>>(function Input(
  { className = "", ...props },
  ref,
) {
  return <input ref={ref} className={`${inputClass} ${className}`} {...props} />;
});
