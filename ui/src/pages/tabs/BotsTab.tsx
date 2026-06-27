import { useState } from "react";
import { AlertTriangle, Check, Info, Plus, Trash2 } from "lucide-react";
import { type Abi, type Address, parseEther } from "viem";
import { useBalance, useReadContracts, useWriteContract } from "wagmi";
import { APP_CHAINS } from "../../config/chains";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { Input } from "../../components/Input";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { Tooltip } from "../../components/Tooltip";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { useHasRole } from "../../hooks/useHasRole";
import { roleHash, type RoleName } from "../../lib/roles";
import { shortRoleName } from "../../lib/format";

// A fully-provisioned bot holds all three roles; an admin can revoke any of
// them individually, so we surface exactly which ones each address has.
const BOT_ROLES: RoleName[] = ["MESSENGER_BOT_ROLE", "BOOKING_OPERATOR_ROLE", "GAS_WITHDRAWER_ROLE"];

function BotBalance({ bot }: { bot: Address }) {
  const { chainId } = useActiveContracts();
  const { data } = useBalance({ address: bot, chainId });
  if (!data) return <span className="text-xs text-tarmac-400">…</span>;
  const isZero = data.value === 0n;
  return (
    <span
      className={`rounded px-2 py-0.5 text-xs font-medium ${
        isZero
          ? "bg-red-100 text-red-700 dark:bg-red-950 dark:text-red-300"
          : "bg-camino-100 text-camino-700 dark:bg-camino-950 dark:text-camino-300"
      }`}
      title={isZero ? "Bot has no funds for transaction fees" : undefined}
    >
      {data.formatted} {data.symbol}
    </span>
  );
}

function BotRoles({ account, bot, abi }: { account: Address; bot: Address; abi: Abi }) {
  const { chainId } = useActiveContracts();
  const { data, isLoading } = useReadContracts({
    contracts: BOT_ROLES.map((r) => ({
      chainId,
      address: account,
      abi,
      functionName: "hasRole",
      args: [roleHash(r), bot],
    })),
    allowFailure: true,
  });

  if (isLoading) return <span className="text-xs text-tarmac-400">Checking roles…</span>;

  return (
    <span className="flex flex-wrap items-center gap-1">
      {BOT_ROLES.map((r, i) => {
        const has = data?.[i]?.result === true;
        return (
          <Tooltip
            key={r}
            content={
              has ? `Has ${r}` : `Missing ${r} — this bot may not function fully. Re-add or grant it in the Roles tab.`
            }
          >
            <span
              className={`inline-flex items-center gap-1 rounded-full px-2 py-0.5 text-xs ${
                has
                  ? "bg-camino-100 text-camino-700 dark:bg-camino-950 dark:text-camino-300"
                  : "bg-amber-100 text-amber-700 dark:bg-amber-950 dark:text-amber-300"
              }`}
            >
              {has ? <Check className="h-3 w-3" /> : <AlertTriangle className="h-3 w-3" />}
              {shortRoleName(r)}
            </span>
          </Tooltip>
        );
      })}
    </span>
  );
}

function BotRow({
  account,
  bot,
  abi,
  hasRole,
  onChanged,
}: {
  account: Address;
  bot: Address;
  abi: Abi;
  hasRole: boolean;
  onChanged: () => void;
}) {
  const { writeContractAsync } = useWriteContract();
  return (
    <li className="group flex flex-col gap-2 py-2.5">
      <div className="flex flex-wrap items-center justify-between gap-2">
        <span className="flex items-center gap-3">
          <AddressDisplay address={bot} className="text-sm" />
          <BotBalance bot={bot} />
        </span>
        {hasRole && (
          <RowAction>
            <TxButton
              label="Remove"
              variant="danger"
              icon={<Trash2 className="h-4 w-4" />}
              tooltip="Revokes all three bot roles from this address — sends a transaction to your wallet."
              write={() =>
                writeContractAsync({ address: account, abi, functionName: "removeMessengerBot", args: [bot] })
              }
              onConfirmed={onChanged}
            />
          </RowAction>
        )}
      </div>
      <BotRoles account={account} bot={bot} abi={abi} />
    </li>
  );
}

export function BotsTab({ account }: { account: Address }) {
  const { cmAccountAbi, chainId } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, abi, "MESSENGER_BOT_ROLE");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, abi, "BOT_ADMIN_ROLE");
  const { data: accountBalance } = useBalance({ address: account, chainId });
  const symbol = accountBalance?.symbol ?? APP_CHAINS.find((c) => c.id === chainId)?.nativeCurrency.symbol ?? "";
  const [bot, setBot] = useState("");
  const [gas, setGas] = useState("");

  // The gas money is paid out of the CM Account's own balance, so validate
  // against that and warn (without blocking) when it would exceed it.
  let gasError: string | undefined;
  let gasValue = 0n;
  if (gas.trim()) {
    try {
      gasValue = parseEther(gas.trim());
    } catch {
      gasError = "Enter a valid amount.";
    }
  }
  const insufficient = !gasError && accountBalance !== undefined && gasValue > accountBalance.value;
  const gasInvalid = Boolean(gasError) || insufficient;

  return (
    <Card title="Messenger Bots">
      <p className="mb-3 text-xs text-tarmac-500 dark:text-tarmac-400">
        A bot is an address granted three roles — Messenger Bot, Booking Operator and Gas Withdrawer — and funded with
        native tokens for transaction fees. Each bot's roles are shown below; an amber badge means that role is missing.
      </p>
      {isLoading || roleLoading ? (
        <p>Loading…</p>
      ) : (
        <ul className="mb-4 divide-y dark:divide-tarmac-700">
          {members.length === 0 && <li className="py-2 text-sm text-tarmac-400">None</li>}
          {members.map((b) => (
            <BotRow key={b} account={account} bot={b as Address} abi={abi} hasRole={hasRole} onChanged={refetch} />
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="BOT_ADMIN_ROLE" action="Add bot">
        <div className="space-y-1.5">
          <div className="flex flex-wrap items-end gap-2">
            <Input
              className="min-w-[14rem] flex-1 font-mono"
              placeholder="Bot address 0x…"
              value={bot}
              onChange={(e) => setBot(e.target.value)}
            />
            <div className="flex items-center gap-1.5">
              <Input
                className={`w-36 ${gasInvalid ? "!border-red-500 focus:!border-red-500 focus:!ring-red-500" : ""}`}
                inputMode="decimal"
                placeholder={`Gas money (${symbol})`}
                value={gas}
                onChange={(e) => setGas(e.target.value)}
              />
              <Tooltip
                content={`This amount is sent from the CM Account's own balance to the bot to cover its transaction fees${
                  symbol ? ` (in ${symbol})` : ""
                }.`}
              >
                <button
                  type="button"
                  aria-label="Gas money info"
                  className="text-tarmac-400 hover:text-tarmac-600 dark:hover:text-tarmac-200"
                >
                  <Info className="h-4 w-4" />
                </button>
              </Tooltip>
            </div>
            <TxButton
              label="Add bot"
              icon={<Plus className="h-4 w-4" />}
              disabled={!bot || Boolean(gasError)}
              write={() =>
                writeContractAsync({
                  address: account,
                  abi,
                  functionName: "addMessengerBot",
                  args: [bot as Address, gasValue],
                })
              }
              onConfirmed={() => {
                setBot("");
                setGas("");
                refetch();
              }}
            />
          </div>
          {gasError && <p className="text-xs text-red-600">{gasError}</p>}
          {insufficient && (
            <p className="text-xs text-red-600">
              Exceeds the CM Account balance ({accountBalance?.formatted} {symbol}). Fund the account first or the
              transaction will revert.
            </p>
          )}
        </div>
      </RoleGate>
    </Card>
  );
}
