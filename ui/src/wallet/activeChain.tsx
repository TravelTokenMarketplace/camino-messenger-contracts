import { createContext, useContext, useEffect, useState, type ReactNode } from "react";
import { useAccount } from "wagmi";
import { ENABLED_CHAINS } from "../config/chains";

interface ActiveChainCtx {
  activeChainId: number;
  setActiveChainId: (id: number) => void;
}

const Ctx = createContext<ActiveChainCtx | undefined>(undefined);

// Prefer Base Sepolia (the cheap testnet) as the default when no wallet is
// connected; fall back to the first enabled chain if it isn't available. Stay
// strictly within ENABLED_CHAINS — never default to a chain with no contracts.
const PREFERRED_DEFAULT_CHAIN_ID = 84532; // Base Sepolia
const DEFAULT_CHAIN_ID = ENABLED_CHAINS.find((c) => c.id === PREFERRED_DEFAULT_CHAIN_ID)?.id ?? ENABLED_CHAINS[0]?.id;

if (DEFAULT_CHAIN_ID === undefined) {
  throw new Error("No enabled chains are configured (see config/chains.ts).");
}

/**
 * Tracks the chain the app reads from. When no wallet is connected the user can
 * freely select any enabled chain; when a wallet is connected the active chain
 * follows the wallet (and the selector switches the wallet, see NetworkSelector).
 */
export function ActiveChainProvider({ children }: { children: ReactNode }) {
  const { chainId: walletChainId } = useAccount();
  const [selected, setSelected] = useState<number>(DEFAULT_CHAIN_ID);

  useEffect(() => {
    if (walletChainId && ENABLED_CHAINS.some((c) => c.id === walletChainId)) {
      setSelected(walletChainId);
    }
  }, [walletChainId]);

  return <Ctx.Provider value={{ activeChainId: selected, setActiveChainId: setSelected }}>{children}</Ctx.Provider>;
}

export function useActiveChain(): ActiveChainCtx {
  const ctx = useContext(Ctx);
  if (!ctx) throw new Error("useActiveChain must be used within ActiveChainProvider");
  return ctx;
}
