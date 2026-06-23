import { type Abi, type Address } from "viem";
import { useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { ListManager } from "../../components/ListManager";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

export function PaymentTokensTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { items, isLoading, refetch } = useContractList(account, abi, "getSupportedTokens");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, abi, "SERVICE_ADMIN_ROLE");

  return (
    <ListManager
      title="Payment Tokens"
      items={items}
      isLoading={isLoading || roleLoading}
      roleName="SERVICE_ADMIN_ROLE"
      hasRole={hasRole}
      addLabel="Add token"
      addPlaceholder="Token address 0x…"
      onAdd={(v) => writeContractAsync({ address: account, abi, functionName: "addSupportedToken", args: [v as Address] })}
      onRemove={(v) => writeContractAsync({ address: account, abi, functionName: "removeSupportedToken", args: [v as Address] })}
      onChanged={refetch}
      renderItem={(v) => <AddressDisplay address={v} />}
    />
  );
}
