import { Card } from "../components/Card";
import { RefreshButton } from "../components/RefreshButton";
import { ActivityList } from "../components/activity/ActivityList";
import { useEcosystemActivity } from "../hooks/useEcosystemActivity";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { explorerUrlFor } from "../config/chains";

export function Activity() {
  const { supported, chainId } = useActiveContracts();
  const activity = useEcosystemActivity();

  if (!supported) return <Card title="Activity">Connect to a supported network.</Card>;

  return (
    <Card title="Ecosystem activity" actions={<RefreshButton />}>
      <ActivityList
        showFilters
        explorerUrl={explorerUrlFor(chainId)}
        events={activity.events}
        isLoading={activity.isLoading}
        error={activity.error}
        hasNextPage={activity.hasNextPage}
        isFetchingNextPage={activity.isFetchingNextPage}
        onLoadOlder={activity.fetchNextPage}
        oldestBlockLoaded={activity.oldestBlockLoaded}
      />
    </Card>
  );
}
