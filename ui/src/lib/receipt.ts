import { decodeEventLog, type Abi, type Address, type Log } from "viem";

export function findCreatedAccount(logs: Pick<Log, "data" | "topics">[], abi: Abi): Address | undefined {
  for (const log of logs) {
    try {
      const decoded = decodeEventLog({ abi, data: log.data, topics: log.topics });
      if (decoded.eventName === "CMAccountCreated") {
        return (decoded.args as unknown as { account: Address }).account;
      }
    } catch {
      // not our event, skip
    }
  }
  return undefined;
}
