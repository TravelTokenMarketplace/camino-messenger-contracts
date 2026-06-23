import { type Address } from "viem";
import { useBalance } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { useActiveContracts } from "../../hooks/useActiveContracts";

export function OverviewTab({ account }: { account: Address }) {
  const { chainId } = useActiveContracts();
  const { data } = useBalance({ address: account, chainId });
  return (
    <Card title="Overview">
      <dl className="grid grid-cols-2 gap-2 text-sm">
        <dt className="text-gray-500 dark:text-gray-400">Address</dt><dd><AddressDisplay address={account} /></dd>
        <dt className="text-gray-500 dark:text-gray-400">Native balance</dt><dd>{data ? `${data.formatted} ${data.symbol}` : "—"}</dd>
      </dl>
    </Card>
  );
}
