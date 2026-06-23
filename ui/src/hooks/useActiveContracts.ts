import { type Address } from "viem";
import { useChainId } from "wagmi";
import {
  BOOKINGTOKEN_ABI,
  CMACCOUNT_ABI,
  MANAGER_ABI,
  getContractsForChain,
} from "../contracts";

export function useActiveContracts() {
  const chainId = useChainId();
  const resolved = chainId ? getContractsForChain(chainId) : undefined;
  return {
    chainId,
    supported: Boolean(resolved),
    manager: resolved?.manager as Address | undefined,
    bookingToken: resolved?.bookingToken as Address | undefined,
    cmAccountImpl: resolved?.cmAccountImpl as Address | undefined,
    managerAbi: MANAGER_ABI,
    cmAccountAbi: CMACCOUNT_ABI,
    bookingTokenAbi: BOOKINGTOKEN_ABI,
  };
}
