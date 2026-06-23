import { Link, Outlet } from "react-router-dom";
import { ConnectButton } from "./ConnectButton";
import { NetworkSelector } from "./NetworkSelector";
import { ThemeToggle } from "./ThemeToggle";

export function Layout() {
  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <header className="flex flex-wrap items-center justify-between gap-3 border-b bg-white px-6 py-3 dark:border-gray-800 dark:bg-gray-950">
        <div className="flex items-center gap-4">
          <Link to="/" className="font-bold">Camino Messenger</Link>
          <Link to="/" className="text-sm text-gray-600 dark:text-gray-400">Dashboard</Link>
          <Link to="/create" className="text-sm text-gray-600 dark:text-gray-400">Create Account</Link>
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
