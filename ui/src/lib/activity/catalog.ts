import { type AbiEvent, type Log, formatEther, parseAbiItem } from "viem";
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

function entry(
  source: ActivitySource,
  signature: string,
  category: CatalogEntry["category"],
  icon: CatalogEntry["icon"],
  render: CatalogEntry["render"],
): CatalogEntry {
  const event = parseAbiItem(signature) as AbiEvent;
  return { source, eventName: event.name, event, category, icon, render };
}

/**
 * The curated catalog: the single source of truth for which events the feed
 * shows and how each renders. Ordering is irrelevant — lookups are by
 * `${source}:${eventName}`.
 */
export const CATALOG: CatalogEntry[] = [
  // ── Manager (ecosystem, contract-level) ──────────────────────────────────
  entry(
    "manager",
    "event CMAccountCreated(address indexed account)",
    "Accounts",
    UserPlus,
    (a) => `CM Account ${addr(a.account)} created`,
  ),
  entry(
    "manager",
    "event ServiceRegistered(string serviceName, bytes32 serviceHash)",
    "Services",
    Server,
    (a) => `Service "${str(a.serviceName)}" registered`,
  ),
  entry(
    "manager",
    "event ServiceUnregistered(string serviceName, bytes32 serviceHash)",
    "Services",
    Server,
    (a) => `Service "${str(a.serviceName)}" unregistered`,
  ),

  // ── BookingToken (ecosystem, contract-level) ─────────────────────────────
  entry(
    "bookingToken",
    "event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable)",
    "Bookings",
    Ticket,
    (a) => `Booking token ${id(a.tokenId)} reserved for ${addr(a.reservedFor)}`,
  ),
  entry(
    "bookingToken",
    "event TokenBought(uint256 indexed tokenId, address indexed buyer)",
    "Bookings",
    Ticket,
    (a) => `Booking token ${id(a.tokenId)} bought by ${addr(a.buyer)}`,
  ),
  entry(
    "bookingToken",
    "event TokenReservationExpired(uint256 indexed tokenId)",
    "Bookings",
    Ticket,
    (a) => `Booking token ${id(a.tokenId)} reservation expired`,
  ),
  entry(
    "bookingToken",
    "event CancellationPending(uint256 indexed tokenId, address indexed initialProposer, address indexed currentProposer, uint256 refundAmount, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)",
    "Cancellations",
    XCircle,
    (a) => `Cancellation proposed for booking token ${id(a.tokenId)}`,
  ),
  entry(
    "bookingToken",
    "event CancellationFinalized(uint256 indexed tokenId)",
    "Cancellations",
    XCircle,
    (a) => `Cancellation finalized for booking token ${id(a.tokenId)}`,
  ),
  entry(
    "bookingToken",
    "event CancellationWithdrawn(uint256 indexed tokenId, uint16 withdrawalReason, uint16 withdrawalVersion)",
    "Cancellations",
    XCircle,
    (a) => `Cancellation withdrawn for booking token ${id(a.tokenId)}`,
  ),
  entry(
    "bookingToken",
    "event CancellationRejected(uint256 indexed tokenId, uint16 rejectionReason, uint16 rejectionVersion)",
    "Cancellations",
    XCircle,
    (a) => `Cancellation rejected for booking token ${id(a.tokenId)}`,
  ),

  // ── CM Account (account detail tab — everything) ─────────────────────────
  entry(
    "account",
    "event MessengerBotAdded(address indexed bot)",
    "Bots",
    Bot,
    (a) => `Messenger bot ${addr(a.bot)} added`,
  ),
  entry(
    "account",
    "event MessengerBotRemoved(address indexed bot)",
    "Bots",
    Bot,
    (a) => `Messenger bot ${addr(a.bot)} removed`,
  ),
  entry(
    "account",
    "event ServiceAdded(string indexed serviceName)",
    "Services",
    Server,
    (a) => `Supported service ${serviceLabel(a)} added`,
  ),
  entry(
    "account",
    "event ServiceRemoved(string indexed serviceName)",
    "Services",
    Server,
    (a) => `Supported service ${serviceLabel(a)} removed`,
  ),
  entry(
    "account",
    "event WantedServiceAdded(string indexed serviceName)",
    "Services",
    Server,
    (a) => `Wanted service ${serviceLabel(a)} added`,
  ),
  entry(
    "account",
    "event WantedServiceRemoved(string indexed serviceName)",
    "Services",
    Server,
    (a) => `Wanted service ${serviceLabel(a)} removed`,
  ),
  entry(
    "account",
    "event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} restricted rate ${a.restrictedRate ? "enabled" : "disabled"}`,
  ),
  entry(
    "account",
    "event ServiceCapabilitiesUpdated(string indexed serviceName)",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} capabilities updated`,
  ),
  entry(
    "account",
    "event ServiceCapabilityAdded(string indexed serviceName, string capability)",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} capability "${str(a.capability)}" added`,
  ),
  entry(
    "account",
    "event ServiceCapabilityRemoved(string indexed serviceName, string capability)",
    "Services",
    Server,
    (a) => `Service ${serviceLabel(a)} capability "${str(a.capability)}" removed`,
  ),
  entry(
    "account",
    "event PaymentTokenAdded(address indexed token)",
    "Tokens",
    Coins,
    (a) => `Payment token ${addr(a.token)} added`,
  ),
  entry(
    "account",
    "event PaymentTokenRemoved(address indexed token)",
    "Tokens",
    Coins,
    (a) => `Payment token ${addr(a.token)} removed`,
  ),
  entry(
    "account",
    "event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)",
    "Config",
    ShieldCheck,
    (a) => `Off-chain payment support ${a.supportsOffChainPayment ? "enabled" : "disabled"}`,
  ),
  entry(
    "account",
    "event PublicKeyAdded(address indexed pubKeyAddress)",
    "Pubkeys",
    KeyRound,
    (a) => `Public key ${addr(a.pubKeyAddress)} added`,
  ),
  entry(
    "account",
    "event PublicKeyRemoved(address indexed pubKeyAddress)",
    "Pubkeys",
    KeyRound,
    (a) => `Public key ${addr(a.pubKeyAddress)} removed`,
  ),
  entry(
    "account",
    "event Deposit(address indexed sender, uint256 amount)",
    "Funds",
    ArrowDownToLine,
    (a) => `Deposit of ${ether(a.amount)} from ${addr(a.sender)}`,
  ),
  entry(
    "account",
    "event Withdraw(address indexed receiver, uint256 amount)",
    "Funds",
    ArrowUpFromLine,
    (a) => `Withdrawal of ${ether(a.amount)} to ${addr(a.receiver)}`,
  ),
  entry(
    "account",
    "event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)",
    "Funds",
    Fuel,
    (a) => `Gas money withdrawal of ${ether(a.amount)} by ${addr(a.withdrawer)}`,
  ),
  entry(
    "account",
    "event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)",
    "Config",
    Fuel,
    (a) => `Gas money limit updated to ${ether(a.limit)} per ${str(a.period)}s`,
  ),
  entry(
    "account",
    "event CMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)",
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
