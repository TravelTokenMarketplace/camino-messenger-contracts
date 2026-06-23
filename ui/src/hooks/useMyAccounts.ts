import { type Address } from "viem";
import { usePublicClient } from "wagmi";
import { useEffect, useState } from "react";
import { useActiveContracts } from "./useActiveContracts";

export function uniqueAddresses(addrs: string[]): string[] {
  const seen = new Set<string>();
  const out: string[] = [];
  for (const a of addrs) {
    const k = a.toLowerCase();
    if (!seen.has(k)) {
      seen.add(k);
      out.push(a);
    }
  }
  return out;
}

export function useMyAccounts() {
  const client = usePublicClient();
  const { manager, managerAbi } = useActiveContracts();
  const [accounts, setAccounts] = useState<Address[]>([]);
  const [isLoading, setLoading] = useState(false);

  useEffect(() => {
    if (!client || !manager) return;
    let cancelled = false;
    (async () => {
      setLoading(true);
      const event = (managerAbi as readonly unknown[]).find(
        (x) => (x as { type: string; name?: string }).type === "event" &&
          (x as { name?: string }).name === "CMAccountCreated",
      );
      const logs = await client.getLogs({ address: manager, event: event as never, fromBlock: 0n, toBlock: "latest" });
      const addrs = logs.map((l) => String((l as { args: { account: string } }).args.account));
      if (!cancelled) {
        setAccounts(uniqueAddresses(addrs) as Address[]);
        setLoading(false);
      }
    })();
    return () => { cancelled = true; };
  }, [client, manager, managerAbi]);

  return { accounts, isLoading };
}
