import { useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { type Abi } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { CopyButton } from "../../components/CopyButton";
import { inputClass } from "../../components/Input";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { groupServicesByPackage, parseServiceName } from "../../lib/serviceName";

export function ServiceRegistryTab() {
  const { manager, managerAbi, chainId, supported } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { hasRole } = useHasRole(manager, abi, "SERVICE_REGISTRY_ADMIN_ROLE");
  const { writeContractAsync } = useWriteContract();
  const { data, isLoading, refetch } = useReadContract({
    chainId, address: manager, abi, functionName: "getAllRegisteredServiceNames", query: { enabled: supported },
  });
  const names = (data as string[] | undefined) ?? [];
  const groups = groupServicesByPackage(names.map((n) => ({ name: n })));
  const [name, setName] = useState("");

  return (
    <Card title="Service Registry">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Services must be registered here before any CM Account can support or want them.
      </p>
      {isLoading ? <p className="py-2 text-sm text-gray-400">Loading…</p> : names.length === 0 ? (
        <p className="mb-4 py-2 text-sm text-gray-400">No services registered.</p>
      ) : (
        <div className="mb-4 space-y-5">
          {groups.map((g) => (
            <div key={g.pkg}>
              <h4 className="mb-1.5 text-xs font-semibold uppercase tracking-wide text-gray-500 dark:text-gray-400">{g.pkg} <span className="text-gray-400">{g.items.length}</span></h4>
              <ul className="space-y-2">
                {g.items.map((s) => {
                  const parsed = parseServiceName(s.name);
                  return (
                    <li key={s.name} className="group flex items-center justify-between gap-3 rounded-md border border-gray-100 px-3 py-2 dark:border-gray-700/60">
                      <span className="flex min-w-0 items-baseline gap-2">
                        {parsed.version && <span className="rounded bg-indigo-50 px-1.5 py-0.5 font-mono text-xs font-medium text-indigo-700 dark:bg-indigo-950 dark:text-indigo-300">{parsed.version}</span>}
                        <span className="break-all font-mono text-sm font-medium text-gray-900 dark:text-gray-100">{parsed.name}</span>
                        <CopyButton value={s.name} label="Copy full service name" />
                      </span>
                      {hasRole && (
                        <RowAction>
                          <TxButton
                            label="Unregister" variant="danger" icon={<Trash2 className="h-4 w-4" />}
                            tooltip="Unregisters this service from the manager — sends a transaction to your wallet."
                            write={() => writeContractAsync({ address: manager!, abi, functionName: "unregisterService", args: [s.name] })}
                            onConfirmed={() => refetch()}
                          />
                        </RowAction>
                      )}
                    </li>
                  );
                })}
              </ul>
            </div>
          ))}
        </div>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_REGISTRY_ADMIN_ROLE" action="Register service">
        <div className="flex items-end gap-2">
          <input
            className={`flex-1 ${inputClass}`}
            placeholder="cmp.services.<package>.<version>.<Name>"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <TxButton
            label="Register" icon={<Plus className="h-4 w-4" />} disabled={!name.trim()}
            tooltip="Registers a new service name in the manager — sends a transaction to your wallet."
            write={() => writeContractAsync({ address: manager!, abi, functionName: "registerService", args: [name.trim()] })}
            onConfirmed={() => { setName(""); refetch(); }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}
