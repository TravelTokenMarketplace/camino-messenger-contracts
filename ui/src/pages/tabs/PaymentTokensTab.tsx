import { type Address } from "viem";
import { Card } from "../../components/Card";

export function PaymentTokensTab({ account }: { account: Address }) {
  return <Card title="Payment Tokens">Coming soon for {account}.</Card>;
}
