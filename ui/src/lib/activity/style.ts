import { type ActivityCategory } from "./types";

/**
 * Per-category colors so different event types are easy to spot at a glance.
 * `icon` styles the round icon badge on each row; `dot` is a small swatch used
 * on the Activity-page filter chips. Class strings are full literals so
 * Tailwind's content scanner keeps them.
 */
export const CATEGORY_STYLE: Record<ActivityCategory, { icon: string; dot: string }> = {
  Bookings: { icon: "bg-indigo-500/10 text-indigo-600 dark:text-indigo-400", dot: "bg-indigo-500" },
  Cancellations: { icon: "bg-rose-500/10 text-rose-600 dark:text-rose-400", dot: "bg-rose-500" },
  Accounts: { icon: "bg-emerald-500/10 text-emerald-600 dark:text-emerald-400", dot: "bg-emerald-500" },
  Services: { icon: "bg-sky-500/10 text-sky-600 dark:text-sky-400", dot: "bg-sky-500" },
  Bots: { icon: "bg-violet-500/10 text-violet-600 dark:text-violet-400", dot: "bg-violet-500" },
  Tokens: { icon: "bg-amber-500/10 text-amber-600 dark:text-amber-400", dot: "bg-amber-500" },
  Pubkeys: { icon: "bg-cyan-500/10 text-cyan-600 dark:text-cyan-400", dot: "bg-cyan-500" },
  Funds: { icon: "bg-green-500/10 text-green-600 dark:text-green-400", dot: "bg-green-500" },
  Config: { icon: "bg-gray-500/10 text-gray-500 dark:text-gray-400", dot: "bg-gray-400" },
};
