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

// ENABLED_CHAINS is empty when no enabled chain has synced contract addresses
// (e.g. a missing/stale generated/ directory). Fail loudly here instead of
// letting the non-empty-tuple cast below hide it as an opaque wagmi error.
if (viemChains.length === 0) {
  throw new Error(
    "No usable chains: ENABLED_CHAINS is empty. Run `yarn sync` to generate contract addresses, or enable a chain with a deployment in config/chains.ts.",
  );
}

export const wagmiConfig = createConfig({
  chains: viemChains as [Chain, ...Chain[]],
  connectors: [injected()],
  transports: Object.fromEntries(
    ENABLED_CHAINS.map((c) => [c.id, http(c.rpcUrl)]),
  ),
});
