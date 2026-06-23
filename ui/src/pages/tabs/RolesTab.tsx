import { useState } from "react";
import { ChevronRight, ShieldPlus, Trash2 } from "lucide-react";
import { type Abi, type Address } from "viem";
import { useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { ACCOUNT_ROLES, roleHash, type RoleName } from "../../lib/roles";
import { shortRoleName } from "../../lib/format";

const inputClass =
  "rounded border border-gray-300 bg-white px-2 py-1.5 text-sm focus:border-indigo-500 focus:outline-none dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100";

function RoleRow({
  account,
  abi,
  role,
  hasAdmin,
  open,
  onToggle,
}: {
  account: Address;
  abi: Abi;
  role: RoleName;
  hasAdmin: boolean;
  open: boolean;
  onToggle: () => void;
}) {
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, abi, role);
  const [grantee, setGrantee] = useState("");
  const label = shortRoleName(role);

  return (
    <li className="rounded-md border border-gray-100 dark:border-gray-700/60">
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-center gap-2 px-3 py-2 text-left"
      >
        <ChevronRight className={`h-4 w-4 shrink-0 text-gray-400 transition-transform ${open ? "rotate-90" : ""}`} />
        <span className="flex-1 text-sm font-medium text-gray-800 dark:text-gray-100">{label}</span>
        <span className="font-mono text-[11px] text-gray-400">{role}</span>
        <span
          className={`rounded-full px-2 py-0.5 text-xs font-medium ${
            !isLoading && members.length > 0
              ? "bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300"
              : "bg-gray-100 text-gray-500 dark:bg-gray-700 dark:text-gray-400"
          }`}
          title={`${members.length} address${members.length === 1 ? "" : "es"} with this role`}
        >
          {isLoading ? "…" : members.length}
        </span>
      </button>

      {open && (
        <div className="space-y-3 border-t border-gray-100 px-3 py-3 dark:border-gray-700/60">
          <ul className="divide-y dark:divide-gray-700">
            {isLoading && <li className="py-2 text-sm text-gray-400">Loading…</li>}
            {!isLoading && members.length === 0 && <li className="py-2 text-sm text-gray-400">No members</li>}
            {members.map((m) => (
              <li key={m} className="group flex items-center justify-between gap-3 py-2">
                <AddressDisplay address={m} className="text-sm" />
                {hasAdmin && (
                  <RowAction>
                    <TxButton
                      label="Revoke"
                      variant="danger"
                      icon={<Trash2 className="h-4 w-4" />}
                      tooltip={`Revokes ${label} from this address — sends a transaction to your wallet.`}
                      write={() => writeContractAsync({ address: account, abi, functionName: "revokeRole", args: [roleHash(role), m as Address] })}
                      onConfirmed={refetch}
                    />
                  </RowAction>
                )}
              </li>
            ))}
          </ul>

          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`grant ${label}`}>
            <div className="flex items-end gap-2">
              <input
                className={`flex-1 ${inputClass}`}
                placeholder="Address 0x…"
                value={grantee}
                onChange={(e) => setGrantee(e.target.value)}
              />
              <TxButton
                label="Grant"
                icon={<ShieldPlus className="h-4 w-4" />}
                disabled={!grantee.trim()}
                tooltip={`Grants ${label} to this address — sends a transaction to your wallet.`}
                write={() => writeContractAsync({ address: account, abi, functionName: "grantRole", args: [roleHash(role), grantee.trim() as Address] })}
                onConfirmed={() => { setGrantee(""); refetch(); }}
              />
            </div>
          </RoleGate>
        </div>
      )}
    </li>
  );
}

export function RolesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { hasRole } = useHasRole(account, abi, "DEFAULT_ADMIN_ROLE");
  const [openRole, setOpenRole] = useState<RoleName | null>(null);

  return (
    <Card title="Roles">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Expand a role to see its members and grant or revoke it. Only one role is open at a time.
      </p>
      <ul className="space-y-2">
        {ACCOUNT_ROLES.map((r) => (
          <RoleRow
            key={r}
            account={account}
            abi={abi}
            role={r}
            hasAdmin={hasRole}
            open={openRole === r}
            onToggle={() => setOpenRole((cur) => (cur === r ? null : r))}
          />
        ))}
      </ul>
    </Card>
  );
}
