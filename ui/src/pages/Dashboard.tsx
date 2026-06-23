import { Link } from "react-router-dom";
import { useReadContract } from "wagmi";
import { Card } from "../components/Card";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useMyAccounts } from "../hooks/useMyAccounts";
import { shortAddress } from "../lib/format";

export function Dashboard() {
  const { manager, managerAbi, supported } = useActiveContracts();
  const { accounts, isLoading } = useMyAccounts();
  const { data: paused } = useReadContract({ address: manager, abi: managerAbi, functionName: "paused" });
  const { data: impl } = useReadContract({ address: manager, abi: managerAbi, functionName: "getAccountImplementation" });

  if (!supported) return <Card title="Dashboard">Connect to a supported network.</Card>;

  return (
    <div className="grid gap-4">
      <Card title="Network status">
        <dl className="grid grid-cols-2 gap-2 text-sm">
          <dt className="text-gray-500">Manager</dt><dd className="break-all">{manager}</dd>
          <dt className="text-gray-500">Paused</dt><dd>{paused ? "Yes" : "No"}</dd>
          <dt className="text-gray-500">Account implementation</dt><dd className="break-all">{impl as string}</dd>
        </dl>
      </Card>
      <Card title="Created accounts">
        {isLoading ? "Loading…" : (
          <ul className="divide-y">
            {accounts.map((a) => (
              <li key={a} className="py-2">
                <Link className="text-indigo-600 underline" to={`/account/${a}`}>{shortAddress(a)}</Link>
              </li>
            ))}
          </ul>
        )}
      </Card>
    </div>
  );
}
