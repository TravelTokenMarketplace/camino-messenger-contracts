export function shortAddress(a: string): string {
  if (a.length < 10) return a;
  return `${a.slice(0, 6)}…${a.slice(-4)}`;
}

export function explorerTxUrl(base: string, hash: string): string {
  return `${base}/tx/${hash}`;
}

export function explorerAddrUrl(base: string, addr: string): string {
  return `${base}/address/${addr}`;
}

/** Compact, human-friendly role label: "SERVICE_ADMIN_ROLE" -> "Service Admin". */
export function shortRoleName(role: string): string {
  if (role === "DEFAULT_ADMIN_ROLE") return "Admin";
  return role
    .replace(/_ROLE$/, "")
    .split("_")
    .map((w) => w.charAt(0) + w.slice(1).toLowerCase())
    .join(" ");
}
