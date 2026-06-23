import { type ReactNode } from "react";

interface RoleGateProps {
  hasRole: boolean;
  isLoading?: boolean;
  roleName: string;
  children: ReactNode;
}

export function RoleGate({ hasRole, isLoading, roleName, children }: RoleGateProps) {
  if (isLoading) return <span className="text-xs text-gray-400">Checking permissions…</span>;
  if (!hasRole)
    return (
      <p className="text-xs text-amber-600">
        Requires <code>{roleName}</code> on the connected account.
      </p>
    );
  return <>{children}</>;
}
