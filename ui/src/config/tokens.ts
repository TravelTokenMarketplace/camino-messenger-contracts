import { type Address } from "viem";

// Curated ERC20 token addresses to always display balances for, per chainId.
// Only addresses live here — symbol/decimals/balance are read on-chain.
// Merged with each account's on-chain getSupportedTokens() at render time.
export const EXTRA_TOKENS: Record<number, Address[]> = {
  500: [], // Camino
  8453: [], // Base
  84532: [], // Base Sepolia
};
