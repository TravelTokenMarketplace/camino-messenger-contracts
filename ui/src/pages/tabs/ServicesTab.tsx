import { useId, useState } from "react";
import { ChevronRight, Plus, Trash2, X } from "lucide-react";
import { type Abi, type Address, type Hex } from "viem";
import { useReadContract, useReadContracts, useWriteContract } from "wagmi";
import { Autocomplete } from "../../components/Autocomplete";
import { Card } from "../../components/Card";
import { Checkbox } from "../../components/Checkbox";
import { CopyButton } from "../../components/CopyButton";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { Tooltip } from "../../components/Tooltip";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";
import { type ParsedService, groupServicesByPackage } from "../../lib/serviceName";
import { useTx } from "../../tx/TxProvider";

// Deterministic accent colour per package, for quick visual grouping.
const PKG_DOTS = [
  "bg-camino-500",
  "bg-camino-500",
  "bg-amber-500",
  "bg-sky-500",
  "bg-rose-500",
  "bg-violet-500",
  "bg-teal-500",
  "bg-orange-500",
];
function pkgColor(pkg: string): string {
  let h = 0;
  for (let i = 0; i < pkg.length; i++) h = (h * 31 + pkg.charCodeAt(i)) >>> 0;
  return PKG_DOTS[h % PKG_DOTS.length];
}

function PackageHeader({ pkg, count }: { pkg: string; count: number }) {
  return (
    <div className="mb-1.5 flex items-center gap-2">
      <span className={`h-2 w-2 rounded-full ${pkgColor(pkg)}`} aria-hidden />
      <h4 className="text-xs font-semibold uppercase tracking-wide text-tarmac-500 dark:text-tarmac-400">{pkg}</h4>
      <span className="text-xs text-tarmac-400">{count}</span>
    </div>
  );
}

/** A service name shown as an optional version pill + the service name. */
function ServiceLabel({ parsed }: { parsed: ParsedService }) {
  return (
    <span className="flex min-w-0 flex-wrap items-baseline gap-x-2 gap-y-1">
      {parsed.version && (
        <span className="rounded bg-camino-50 px-1.5 py-0.5 font-mono text-xs font-medium text-camino-700 dark:bg-camino-950 dark:text-camino-300">
          {parsed.version}
        </span>
      )}
      <span className="break-all font-mono text-sm font-medium text-tarmac-900 dark:text-tarmac-100">
        {parsed.name}
      </span>
    </span>
  );
}

// Explicit single-overload fragments: the CMAccount ABI overloads these by
// (string) and (bytes32), which makes viem's overload resolution ambiguous.
const RESTRICTED_RATE_ABI = [
  {
    type: "function",
    name: "getServiceRestrictedRate",
    stateMutability: "view",
    inputs: [{ type: "bytes32" }],
    outputs: [{ type: "bool" }],
  },
] as const;
const CAPABILITIES_ABI = [
  {
    type: "function",
    name: "getServiceCapabilities",
    stateMutability: "view",
    inputs: [{ type: "bytes32" }],
    outputs: [{ type: "string[]" }],
  },
] as const;

interface ServiceInfo {
  hash: Hex;
  name: string;
  restricted: boolean;
  capabilities: string[];
}

const inputClass =
  "rounded border border-tarmac-300 bg-paper-raised px-2 py-1.5 text-sm focus:border-camino-500 focus:outline-none dark:border-tarmac-600 dark:bg-tarmac-900 dark:text-tarmac-100";

function SupportedServiceRow({
  account,
  abi,
  service,
  parsed,
  hasRole,
  open,
  onToggle,
  onChanged,
}: {
  account: Address;
  abi: Abi;
  service: ServiceInfo;
  parsed: ParsedService;
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
    <li className="rounded-md border border-tarmac-100 dark:border-tarmac-700/60">
      {/* Header — click to expand the editing controls for this service. */}
      <button
        type="button"
        onClick={onToggle}
        aria-expanded={open}
        className="flex w-full items-start gap-2 px-3 py-2 text-left"
      >
        <ChevronRight
          className={`mt-0.5 h-4 w-4 shrink-0 text-tarmac-400 transition-transform ${open ? "rotate-90" : ""}`}
        />
        <span className="min-w-0 flex-1">
          <ServiceLabel parsed={parsed} />
          {(service.restricted || service.capabilities.length > 0) && (
            <span className="mt-1 flex flex-wrap items-center gap-1">
              {service.restricted && (
                <span className="rounded-full bg-amber-100 px-2 py-0.5 text-xs text-amber-700 dark:bg-amber-950 dark:text-amber-300">
                  Restricted rate
                </span>
              )}
              {service.capabilities.map((c) => (
                <span
                  key={c}
                  className="rounded bg-tarmac-100 px-2 py-0.5 text-xs text-tarmac-600 dark:bg-tarmac-700 dark:text-tarmac-300"
                >
                  {c}
                </span>
              ))}
            </span>
          )}
        </span>
      </button>

      {open && (
        <div className="space-y-3 border-t border-tarmac-100 px-3 py-3 dark:border-tarmac-700/60">
          <div className="flex items-start gap-2">
            <code className="min-w-0 break-all font-mono text-xs text-tarmac-500 dark:text-tarmac-400">
              {service.name}
            </code>
            <CopyButton value={service.name} label="Copy full service name" />
          </div>
          {hasRole ? (
            <>
              <div className="flex flex-wrap items-center gap-2">
                <span className="text-xs font-medium text-tarmac-500 dark:text-tarmac-400">Rate</span>
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
                        : "bg-tarmac-100 text-tarmac-500 hover:bg-tarmac-200 dark:bg-tarmac-700 dark:text-tarmac-400"
                    }`}
                  >
                    Restricted rate: {service.restricted ? "on" : "off"}
                  </button>
                </Tooltip>
              </div>

              <div>
                <span className="text-xs font-medium text-tarmac-500 dark:text-tarmac-400">Capabilities</span>
                <div className="mt-1 flex flex-wrap items-center gap-1.5">
                  {service.capabilities.length === 0 && <span className="text-xs text-tarmac-400">None</span>}
                  {service.capabilities.map((c) => (
                    <span
                      key={c}
                      className="inline-flex items-center gap-1 rounded bg-tarmac-100 px-2 py-0.5 text-xs text-tarmac-600 dark:bg-tarmac-700 dark:text-tarmac-300"
                    >
                      {c}
                      <Tooltip content={`Remove capability "${c}" — sends a transaction to your wallet.`}>
                        <button
                          type="button"
                          disabled={busy}
                          onClick={() =>
                            run(`Remove capability "${c}" · ${service.name}`, "removeServiceCapability", [
                              service.name,
                              c,
                            ])
                          }
                          className="text-tarmac-400 hover:text-red-500 disabled:opacity-50"
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
                        run(
                          `Add capability "${newCap.trim()}" · ${service.name}`,
                          "addServiceCapability",
                          [service.name, newCap.trim()],
                          () => setNewCap(""),
                        );
                      }
                    }}
                  />
                  <Tooltip content="Add this capability to the service — sends a transaction to your wallet.">
                    <button
                      type="button"
                      disabled={busy || !newCap.trim()}
                      onClick={() =>
                        run(
                          `Add capability "${newCap.trim()}" · ${service.name}`,
                          "addServiceCapability",
                          [service.name, newCap.trim()],
                          () => setNewCap(""),
                        )
                      }
                      className="rounded bg-tarmac-100 px-2 py-0.5 text-xs text-tarmac-600 hover:bg-tarmac-200 disabled:opacity-50 dark:bg-tarmac-700 dark:text-tarmac-300"
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
                  write={() =>
                    writeContractAsync({ address: account, abi, functionName: "removeService", args: [service.name] })
                  }
                  onConfirmed={onChanged}
                />
              </div>
            </>
          ) : (
            <p className="text-xs text-tarmac-400">You need the SERVICE_ADMIN_ROLE to edit this service.</p>
          )}
        </div>
      )}
    </li>
  );
}

function SupportedServices({
  account,
  abi,
  hasRole,
  registered,
}: {
  account: Address;
  abi: Abi;
  hasRole: boolean;
  registered: string[];
}) {
  const { manager, managerAbi, chainId } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const serviceInputId = useId();
  // getSupportedServices() returns a (uint256,bool,string[])[] tuple that viem
  // cannot reliably decode, so list service hashes and resolve names + config
  // via per-hash getters instead.
  const {
    data: hashesData,
    isLoading: hashesLoading,
    refetch: refetchHashes,
  } = useReadContract({
    chainId,
    address: account,
    abi,
    functionName: "getAllServiceHashes",
  });
  const hashes = (hashesData as Hex[] | undefined) ?? [];
  const {
    data: nameResults,
    isLoading: namesLoading,
    refetch: refetchNames,
  } = useReadContracts({
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
  const {
    data: configResults,
    isLoading: configLoading,
    refetch: refetchConfig,
  } = useReadContracts({
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
  // Gate on config too: rows expose restricted-rate / capability controls whose
  // defaults (false / []) would otherwise render before config settles and could
  // drive the wrong action (e.g. inverted restricted-rate toggle).
  const isLoading = hashesLoading || (hashes.length > 0 && (namesLoading || configLoading));
  const refetch = () => {
    void refetchHashes();
    void refetchNames();
    void refetchConfig();
  };
  const [openHash, setOpenHash] = useState<Hex | null>(null);
  const [name, setName] = useState("");
  const [restricted, setRestricted] = useState(false);
  const [caps, setCaps] = useState("");

  // Group by package so related services cluster together.
  const groups = groupServicesByPackage(services);

  return (
    <Card title="Supported Services">
      {isLoading ? (
        <p>Loading…</p>
      ) : services.length === 0 ? (
        <p className="mb-4 py-2 text-sm text-tarmac-400">None</p>
      ) : (
        <div className="mb-4 space-y-5">
          {groups.map((g) => (
            <div key={g.pkg}>
              <PackageHeader pkg={g.pkg} count={g.items.length} />
              <ul className="space-y-2">
                {g.items.map((s) => (
                  <SupportedServiceRow
                    key={s.hash}
                    account={account}
                    abi={abi}
                    service={s}
                    parsed={s.parsed}
                    hasRole={hasRole}
                    open={openHash === s.hash}
                    onToggle={() => setOpenHash((cur) => (cur === s.hash ? null : s.hash))}
                    onChanged={refetch}
                  />
                ))}
              </ul>
            </div>
          ))}
        </div>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE" action="Add service">
        <div className="rounded-lg border border-tarmac-200 p-3 dark:border-tarmac-700">
          <h3 className="mb-3 text-sm font-medium text-tarmac-700 dark:text-tarmac-200">Add a service</h3>
          <div className="grid gap-3">
            <label className="block">
              <span className="mb-1 block text-xs font-medium text-tarmac-500 dark:text-tarmac-400">
                Service name <span className="font-normal text-tarmac-400">(must be registered in the manager)</span>
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
              <span className="mb-1 block text-xs font-medium text-tarmac-500 dark:text-tarmac-400">
                Capabilities <span className="font-normal text-tarmac-400">(optional, comma separated)</span>
              </span>
              <input
                className={`w-full ${inputClass}`}
                placeholder="e.g. search, availability"
                value={caps}
                onChange={(e) => setCaps(e.target.value)}
              />
            </label>
            <Checkbox checked={restricted} onChange={setRestricted} label="Restricted rate" />
            <div className="flex justify-end">
              <TxButton
                label="Add service"
                icon={<Plus className="h-4 w-4" />}
                disabled={!name.trim()}
                tooltip="Adds a supported service to the account — sends a transaction to your wallet."
                write={() =>
                  writeContractAsync({
                    address: account,
                    abi,
                    functionName: "addService",
                    args: [
                      name.trim(),
                      restricted,
                      caps
                        .split(",")
                        .map((c) => c.trim())
                        .filter(Boolean),
                    ],
                  })
                }
                onConfirmed={() => {
                  setName("");
                  setRestricted(false);
                  setCaps("");
                  void refetch();
                }}
              />
            </div>
          </div>
        </div>
      </RoleGate>
    </Card>
  );
}

function WantedServices({
  account,
  abi,
  hasRole,
  registered,
}: {
  account: Address;
  abi: Abi;
  hasRole: boolean;
  registered: string[];
}) {
  const { writeContractAsync } = useWriteContract();
  const { items, isLoading, refetch } = useContractList(account, abi, "getWantedServices");
  const [name, setName] = useState("");
  const groups = groupServicesByPackage(items.map((n) => ({ name: n })));

  return (
    <Card title="Wanted Services">
      {isLoading ? (
        <p>Loading…</p>
      ) : items.length === 0 ? (
        <p className="mb-4 py-2 text-sm text-tarmac-400">None</p>
      ) : (
        <div className="mb-4 space-y-5">
          {groups.map((g) => (
            <div key={g.pkg}>
              <PackageHeader pkg={g.pkg} count={g.items.length} />
              <ul className="space-y-2">
                {g.items.map((s) => (
                  <li
                    key={s.name}
                    className="group flex items-center justify-between gap-3 rounded-md border border-tarmac-100 px-3 py-2 dark:border-tarmac-700/60"
                  >
                    <span className="flex min-w-0 items-center gap-2">
                      <ServiceLabel parsed={s.parsed} />
                      <CopyButton value={s.name} label="Copy full service name" />
                    </span>
                    {hasRole && (
                      <RowAction>
                        <TxButton
                          label="Remove"
                          variant="danger"
                          icon={<Trash2 className="h-4 w-4" />}
                          tooltip="Removes this wanted service — sends a transaction to your wallet."
                          write={() =>
                            writeContractAsync({
                              address: account,
                              abi,
                              functionName: "removeWantedServices",
                              args: [[s.name]],
                            })
                          }
                          onConfirmed={refetch}
                        />
                      </RowAction>
                    )}
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE" action="Add wanted service">
        <div className="flex items-end gap-2">
          <Autocomplete
            className="flex-1"
            value={name}
            onChange={setName}
            options={registered.filter((n) => !items.includes(n))}
            placeholder="Pick or type a registered service…"
          />
          <TxButton
            label="Add wanted"
            icon={<Plus className="h-4 w-4" />}
            disabled={!name.trim()}
            tooltip="Adds a wanted service to the account — sends a transaction to your wallet."
            write={() =>
              writeContractAsync({ address: account, abi, functionName: "addWantedServices", args: [[name.trim()]] })
            }
            onConfirmed={() => {
              setName("");
              refetch();
            }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}

export function ServicesTab({ account }: { account: Address }) {
  const { cmAccountAbi, manager, managerAbi, chainId } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
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
      <WantedServices account={account} abi={abi} hasRole={hasRole} registered={registered} />
    </div>
  );
}
