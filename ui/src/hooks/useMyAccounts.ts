import { type Abi, type Address } from "viem";
import { useReadContract, useReadContracts } from "wagmi";
import { useActiveContracts } from "./useActiveContracts";
import { ACCOUNT_ROLES, roleHash } from "../lib/roles";
import { useActiveChain } from "../wallet/activeChain";

export function uniqueAddresses(addrs: string[]): string[] {
  const seen = new Set<string>();
  const out: string[] = [];
  for (const a of addrs) {
    const k = a.toLowerCase();
    if (!seen.has(k)) {
      seen.add(k);
      out.push(a);
    }
  }
  return out;
}

/**
 * Lists every CM Account by enumerating members of the manager's CMACCOUNT_ROLE.
 * This mirrors the `account find` CLI task and avoids eth_getLogs, which
 * free-tier RPCs reject for wide block ranges.
 */
export function useManagerAccounts() {
  const { manager, managerAbi } = useActiveContracts();
  const { activeChainId } = useActiveChain();
  const abi = managerAbi as Abi;

  const { data: cmRole } = useReadContract({
    chainId: activeChainId,
    address: manager,
    abi,
    functionName: "CMACCOUNT_ROLE",
    query: { enabled: Boolean(manager) },
  });

  const { data, isLoading } = useReadContract({
    chainId: activeChainId,
    address: manager,
    abi,
    functionName: "getRoleMembers",
    args: cmRole ? [cmRole] : undefined,
    query: { enabled: Boolean(manager && cmRole) },
  });

  const accounts = uniqueAddresses((data as string[]) ?? []) as Address[];
  return { accounts, isLoading };
}

/**
 * Footprint counts for a single CM Account, batched in one multicall: how many
 * services it supports, payment tokens it accepts, and public keys it has
 * registered. Gives each dashboard row substance at a glance. Counts only — the
 * array getters decode reliably for length even where tuple getters don't (see
 * CLAUDE.md), so we read the hash/address arrays and take `.length`.
 */
export function useAccountStats(account: Address) {
  const { cmAccountAbi } = useActiveContracts();
  const { activeChainId } = useActiveChain();
  const abi = cmAccountAbi as Abi;

  const { data, isLoading } = useReadContracts({
    contracts: [
      { chainId: activeChainId, address: account, abi, functionName: "getAllServiceHashes" },
      { chainId: activeChainId, address: account, abi, functionName: "getSupportedTokens" },
      { chainId: activeChainId, address: account, abi, functionName: "getPublicKeysAddresses" },
    ],
    allowFailure: true,
  });

  const len = (i: number) => {
    const r = data?.[i];
    return r?.status === "success" && Array.isArray(r.result) ? (r.result as unknown[]).length : undefined;
  };

  return { services: len(0), tokens: len(1), pubkeys: len(2), isLoading };
}

/**
 * For a single CM Account, returns which account-level roles the given address
 * holds. Uses a multicall batch of hasRole() reads (plain eth_call). Returns
 * `isLoading` so callers can distinguish "no roles" from "not resolved yet" —
 * filtering on `roles.length === 0` before the read settles would wrongly drop
 * every row.
 */
export function useAccountRolesFor(account: Address, address: Address | undefined) {
  const { cmAccountAbi } = useActiveContracts();
  const { activeChainId } = useActiveChain();
  const abi = cmAccountAbi as Abi;

  const { data, isLoading } = useReadContracts({
    contracts: ACCOUNT_ROLES.map((r) => ({
      chainId: activeChainId,
      address: account,
      abi,
      functionName: "hasRole",
      args: [roleHash(r), address as Address],
    })),
    allowFailure: true,
    query: { enabled: Boolean(address) },
  });

  const roles = ACCOUNT_ROLES.filter((_, i) => data?.[i]?.result === true);
  return { roles, isLoading };
}
