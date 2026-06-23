import { type ReactNode } from "react";

/**
 * Lightweight rich tooltip that appears instantly on hover (and keyboard focus),
 * unlike the native `title` attribute which is delayed and unstyled. Wrap any
 * focusable/hoverable element; the tooltip is positioned relative to it.
 */
export function Tooltip({
  content,
  children,
  side = "top",
}: {
  content: ReactNode;
  children: ReactNode;
  side?: "top" | "bottom";
}) {
  if (!content) return <>{children}</>;
  const pos = side === "top" ? "bottom-full mb-1.5" : "top-full mt-1.5";
  return (
    <span className="group/tt relative inline-flex">
      {children}
      <span
        role="tooltip"
        className={`pointer-events-none absolute left-1/2 z-30 hidden w-max max-w-xs -translate-x-1/2 ${pos} rounded-md bg-gray-900 px-2.5 py-1.5 text-left text-xs font-normal leading-snug text-white shadow-lg group-hover/tt:block group-focus-within/tt:block dark:bg-gray-700`}
      >
        {content}
      </span>
    </span>
  );
}
