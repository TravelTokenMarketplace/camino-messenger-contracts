import { useMemo } from "react";
import { type Address } from "viem";
import { type ActivitySourceInput, useActivity } from "./useActivity";
import { useBlockTimestamps } from "./useBlockTimestamps";
import { useActiveContracts } from "./useActiveContracts";
import { ACCOUNT_EVENTS } from "../lib/activity/catalog";

/**
 * All activity emitted by a single CM Account proxy (bots, services, tokens,
 * pubkeys, funds, config). One address, so a single getLogs filter covers it.
 */
export function useAccountActivity(account: Address) {
  const { chainId } = useActiveContracts();

  const sources = useMemo<ActivitySourceInput[]>(
    () => [{ source: "account", address: account, events: ACCOUNT_EVENTS }],
    [account],
  );

  const activity = useActivity({ sources, chainId });
  const timestamps = useBlockTimestamps(
    chainId,
    activity.events.map((e) => e.blockNumber),
  );

  const events = useMemo(
    () => activity.events.map((e) => ({ ...e, timestamp: timestamps.get(e.blockNumber) })),
    [activity.events, timestamps],
  );

  return { ...activity, events };
}
