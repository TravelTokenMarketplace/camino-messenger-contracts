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
