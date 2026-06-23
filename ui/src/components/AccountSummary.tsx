import { type Address } from "viem";
import { useAccount, useBalance } from "wagmi";
import { APP_CHAINS } from "../config/chains";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useAccountRolesFor } from "../hooks/useMyAccounts";
import { AddressDisplay } from "./AddressDisplay";
import { RoleBadge } from "./RoleBadge";

export function AccountSummary({ account }: { account: Address }) {
  const { chainId } = useActiveContracts();
  const { address } = useAccount();
  const { data: bal } = useBalance({ address: account, chainId });
  const roles = useAccountRolesFor(account, address);
  const chainName = APP_CHAINS.find((c) => c.id === chainId)?.name;

  return (
    <aside className="h-fit rounded-lg border border-gray-200 bg-white p-4 shadow-sm dark:border-gray-800 dark:bg-gray-800">
      <h2 className="mb-3 text-xs font-semibold uppercase tracking-wide text-gray-400">Active account</h2>

      <div className="mb-4 text-sm">
        <AddressDisplay address={account} />
      </div>

      <dl className="grid grid-cols-1 gap-3 text-sm">
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Network</dt>
          <dd>{chainName ?? "—"}</dd>
        </div>
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Native balance</dt>
          <dd>{bal ? `${bal.formatted} ${bal.symbol}` : "—"}</dd>
        </div>
        <div>
          <dt className="mb-1 text-gray-500 dark:text-gray-400">Your roles</dt>
          <dd className="flex flex-wrap gap-1">
            {address ? (roles.length ? roles.map((r) => <RoleBadge key={r} role={r} />) : <span className="text-gray-400">None</span>) : <span className="text-gray-400">Connect wallet</span>}
          </dd>
        </div>
      </dl>
    </aside>
  );
}
