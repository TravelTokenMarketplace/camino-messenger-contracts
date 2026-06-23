import { useState } from "react";
import { ArrowUpFromLine } from "lucide-react";
import { type Abi, type Address, parseEther } from "viem";
import { useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { Input } from "../../components/Input";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";

export function WithdrawalsTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { hasRole } = useHasRole(account, abi, "WITHDRAWER_ROLE");
  const [recipient, setRecipient] = useState("");
  const [amount, setAmount] = useState("0");

  return (
    <Card title="Withdraw native funds">
      <RoleGate hasRole={hasRole} roleName="WITHDRAWER_ROLE" action="Withdraw">
        <div className="flex items-end gap-2">
          <Input className="flex-1 font-mono" placeholder="Recipient 0x…" value={recipient} onChange={(e) => setRecipient(e.target.value)} />
          <Input className="w-32" placeholder="Amount" value={amount} onChange={(e) => setAmount(e.target.value)} />
          <TxButton label="Withdraw" icon={<ArrowUpFromLine className="h-4 w-4" />} disabled={!recipient} write={() => writeContractAsync({ address: account, abi, functionName: "withdraw", args: [recipient as Address, parseEther(amount || "0")] })} onConfirmed={() => { setRecipient(""); setAmount("0"); }} />
        </div>
      </RoleGate>
    </Card>
  );
}
