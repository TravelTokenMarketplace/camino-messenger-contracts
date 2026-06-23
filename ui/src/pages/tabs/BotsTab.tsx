import { useState } from "react";
import { type Abi, type Address, parseEther } from "viem";
import { useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { useHasRole } from "../../hooks/useHasRole";

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
      <p className="mb-3 text-xs text-gray-500">A bot is an address granted MESSENGER_BOT_ROLE (plus booking/gas roles) on this account.</p>
      {isLoading || roleLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y">
          {members.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {members.map((b) => (
            <li key={b} className="flex items-center justify-between py-2">
              <span className="font-mono text-sm">{b}</span>
              <RoleGate hasRole={hasRole} roleName="BOT_ADMIN_ROLE">
                <TxButton label="Remove" write={() => writeContractAsync({ address: account, abi, functionName: "removeMessengerBot", args: [b as Address] })} onConfirmed={refetch} />
              </RoleGate>
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="BOT_ADMIN_ROLE">
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Bot address 0x…" value={bot} onChange={(e) => setBot(e.target.value)} />
          <input className="w-32 rounded border px-2 py-1" placeholder="Gas money (CAM)" value={gas} onChange={(e) => setGas(e.target.value)} />
          <TxButton label="Add bot" disabled={!bot} write={() => writeContractAsync({ address: account, abi, functionName: "addMessengerBot", args: [bot as Address, parseEther(gas || "0")] })} onConfirmed={() => { setBot(""); setGas("0"); refetch(); }} />
        </div>
      </RoleGate>
    </Card>
  );
}
