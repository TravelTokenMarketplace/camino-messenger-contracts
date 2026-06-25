import { type Address, formatUnits } from "viem";
import { useReadContract, useReadContracts } from "wagmi";
import { ERC20_ABI } from "../lib/erc20";
import { shortAddress } from "../lib/format";
import { EXTRA_TOKENS } from "../config/tokens";
import { useActiveContracts } from "./useActiveContracts";

export interface TokenBalance {
  address: Address;
  symbol: string;
  decimals: number;
  balance: bigint;
  formatted: string;
  isZero: boolean;
}

export function useErc20Balances(account: Address): { tokens: TokenBalance[]; isLoading: boolean } {
  const { chainId, cmAccountAbi } = useActiveContracts();

  // On-chain payment tokens the account supports.
  const { data: supportedRaw, isLoading: supportedLoading } = useReadContract({
    chainId,
    address: account,
    abi: cmAccountAbi,
    functionName: "getSupportedTokens",
  });
  const supported = (supportedRaw as Address[] | undefined) ?? [];
  const configured = EXTRA_TOKENS[chainId] ?? [];

  // Merge + dedupe case-insensitively, keeping the first-seen casing.
  const seen = new Set<string>();
  const addresses: Address[] = [];
  for (const a of [...configured, ...supported]) {
    const key = a.toLowerCase();
    if (seen.has(key)) continue;
    seen.add(key);
    addresses.push(a);
  }

  const { data: multi, isLoading: multiLoading } = useReadContracts({
    allowFailure: true,
    contracts: addresses.flatMap((address) => [
      { chainId, address, abi: ERC20_ABI, functionName: "symbol" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "decimals" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "balanceOf", args: [account] } as const,
    ]),
    query: { enabled: addresses.length > 0 },
  });

  const tokens: TokenBalance[] = [];
  if (multi) {
    addresses.forEach((address, i) => {
      const symbolRes = multi[i * 3];
      const decimalsRes = multi[i * 3 + 1];
      const balanceRes = multi[i * 3 + 2];

      // Require balanceOf AND decimals: without a trustworthy decimals we'd
      // misformat the balance (e.g. a 6-decimal token shown as 18), so drop
      // rather than guess. symbol is cosmetic and may fall back.
      if (balanceRes?.status !== "success" || decimalsRes?.status !== "success") {
        if (import.meta.env.DEV) {
          // eslint-disable-next-line no-console
          console.warn(`[useErc20Balances] ${address} is not a usable ERC20 on chain ${chainId} — dropping.`);
        }
        return;
      }

      const symbol = symbolRes?.status === "success" ? (symbolRes.result as string) : shortAddress(address);
      const decimals = Number(decimalsRes.result);
      const balance = balanceRes.result as bigint;
      tokens.push({
        address,
        symbol,
        decimals,
        balance,
        formatted: formatUnits(balance, decimals),
        isZero: balance === 0n,
      });
    });
  }

  return { tokens, isLoading: supportedLoading || multiLoading };
}
