import { useState } from "react";
import { ChevronRight, Coins, KeyRound, Layers, ShieldCheck } from "lucide-react";
import { Link, useNavigate } from "react-router-dom";
import { type Abi, type Address } from "viem";
import { useAccount, useReadContract } from "wagmi";
import { AddressDisplay } from "../components/AddressDisplay";
import { Card } from "../components/Card";
import { GoToAccount } from "../components/GoToAccount";
import { Switch } from "../components/Switch";
import { Tooltip } from "../components/Tooltip";
import { ActivityList } from "../components/activity/ActivityList";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useEcosystemActivity } from "../hooks/useEcosystemActivity";
import { useAccountRolesFor, useAccountStats, useManagerAccounts } from "../hooks/useMyAccounts";
import { APP_CHAINS, explorerUrlFor } from "../config/chains";
import { shortRoleName } from "../lib/format";

// Shared column template so the header and every row line up like a board.
// Every column is a fixed width (not `auto`) so the grid is identical on each
// row — otherwise a roles chip on one row and a dash on another would size the
// column differently and knock the numeric columns out of alignment.
// Mobile collapses to address · roles · chevron; ≥sm adds the three stat columns.
const ROW_GRID =
  "grid items-center gap-x-3 grid-cols-[minmax(0,1fr)_3.25rem_1rem] sm:grid-cols-[minmax(0,1fr)_2.25rem_2.25rem_2.25rem_3.25rem_1rem]";

// Inner links navigate to a specific account tab; stop the event so the row's
// own click/Enter (which opens the account at its first tab) doesn't also fire.
const stopRow = {
  onClick: (e: React.MouseEvent) => e.stopPropagation(),
  onKeyDown: (e: React.KeyboardEvent) => e.stopPropagation(),
};

/** One right-aligned stat number, linking to its tab; faint when zero/loading. */
function Stat({ value, icon: Icon, label, to }: { value?: number; icon: typeof Layers; label: string; to: string }) {
  const dim = value === undefined || value === 0;
  // `hidden sm:flex` sits on the grid cell itself so the column drops cleanly on
  // mobile instead of leaving an empty placeholder behind.
  return (
    <Link to={to} {...stopRow} aria-label={`${label}: ${value ?? 0}`} className="hidden justify-end rounded-sm sm:flex">
      <Tooltip content={label}>
        <span
          className={`inline-flex items-center gap-1 font-mono text-xs tabular-nums transition-colors hover:text-camino-600 dark:hover:text-camino-400 ${
            dim ? "text-tarmac-300 dark:text-tarmac-600" : "text-tarmac-600 dark:text-tarmac-300"
          }`}
        >
          <Icon className="h-3.5 w-3.5 opacity-70" aria-hidden />
          {value ?? "·"}
        </span>
      </Tooltip>
    </Link>
  );
}

/** Compact "you hold N roles" chip linking to the Roles tab; list on hover. */
function RolesIndicator({ roles, to }: { roles: string[]; to: string }) {
  if (roles.length === 0) return <span className="text-tarmac-300 dark:text-tarmac-700">—</span>;
  return (
    <Tooltip
      content={
        <div className="flex flex-col gap-1.5">
          <span className="text-[0.625rem] font-medium uppercase tracking-[0.14em] text-tarmac-300">Your roles</span>
          <ul className="flex flex-col gap-1">
            {roles.map((r) => (
              <li key={r} className="flex items-center gap-1.5 whitespace-nowrap text-xs text-white">
                <span className="h-1.5 w-1.5 rounded-full bg-camino-400" aria-hidden />
                {shortRoleName(r)}
              </li>
            ))}
          </ul>
        </div>
      }
    >
      <Link
        to={to}
        {...stopRow}
        aria-label={`You hold ${roles.length} role${roles.length === 1 ? "" : "s"}`}
        className="inline-flex items-center gap-1 rounded-[3px] border border-camino-300 bg-camino-50 px-1.5 py-0.5 font-mono text-[0.625rem] font-medium tabular-nums text-camino-700 transition-colors hover:border-camino-400 hover:bg-camino-100 dark:border-camino-800 dark:bg-camino-950 dark:text-camino-300 dark:hover:bg-camino-900"
      >
        <ShieldCheck className="h-3 w-3" aria-hidden />
        {roles.length}
      </Link>
    </Tooltip>
  );
}

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
  const { roles, isLoading: rolesLoading } = useAccountRolesFor(account, connected);
  const stats = useAccountStats(account);
  // Only hide once the role read has settled — `roles` is `[]` while pending,
  // which would otherwise drop every row on the first pass when the filter is on.
  if (onlyMine && !rolesLoading && roles.length === 0) return null;

  return (
    <li
      onClick={() => navigate(`/account/${account}`)}
      onKeyDown={(e) => {
        if (e.key === "Enter") navigate(`/account/${account}`);
      }}
      role="link"
      tabIndex={0}
      className={`group cursor-pointer rounded-sm px-2 py-2.5 hover:bg-tarmac-50 dark:hover:bg-tarmac-800/60 ${ROW_GRID}`}
    >
      <AddressDisplay address={account} truncate lead={10} tail={8} className="min-w-0 text-sm" />
      <Stat value={stats.services} icon={Layers} label="Supported services" to={`/account/${account}?tab=services`} />
      <Stat value={stats.tokens} icon={Coins} label="Payment tokens" to={`/account/${account}?tab=tokens`} />
      <Stat value={stats.pubkeys} icon={KeyRound} label="Public keys" to={`/account/${account}?tab=pubkeys`} />
      <div className="flex justify-end">
        <RolesIndicator roles={roles} to={`/account/${account}?tab=roles`} />
      </div>
      <ChevronRight className="h-4 w-4 justify-self-end text-tarmac-300 transition-colors group-hover:text-camino-500 dark:text-tarmac-600 dark:group-hover:text-camino-400" />
    </li>
  );
}

/** Column headers above the account rows, aligned to the same grid. */
function AccountListHeader() {
  return (
    <div className={`px-2 pb-1.5 ${ROW_GRID}`}>
      <span className="eyebrow">Account</span>
      <span className="eyebrow hidden text-right sm:block">Svc</span>
      <span className="eyebrow hidden text-right sm:block">Tok</span>
      <span className="eyebrow hidden text-right sm:block">Key</span>
      <span className="eyebrow text-right">Roles</span>
      <span aria-hidden />
    </div>
  );
}

/** A single labelled cell of the manifest board. */
function ManifestCell({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <div className="flex flex-col gap-1 px-4 py-3">
      <dt className="eyebrow">{label}</dt>
      <dd className="font-display text-2xl font-semibold tabular-nums text-tarmac-900 dark:text-tarmac-50">
        {children}
      </dd>
    </div>
  );
}

/**
 * The manifest — the dashboard hero. A departures-board header strip (network as
 * the line, a status lamp) over a row of labelled board cells. This is the
 * thesis: the state of the ecosystem read at a glance.
 */
function Manifest({
  chainId,
  manager,
  accountCount,
  loadingAccounts,
  paused,
  impl,
}: {
  chainId: number;
  manager?: Address;
  accountCount: number;
  loadingAccounts: boolean;
  paused?: boolean;
  impl?: string;
}) {
  const chain = APP_CHAINS.find((c) => c.id === chainId);
  const active = paused === false;
  return (
    <section className="board overflow-hidden rounded-md">
      <div className="board-grid flex flex-wrap items-end justify-between gap-4 border-b border-tarmac-200/80 bg-paper px-5 py-5 dark:border-tarmac-800 dark:bg-tarmac-950">
        <div>
          <span className="eyebrow">Network manifest</span>
          <h1 className="mt-1 font-display text-3xl font-bold uppercase tracking-tight text-tarmac-900 dark:text-tarmac-50">
            {chain?.name ?? "—"}
          </h1>
          <span className="mt-1 block font-mono text-xs text-tarmac-400">chain · {chainId}</span>
        </div>
        <div
          className={`inline-flex items-center gap-2 rounded-[3px] border px-2.5 py-1 font-mono text-[0.6875rem] uppercase tracking-[0.12em] ${
            paused === undefined
              ? "border-tarmac-200 text-tarmac-400 dark:border-tarmac-700"
              : active
                ? "border-camino-300 bg-camino-50 text-camino-700 dark:border-camino-800 dark:bg-camino-950 dark:text-camino-300"
                : "border-departure-300 bg-departure-50 text-departure-700 dark:border-departure-800 dark:bg-departure-900/40 dark:text-departure-300"
          }`}
        >
          <span
            className={`h-1.5 w-1.5 rounded-full ${
              paused === undefined ? "bg-tarmac-300" : active ? "bg-camino-500" : "bg-departure-500 animate-lamp"
            }`}
            aria-hidden
          />
          {paused === undefined ? "Checking…" : active ? "Accepting accounts" : "Creation paused"}
        </div>
      </div>
      <dl className="grid grid-cols-2 divide-tarmac-200/70 dark:divide-tarmac-800 sm:grid-cols-4 sm:divide-x [&>div]:border-t [&>div]:border-tarmac-200/70 dark:[&>div]:border-tarmac-800 sm:[&>div]:border-t-0">
        <ManifestCell label="Accounts">{loadingAccounts ? "…" : accountCount}</ManifestCell>
        <ManifestCell label="Status">
          <span className="text-base">{paused === undefined ? "…" : active ? "Active" : "Paused"}</span>
        </ManifestCell>
        <div className="flex flex-col gap-1 px-4 py-3">
          <dt className="eyebrow">Manager</dt>
          <dd className="text-sm">{manager ? <AddressDisplay address={manager} /> : "—"}</dd>
        </div>
        <div className="flex flex-col gap-1 px-4 py-3">
          <dt className="eyebrow">Implementation</dt>
          <dd className="text-sm">{impl ? <AddressDisplay address={impl} /> : "—"}</dd>
        </div>
      </dl>
    </section>
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
          className="font-mono text-xs uppercase tracking-[0.1em] text-camino-600 transition-colors hover:text-camino-700 dark:text-camino-400 dark:hover:text-camino-300"
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
    <div className="grid gap-6">
      <Manifest
        chainId={chainId}
        manager={manager}
        accountCount={accounts.length}
        loadingAccounts={isLoading}
        paused={paused as boolean | undefined}
        impl={impl as string | undefined}
      />

      <div className="grid gap-6 lg:grid-cols-5">
        <div className="grid gap-6 lg:col-span-3">
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
              <p className="py-2 font-mono text-sm text-tarmac-400">Loading…</p>
            ) : accounts.length === 0 ? (
              <p className="py-2 font-mono text-sm text-tarmac-400">No accounts found.</p>
            ) : (
              <div className="-mx-2">
                <AccountListHeader />
                <ul className="divide-y divide-tarmac-200/50 dark:divide-tarmac-800/70">
                  {accounts.map((a) => (
                    <AccountRow key={a} account={a} connected={address} onlyMine={onlyMine} />
                  ))}
                </ul>
              </div>
            )}
          </Card>
        </div>

        <div className="grid content-start gap-6 lg:col-span-2">
          <Card title="Open an account">
            <GoToAccount />
          </Card>
          <RecentActivityCard chainId={chainId} />
        </div>
      </div>
    </div>
  );
}
