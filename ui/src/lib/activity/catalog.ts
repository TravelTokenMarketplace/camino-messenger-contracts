import { type Abi, type AbiEvent, type Log, formatEther } from "viem";
import {
  ArrowDownToLine,
  ArrowUpFromLine,
  Bot,
  Coins,
  Fuel,
  KeyRound,
  Server,
  ShieldCheck,
  Ticket,
  UserPlus,
  XCircle,
} from "lucide-react";
import { BOOKINGTOKEN_ABI, CMACCOUNT_ABI, MANAGER_ABI } from "../../contracts";
import { shortAddress } from "../format";
import { type ActivityEvent, type ActivitySource, type CatalogEntry } from "./types";

// Helpers for rendering args. Logs are decoded by viem, so args carry their
// solidity types: uint -> bigint, address -> string, bool -> boolean. Indexed
// `string` params (e.g. a service name) arrive as a keccak hash, not the
// original string. We render a short hash by default; useAccountActivity can
// resolve the real name via the manager and inject it as `serviceLabel`, after
// which renderSentence() is called again to upgrade the text.
const addr = (v: unknown) => shortAddress(String(v));
const id = (v: unknown) => `#${String(v)}`;
const ether = (v: unknown) => formatEther(BigInt(v as bigint | number | string));
const str = (v: unknown) => String(v);

/** Human service label: the resolved name in quotes, else the short hash. */
function serviceLabel(args: Record<string, unknown>): string {
  const name = args.serviceLabel as string | undefined;
  if (name) return `"${name}"`;
  return `(${addr(args.serviceName)})`;
}

// The catalog references events by name; the ABI surface itself is the single
// source of truth. Resolving each event from the generated ABIs (rather than
// re-typing `parseAbiItem(...)` signatures) keeps the getLogs filter and decode
// aligned with the contracts — if a signature changes upstream, `yarn sync`
// regenerates the ABI and this picks it up instead of drifting silently.
const ABI_BY_SOURCE: Record<ActivitySource, Abi> = {
  manager: MANAGER_ABI as Abi,
  bookingToken: BOOKINGTOKEN_ABI as Abi,
  account: CMACCOUNT_ABI as Abi,
};

function abiEvent(source: ActivitySource, eventName: string): AbiEvent {
  const found = ABI_BY_SOURCE[source].find(
    (item): item is AbiEvent => item.type === "event" && item.name === eventName,
  );
  if (!found) throw new Error(`Activity catalog: event "${eventName}" not found in ${source} ABI`);
  return found;
}

function entry(
  source: ActivitySource,
  eventName: string,
  category: CatalogEntry["category"],
  icon: CatalogEntry["icon"],
  render: CatalogEntry["render"],
): CatalogEntry {
  return { source, eventName, event: abiEvent(source, eventName), category, icon, render };
}

/**
 * The curated catalog: the single source of truth for which events the feed
 * shows and how each renders. Ordering is irrelevant — lookups are by
 * `${source}:${eventName}`.
 */
export const CATALOG: CatalogEntry[] = [
  // ── Manager (ecosystem, contract-level) ──────────────────────────────────
  entry("manager", "CMAccountCreated", "Accounts", UserPlus, (a) => `CM Account ${addr(a.account)} created`),
  entry("manager", "ServiceRegistered", "Services", Server, (a) => `Service "${str(a.serviceName)}" registered`),
  entry("manager", "ServiceUnregistered", "Services", Server, (a) => `Service "${str(a.serviceName)}" unregistered`),

  // ── BookingToken (ecosystem, contract-level) ─────────────────────────────
  entry(
    "bookingToken",
    "TokenReserved",
    "Bookings",
    Ticket,
    (a) => `Booking token ${id(a.tokenId)} reserved for ${addr(a.reservedFor)}`,
  ),
  entry(
    "bookingToken",
    "TokenBought",
    "Bookings",
    Ticket,
    (a) => `Booking token ${id(a.tokenId)} bought by ${addr(a.buyer)}`,
  ),
  entry(
    "bookingToken",
    "TokenReservationExpired",
    "Bookings",
    Ticket,
    (a) => `Booking token ${id(a.tokenId)} reservation expired`,
  ),
  entry(
    "bookingToken",
    "CancellationPending",
    "Cancellations",
    XCircle,
    (a) => `Cancellation proposed for booking token ${id(a.tokenId)}`,
  ),
  entry(
    "bookingToken",
    "CancellationFinalized",
    "Cancellations",
    XCircle,
    (a) => `Cancellation finalized for booking token ${id(a.tokenId)}`,
  ),
  entry(
    "bookingToken",
    "CancellationWithdrawn",
    "Cancellations",
    XCircle,
    (a) => `Cancellation withdrawn for booking token ${id(a.tokenId)}`,
  ),
  entry(
    "bookingToken",
    "CancellationRejected",
    "Cancellations",
    XCircle,
    (a) => `Cancellation rejected for booking token ${id(a.tokenId)}`,
  ),

  // ── CM Account (account detail tab — everything) ─────────────────────────
  entry("account", "MessengerBotAdded", "Bots", Bot, (a) => `Messenger bot ${addr(a.bot)} added`),
  entry("account", "MessengerBotRemoved", "Bots", Bot, (a) => `Messenger bot ${addr(a.bot)} removed`),
  entry("account", "ServiceAdded", "Services", Server, (a) => `Supported service ${serviceLabel(a)} added`),
  entry("account", "ServiceRemoved", "Services", Server, (a) => `Supported service ${serviceLabel(a)} removed`),
  entry("account", "WantedServiceAdded", "Services", Server, (a) => `Wanted service ${serviceLabel(a)} added`),
  entry("account", "WantedServiceRemoved", "Services", Server, (a) => `Wanted service ${serviceLabel(a)} removed`),
  entry(
    "account",
    "ServiceRestrictedRateUpdated",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} restricted rate ${a.restrictedRate ? "enabled" : "disabled"}`,
  ),
  entry(
    "account",
    "ServiceCapabilitiesUpdated",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} capabilities updated`,
  ),
  entry(
    "account",
    "ServiceCapabilityAdded",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} capability "${str(a.capability)}" added`,
  ),
  entry(
    "account",
    "ServiceCapabilityRemoved",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} capability "${str(a.capability)}" removed`,
  ),
  entry("account", "PaymentTokenAdded", "Tokens", Coins, (a) => `Payment token ${addr(a.token)} added`),
  entry("account", "PaymentTokenRemoved", "Tokens", Coins, (a) => `Payment token ${addr(a.token)} removed`),
  entry(
    "account",
    "OffChainPaymentSupportUpdated",
    "Config",
    ShieldCheck,
    (a) => `Off-chain payment support ${a.supportsOffChainPayment ? "enabled" : "disabled"}`,
  ),
  entry("account", "PublicKeyAdded", "Pubkeys", KeyRound, (a) => `Public key ${addr(a.pubKeyAddress)} added`),
  entry("account", "PublicKeyRemoved", "Pubkeys", KeyRound, (a) => `Public key ${addr(a.pubKeyAddress)} removed`),
  entry("account", "Deposit", "Funds", ArrowDownToLine, (a) => `Deposit of ${ether(a.amount)} from ${addr(a.sender)}`),
  entry(
    "account",
    "Withdraw",
    "Funds",
    ArrowUpFromLine,
    (a) => `Withdrawal of ${ether(a.amount)} to ${addr(a.receiver)}`,
  ),
  entry(
    "account",
    "GasMoneyWithdrawal",
    "Funds",
    Fuel,
    (a) => `Gas money withdrawal of ${ether(a.amount)} by ${addr(a.withdrawer)}`,
  ),
  entry(
    "account",
    "GasMoneyWithdrawalUpdated",
    "Config",
    Fuel,
    (a) => `Gas money limit updated to ${ether(a.limit)} per ${str(a.period)}s`,
  ),
  entry(
    "account",
    "CMAccountUpgraded",
    "Config",
    ShieldCheck,
    (a) => `Account upgraded to implementation ${addr(a.newImplementation)}`,
  ),
];

const BY_KEY = new Map(CATALOG.map((e) => [`${e.source}:${e.eventName}`, e]));

export function lookupEntry(source: ActivitySource, eventName: string): CatalogEntry | undefined {
  return BY_KEY.get(`${source}:${eventName}`);
}

/** Render an event's sentence from (possibly enriched) args. */
export function renderSentence(source: ActivitySource, eventName: string, args: Record<string, unknown>): string {
  return lookupEntry(source, eventName)?.render(args) ?? eventName;
}

/** Account events whose indexed `serviceName` is a hash we can resolve to a name. */
export const SERVICE_HASH_EVENTS = new Set(
  CATALOG.filter((e) => e.source === "account" && e.category === "Services").map((e) => e.eventName),
);

/** The viem ABI events for a source, ready to pass as `getLogs({ events })`. */
export function eventsForSource(source: ActivitySource): AbiEvent[] {
  return CATALOG.filter((e) => e.source === source).map((e) => e.event);
}

export const MANAGER_EVENTS = eventsForSource("manager");
export const BOOKING_TOKEN_EVENTS = eventsForSource("bookingToken");
export const ACCOUNT_EVENTS = eventsForSource("account");

type DecodedLog = Log<bigint, number, false> & { eventName?: string; args?: Record<string, unknown> };

/**
 * Map a viem-decoded log to an ActivityEvent via the catalog. Returns undefined
 * for logs whose event isn't in the catalog (defensive — getLogs is already
 * filtered to catalog events).
 */
export function toActivityEvent(log: DecodedLog, source: ActivitySource): ActivityEvent | undefined {
  const name = log.eventName;
  if (!name) return undefined;
  const found = lookupEntry(source, name);
  if (!found || log.blockNumber == null || log.logIndex == null || !log.transactionHash) return undefined;
  const args = (log.args ?? {}) as Record<string, unknown>;
  return {
    id: `${log.transactionHash}#${log.logIndex}`,
    source,
    category: found.category,
    contract: log.address,
    blockNumber: log.blockNumber,
    logIndex: log.logIndex,
    txHash: log.transactionHash,
    eventName: name,
    args,
    sentence: found.render(args),
  };
}
