import { type ReactNode } from "react";
import { PermissionHint } from "./PermissionHint";

interface RoleGateProps {
  hasRole: boolean;
  isLoading?: boolean;
  roleName: string;
  /** Short description of the gated action, e.g. "Add bot". Shown to the user. */
  action?: string;
  children: ReactNode;
}

export function RoleGate({ hasRole, isLoading, roleName, action, children }: RoleGateProps) {
  if (isLoading) return <span className="text-xs text-gray-400">Checking permissions…</span>;
  if (!hasRole) return <PermissionHint roleName={roleName} action={action} />;
  return <>{children}</>;
}
