import { useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { type Abi, type Address, type Hex } from "viem";
import { useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { Input } from "../../components/Input";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
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
      {isLoading ? (
        <p>Loading…</p>
      ) : (
        <ul className="mb-4 divide-y">
          {items.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {items.map((k) => (
            <li key={k} className="group flex items-center justify-between gap-3 py-2">
              <AddressDisplay address={k} className="text-sm" />
              {hasRole && (
                <RowAction>
                  <TxButton
                    label="Remove"
                    variant="danger"
                    icon={<Trash2 className="h-4 w-4" />}
                    write={() =>
                      writeContractAsync({
                        address: account,
                        abi,
                        functionName: "removePublicKey",
                        args: [k as Address],
                      })
                    }
                    onConfirmed={refetch}
                  />
                </RowAction>
              )}
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_ADMIN_ROLE" action="Add public key">
        <div className="flex items-end gap-2">
          <Input
            className="flex-1 font-mono"
            placeholder="Address 0x…"
            value={addr}
            onChange={(e) => setAddr(e.target.value)}
          />
          <Input
            className="flex-1 font-mono"
            placeholder="Pubkey data (hex 0x…)"
            value={data}
            onChange={(e) => setData(e.target.value)}
          />
          <TxButton
            label="Add"
            icon={<Plus className="h-4 w-4" />}
            disabled={!addr || !data}
            write={() =>
              writeContractAsync({
                address: account,
                abi,
                functionName: "addPublicKey",
                args: [addr as Address, data as Hex],
              })
            }
            onConfirmed={() => {
              setAddr("");
              setData("");
              refetch();
            }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}
