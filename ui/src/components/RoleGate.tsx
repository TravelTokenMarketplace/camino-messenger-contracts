import { type ReactNode } from "react";
import { PermissionHint } from "./PermissionHint";

interface RoleGateProps {
  hasRole: boolean;
  isLoading?: boolean;
  roleName: string;
  children: ReactNode;
}

export function RoleGate({ hasRole, isLoading, roleName, children }: RoleGateProps) {
  if (isLoading) return <span className="text-xs text-gray-400">Checking permissions…</span>;
  if (!hasRole) return <PermissionHint roleName={roleName} />;
  return <>{children}</>;
}
