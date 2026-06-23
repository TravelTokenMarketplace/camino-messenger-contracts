import { useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { type Abi, type Address, type Hex } from "viem";
import { useReadContract, useReadContracts, useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { ListManager } from "../../components/ListManager";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

// Explicit single-overload fragments: the CMAccount ABI overloads these by
// (string) and (bytes32), which makes viem's overload resolution ambiguous.
const RESTRICTED_RATE_ABI = [{ type: "function", name: "getServiceRestrictedRate", stateMutability: "view", inputs: [{ type: "bytes32" }], outputs: [{ type: "bool" }] }] as const;
const CAPABILITIES_ABI = [{ type: "function", name: "getServiceCapabilities", stateMutability: "view", inputs: [{ type: "bytes32" }], outputs: [{ type: "string[]" }] }] as const;

function SupportedServices({ account, abi, hasRole }: { account: Address; abi: Abi; hasRole: boolean }) {
  const { manager, managerAbi, chainId } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  // getSupportedServices() returns a (uint256,bool,string[])[] tuple that viem
  // cannot reliably decode, so list service hashes and resolve names + config
  // via per-hash getters instead.
  const { data: hashesData, isLoading: hashesLoading, refetch: refetchHashes } = useReadContract({
    chainId,
    address: account,
    abi,
    functionName: "getAllServiceHashes",
  });
  const hashes = (hashesData as Hex[] | undefined) ?? [];
  const { data: nameResults, isLoading: namesLoading, refetch: refetchNames } = useReadContracts({
    contracts: hashes.map((h) => ({
      chainId,
      address: manager,
      abi: managerAbi as Abi,
      functionName: "getRegisteredServiceNameByHash",
      args: [h],
    })),
    allowFailure: true,
    query: { enabled: hashes.length > 0 },
  });
  // Per-service config (restricted rate + capabilities), best-effort.
  const { data: configResults, refetch: refetchConfig } = useReadContracts({
    contracts: hashes.flatMap((h) => [
      { chainId, address: account, abi: RESTRICTED_RATE_ABI, functionName: "getServiceRestrictedRate", args: [h] },
      { chainId, address: account, abi: CAPABILITIES_ABI, functionName: "getServiceCapabilities", args: [h] },
    ]),
    allowFailure: true,
    query: { enabled: hashes.length > 0 },
  });
  const services = hashes.map((h, i) => ({
    hash: h,
    name: (nameResults?.[i]?.result as string | undefined) ?? h,
    restricted: configResults?.[i * 2]?.result === true,
    capabilities: (configResults?.[i * 2 + 1]?.result as string[] | undefined) ?? [],
  }));
  const isLoading = hashesLoading || (hashes.length > 0 && namesLoading);
  const refetch = () => { void refetchHashes(); void refetchNames(); void refetchConfig(); };
  const [name, setName] = useState("");
  const [restricted, setRestricted] = useState(false);
  const [caps, setCaps] = useState("");

  return (
    <Card title="Supported Services">
      {isLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y dark:divide-gray-700">
          {services.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {services.map((s) => (
            <li key={s.hash} className="group flex items-center justify-between gap-3 py-2">
              <span className="min-w-0">
                <span className="block break-all font-mono text-sm">{s.name}</span>
                <span className="mt-1 flex flex-wrap items-center gap-1">
                  {s.restricted && (
                    <span className="rounded bg-amber-100 px-2 py-0.5 text-xs text-amber-700 dark:bg-amber-950 dark:text-amber-300">Restricted rate</span>
                  )}
                  {s.capabilities.map((c) => (
                    <span key={c} className="rounded bg-gray-100 px-2 py-0.5 text-xs text-gray-600 dark:bg-gray-700 dark:text-gray-300">{c}</span>
                  ))}
                </span>
              </span>
              {hasRole && (
                <RowAction>
                  <TxButton label="Remove" variant="danger" icon={<Trash2 className="h-4 w-4" />} write={() => writeContractAsync({ address: account, abi, functionName: "removeService", args: [s.name] })} onConfirmed={() => void refetch()} />
                </RowAction>
              )}
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
            icon={<Plus className="h-4 w-4" />}
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
