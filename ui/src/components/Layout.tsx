import { Link, NavLink, Outlet } from "react-router-dom";
import { ConnectButton } from "./ConnectButton";
import { NetworkSelector } from "./NetworkSelector";
import { RefreshButton } from "./RefreshButton";
import { ThemeToggle } from "./ThemeToggle";

const navLinkClass = ({ isActive }: { isActive: boolean }) =>
  `relative font-mono text-xs uppercase tracking-[0.14em] transition-colors ${
    isActive
      ? "text-camino-700 dark:text-camino-300"
      : "text-tarmac-500 hover:text-tarmac-900 dark:text-tarmac-400 dark:hover:text-tarmac-100"
  }`;

/** Underline rule that lights up teal under the active board section. */
function NavItem({ to, end, children }: { to: string; end?: boolean; children: React.ReactNode }) {
  return (
    <NavLink to={to} end={end} className={navLinkClass}>
      {({ isActive }) => (
        <span className="relative inline-block py-0.5">
          {children}
          {isActive && <span className="absolute -bottom-[3px] left-0 right-0 h-[2px] bg-camino-500" aria-hidden />}
        </span>
      )}
    </NavLink>
  );
}

export function Layout() {
  return (
    <div className="min-h-screen bg-paper dark:bg-tarmac-950">
      <header className="border-b border-tarmac-200 bg-paper-raised/90 backdrop-blur dark:border-tarmac-800 dark:bg-tarmac-900/90">
        <div className="mx-auto flex max-w-6xl flex-wrap items-center justify-between gap-x-6 gap-y-3 px-6 py-3">
          <div className="flex items-center gap-6">
            <Link to="/" className="group flex items-center gap-2">
              <span
                className="grid h-6 w-6 place-items-center rounded-sm bg-tarmac-900 font-display text-sm font-bold text-camino-400 dark:bg-camino-500 dark:text-tarmac-950"
                aria-hidden
              >
                ◆
              </span>
              <span className="font-display text-sm font-bold uppercase tracking-[0.12em] text-tarmac-900 dark:text-tarmac-50">
                Camino<span className="text-camino-600 dark:text-camino-400"> Messenger</span>
              </span>
            </Link>
            <nav className="flex items-center gap-5">
              <NavItem to="/" end>
                Dashboard
              </NavItem>
              <NavItem to="/activity">Activity</NavItem>
              <NavItem to="/manager">Manager</NavItem>
              <NavItem to="/create">Create</NavItem>
            </nav>
          </div>
          <div className="flex items-center gap-2">
            <RefreshButton />
            <NetworkSelector />
            <ThemeToggle />
            <ConnectButton />
          </div>
        </div>
      </header>
      <main className="mx-auto max-w-6xl px-6 py-8">
        <Outlet />
      </main>
    </div>
  );
}
