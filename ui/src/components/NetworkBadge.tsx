import { useActiveContracts } from "../hooks/useActiveContracts";
import { APP_CHAINS } from "../config/chains";

export function NetworkBadge() {
  const { chainId, supported } = useActiveContracts();
  const name = APP_CHAINS.find((c) => c.id === chainId)?.name;
  return (
    <span
      className={`rounded px-2 py-1 text-xs ${supported ? "bg-green-100 text-green-800" : "bg-red-100 text-red-800"}`}
    >
      {supported ? name : "Unsupported network"}
    </span>
  );
}
