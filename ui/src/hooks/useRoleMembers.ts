import { type Abi, type Address } from "viem";
import { useReadContract } from "wagmi";
import { roleHash, type RoleName } from "../lib/roles";

export function toMemberList(data: unknown): string[] {
  return ((data as unknown[]) ?? []).map((x) => String(x));
}

export function useRoleMembers(account: Address | undefined, abi: Abi, role: RoleName) {
  const { data, isLoading, refetch } = useReadContract({
    address: account,
    abi,
    functionName: "getRoleMembers",
    args: [roleHash(role)],
    query: { enabled: Boolean(account) },
  });
  return { members: toMemberList(data), isLoading, refetch: () => void refetch() };
}
