import { type Address } from "viem";
import { Card } from "../../components/Card";

export function BotsTab({ account }: { account: Address }) {
  return <Card title="Bots">Coming soon for {account}.</Card>;
}
