import { type Address } from "viem";
import { useBalance } from "wagmi";
import { Card } from "../../components/Card";

export function OverviewTab({ account }: { account: Address }) {
  const { data } = useBalance({ address: account });
  return (
    <Card title="Overview">
      <dl className="grid grid-cols-2 gap-2 text-sm">
        <dt className="text-gray-500">Address</dt><dd className="break-all">{account}</dd>
        <dt className="text-gray-500">Native balance</dt><dd>{data ? `${data.formatted} ${data.symbol}` : "—"}</dd>
      </dl>
    </Card>
  );
}
