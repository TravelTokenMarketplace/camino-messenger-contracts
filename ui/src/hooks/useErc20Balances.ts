import { type Address, formatUnits } from "viem";
import { useReadContract, useReadContracts } from "wagmi";
import { ERC20_ABI } from "../lib/erc20";
import { shortAddress } from "../lib/format";
import { EXTRA_TOKENS } from "../config/tokens";
import { useActiveContracts } from "./useActiveContracts";
import { useTokenMetadata } from "./useTokenMetadata";

export interface TokenBalance {
  address: Address;
  symbol: string;
  name?: string;
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

  const { meta, isLoading: metaLoading } = useTokenMetadata(addresses);

  const { data: balances, isLoading: balLoading } = useReadContracts({
    allowFailure: true,
    contracts: addresses.map(
      (address) => ({ chainId, address, abi: ERC20_ABI, functionName: "balanceOf", args: [account] }) as const,
    ),
    query: { enabled: addresses.length > 0 },
  });

  const tokens: TokenBalance[] = [];
  if (balances) {
    addresses.forEach((address, i) => {
      const m = meta.get(address.toLowerCase());
      const balanceRes = balances[i];

      // Require balanceOf AND a trustworthy decimals: without decimals we'd
      // misformat the balance (e.g. a 6-decimal token shown as 18), so drop
      // rather than guess. symbol/name are cosmetic and may fall back.
      if (balanceRes?.status !== "success" || m?.decimals === undefined) {
        if (import.meta.env.DEV) {
          // eslint-disable-next-line no-console
          console.warn(`[useErc20Balances] ${address} is not a usable ERC20 on chain ${chainId} — dropping.`);
        }
        return;
      }

      const balance = balanceRes.result as bigint;
      tokens.push({
        address,
        symbol: m.symbol ?? shortAddress(address),
        name: m.name,
        decimals: m.decimals,
        balance,
        formatted: formatUnits(balance, m.decimals),
        isZero: balance === 0n,
      });
    });
  }

  return { tokens, isLoading: supportedLoading || metaLoading || balLoading };
}
