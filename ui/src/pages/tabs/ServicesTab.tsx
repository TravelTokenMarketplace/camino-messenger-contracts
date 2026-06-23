import { useState } from "react";
import { Plus, Trash2, X } from "lucide-react";
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
import { useTx } from "../../tx/TxProvider";

// Explicit single-overload fragments: the CMAccount ABI overloads these by
// (string) and (bytes32), which makes viem's overload resolution ambiguous.
const RESTRICTED_RATE_ABI = [{ type: "function", name: "getServiceRestrictedRate", stateMutability: "view", inputs: [{ type: "bytes32" }], outputs: [{ type: "bool" }] }] as const;
const CAPABILITIES_ABI = [{ type: "function", name: "getServiceCapabilities", stateMutability: "view", inputs: [{ type: "bytes32" }], outputs: [{ type: "string[]" }] }] as const;

interface ServiceInfo {
  hash: Hex;
  name: string;
  restricted: boolean;
  capabilities: string[];
}

const inputClass =
  "rounded border border-gray-300 bg-white px-2 py-1.5 text-sm focus:border-indigo-500 focus:outline-none dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100";

function SupportedServiceRow({
  account,
  abi,
  service,
  hasRole,
  onChanged,
}: {
  account: Address;
  abi: Abi;
  service: ServiceInfo;
  hasRole: boolean;
  onChanged: () => void;
}) {
  const { writeContractAsync } = useWriteContract();
  const { track } = useTx();
  const [busy, setBusy] = useState(false);
  const [newCap, setNewCap] = useState("");

  async function run(label: string, functionName: string, args: unknown[], after?: () => void) {
    setBusy(true);
    try {
      await track({
        label,
        write: () => writeContractAsync({ address: account, abi, functionName, args }),
        onConfirmed: () => {
          after?.();
          onChanged();
        },
      });
    } catch {
      // Submission errors are surfaced by the transaction panel / wallet.
    } finally {
      setBusy(false);
    }
  }

  return (
    <li className="group rounded-md border border-gray-100 px-3 py-2 dark:border-gray-700/60">
      <div className="flex items-start justify-between gap-3">
        <span className="block min-w-0 break-all font-mono text-sm">{service.name}</span>
        {hasRole && (
          <RowAction>
            <TxButton
              label="Remove"
              variant="danger"
              icon={<Trash2 className="h-4 w-4" />}
              write={() => writeContractAsync({ address: account, abi, functionName: "removeService", args: [service.name] })}
              onConfirmed={onChanged}
            />
          </RowAction>
        )}
      </div>

      <div className="mt-2 flex flex-wrap items-center gap-1.5">
        {/* Restricted rate: an interactive pill when the user can edit, else a badge. */}
        {hasRole ? (
          <button
            type="button"
            disabled={busy}
            onClick={() =>
              run(
                `${service.restricted ? "Disable" : "Enable"} restricted rate · ${service.name}`,
                "setServiceRestrictedRate",
                [service.name, !service.restricted],
              )
            }
            className={`inline-flex items-center gap-1 rounded-full px-2 py-0.5 text-xs transition-colors disabled:opacity-50 ${
              service.restricted
                ? "bg-amber-100 text-amber-700 hover:bg-amber-200 dark:bg-amber-950 dark:text-amber-300"
                : "bg-gray-100 text-gray-500 hover:bg-gray-200 dark:bg-gray-700 dark:text-gray-400"
            }`}
            title="Toggle restricted rate"
          >
            Restricted rate: {service.restricted ? "on" : "off"}
          </button>
        ) : (
          service.restricted && (
            <span className="rounded-full bg-amber-100 px-2 py-0.5 text-xs text-amber-700 dark:bg-amber-950 dark:text-amber-300">Restricted rate</span>
          )
        )}

        {service.capabilities.map((c) => (
          <span key={c} className="inline-flex items-center gap-1 rounded bg-gray-100 px-2 py-0.5 text-xs text-gray-600 dark:bg-gray-700 dark:text-gray-300">
            {c}
            {hasRole && (
              <button
                type="button"
                disabled={busy}
                onClick={() => run(`Remove capability "${c}" · ${service.name}`, "removeServiceCapability", [service.name, c])}
                className="text-gray-400 hover:text-red-500 disabled:opacity-50"
                aria-label={`Remove capability ${c}`}
              >
                <X className="h-3 w-3" />
              </button>
            )}
          </span>
        ))}

        {hasRole && (
          <span className="inline-flex items-center gap-1">
            <input
              className={`w-32 ${inputClass} py-0.5`}
              placeholder="+ capability"
              value={newCap}
              onChange={(e) => setNewCap(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === "Enter" && newCap.trim()) {
                  run(`Add capability "${newCap.trim()}" · ${service.name}`, "addServiceCapability", [service.name, newCap.trim()], () => setNewCap(""));
                }
              }}
            />
            <button
              type="button"
              disabled={busy || !newCap.trim()}
              onClick={() => run(`Add capability "${newCap.trim()}" · ${service.name}`, "addServiceCapability", [service.name, newCap.trim()], () => setNewCap(""))}
              className="rounded bg-gray-100 px-2 py-0.5 text-xs text-gray-600 hover:bg-gray-200 disabled:opacity-50 dark:bg-gray-700 dark:text-gray-300"
            >
              Add
            </button>
          </span>
        )}
      </div>
    </li>
  );
}

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
  const services: ServiceInfo[] = hashes.map((h, i) => ({
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
        <ul className="mb-4 space-y-2">
          {services.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {services.map((s) => (
            <SupportedServiceRow key={s.hash} account={account} abi={abi} service={s} hasRole={hasRole} onChanged={refetch} />
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE">
        <div className="rounded-lg border border-gray-200 p-3 dark:border-gray-700">
          <h3 className="mb-3 text-sm font-medium text-gray-700 dark:text-gray-200">Add a service</h3>
          <div className="grid gap-3">
            <label className="block">
              <span className="mb-1 block text-xs font-medium text-gray-500 dark:text-gray-400">Service name</span>
              <input
                className={`w-full ${inputClass}`}
                placeholder="cmp.services.accommodation.v2.AccommodationSearchService"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </label>
            <label className="block">
              <span className="mb-1 block text-xs font-medium text-gray-500 dark:text-gray-400">
                Capabilities <span className="font-normal text-gray-400">(optional, comma separated)</span>
              </span>
              <input
                className={`w-full ${inputClass}`}
                placeholder="e.g. search, availability"
                value={caps}
                onChange={(e) => setCaps(e.target.value)}
              />
            </label>
            <label className="flex items-center gap-2 text-sm text-gray-700 dark:text-gray-200">
              <input type="checkbox" checked={restricted} onChange={(e) => setRestricted(e.target.checked)} />
              Restricted rate
            </label>
            <div className="flex justify-end">
              <TxButton
                label="Add service"
                icon={<Plus className="h-4 w-4" />}
                disabled={!name.trim()}
                write={() => writeContractAsync({
                  address: account,
                  abi,
                  functionName: "addService",
                  args: [name.trim(), restricted, caps.split(",").map((c) => c.trim()).filter(Boolean)],
                })}
                onConfirmed={() => { setName(""); setRestricted(false); setCaps(""); void refetch(); }}
              />
            </div>
          </div>
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
