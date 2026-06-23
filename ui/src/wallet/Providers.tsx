import { type ReactNode } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { WagmiProvider } from "wagmi";
import { wagmiConfig } from "./wagmi";
import { ActiveChainProvider } from "./activeChain";
import { TxProvider } from "../tx/TxProvider";

const queryClient = new QueryClient();

export function Providers({ children }: { children: ReactNode }) {
  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <ActiveChainProvider>
          <TxProvider>{children}</TxProvider>
        </ActiveChainProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
