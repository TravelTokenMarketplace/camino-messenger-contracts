import { useState } from "react";
import { PlusCircle } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { type Abi, type Address, type TransactionReceipt, parseEther } from "viem";
import { useAccount, useWriteContract } from "wagmi";
import { APP_CHAINS } from "../config/chains";
import { Card } from "../components/Card";
import { Input } from "../components/Input";
import { TxButton } from "../components/TxButton";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { findCreatedAccount } from "../lib/receipt";

export function CreateAccount() {
  const { address } = useAccount();
  const { manager, managerAbi, cmAccountAbi, supported, chainId } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const navigate = useNavigate();
  const [admin, setAdmin] = useState("");
  const [upgrader, setUpgrader] = useState("");
  const [prefund, setPrefund] = useState("");

  const adminVal = (admin || address || "") as Address;
  const upgraderVal = (upgrader || address || "") as Address;
  const symbol = APP_CHAINS.find((c) => c.id === chainId)?.nativeCurrency.symbol ?? "";

  // Reject obviously-malformed amounts before we try to send the transaction.
  let amountError: string | undefined;
  let value = 0n;
  if (prefund.trim()) {
    try {
      value = parseEther(prefund.trim());
    } catch {
      amountError = "Enter a valid amount.";
    }
  }

  // Return the hash as soon as it's submitted so TxProvider can register the
  // pending entry and own the single receipt wait. Navigation to the new
  // account happens in onConfirmed, which receives the mined receipt.
  function write() {
    return writeContractAsync({
      address: manager!,
      abi: managerAbi as Abi,
      functionName: "createCMAccount",
      args: [adminVal, upgraderVal],
      value,
    });
  }

  function onConfirmed(receipt: TransactionReceipt) {
    const created = findCreatedAccount(receipt.logs, cmAccountAbi as Abi);
    if (created) navigate(`/account/${created}`);
  }

  if (!supported) return <Card title="Create CM Account">Connect to a supported network.</Card>;

  return (
    <Card title="Create CM Account">
      <div className="grid max-w-md gap-3">
        <label className="text-sm">Admin address
          <Input className="mt-1 w-full font-mono" placeholder={address} value={admin} onChange={(e) => setAdmin(e.target.value)} />
        </label>
        <label className="text-sm">Upgrader address
          <Input className="mt-1 w-full font-mono" placeholder={address} value={upgrader} onChange={(e) => setUpgrader(e.target.value)} />
        </label>
        <label className="text-sm">
          Initial funding <span className="text-gray-400">(optional)</span>
          <Input
            className="mt-1 w-full"
            inputMode="decimal"
            placeholder={`0.0 ${symbol}`.trim()}
            value={prefund}
            onChange={(e) => setPrefund(e.target.value)}
          />
          <span className="mt-1 block text-xs text-gray-500 dark:text-gray-400">
            Sent to the new account on creation{symbol && ` (in ${symbol})`}.
          </span>
          {amountError && <span className="mt-1 block text-xs text-red-600">{amountError}</span>}
        </label>
        <TxButton label="Create account" icon={<PlusCircle className="h-4 w-4" />} disabled={!address || !!amountError} write={write} onConfirmed={onConfirmed} />
      </div>
    </Card>
  );
}
