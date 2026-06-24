import { type Abi, type Address } from "viem";
import { Card } from "../../components/Card";
import { RolesPanel } from "../../components/RolesPanel";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { ACCOUNT_ROLES } from "../../lib/roles";

export function RolesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  return (
    <Card title="Roles">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Expand a role to see its members and grant or revoke it. Only one role is open at a time.
      </p>
      <RolesPanel address={account} abi={cmAccountAbi as Abi} roles={ACCOUNT_ROLES} enumerable />
    </Card>
  );
}
