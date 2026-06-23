import { useState } from "react";
import { PlusCircle } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { type Abi, type Address } from "viem";
import { useAccount, usePublicClient, useWriteContract } from "wagmi";
import { Card } from "../components/Card";
import { TxButton } from "../components/TxButton";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { findCreatedAccount } from "../lib/receipt";

export function CreateAccount() {
  const { address } = useAccount();
  const { manager, managerAbi, cmAccountAbi, supported } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const client = usePublicClient();
  const navigate = useNavigate();
  const [admin, setAdmin] = useState("");
  const [upgrader, setUpgrader] = useState("");

  const adminVal = (admin || address || "") as Address;
  const upgraderVal = (upgrader || address || "") as Address;

  async function write() {
    const hash = await writeContractAsync({
      address: manager!,
      abi: managerAbi as Abi,
      functionName: "createCMAccount",
      args: [adminVal, upgraderVal],
    });
    const receipt = await client!.waitForTransactionReceipt({ hash });
    const created = findCreatedAccount(receipt.logs, cmAccountAbi as Abi);
    if (created) navigate(`/account/${created}`);
    return hash;
  }

  if (!supported) return <Card title="Create CM Account">Connect to a supported network.</Card>;

  return (
    <Card title="Create CM Account">
      <div className="grid max-w-md gap-3">
        <label className="text-sm">Admin address
          <input className="mt-1 w-full rounded border px-2 py-1" placeholder={address} value={admin} onChange={(e) => setAdmin(e.target.value)} />
        </label>
        <label className="text-sm">Upgrader address
          <input className="mt-1 w-full rounded border px-2 py-1" placeholder={address} value={upgrader} onChange={(e) => setUpgrader(e.target.value)} />
        </label>
        <TxButton label="Create account" icon={<PlusCircle className="h-4 w-4" />} disabled={!address} write={write} />
      </div>
    </Card>
  );
}
