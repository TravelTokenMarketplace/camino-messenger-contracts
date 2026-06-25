import { type Address } from "viem";
import { useReadContracts } from "wagmi";
import { ERC20_ABI } from "../lib/erc20";
import { useActiveContracts } from "./useActiveContracts";

export interface TokenMeta {
  address: Address;
  symbol?: string;
  name?: string;
  decimals?: number;
}

/**
 * Resolve ERC20 symbol/name/decimals for a list of token addresses via a single
 * multicall. Returns a map keyed by lowercase address; any field whose on-chain
 * read fails is left undefined (consumers fall back to the address).
 */
export function useTokenMetadata(addresses: Address[]): { meta: Map<string, TokenMeta>; isLoading: boolean } {
  const { chainId } = useActiveContracts();

  const seen = new Set<string>();
  const list: Address[] = [];
  for (const a of addresses) {
    const key = a.toLowerCase();
    if (seen.has(key)) continue;
    seen.add(key);
    list.push(a);
  }

  const { data, isLoading } = useReadContracts({
    allowFailure: true,
    contracts: list.flatMap((address) => [
      { chainId, address, abi: ERC20_ABI, functionName: "symbol" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "name" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "decimals" } as const,
    ]),
    query: { enabled: list.length > 0 },
  });

  const meta = new Map<string, TokenMeta>();
  if (data) {
    list.forEach((address, i) => {
      const s = data[i * 3];
      const n = data[i * 3 + 1];
      const d = data[i * 3 + 2];
      meta.set(address.toLowerCase(), {
        address,
        symbol: s?.status === "success" ? (s.result as string) : undefined,
        name: n?.status === "success" ? (n.result as string) : undefined,
        decimals: d?.status === "success" ? Number(d.result) : undefined,
      });
    });
  }

  return { meta, isLoading };
}
