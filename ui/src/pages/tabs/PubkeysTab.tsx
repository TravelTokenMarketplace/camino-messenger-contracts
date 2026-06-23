import { useState } from "react";
import { type Abi, type Address, type Hex } from "viem";
import { useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

export function PubkeysTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { items, isLoading, refetch } = useContractList(account, abi, "getPublicKeysAddresses");
  const { hasRole } = useHasRole(account, abi, "SERVICE_ADMIN_ROLE");
  const [addr, setAddr] = useState("");
  const [data, setData] = useState("");

  return (
    <Card title="Encryption Public Keys">
      {isLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y">
          {items.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {items.map((k) => (
            <li key={k} className="flex items-center justify-between py-2">
              <span className="font-mono text-sm">{k}</span>
              <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE">
                <TxButton label="Remove" write={() => writeContractAsync({ address: account, abi, functionName: "removePublicKey", args: [k as Address] })} onConfirmed={refetch} />
              </RoleGate>
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE">
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Address 0x…" value={addr} onChange={(e) => setAddr(e.target.value)} />
          <input className="flex-1 rounded border px-2 py-1" placeholder="Pubkey data (hex 0x…)" value={data} onChange={(e) => setData(e.target.value)} />
          <TxButton label="Add" disabled={!addr || !data} write={() => writeContractAsync({ address: account, abi, functionName: "addPublicKey", args: [addr as Address, data as Hex] })} onConfirmed={() => { setAddr(""); setData(""); refetch(); }} />
        </div>
      </RoleGate>
    </Card>
  );
}
