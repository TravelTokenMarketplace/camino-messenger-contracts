import { useState } from "react";
import { ChevronRight, ShieldCheck, ShieldPlus, Trash2 } from "lucide-react";
import { type Abi, type Address, isAddress } from "viem";
import { useWriteContract } from "wagmi";
import { AddressDisplay } from "./AddressDisplay";
import { inputClass } from "./Input";
import { RoleGate } from "./RoleGate";
import { RowAction } from "./RowAction";
import { TxButton } from "./TxButton";
import { useHasRole } from "../hooks/useHasRole";
import { useRoleMembers } from "../hooks/useRoleMembers";
import { roleHash, type RoleName } from "../lib/roles";
import { shortRoleName } from "../lib/format";

function RoleHeader({
  label,
  role,
  open,
  badge,
}: {
  label: string;
  role: RoleName;
  open: boolean;
  badge: React.ReactNode;
}) {
  return (
    <>
      <ChevronRight className={`h-4 w-4 shrink-0 text-gray-400 transition-transform ${open ? "rotate-90" : ""}`} />
      <span className="flex-1 text-sm font-medium text-gray-800 dark:text-gray-100">{label}</span>
      <span className="font-mono text-[11px] text-gray-400">{role}</span>
      {badge}
    </>
  );
}

function GrantForm({
  account,
  abi,
  role,
  label,
  onDone,
}: {
  account: Address;
  abi: Abi;
  role: RoleName;
  label: string;
  onDone: () => void;
}) {
  const { writeContractAsync } = useWriteContract();
  const [grantee, setGrantee] = useState("");
  const trimmed = grantee.trim();
  const valid = isAddress(trimmed);
  return (
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
        disabled={!valid}
        tooltip={`Grants ${label} to this address — sends a transaction to your wallet.`}
        write={() =>
          writeContractAsync({
            address: account,
            abi,
            functionName: "grantRole",
            args: [roleHash(role), trimmed as Address],
          })
        }
        onConfirmed={() => {
          setGrantee("");
          onDone();
        }}
      />
    </div>
  );
}

function EnumerableRoleRow({
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
  const label = shortRoleName(role);

  return (
    <li className="rounded-md border border-gray-100 dark:border-gray-700/60">
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-center gap-2 px-3 py-2 text-left"
      >
        <RoleHeader
          label={label}
          role={role}
          open={open}
          badge={
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
          }
        />
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
                      write={() =>
                        writeContractAsync({
                          address: account,
                          abi,
                          functionName: "revokeRole",
                          args: [roleHash(role), m as Address],
                        })
                      }
                      onConfirmed={refetch}
                    />
                  </RowAction>
                )}
              </li>
            ))}
          </ul>
          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`grant ${label}`}>
            <GrantForm account={account} abi={abi} role={role} label={label} onDone={refetch} />
          </RoleGate>
        </div>
      )}
    </li>
  );
}

function NonEnumerableRoleRow({
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
  const { hasRole: youHold } = useHasRole(account, abi, role);
  const [revokee, setRevokee] = useState("");
  const label = shortRoleName(role);

  return (
    <li className="rounded-md border border-gray-100 dark:border-gray-700/60">
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-center gap-2 px-3 py-2 text-left"
      >
        <RoleHeader
          label={label}
          role={role}
          open={open}
          badge={
            youHold ? (
              <span className="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2 py-0.5 text-xs font-medium text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300">
                <ShieldCheck className="h-3 w-3" /> You
              </span>
            ) : (
              <span className="w-px" />
            )
          }
        />
      </button>
      {open && (
        <div className="space-y-3 border-t border-gray-100 px-3 py-3 dark:border-gray-700/60">
          <p className="text-xs text-gray-400">
            This contract cannot list role members on-chain. Grant or revoke by address.
          </p>
          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`manage ${label}`}>
            <div className="space-y-2">
              <GrantForm
                account={account}
                abi={abi}
                role={role}
                label={label}
                onDone={() => {
                  /* no member list to refetch in non-enumerable mode */
                }}
              />
              <div className="flex items-end gap-2">
                <input
                  className={`flex-1 ${inputClass}`}
                  placeholder="Address 0x… to revoke"
                  value={revokee}
                  onChange={(e) => setRevokee(e.target.value)}
                />
                <TxButton
                  label="Revoke"
                  variant="danger"
                  icon={<Trash2 className="h-4 w-4" />}
                  disabled={!isAddress(revokee.trim())}
                  tooltip={`Revokes ${label} from this address — sends a transaction to your wallet.`}
                  write={() =>
                    writeContractAsync({
                      address: account,
                      abi,
                      functionName: "revokeRole",
                      args: [roleHash(role), revokee.trim() as Address],
                    })
                  }
                  onConfirmed={() => setRevokee("")}
                />
              </div>
            </div>
          </RoleGate>
        </div>
      )}
    </li>
  );
}

export function RolesPanel({
  address,
  abi,
  roles,
  enumerable,
}: {
  address: Address;
  abi: Abi;
  roles: readonly RoleName[];
  enumerable: boolean;
}) {
  const { hasRole: hasAdmin } = useHasRole(address, abi, "DEFAULT_ADMIN_ROLE");
  const [openRole, setOpenRole] = useState<RoleName | null>(null);
  const Row = enumerable ? EnumerableRoleRow : NonEnumerableRoleRow;

  return (
    <ul className="space-y-2">
      {roles.map((r) => (
        <Row
          key={r}
          account={address}
          abi={abi}
          role={r}
          hasAdmin={hasAdmin}
          open={openRole === r}
          onToggle={() => setOpenRole((cur) => (cur === r ? null : r))}
        />
      ))}
    </ul>
  );
}
