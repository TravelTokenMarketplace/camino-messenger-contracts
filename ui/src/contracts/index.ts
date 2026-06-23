import { ADDRESSES, type ResolvedAddresses } from "./generated/addresses";
export { MANAGER_ABI, CMACCOUNT_ABI, BOOKINGTOKEN_ABI } from "./generated/abis";
export type { ResolvedAddresses };

export function getContractsForChain(chainId: number): ResolvedAddresses | undefined {
  return ADDRESSES[chainId];
}

export function hasContracts(chainId: number): boolean {
  return ADDRESSES[chainId] !== undefined;
}
