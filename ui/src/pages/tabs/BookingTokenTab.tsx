import { useState } from "react";
import { Save } from "lucide-react";
import { type Abi, type Address, isAddress } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { inputClass } from "../../components/Input";
import { RoleGate } from "../../components/RoleGate";
import { RolesPanel } from "../../components/RolesPanel";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { BOOKINGTOKEN_ROLES } from "../../lib/roles";

function formatVersion(v: unknown): string {
  if (!Array.isArray(v) || v.length < 3) return "—";
  return `${v[0]}.${v[1]}.${v[2]}`;
}

export function BookingTokenTab() {
  const { bookingToken, bookingTokenAbi, chainId, supported } = useActiveContracts();
  const abi = bookingTokenAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const enabled = { query: { enabled: supported } };

  const { data: name } = useReadContract({ chainId, address: bookingToken, abi, functionName: "name", ...enabled });
  const { data: symbol } = useReadContract({ chainId, address: bookingToken, abi, functionName: "symbol", ...enabled });
  const { data: version } = useReadContract({
    chainId,
    address: bookingToken,
    abi,
    functionName: "version",
    ...enabled,
  });
  const { data: managerAddr, refetch: refetchManager } = useReadContract({
    chainId,
    address: bookingToken,
    abi,
    functionName: "getManagerAddress",
    ...enabled,
  });
  const { data: minDiff, refetch: refetchMinDiff } = useReadContract({
    chainId,
    address: bookingToken,
    abi,
    functionName: "getMinExpirationTimestampDiff",
    ...enabled,
  });

  const { hasRole: isAdmin } = useHasRole(bookingToken, abi, "DEFAULT_ADMIN_ROLE");
  const { hasRole: canSetMin } = useHasRole(bookingToken, abi, "MIN_EXPIRATION_ADMIN_ROLE");

  const [newManager, setNewManager] = useState("");
  const [newMin, setNewMin] = useState("");

  if (!bookingToken) return <Card title="Booking Token">Connect to a supported network.</Card>;

  return (
    <div className="grid gap-4">
      <Card title="Booking Token">
        <dl className="grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-tarmac-500 dark:text-tarmac-400">Address</dt>
          <dd>
            <AddressDisplay address={bookingToken} />
          </dd>
          <dt className="text-tarmac-500 dark:text-tarmac-400">Name</dt>
          <dd>{(name as string) ?? "—"}</dd>
          <dt className="text-tarmac-500 dark:text-tarmac-400">Symbol</dt>
          <dd>{(symbol as string) ?? "—"}</dd>
          <dt className="text-tarmac-500 dark:text-tarmac-400">Version</dt>
          <dd>{formatVersion(version)}</dd>
        </dl>
      </Card>

      <Card title="Manager address">
        <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-tarmac-500 dark:text-tarmac-400">Current</dt>
          <dd>{managerAddr ? <AddressDisplay address={managerAddr as Address} /> : "—"}</dd>
        </dl>
        <RoleGate hasRole={isAdmin} roleName="DEFAULT_ADMIN_ROLE" action="set manager address">
          <div className="flex items-end gap-2">
            <input
              className={`flex-1 ${inputClass}`}
              placeholder="0x…"
              value={newManager}
              onChange={(e) => setNewManager(e.target.value)}
            />
            <TxButton
              label="Save"
              icon={<Save className="h-4 w-4" />}
              disabled={!isAddress(newManager.trim())}
              tooltip="Sets the manager address on the booking token — sends a transaction to your wallet."
              write={() =>
                writeContractAsync({
                  address: bookingToken,
                  abi,
                  functionName: "setManagerAddress",
                  args: [newManager.trim() as Address],
                })
              }
              onConfirmed={() => {
                setNewManager("");
                refetchManager();
              }}
            />
          </div>
        </RoleGate>
      </Card>

      <Card title="Min expiration timestamp diff">
        <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-tarmac-500 dark:text-tarmac-400">Current</dt>
          <dd>{minDiff !== undefined ? `${minDiff} seconds` : "—"}</dd>
        </dl>
        <RoleGate hasRole={canSetMin} roleName="MIN_EXPIRATION_ADMIN_ROLE" action="set min expiration diff">
          <div className="flex items-end gap-2">
            <input
              className={`w-40 ${inputClass}`}
              type="number"
              min="0"
              step="1"
              placeholder="seconds"
              value={newMin}
              onChange={(e) => setNewMin(e.target.value)}
            />
            <TxButton
              label="Save"
              icon={<Save className="h-4 w-4" />}
              disabled={!/^\d+$/.test(newMin.trim())}
              tooltip="Sets the minimum reservation expiration difference — sends a transaction to your wallet."
              write={() =>
                writeContractAsync({
                  address: bookingToken,
                  abi,
                  functionName: "setMinExpirationTimestampDiff",
                  args: [BigInt(newMin.trim())],
                })
              }
              onConfirmed={() => {
                setNewMin("");
                refetchMinDiff();
              }}
            />
          </div>
        </RoleGate>
      </Card>

      <Card title="Booking Token Roles">
        <RolesPanel address={bookingToken} abi={abi} roles={BOOKINGTOKEN_ROLES} enumerable={false} />
      </Card>
    </div>
  );
}
