import { type Address } from "viem";
import { Card } from "../../components/Card";

export function PubkeysTab({ account }: { account: Address }) {
  return <Card title="Pubkeys">Coming soon for {account}.</Card>;
}
