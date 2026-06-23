import { useState } from "react";
import { type Abi, type Address } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { ListManager } from "../../components/ListManager";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

function SupportedServices({ account, abi, hasRole }: { account: Address; abi: Abi; hasRole: boolean }) {
  const { writeContractAsync } = useWriteContract();
  // getSupportedServices() returns [names: string[], configs: tuple[]]; we list the names.
  const { data, isLoading, refetch } = useReadContract({ address: account, abi, functionName: "getSupportedServices" });
  const names = (((data as unknown[]) ?? [])[0] as string[] | undefined) ?? [];
  const [name, setName] = useState("");
  const [restricted, setRestricted] = useState(false);
  const [caps, setCaps] = useState("");

  return (
    <Card title="Supported Services">
      {isLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y">
          {names.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {names.map((s) => (
            <li key={s} className="flex items-center justify-between py-2">
              <span className="font-mono text-sm">{s}</span>
              <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE">
                <TxButton label="Remove" write={() => writeContractAsync({ address: account, abi, functionName: "removeService", args: [s] })} onConfirmed={() => void refetch()} />
              </RoleGate>
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE">
        <div className="flex flex-wrap items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Service name (e.g. cmp.services.accommodation.v2.AccommodationSearchService)" value={name} onChange={(e) => setName(e.target.value)} />
          <label className="flex items-center gap-1 text-sm"><input type="checkbox" checked={restricted} onChange={(e) => setRestricted(e.target.checked)} /> Restricted rate</label>
          <input className="w-48 rounded border px-2 py-1" placeholder="Capabilities (comma separated)" value={caps} onChange={(e) => setCaps(e.target.value)} />
          <TxButton
            label="Add service"
            disabled={!name}
            write={() => writeContractAsync({
              address: account,
              abi,
              functionName: "addService",
              args: [name, restricted, caps.split(",").map((c) => c.trim()).filter(Boolean)],
            })}
            onConfirmed={() => { setName(""); setRestricted(false); setCaps(""); void refetch(); }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}

export function ServicesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const wanted = useContractList(account, abi, "getWantedServices");
  const { hasRole } = useHasRole(account, abi, "SERVICE_ADMIN_ROLE");

  return (
    <div className="grid gap-4">
      <SupportedServices account={account} abi={abi} hasRole={hasRole} />
      <ListManager
        title="Wanted Services"
        items={wanted.items}
        isLoading={wanted.isLoading}
        roleName="SERVICE_ADMIN_ROLE"
        hasRole={hasRole}
        addLabel="Add wanted"
        addPlaceholder="Service name"
        onAdd={(v) => writeContractAsync({ address: account, abi, functionName: "addWantedServices", args: [[v]] })}
        onRemove={(v) => writeContractAsync({ address: account, abi, functionName: "removeWantedServices", args: [[v]] })}
        onChanged={wanted.refetch}
      />
    </div>
  );
}
