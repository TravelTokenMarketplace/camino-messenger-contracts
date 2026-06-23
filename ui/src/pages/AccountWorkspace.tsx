import { Link, useParams, useSearchParams } from "react-router-dom";
import { ArrowUpFromLine, Bot, Coins, KeyRound, Server, Users } from "lucide-react";
import { type Address } from "viem";
import { AccountSummary } from "../components/AccountSummary";
import { AccountValidityNotice } from "../components/AccountValidityNotice";
import { RefreshButton } from "../components/RefreshButton";
import { TxPanel } from "../components/TxPanel";
import { BotsTab } from "./tabs/BotsTab";
import { PaymentTokensTab } from "./tabs/PaymentTokensTab";
import { ServicesTab } from "./tabs/ServicesTab";
import { RolesTab } from "./tabs/RolesTab";
import { PubkeysTab } from "./tabs/PubkeysTab";
import { WithdrawalsTab } from "./tabs/WithdrawalsTab";

const TABS = [
  { id: "bots", label: "Bots", Icon: Bot, Component: BotsTab },
  { id: "tokens", label: "Payment Tokens", Icon: Coins, Component: PaymentTokensTab },
  { id: "services", label: "Services", Icon: Server, Component: ServicesTab },
  { id: "roles", label: "Roles", Icon: Users, Component: RolesTab },
  { id: "pubkeys", label: "Pubkeys", Icon: KeyRound, Component: PubkeysTab },
  { id: "withdrawals", label: "Withdrawals", Icon: ArrowUpFromLine, Component: WithdrawalsTab },
] as const;

export function AccountWorkspace() {
  const { address } = useParams();
  const [params] = useSearchParams();
  const active = params.get("tab") ?? TABS[0].id;
  const account = address as Address;
  const Active = (TABS.find((t) => t.id === active) ?? TABS[0]).Component;

  return (
    <div className="grid items-start gap-6 md:grid-cols-[260px_1fr]">
      <div className="flex flex-col gap-4">
        <AccountValidityNotice account={account} />
        <AccountSummary account={account} />
        <TxPanel />
      </div>
      <div className="flex min-w-0 flex-col gap-4">
        <div className="flex items-end justify-between gap-3 border-b text-sm dark:border-gray-800">
          <nav className="flex flex-wrap gap-3">
            {TABS.map((t) => (
              <Link
                key={t.id}
                to={`?tab=${t.id}`}
                className={`-mb-px inline-flex items-center gap-1.5 border-b-2 pb-2 ${
                  active === t.id
                    ? "border-indigo-600 font-medium text-gray-900 dark:text-gray-100"
                    : "border-transparent text-gray-500 dark:text-gray-400"
                }`}
              >
                <t.Icon className="h-4 w-4" /> {t.label}
              </Link>
            ))}
          </nav>
          <div className="pb-1.5">
            <RefreshButton />
          </div>
        </div>
        <Active account={account} />
      </div>
    </div>
  );
}
