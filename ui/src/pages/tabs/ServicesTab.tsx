import { useId, useState } from "react";
import { ChevronRight, Plus, Trash2, X } from "lucide-react";
import { type Abi, type Address, type Hex } from "viem";
import { useReadContract, useReadContracts, useWriteContract } from "wagmi";
import { Autocomplete } from "../../components/Autocomplete";
import { Card } from "../../components/Card";
import { ListManager } from "../../components/ListManager";
import { RoleGate } from "../../components/RoleGate";
import { Tooltip } from "../../components/Tooltip";
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
  open,
  onToggle,
  onChanged,
}: {
  account: Address;
  abi: Abi;
  service: ServiceInfo;
  hasRole: boolean;
  open: boolean;
  onToggle: () => void;
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
    <li className="rounded-md border border-gray-100 dark:border-gray-700/60">
      {/* Header — click to expand the editing controls for this service. */}
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-start gap-2 px-3 py-2 text-left"
      >
        <ChevronRight className={`mt-0.5 h-4 w-4 shrink-0 text-gray-400 transition-transform ${open ? "rotate-90" : ""}`} />
        <span className="min-w-0 flex-1">
          <span className="block break-all font-mono text-sm">{service.name}</span>
          <span className="mt-1 flex flex-wrap items-center gap-1">
            {service.restricted && (
              <span className="rounded-full bg-amber-100 px-2 py-0.5 text-xs text-amber-700 dark:bg-amber-950 dark:text-amber-300">Restricted rate</span>
            )}
            {service.capabilities.map((c) => (
              <span key={c} className="rounded bg-gray-100 px-2 py-0.5 text-xs text-gray-600 dark:bg-gray-700 dark:text-gray-300">{c}</span>
            ))}
            {open && !service.restricted && service.capabilities.length === 0 && (
              <span className="text-xs text-gray-400">No restrictions or capabilities</span>
            )}
          </span>
        </span>
      </button>

      {open && (
        <div className="space-y-3 border-t border-gray-100 px-3 py-3 dark:border-gray-700/60">
          {hasRole ? (
            <>
              <div className="flex flex-wrap items-center gap-2">
                <span className="text-xs font-medium text-gray-500 dark:text-gray-400">Rate</span>
                <Tooltip
                  content={
                    service.restricted
                      ? "Restricted rate is ON. Click to disable it — sends a transaction to your wallet."
                      : "Restricted rate is OFF. Click to enable it — sends a transaction to your wallet."
                  }
                >
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
                  >
                    Restricted rate: {service.restricted ? "on" : "off"}
                  </button>
                </Tooltip>
              </div>

              <div>
                <span className="text-xs font-medium text-gray-500 dark:text-gray-400">Capabilities</span>
                <div className="mt-1 flex flex-wrap items-center gap-1.5">
                  {service.capabilities.length === 0 && <span className="text-xs text-gray-400">None</span>}
                  {service.capabilities.map((c) => (
                    <span key={c} className="inline-flex items-center gap-1 rounded bg-gray-100 px-2 py-0.5 text-xs text-gray-600 dark:bg-gray-700 dark:text-gray-300">
                      {c}
                      <Tooltip content={`Remove capability "${c}" — sends a transaction to your wallet.`}>
                        <button
                          type="button"
                          disabled={busy}
                          onClick={() => run(`Remove capability "${c}" · ${service.name}`, "removeServiceCapability", [service.name, c])}
                          className="text-gray-400 hover:text-red-500 disabled:opacity-50"
                          aria-label={`Remove capability ${c}`}
                        >
                          <X className="h-3 w-3" />
                        </button>
                      </Tooltip>
                    </span>
                  ))}
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
                  <Tooltip content="Add this capability to the service — sends a transaction to your wallet.">
                    <button
                      type="button"
                      disabled={busy || !newCap.trim()}
                      onClick={() => run(`Add capability "${newCap.trim()}" · ${service.name}`, "addServiceCapability", [service.name, newCap.trim()], () => setNewCap(""))}
                      className="rounded bg-gray-100 px-2 py-0.5 text-xs text-gray-600 hover:bg-gray-200 disabled:opacity-50 dark:bg-gray-700 dark:text-gray-300"
                    >
                      Add
                    </button>
                  </Tooltip>
                </div>
              </div>

              <div className="flex justify-end">
                <TxButton
                  label="Remove service"
                  variant="danger"
                  icon={<Trash2 className="h-4 w-4" />}
                  tooltip="Removes this service from the account — sends a transaction to your wallet."
                  write={() => writeContractAsync({ address: account, abi, functionName: "removeService", args: [service.name] })}
                  onConfirmed={onChanged}
                />
              </div>
            </>
          ) : (
            <p className="text-xs text-gray-400">You need the SERVICE_ADMIN_ROLE to edit this service.</p>
          )}
        </div>
      )}
    </li>
  );
}

function SupportedServices({ account, abi, hasRole, registered }: { account: Address; abi: Abi; hasRole: boolean; registered: string[] }) {
  const { manager, managerAbi, chainId } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const serviceInputId = useId();
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
  const [openHash, setOpenHash] = useState<Hex | null>(null);
  const [name, setName] = useState("");
  const [restricted, setRestricted] = useState(false);
  const [caps, setCaps] = useState("");

  return (
    <Card title="Supported Services">
      {isLoading ? <p>Loading…</p> : (
        <ul className="mb-4 space-y-2">
          {services.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {services.map((s) => (
            <SupportedServiceRow
              key={s.hash}
              account={account}
              abi={abi}
              service={s}
              hasRole={hasRole}
              open={openHash === s.hash}
              onToggle={() => setOpenHash((cur) => (cur === s.hash ? null : s.hash))}
              onChanged={refetch}
            />
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE" action="Add service">
        <div className="rounded-lg border border-gray-200 p-3 dark:border-gray-700">
          <h3 className="mb-3 text-sm font-medium text-gray-700 dark:text-gray-200">Add a service</h3>
          <div className="grid gap-3">
            <label className="block">
              <span className="mb-1 block text-xs font-medium text-gray-500 dark:text-gray-400">
                Service name <span className="font-normal text-gray-400">(must be registered in the manager)</span>
              </span>
              <Autocomplete
                id={serviceInputId}
                value={name}
                onChange={setName}
                options={registered.filter((n) => !services.some((s) => s.name === n))}
                placeholder="Click to pick a registered service…"
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
                tooltip="Adds a supported service to the account — sends a transaction to your wallet."
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
  const { cmAccountAbi, manager, managerAbi, chainId } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const wanted = useContractList(account, abi, "getWantedServices");
  const { hasRole } = useHasRole(account, abi, "SERVICE_ADMIN_ROLE");
  // Services can only reference names registered in the manager — surface them
  // as autocomplete suggestions so users don't have to know the exact string.
  const { data: registeredData } = useReadContract({
    chainId,
    address: manager,
    abi: managerAbi as Abi,
    functionName: "getAllRegisteredServiceNames",
  });
  const registered = (registeredData as string[] | undefined) ?? [];

  return (
    <div className="grid gap-4">
      <SupportedServices account={account} abi={abi} hasRole={hasRole} registered={registered} />
      <ListManager
        title="Wanted Services"
        items={wanted.items}
        isLoading={wanted.isLoading}
        roleName="SERVICE_ADMIN_ROLE"
        hasRole={hasRole}
        addLabel="Add wanted"
        addPlaceholder="Service name"
        suggestions={registered.filter((n) => !wanted.items.includes(n))}
        onAdd={(v) => writeContractAsync({ address: account, abi, functionName: "addWantedServices", args: [[v]] })}
        onRemove={(v) => writeContractAsync({ address: account, abi, functionName: "removeWantedServices", args: [[v]] })}
        onChanged={wanted.refetch}
      />
    </div>
  );
}
