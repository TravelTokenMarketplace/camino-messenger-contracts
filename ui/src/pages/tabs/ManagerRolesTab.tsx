import { type Abi } from "viem";
import { Card } from "../../components/Card";
import { RolesPanel } from "../../components/RolesPanel";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { MANAGER_ROLES } from "../../lib/roles";

export function ManagerRolesTab() {
  const { manager, managerAbi } = useActiveContracts();
  if (!manager) return <Card title="Manager Roles">Connect to a supported network.</Card>;
  return (
    <Card title="Manager Roles">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Expand a role to see its members and grant or revoke it. Grant/revoke requires the Admin role.
      </p>
      <RolesPanel address={manager} abi={managerAbi as Abi} roles={MANAGER_ROLES} enumerable />
    </Card>
  );
}
