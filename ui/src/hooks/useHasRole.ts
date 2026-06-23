import { type Abi, type Address } from "viem";
import { useAccount, useReadContract } from "wagmi";
import { roleHash, type RoleName } from "../lib/roles";
import { useActiveChain } from "../wallet/activeChain";

export function useHasRole(contractAddress: Address | undefined, abi: Abi, role: RoleName) {
  const { address } = useAccount();
  const { activeChainId } = useActiveChain();
  const { data, isLoading } = useReadContract({
    chainId: activeChainId,
    address: contractAddress,
    abi,
    functionName: "hasRole",
    args: address ? [roleHash(role), address] : undefined,
    query: { enabled: Boolean(contractAddress && address) },
  });
  return { hasRole: Boolean(data), isLoading };
}
