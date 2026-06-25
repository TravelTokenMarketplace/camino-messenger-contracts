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

/**
 * Format a decimal amount string for display: thousands separators on the
 * integer part, and the fractional part trimmed to `sigFractionDigits`
 * significant digits (skipping leading zeros) with trailing zeros stripped.
 * Returns the trimmed `display` and the untouched `full` value (for tooltips).
 */
export function formatAmount(value: string, sigFractionDigits = 7): { display: string; full: string } {
  const full = value;
  const neg = value.startsWith("-");
  const v = neg ? value.slice(1) : value;
  const [intRaw, fracRaw = ""] = v.split(".");
  const intGrouped = intRaw.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  let frac = fracRaw;
  if (frac) {
    let lead = 0;
    while (lead < frac.length && frac[lead] === "0") lead++;
    frac = frac.slice(0, lead + sigFractionDigits).replace(/0+$/, "");
  }
  const display = (neg ? "-" : "") + intGrouped + (frac ? "." + frac : "");
  return { display, full };
}
