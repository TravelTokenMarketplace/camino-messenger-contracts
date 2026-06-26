import { useState } from "react";
import { ChevronRight } from "lucide-react";
import { Link, useNavigate } from "react-router-dom";
import { type Abi, type Address } from "viem";
import { useAccount, useReadContract } from "wagmi";
import { AddressDisplay } from "../components/AddressDisplay";
import { Card } from "../components/Card";
import { GoToAccount } from "../components/GoToAccount";
import { Switch } from "../components/Switch";
import { RoleBadge } from "../components/RoleBadge";
import { ActivityList } from "../components/activity/ActivityList";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useEcosystemActivity } from "../hooks/useEcosystemActivity";
import { useAccountRolesFor, useManagerAccounts } from "../hooks/useMyAccounts";
import { explorerUrlFor } from "../config/chains";
import { shortRoleName } from "../lib/format";

function AccountRow({
  account,
  connected,
  onlyMine,
}: {
  account: Address;
  connected: Address | undefined;
  onlyMine: boolean;
}) {
  const navigate = useNavigate();
  const roles = useAccountRolesFor(account, connected);
  if (onlyMine && roles.length === 0) return null;

  return (
    <li
      onClick={() => navigate(`/account/${account}`)}
      onKeyDown={(e) => {
        if (e.key === "Enter") navigate(`/account/${account}`);
      }}
      role="link"
      tabIndex={0}
      className="group flex cursor-pointer items-center gap-3 rounded-md px-2 py-2.5 hover:bg-gray-50 dark:hover:bg-gray-700/50"
    >
      <AddressDisplay address={account} className="text-sm" />
      <span className="ml-auto flex items-center gap-1.5">
        {roles.slice(0, 3).map((r) => (
          <RoleBadge key={r} role={shortRoleName(r)} />
        ))}
        {roles.length > 3 && (
          <span className="text-xs text-gray-400" title={roles.map(shortRoleName).join(", ")}>
            +{roles.length - 3}
          </span>
        )}
        <ChevronRight className="h-4 w-4 text-gray-300 transition-colors group-hover:text-gray-500 dark:text-gray-600 dark:group-hover:text-gray-300" />
      </span>
    </li>
  );
}

function RecentActivityCard({ chainId }: { chainId: number }) {
  const activity = useEcosystemActivity();
  return (
    <Card
      title="Recent activity"
      actions={
        <Link
          to="/activity"
          className="text-sm text-indigo-600 transition-colors hover:text-indigo-800 dark:text-indigo-400 dark:hover:text-indigo-300"
        >
          View all →
        </Link>
      }
    >
      <ActivityList
        limit={5}
        explorerUrl={explorerUrlFor(chainId)}
        events={activity.events}
        isLoading={activity.isLoading}
        error={activity.error}
      />
    </Card>
  );
}

export function Dashboard() {
  const { manager, managerAbi, supported, chainId } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { address } = useAccount();
  const { accounts, isLoading } = useManagerAccounts();
  const [onlyMine, setOnlyMine] = useState(false);
  const { data: paused } = useReadContract({ chainId, address: manager, abi, functionName: "paused" });
  const { data: impl } = useReadContract({ chainId, address: manager, abi, functionName: "getAccountImplementation" });

  if (!supported) return <Card title="Dashboard">Connect to a supported network.</Card>;

  return (
    <div className="grid gap-4">
      <Card title="Open an account">
        <GoToAccount />
      </Card>

      <Card title="Network status">
        <dl className="grid grid-cols-2 gap-2 text-sm">
          <dt className="text-gray-500 dark:text-gray-400">Manager</dt>
          <dd>{manager && <AddressDisplay address={manager} />}</dd>
          <dt className="text-gray-500 dark:text-gray-400">Paused</dt>
          <dd>{paused ? "Yes" : "No"}</dd>
          <dt className="text-gray-500 dark:text-gray-400">Account implementation</dt>
          <dd>{impl ? <AddressDisplay address={impl as string} /> : "—"}</dd>
        </dl>
      </Card>

      <RecentActivityCard chainId={chainId} />

      <Card
        title="CM Accounts"
        actions={
          <Switch
            checked={onlyMine}
            disabled={!address}
            onChange={setOnlyMine}
            label="Only accounts where I hold a role"
          />
        }
      >
        {isLoading ? (
          <p className="py-2 text-sm text-gray-400">Loading…</p>
        ) : (
          <ul className="-mx-2">
            {accounts.length === 0 && <li className="px-2 py-2 text-sm text-gray-400">No accounts found.</li>}
            {accounts.map((a) => (
              <AccountRow key={a} account={a} connected={address} onlyMine={onlyMine} />
            ))}
          </ul>
        )}
      </Card>
    </div>
  );
}
