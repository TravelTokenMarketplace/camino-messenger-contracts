import { useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { type Abi, type Address, parseEther } from "viem";
import { useBalance, useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { useHasRole } from "../../hooks/useHasRole";

function BotBalance({ bot }: { bot: Address }) {
  const { chainId } = useActiveContracts();
  const { data } = useBalance({ address: bot, chainId });
  if (!data) return <span className="text-xs text-gray-400">…</span>;
  const isZero = data.value === 0n;
  return (
    <span
      className={`rounded px-2 py-0.5 text-xs font-medium ${
        isZero
          ? "bg-red-100 text-red-700 dark:bg-red-950 dark:text-red-300"
          : "bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300"
      }`}
      title={isZero ? "Bot has no funds for transaction fees" : undefined}
    >
      {data.formatted} {data.symbol}
    </span>
  );
}

function BotRow({ account, bot, abi, hasRole, onChanged }: { account: Address; bot: Address; abi: Abi; hasRole: boolean; onChanged: () => void }) {
  const { writeContractAsync } = useWriteContract();
  return (
    <li className="group flex flex-wrap items-center justify-between gap-2 py-2">
      <span className="flex items-center gap-3">
        <AddressDisplay address={bot} className="text-sm" />
        <BotBalance bot={bot} />
      </span>
      {hasRole && (
        <RowAction>
          <TxButton label="Remove" variant="danger" icon={<Trash2 className="h-4 w-4" />} write={() => writeContractAsync({ address: account, abi, functionName: "removeMessengerBot", args: [bot] })} onConfirmed={onChanged} />
        </RowAction>
      )}
    </li>
  );
}

export function BotsTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, abi, "MESSENGER_BOT_ROLE");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, abi, "BOT_ADMIN_ROLE");
  const [bot, setBot] = useState("");
  const [gas, setGas] = useState("0");

  return (
    <Card title="Messenger Bots">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">A bot is an address granted MESSENGER_BOT_ROLE (plus booking/gas roles). It needs native funds to pay transaction fees.</p>
      {isLoading || roleLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y dark:divide-gray-700">
          {members.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {members.map((b) => (
            <BotRow key={b} account={account} bot={b as Address} abi={abi} hasRole={hasRole} onChanged={refetch} />
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="BOT_ADMIN_ROLE" action="Add bot">
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Bot address 0x…" value={bot} onChange={(e) => setBot(e.target.value)} />
          <input className="w-32 rounded border px-2 py-1" placeholder="Gas money (CAM)" value={gas} onChange={(e) => setGas(e.target.value)} />
          <TxButton label="Add bot" icon={<Plus className="h-4 w-4" />} disabled={!bot} write={() => writeContractAsync({ address: account, abi, functionName: "addMessengerBot", args: [bot as Address, parseEther(gas || "0")] })} onConfirmed={() => { setBot(""); setGas("0"); refetch(); }} />
        </div>
      </RoleGate>
    </Card>
  );
}
