import { Link, useSearchParams } from "react-router-dom";
import { KeyRound, Server, Settings, Ticket } from "lucide-react";
import { Card } from "../components/Card";
import { ManagerSummary } from "../components/ManagerSummary";
import { RefreshButton } from "../components/RefreshButton";
import { TxPanel } from "../components/TxPanel";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { ManagerConfigTab } from "./tabs/ManagerConfigTab";
import { ServiceRegistryTab } from "./tabs/ServiceRegistryTab";
import { ManagerRolesTab } from "./tabs/ManagerRolesTab";
import { BookingTokenTab } from "./tabs/BookingTokenTab";

const TABS = [
  { id: "config", label: "Config", Icon: Settings, Component: ManagerConfigTab },
  { id: "services", label: "Service Registry", Icon: Server, Component: ServiceRegistryTab },
  { id: "roles", label: "Manager Roles", Icon: KeyRound, Component: ManagerRolesTab },
  { id: "booking-token", label: "Booking Token", Icon: Ticket, Component: BookingTokenTab },
] as const;

export function ManagerWorkspace() {
  const { supported } = useActiveContracts();
  const [params] = useSearchParams();
  const active = params.get("tab") ?? TABS[0].id;
  const Active = (TABS.find((t) => t.id === active) ?? TABS[0]).Component;

  if (!supported) return <Card title="Manager">Connect to a supported network.</Card>;

  return (
    <div className="grid items-start gap-6 md:grid-cols-[260px_1fr]">
      <div className="flex flex-col gap-4">
        <ManagerSummary />
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
                    ? "border-indigo-600 dark:border-indigo-400 font-medium text-gray-900 dark:text-gray-100"
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
        <Active />
      </div>
    </div>
  );
}
