import { type Abi, type Address } from "viem";
import { useWriteContract } from "wagmi";
import { ListManager } from "../../components/ListManager";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { ACCOUNT_ROLES, roleHash, type RoleName } from "../../lib/roles";

function RoleSection({ account, role, hasAdmin }: { account: Address; role: RoleName; hasAdmin: boolean }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, abi, role);
  return (
    <ListManager
      title={role}
      items={members}
      isLoading={isLoading}
      roleName="DEFAULT_ADMIN_ROLE"
      hasRole={hasAdmin}
      addLabel="Grant"
      addPlaceholder="Account address 0x…"
      onAdd={(v) => writeContractAsync({ address: account, abi, functionName: "grantRole", args: [roleHash(role), v as Address] })}
      onRemove={(v) => writeContractAsync({ address: account, abi, functionName: "revokeRole", args: [roleHash(role), v as Address] })}
      onChanged={refetch}
    />
  );
}

export function RolesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { hasRole } = useHasRole(account, cmAccountAbi as Abi, "DEFAULT_ADMIN_ROLE");
  return (
    <div className="grid gap-4">
      {ACCOUNT_ROLES.map((r) => (
        <RoleSection key={r} account={account} role={r} hasAdmin={hasRole} />
      ))}
    </div>
  );
}
