import { type Abi, type Address } from "viem";
import { useReadContract } from "wagmi";

export function useContractList(account: Address, abi: Abi, functionName: string) {
  const { data, isLoading, refetch } = useReadContract({ address: account, abi, functionName });
  const items = ((data as unknown[]) ?? []).map((x) => String(x));
  return { items, isLoading, refetch: () => void refetch() };
}
