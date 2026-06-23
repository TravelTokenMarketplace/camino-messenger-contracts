import { Link, Outlet } from "react-router-dom";
import { ConnectButton } from "./ConnectButton";
import { NetworkBadge } from "./NetworkBadge";

export function Layout() {
  return (
    <div className="min-h-screen bg-gray-50">
      <header className="flex items-center justify-between border-b bg-white px-6 py-3">
        <div className="flex items-center gap-4">
          <Link to="/" className="font-bold">Camino Messenger</Link>
          <Link to="/" className="text-sm text-gray-600">Dashboard</Link>
          <Link to="/create" className="text-sm text-gray-600">Create Account</Link>
        </div>
        <div className="flex items-center gap-3">
          <NetworkBadge />
          <ConnectButton />
        </div>
      </header>
      <main className="mx-auto max-w-5xl p-6">
        <Outlet />
      </main>
    </div>
  );
}
