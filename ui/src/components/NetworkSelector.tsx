import { useAccount, useSwitchChain } from "wagmi";
import { ENABLED_CHAINS } from "../config/chains";
import { useActiveChain } from "../wallet/activeChain";

export function NetworkSelector() {
  const { activeChainId, setActiveChainId } = useActiveChain();
  const { isConnected, chainId: walletChainId } = useAccount();
  const { switchChain } = useSwitchChain();

  function onChange(id: number) {
    if (isConnected) switchChain({ chainId: id });
    else setActiveChainId(id);
  }

  const walletUnsupported =
    isConnected && walletChainId !== undefined && !ENABLED_CHAINS.some((c) => c.id === walletChainId);

  return (
    <div className="flex items-center gap-2">
      {walletUnsupported && (
        <span className="rounded bg-red-100 px-2 py-1 text-xs text-red-800 dark:bg-red-950 dark:text-red-300">
          Wallet on unsupported network
        </span>
      )}
      <select
        value={activeChainId}
        onChange={(e) => onChange(Number(e.target.value))}
        aria-label="Network"
        className="rounded border border-gray-300 bg-white px-2 py-1.5 text-sm dark:border-gray-700 dark:bg-gray-800 dark:text-gray-100"
      >
        {ENABLED_CHAINS.map((c) => (
          <option key={c.id} value={c.id}>
            {c.name}
          </option>
        ))}
      </select>
    </div>
  );
}
