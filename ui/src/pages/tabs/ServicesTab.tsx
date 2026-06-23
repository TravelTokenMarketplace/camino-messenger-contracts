import { type Address } from "viem";
import { Card } from "../../components/Card";

export function ServicesTab({ account }: { account: Address }) {
  return <Card title="Services">Coming soon for {account}.</Card>;
}
