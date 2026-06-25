import { type Abi, type Address, isAddress } from "viem";
import { AlertTriangle } from "lucide-react";
import { useReadContract } from "wagmi";
import { useActiveContracts } from "../hooks/useActiveContracts";

/**
 * Warns, in the account's left pane, when the opened address isn't actually a
 * CM Account according to the manager (no CMACCOUNT_ROLE) or isn't a valid
 * address at all. The workspace still renders so users can inspect raw state.
 */
export function AccountValidityNotice({ account }: { account: Address }) {
  const { manager, managerAbi, chainId, supported } = useActiveContracts();
  const validAddress = isAddress(account);

  const { data: isCM, isLoading } = useReadContract({
    chainId,
    address: manager,
    abi: managerAbi as Abi,
    functionName: "isCMAccount",
    args: [account],
    query: { enabled: validAddress && !!manager && supported },
  });

  if (!validAddress) {
    return (
      <Notice title="Invalid address">
        This is not a valid EVM address, so the data below may be empty or fail to load.
      </Notice>
    );
  }

  // Only warn once we have a definitive negative answer from the manager.
  if (!supported || isLoading || isCM !== false) return null;

  return (
    <Notice title="Not a CM Account">
      The manager on this network does not recognize this address as a CM Account (it lacks the{" "}
      <code className="font-mono">CMACCOUNT_ROLE</code>). You can still browse, but management actions will likely
      revert.
    </Notice>
  );
}

function Notice({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <aside className="rounded-lg border border-amber-300 bg-amber-50 p-4 text-sm shadow-sm dark:border-amber-800/60 dark:bg-amber-950/40">
      <div className="flex items-center gap-2 font-medium text-amber-800 dark:text-amber-300">
        <AlertTriangle className="h-4 w-4 shrink-0" /> {title}
      </div>
      <p className="mt-1 text-amber-700 dark:text-amber-400/90">{children}</p>
    </aside>
  );
}
