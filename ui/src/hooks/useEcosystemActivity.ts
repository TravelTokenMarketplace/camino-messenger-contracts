import { useMemo } from "react";
import { type ActivitySourceInput, useActivity } from "./useActivity";
import { useBlockTimestamps } from "./useBlockTimestamps";
import { useActiveContracts } from "./useActiveContracts";
import { BOOKING_TOKEN_EVENTS, MANAGER_EVENTS } from "../lib/activity/catalog";

/**
 * Ecosystem-wide activity: contract-level events from the manager and the
 * BookingToken. Used by the Dashboard sneak peek and the Activity page.
 */
export function useEcosystemActivity() {
  const { manager, bookingToken, chainId } = useActiveContracts();

  const sources = useMemo(() => {
    const s: ActivitySourceInput[] = [];
    if (manager) s.push({ source: "manager", address: manager, events: MANAGER_EVENTS });
    if (bookingToken) s.push({ source: "bookingToken", address: bookingToken, events: BOOKING_TOKEN_EVENTS });
    return s;
  }, [manager, bookingToken]);

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
