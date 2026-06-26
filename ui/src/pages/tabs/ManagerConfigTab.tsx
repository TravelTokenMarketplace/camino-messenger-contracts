import { useState } from "react";
import { Pause, Play, Save } from "lucide-react";
import { type Abi, type Address, isAddress } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { inputClass } from "../../components/Input";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";

function AddressSetting({
  title,
  current,
  functionName,
  roleName,
  action,
  isLoading,
  refetch,
}: {
  title: string;
  current: Address | undefined;
  functionName: string;
  roleName: "VERSIONER_ROLE";
  action: string;
  isLoading: boolean;
  refetch: () => void;
}) {
  const { manager, managerAbi } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { hasRole } = useHasRole(manager, abi, roleName);
  const { writeContractAsync } = useWriteContract();
  const [value, setValue] = useState("");
  const trimmed = value.trim();
  const valid = isAddress(trimmed);

  return (
    <Card title={title}>
      <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
        <dt className="text-tarmac-500 dark:text-tarmac-400">Current</dt>
        <dd>{isLoading ? "…" : current ? <AddressDisplay address={current} /> : "—"}</dd>
      </dl>
      <RoleGate hasRole={hasRole} roleName={roleName} action={action}>
        <div className="flex items-end gap-2">
          <input
            className={`flex-1 ${inputClass}`}
            placeholder="0x…"
            value={value}
            onChange={(e) => setValue(e.target.value)}
          />
          <TxButton
            label="Save"
            icon={<Save className="h-4 w-4" />}
            disabled={!valid}
            tooltip={`${action} — sends a transaction to your wallet.`}
            write={() => writeContractAsync({ address: manager!, abi, functionName, args: [trimmed as Address] })}
            onConfirmed={() => {
              setValue("");
              refetch();
            }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}

export function ManagerConfigTab() {
  const { manager, managerAbi, chainId, supported } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { hasRole: canPause } = useHasRole(manager, abi, "PAUSER_ROLE");

  const {
    data: paused,
    isLoading: pausedLoading,
    refetch: refetchPaused,
  } = useReadContract({ chainId, address: manager, abi, functionName: "paused", query: { enabled: supported } });
  const {
    data: impl,
    isLoading: implLoading,
    refetch: refetchImpl,
  } = useReadContract({
    chainId,
    address: manager,
    abi,
    functionName: "getAccountImplementation",
    query: { enabled: supported },
  });
  const {
    data: btoken,
    isLoading: btokenLoading,
    refetch: refetchBtoken,
  } = useReadContract({
    chainId,
    address: manager,
    abi,
    functionName: "getBookingTokenAddress",
    query: { enabled: supported },
  });

  const isPaused = paused === true;

  return (
    <div className="grid gap-4">
      <Card title="Account creation">
        <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-tarmac-500 dark:text-tarmac-400">Status</dt>
          <dd>{pausedLoading ? "…" : isPaused ? "Paused" : "Active"}</dd>
        </dl>
        <RoleGate hasRole={canPause} roleName="PAUSER_ROLE" action={isPaused ? "unpause manager" : "pause manager"}>
          <TxButton
            label={isPaused ? "Unpause" : "Pause"}
            variant={isPaused ? "primary" : "danger"}
            icon={isPaused ? <Play className="h-4 w-4" /> : <Pause className="h-4 w-4" />}
            tooltip={
              isPaused
                ? "Resumes CM account creation — sends a transaction to your wallet."
                : "Pauses CM account creation — sends a transaction to your wallet."
            }
            write={() =>
              writeContractAsync({ address: manager!, abi, functionName: isPaused ? "unpause" : "pause", args: [] })
            }
            onConfirmed={() => refetchPaused()}
          />
        </RoleGate>
      </Card>

      <AddressSetting
        title="Account implementation"
        current={impl as Address | undefined}
        functionName="setAccountImplementation"
        roleName="VERSIONER_ROLE"
        action="set account implementation"
        isLoading={implLoading}
        refetch={() => refetchImpl()}
      />

      <AddressSetting
        title="Booking token address"
        current={btoken as Address | undefined}
        functionName="setBookingTokenAddress"
        roleName="VERSIONER_ROLE"
        action="set booking token address"
        isLoading={btokenLoading}
        refetch={() => refetchBtoken()}
      />
    </div>
  );
}
