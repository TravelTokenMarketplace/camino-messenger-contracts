import { useActiveContracts } from "../hooks/useActiveContracts";
import { APP_CHAINS } from "../config/chains";

export function NetworkBadge() {
  const { chainId, supported } = useActiveContracts();
  const name = APP_CHAINS.find((c) => c.id === chainId)?.name;
  return (
    <span
      className={`inline-flex items-center gap-1.5 rounded-[3px] border px-2 py-1 font-mono text-[0.625rem] uppercase tracking-[0.1em] ${
        supported
          ? "border-camino-300 bg-camino-50 text-camino-700 dark:border-camino-800 dark:bg-camino-950 dark:text-camino-300"
          : "border-signal/40 bg-signal/10 text-signal-fg dark:text-signal-dark"
      }`}
    >
      <span
        className={`h-1.5 w-1.5 rounded-full ${supported ? "bg-camino-500" : "bg-signal animate-lamp"}`}
        aria-hidden
      />
      {supported ? name : "Unsupported network"}
    </span>
  );
}
