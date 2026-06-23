import { useState } from "react";
import { Link } from "react-router-dom";
import { type Abi, type Address } from "viem";
import { useAccount, useReadContract } from "wagmi";
import { AddressDisplay } from "../components/AddressDisplay";
import { Card } from "../components/Card";
import { RoleBadge } from "../components/RoleBadge";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useAccountRolesFor, useManagerAccounts } from "../hooks/useMyAccounts";

function AccountRow({ account, connected, onlyMine }: { account: Address; connected: Address | undefined; onlyMine: boolean }) {
  const roles = useAccountRolesFor(account, connected);
  if (onlyMine && roles.length === 0) return null;
  return (
    <li className="flex flex-wrap items-center justify-between gap-2 py-2">
      <span className="flex items-center gap-2">
        <AddressDisplay address={account} truncate />
        <Link className="text-xs text-indigo-600 underline dark:text-indigo-400" to={`/account/${account}`}>Open</Link>
      </span>
      <span className="flex flex-wrap gap-1">
        {roles.map((r) => (
          <RoleBadge key={r} role={r} />
        ))}
      </span>
    </li>
  );
}

export function Dashboard() {
  const { manager, managerAbi, supported } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { address } = useAccount();
  const { accounts, isLoading } = useManagerAccounts();
  const [onlyMine, setOnlyMine] = useState(false);
  const { data: paused } = useReadContract({ address: manager, abi, functionName: "paused" });
  const { data: impl } = useReadContract({ address: manager, abi, functionName: "getAccountImplementation" });

  if (!supported) return <Card title="Dashboard">Connect to a supported network.</Card>;

  return (
    <div className="grid gap-4">
      <Card title="Network status">
        <dl className="grid grid-cols-2 gap-2 text-sm">
          <dt className="text-gray-500 dark:text-gray-400">Manager</dt><dd>{manager && <AddressDisplay address={manager} />}</dd>
          <dt className="text-gray-500 dark:text-gray-400">Paused</dt><dd>{paused ? "Yes" : "No"}</dd>
          <dt className="text-gray-500 dark:text-gray-400">Account implementation</dt><dd>{impl ? <AddressDisplay address={impl as string} /> : "—"}</dd>
        </dl>
      </Card>
      <Card title="CM Accounts">
        <label className="mb-3 flex items-center gap-2 text-sm">
          <input type="checkbox" checked={onlyMine} disabled={!address} onChange={(e) => setOnlyMine(e.target.checked)} />
          Only accounts where I hold a role
        </label>
        {isLoading ? <p>Loading…</p> : (
          <ul className="divide-y">
            {accounts.length === 0 && <li className="py-2 text-sm text-gray-400">No accounts found.</li>}
            {accounts.map((a) => (
              <AccountRow key={a} account={a} connected={address} onlyMine={onlyMine} />
            ))}
          </ul>
        )}
      </Card>
    </div>
  );
}
