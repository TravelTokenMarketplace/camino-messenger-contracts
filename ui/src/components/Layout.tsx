import { Link, NavLink, Outlet } from "react-router-dom";
import { ConnectButton } from "./ConnectButton";
import { NetworkSelector } from "./NetworkSelector";
import { ThemeToggle } from "./ThemeToggle";

const navLinkClass = ({ isActive }: { isActive: boolean }) =>
  `text-sm transition-colors ${
    isActive
      ? "font-medium text-indigo-600 dark:text-indigo-400"
      : "text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-100"
  }`;

export function Layout() {
  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <header className="flex flex-wrap items-center justify-between gap-3 border-b bg-white px-6 py-3 dark:border-gray-800 dark:bg-gray-950">
        <div className="flex items-center gap-4">
          <Link to="/" className="font-bold">
            Camino Messenger
          </Link>
          <NavLink to="/" end className={navLinkClass}>
            Dashboard
          </NavLink>
          <NavLink to="/manager" className={navLinkClass}>
            Manager
          </NavLink>
          <NavLink to="/create" className={navLinkClass}>
            Create Account
          </NavLink>
        </div>
        <div className="flex items-center gap-3">
          <NetworkSelector />
          <ThemeToggle />
          <ConnectButton />
        </div>
      </header>
      <main className="mx-auto max-w-5xl p-6">
        <Outlet />
      </main>
    </div>
  );
}
