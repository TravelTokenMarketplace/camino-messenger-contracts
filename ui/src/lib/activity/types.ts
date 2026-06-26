import { type AbiEvent, type Address, type Hex } from "viem";
import { type LucideIcon } from "lucide-react";

/** Which contract a feed event originates from. */
export type ActivitySource = "manager" | "bookingToken" | "account";

/**
 * Grouping used for the Activity-page filter chips. The ecosystem feed only ever
 * surfaces the first four; the account-level categories are cosmetic (the account
 * tab does not filter). Chips are derived from the categories actually present in
 * the loaded events, so this union can grow without touching the UI.
 */
export type ActivityCategory =
  | "Bookings"
  | "Cancellations"
  | "Accounts"
  | "Services"
  | "Bots"
  | "Tokens"
  | "Pubkeys"
  | "Funds"
  | "Config";

/** A normalized, render-ready activity event decoded from an on-chain log. */
export interface ActivityEvent {
  /** `${txHash}#${logIndex}` — stable dedupe / React key. */
  id: string;
  source: ActivitySource;
  category: ActivityCategory;
  /** Emitting contract address. */
  contract: Address;
  blockNumber: bigint;
  logIndex: number;
  /** Unix seconds; filled in by useBlockTimestamps, absent until then. */
  timestamp?: number;
  txHash: Hex;
  eventName: string;
  args: Record<string, unknown>;
  /** Human-readable one-line sentence (addresses shortened). */
  sentence: string;
}

/** One curated event in the catalog. */
export interface CatalogEntry {
  source: ActivitySource;
  eventName: string;
  /** viem ABI event used as the `getLogs` filter. */
  event: AbiEvent;
  category: ActivityCategory;
  icon: LucideIcon;
  render: (args: Record<string, unknown>) => string;
}
