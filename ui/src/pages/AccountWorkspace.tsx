import { Link, useParams, useSearchParams } from "react-router-dom";
import { type Address } from "viem";
import { AccountSummary } from "../components/AccountSummary";
import { OverviewTab } from "./tabs/OverviewTab";
import { BotsTab } from "./tabs/BotsTab";
import { PaymentTokensTab } from "./tabs/PaymentTokensTab";
import { ServicesTab } from "./tabs/ServicesTab";
import { RolesTab } from "./tabs/RolesTab";
import { PubkeysTab } from "./tabs/PubkeysTab";
import { WithdrawalsTab } from "./tabs/WithdrawalsTab";

const TABS = [
  { id: "overview", label: "Overview", Component: OverviewTab },
  { id: "bots", label: "Bots", Component: BotsTab },
  { id: "tokens", label: "Payment Tokens", Component: PaymentTokensTab },
  { id: "services", label: "Services", Component: ServicesTab },
  { id: "roles", label: "Roles", Component: RolesTab },
  { id: "pubkeys", label: "Pubkeys", Component: PubkeysTab },
  { id: "withdrawals", label: "Withdrawals", Component: WithdrawalsTab },
] as const;

export function AccountWorkspace() {
  const { address } = useParams();
  const [params] = useSearchParams();
  const active = params.get("tab") ?? "overview";
  const account = address as Address;
  const Active = (TABS.find((t) => t.id === active) ?? TABS[0]).Component;

  return (
    <div className="grid gap-6 md:grid-cols-[260px_1fr]">
      <AccountSummary account={account} />
      <div className="grid gap-4">
        <nav className="flex flex-wrap gap-3 border-b text-sm dark:border-gray-800">
          {TABS.map((t) => (
            <Link key={t.id} to={`?tab=${t.id}`} className={`pb-2 ${active === t.id ? "border-b-2 border-indigo-600 font-medium" : "text-gray-500 dark:text-gray-400"}`}>
              {t.label}
            </Link>
          ))}
        </nav>
        <Active account={account} />
      </div>
    </div>
  );
}
