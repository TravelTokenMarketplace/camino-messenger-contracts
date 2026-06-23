import { http, type Chain } from "viem";
import { createConfig } from "wagmi";
import { injected } from "wagmi/connectors";
import { ENABLED_CHAINS, type AppChain } from "../config/chains";

function toViemChain(c: AppChain): Chain {
  return {
    id: c.id,
    name: c.name,
    nativeCurrency: c.nativeCurrency,
    rpcUrls: { default: { http: [c.rpcUrl] } },
    blockExplorers: { default: { name: c.name, url: c.explorerUrl } },
  };
}

const viemChains = ENABLED_CHAINS.map(toViemChain);

export const wagmiConfig = createConfig({
  chains: viemChains as [Chain, ...Chain[]],
  connectors: [injected()],
  transports: Object.fromEntries(
    ENABLED_CHAINS.map((c) => [c.id, http(c.rpcUrl)]),
  ),
});
