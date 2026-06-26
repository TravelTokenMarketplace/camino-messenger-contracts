import { useMemo } from "react";
import { type Abi, type Address } from "viem";
import { useReadContracts } from "wagmi";
import { type ActivitySourceInput, useActivity } from "./useActivity";
import { useBlockTimestamps } from "./useBlockTimestamps";
import { useActiveContracts } from "./useActiveContracts";
import { ACCOUNT_EVENTS, SERVICE_HASH_EVENTS, renderSentence } from "../lib/activity/catalog";

/**
 * All activity emitted by a single CM Account proxy (bots, services, tokens,
 * pubkeys, funds, config). One address, so a single getLogs filter covers it.
 *
 * Service events carry only the keccak hash of the service name (it's an indexed
 * string). We resolve those hashes to human names via the manager's
 * getServiceNameByHash and re-render the affected sentences.
 */
export function useAccountActivity(account: Address) {
  const { chainId, manager, managerAbi } = useActiveContracts();
  const abi = managerAbi as Abi;

  const sources = useMemo<ActivitySourceInput[]>(
    () => [{ source: "account", address: account, events: ACCOUNT_EVENTS }],
    [account],
  );

  const activity = useActivity({ sources, chainId });

  // Unique service-name hashes present in the loaded events.
  const serviceHashes = useMemo(() => {
    const set = new Set<string>();
    for (const e of activity.events) {
      if (SERVICE_HASH_EVENTS.has(e.eventName) && typeof e.args.serviceName === "string") {
        set.add(e.args.serviceName);
      }
    }
    return [...set];
  }, [activity.events]);

  const { data: nameReads } = useReadContracts({
    allowFailure: true,
    contracts: serviceHashes.map(
      (hash) => ({ chainId, address: manager, abi, functionName: "getServiceNameByHash", args: [hash] }) as const,
    ),
    query: { enabled: Boolean(manager) && serviceHashes.length > 0 },
  });

  const nameByHash = useMemo(() => {
    const map = new Map<string, string>();
    serviceHashes.forEach((hash, i) => {
      const name = nameReads?.[i]?.result as string | undefined;
      if (name) map.set(hash, name);
    });
    return map;
  }, [serviceHashes, nameReads]);

  const timestamps = useBlockTimestamps(
    chainId,
    activity.events.map((e) => e.blockNumber),
  );

  const events = useMemo(
    () =>
      activity.events.map((e) => {
        const serviceLabel = typeof e.args.serviceName === "string" ? nameByHash.get(e.args.serviceName) : undefined;
        const sentence = serviceLabel ? renderSentence(e.source, e.eventName, { ...e.args, serviceLabel }) : e.sentence;
        return { ...e, timestamp: timestamps.get(e.blockNumber), sentence };
      }),
    [activity.events, timestamps, nameByHash],
  );

  return { ...activity, events };
}
