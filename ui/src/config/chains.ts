import { hasContracts } from "../contracts";

export interface AppChain {
  id: number;
  name: string;
  enabled: boolean;
  rpcUrl: string;
  explorerUrl: string;
  nativeCurrency: { name: string; symbol: string; decimals: number };
}

export const APP_CHAINS: AppChain[] = [
  {
    id: 500,
    name: "Camino",
    enabled: true,
    rpcUrl: "https://api.camino.network/ext/bc/C/rpc",
    explorerUrl: "https://caminoscan.com",
    nativeCurrency: { name: "Camino", symbol: "CAM", decimals: 18 },
  },
  {
    id: 501,
    name: "Columbus (deprecated)",
    enabled: false,
    rpcUrl: "https://columbus.camino.network/ext/bc/C/rpc",
    explorerUrl: "https://columbus.caminoscan.com",
    nativeCurrency: { name: "Camino", symbol: "CAM", decimals: 18 },
  },
  {
    id: 8453,
    name: "Base",
    enabled: true,
    rpcUrl: "https://mainnet.base.org",
    explorerUrl: "https://basescan.org",
    nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  },
  {
    id: 84532,
    name: "Base Sepolia",
    enabled: true,
    rpcUrl: "https://base-sepolia.drpc.org",
    explorerUrl: "https://sepolia.basescan.org",
    nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  },
];

// A chain is usable only if marked enabled AND its contracts are deployed.
export const ENABLED_CHAINS: AppChain[] = APP_CHAINS.filter(
  (c) => c.enabled && hasContracts(c.id),
);
