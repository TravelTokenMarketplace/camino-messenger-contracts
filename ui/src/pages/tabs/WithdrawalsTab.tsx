import { type Address } from "viem";
import { Card } from "../../components/Card";

export function WithdrawalsTab({ account }: { account: Address }) {
  return <Card title="Withdrawals">Coming soon for {account}.</Card>;
}
