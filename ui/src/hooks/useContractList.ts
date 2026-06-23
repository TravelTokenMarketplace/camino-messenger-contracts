import { type Abi, type Address } from "viem";
import { useReadContract } from "wagmi";
import { useActiveChain } from "../wallet/activeChain";

export function useContractList(account: Address, abi: Abi, functionName: string) {
  const { activeChainId } = useActiveChain();
  const { data, isLoading, refetch } = useReadContract({ chainId: activeChainId, address: account, abi, functionName });
  const items = ((data as unknown[]) ?? []).map((x) => String(x));
  return { items, isLoading, refetch: () => void refetch() };
}
