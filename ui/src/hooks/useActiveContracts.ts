import { type Address } from "viem";
import { BOOKINGTOKEN_ABI, CMACCOUNT_ABI, MANAGER_ABI, getContractsForChain } from "../contracts";
import { useActiveChain } from "../wallet/activeChain";

export function useActiveContracts() {
  const { activeChainId } = useActiveChain();
  const resolved = getContractsForChain(activeChainId);
  return {
    chainId: activeChainId,
    supported: Boolean(resolved),
    manager: resolved?.manager as Address | undefined,
    bookingToken: resolved?.bookingToken as Address | undefined,
    cmAccountImpl: resolved?.cmAccountImpl as Address | undefined,
    managerAbi: MANAGER_ABI,
    cmAccountAbi: CMACCOUNT_ABI,
    bookingTokenAbi: BOOKINGTOKEN_ABI,
  };
}
