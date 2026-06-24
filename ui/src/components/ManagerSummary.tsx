import { type Abi } from "viem";
import { useReadContract } from "wagmi";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { AddressDisplay } from "./AddressDisplay";

export function ManagerSummary() {
  const { manager, bookingToken, managerAbi, chainId, supported } = useActiveContracts();
  const { data: paused, isLoading: pausedLoading } = useReadContract({ chainId, address: manager, abi: managerAbi as Abi, functionName: "paused", query: { enabled: supported } });

  return (
    <aside className="h-fit rounded-lg border border-gray-200 bg-white p-4 shadow-sm dark:border-gray-800 dark:bg-gray-800">
      <h2 className="mb-3 text-xs font-semibold uppercase tracking-wide text-gray-400 dark:text-gray-500">Ecosystem</h2>
      <dl className="grid grid-cols-1 gap-3 text-sm">
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Manager</dt>
          <dd>{manager ? <AddressDisplay address={manager} /> : "—"}</dd>
        </div>
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Booking token</dt>
          <dd>{bookingToken ? <AddressDisplay address={bookingToken} /> : "—"}</dd>
        </div>
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Account creation</dt>
          <dd className="text-gray-900 dark:text-gray-100">{pausedLoading ? "…" : paused ? "Paused" : "Active"}</dd>
        </div>
      </dl>
    </aside>
  );
}
