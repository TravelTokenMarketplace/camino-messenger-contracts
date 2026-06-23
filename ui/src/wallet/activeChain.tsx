import { createContext, useContext, useEffect, useState, type ReactNode } from "react";
import { useAccount } from "wagmi";
import { ENABLED_CHAINS } from "../config/chains";

interface ActiveChainCtx {
  activeChainId: number;
  setActiveChainId: (id: number) => void;
}

const Ctx = createContext<ActiveChainCtx | undefined>(undefined);

const DEFAULT_CHAIN_ID = ENABLED_CHAINS[0]?.id ?? 84532;

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
