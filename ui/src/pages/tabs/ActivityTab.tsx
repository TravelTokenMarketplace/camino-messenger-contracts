import { type Address } from "viem";
import { ActivityList } from "../../components/activity/ActivityList";
import { useAccountActivity } from "../../hooks/useAccountActivity";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { explorerUrlFor } from "../../config/chains";

export function ActivityTab({ account }: { account: Address }) {
  const { chainId } = useActiveContracts();
  const activity = useAccountActivity(account);

  return (
    <ActivityList
      explorerUrl={explorerUrlFor(chainId)}
      events={activity.events}
      isLoading={activity.isLoading}
      error={activity.error}
      hasNextPage={activity.hasNextPage}
      isFetchingNextPage={activity.isFetchingNextPage}
      onLoadOlder={activity.loadOlder}
      oldestBlockLoaded={activity.oldestBlockLoaded}
      emptyHint="No activity for this account in the last 10,000 blocks."
    />
  );
}
