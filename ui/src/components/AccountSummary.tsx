import { AlertTriangle } from "lucide-react";
import { type Address } from "viem";
import { useAccount, useBalance } from "wagmi";
import { APP_CHAINS } from "../config/chains";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useErc20Balances } from "../hooks/useErc20Balances";
import { useAccountRolesFor } from "../hooks/useMyAccounts";
import { formatAmount } from "../lib/format";
import { AddressDisplay } from "./AddressDisplay";
import { RoleBadge } from "./RoleBadge";
import { TokenDisplay } from "./TokenDisplay";
import { Tooltip } from "./Tooltip";

export function AccountSummary({ account }: { account: Address }) {
  const { chainId } = useActiveContracts();
  const { address } = useAccount();
  const { data: bal } = useBalance({ address: account, chainId });
  const { tokens } = useErc20Balances(account);
  const nativeZero = bal ? bal.value === 0n : false;
  const { roles } = useAccountRolesFor(account, address);
  const chainName = APP_CHAINS.find((c) => c.id === chainId)?.name;

  return (
    <aside className="h-fit rounded-lg border border-tarmac-200 bg-paper-raised p-4 shadow-sm dark:border-tarmac-800 dark:bg-tarmac-800">
      <h2 className="mb-3 text-xs font-semibold uppercase tracking-wide text-tarmac-400 dark:text-tarmac-500">
        Active account
      </h2>

      <div className="mb-4 text-sm">
        <AddressDisplay address={account} />
      </div>

      <dl className="grid grid-cols-1 gap-3 text-sm">
        <div>
          <dt className="text-tarmac-500 dark:text-tarmac-400">Network</dt>
          <dd>{chainName ?? "—"}</dd>
        </div>
        <div>
          <dt className="text-tarmac-500 dark:text-tarmac-400">Native balance</dt>
          <dd className="font-num tabular-nums">
            {bal ? (
              <Tooltip content={`${bal.formatted} ${bal.symbol}`}>
                <span>
                  {formatAmount(bal.formatted).display} {bal.symbol}
                </span>
              </Tooltip>
            ) : (
              "—"
            )}
          </dd>
          {nativeZero && (
            <p className="mt-1 flex items-start gap-1 text-xs text-amber-700 dark:text-amber-300">
              <AlertTriangle className="mt-0.5 h-3 w-3 shrink-0" />
              No {bal?.symbol} for gas — this account can't pay gas to buy booking tokens.
            </p>
          )}
        </div>
        {tokens.length > 0 && (
          <div>
            <dt className="mb-1 text-tarmac-500 dark:text-tarmac-400">Token balances</dt>
            <dd className="flex flex-col gap-2">
              {tokens.map((t) => (
                <div key={t.address}>
                  <div className="flex items-center justify-between gap-2">
                    <TokenDisplay address={t.address} symbol={t.symbol} name={t.name} className="min-w-0 text-sm" />
                    <Tooltip content={`${t.formatted} ${t.symbol}`}>
                      <span
                        className={`font-num tabular-nums whitespace-nowrap ${
                          t.isZero ? "text-amber-700 dark:text-amber-300" : ""
                        }`}
                      >
                        {formatAmount(t.formatted).display} {t.symbol}
                      </span>
                    </Tooltip>
                  </div>
                  {t.isZero && (
                    <p className="mt-0.5 flex items-start gap-1 text-xs text-amber-700 dark:text-amber-300">
                      <AlertTriangle className="mt-0.5 h-3 w-3 shrink-0" />0 {t.symbol} — can't buy booking tokens paid
                      in {t.symbol}.
                    </p>
                  )}
                </div>
              ))}
            </dd>
          </div>
        )}
        <div>
          <dt className="mb-1 text-tarmac-500 dark:text-tarmac-400">Your roles</dt>
          <dd className="flex flex-wrap gap-1">
            {address ? (
              roles.length ? (
                roles.map((r) => <RoleBadge key={r} role={r} />)
              ) : (
                <span className="text-tarmac-400 dark:text-tarmac-500">None</span>
              )
            ) : (
              <span className="text-tarmac-400 dark:text-tarmac-500">Connect wallet</span>
            )}
          </dd>
        </div>
      </dl>
    </aside>
  );
}
