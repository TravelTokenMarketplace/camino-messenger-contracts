# Ecosystem Activity feed

**Date:** 2026-06-26
**Area:** `ui/` (React management UI)

## Problem

The management UI today shows only point-in-time state — role members, balances,
service lists, payment tokens. There is no sense of *what just happened* in the
ecosystem: a new partner account being created, booking tokens minted/bought,
a bot added, a service registered. Operators have no at-a-glance "recent activity"
view and must infer change by diffing state over time.

We want a curated, human-readable feed of recent on-chain events, surfaced in
three places: a sneak peek on the Dashboard, a dedicated Activity page, and a
per-account activity view on the CM Account detail page.

## Goals

- Show a curated, human-readable feed of recent ecosystem events.
- Surface it in three places: Dashboard sneak peek, `/activity` page, account detail tab.
- Read events with bounded `eth_getLogs` ranges that public/free-tier RPCs accept.
- Let users page further back through history on demand ("Load older").
- Adapt to per-RPC range limits automatically and gracefully.
- Display relative timestamps ("2h ago"), not just block numbers.

## Non-goals

- **No full historical archive / indexer** (no subgraph, Ponder, Envio). The feed
  is a rolling, on-demand-extendable window read directly from the RPC. An indexer
  is the right tool for permanent full history and is explicitly deferred.
- **No ecosystem-wide aggregation of account-level events** in v1. The Dashboard
  and Activity page show only *contract-level* events (manager + BookingToken).
  Account-level events (bots, services, tokens, deposits) appear only on that
  account's own detail tab, where they are a single-address query.
- No auto-refresh timers — refresh follows the app's existing manual
  `RefreshButton` + query-invalidation convention.
- No changes to Solidity contracts or ABIs.

## Event catalog (decided)

A single curated catalog (`src/lib/activity/catalog.ts`) is the source of truth.
Each entry maps an event to:

- `source`: `"manager" | "bookingToken" | "account"`
- `abiItem`: the viem-parsed ABI event (used as the `getLogs` `events` filter)
- `category`: filter-chip grouping (see below)
- `icon`: a `lucide-react` icon
- `render(args) => string`: a human sentence, e.g.
  *"Booking token #42 bought by 0x12…ab"*, *"CM Account 0xab…cd created"*.

### Ecosystem (contract-level) — Dashboard + Activity page

From the **manager** (`CMAccountManager`, single address):

- `CMAccountCreated(account)` → category **Accounts**
- `ServiceRegistered(serviceName, serviceHash)` → category **Services**
- `ServiceUnregistered(serviceName, serviceHash)` → category **Services**

From the **BookingToken** (single address):

- `TokenReserved(...)` → category **Bookings**
- `TokenBought(tokenId, buyer)` → category **Bookings**
- `TokenReservationExpired(tokenId)` → category **Bookings**
- `CancellationPending(...)` → category **Cancellations**
- `CancellationFinalized(tokenId)` → category **Cancellations**
- `CancellationWithdrawn(...)` → category **Cancellations**
- `CancellationRejected(...)` → category **Cancellations**

### Account-level (everything) — account detail tab only

All emitted by the single CM Account proxy address (covers `CMAccount`,
`PartnerConfiguration`, `GasMoneyManager`, and the account's `ServiceRegistry`
usage), so a single `getLogs` per account captures them:

- `MessengerBotAdded` / `MessengerBotRemoved`
- `ServiceAdded` / `ServiceRemoved` / `WantedServiceAdded` / `WantedServiceRemoved`
- `ServiceRestrictedRateUpdated` / `ServiceCapabilitiesUpdated` /
  `ServiceCapabilityAdded` / `ServiceCapabilityRemoved`
- `PaymentTokenAdded` / `PaymentTokenRemoved`
- `OffChainPaymentSupportUpdated`
- `PublicKeyAdded` / `PublicKeyRemoved`
- `Deposit` / `Withdraw`
- `GasMoneyWithdrawal` / `GasMoneyWithdrawalUpdated`
- `CMAccountUpgraded`

### Filter categories (Activity page chips)

`Bookings`, `Cancellations`, `Accounts`, `Services`. All shown by default;
chips toggle visibility client-side over already-loaded events.

## Normalized event shape

```ts
interface ActivityEvent {
  id: string;            // `${txHash}#${logIndex}` — stable dedupe/React key
  category: ActivityCategory;
  source: "manager" | "bookingToken" | "account";
  contract: Address;     // emitting contract address
  blockNumber: bigint;
  timestamp?: number;    // unix seconds, filled by useBlockTimestamps
  txHash: Hex;
  args: Record<string, unknown>;
  sentence: string;      // from catalog render()
}
```

## RPC layer & pagination (decided)

### Cursor pagination via `useInfiniteQuery`

The feed is **not** a fixed window. It is cursor pagination going backwards:

- **Page 0** = the latest batch: `[latest - batch + 1, latest]`.
- Each subsequent page = the next older batch: `getNextPageParam` returns
  `oldestBlockFetchedSoFar - 1n` as the new `toBlock`; `fromBlock` is
  `toBlock - batch + 1`, floored at `0n`.
- Pages accumulate; the UI renders all loaded events, sorted by
  `(blockNumber, logIndex)` descending.
- A **"Load older"** button advances one page per click (TanStack Query
  `fetchNextPage`). Disabled once `fromBlock` reaches `0n`.

Each page fetch issues one `getLogs({ address, events, fromBlock, toBlock })`
**per distinct address** in scope (manager + bookingToken for the ecosystem feed;
the single account for the account tab), decodes via the catalog, and merges.

Queried through `usePublicClient()` (viem), keyed by `activeChainId`, so the
existing `RefreshButton`/invalidation refreshes it. No polling timers.

### Adaptive batch sizing (decided)

- Per-chain constant `ACTIVITY_BATCH_BLOCKS` (default `10000n`) is the starting
  batch size — the known-safe range for most public RPCs.
- **On a `getLogs` rejection** (range/limit error), halve the range and retry the
  same batch; repeat until it succeeds or the range hits a floor of `500n`. If
  even the floor fails, surface an inline error with retry.
- **Session memory:** the last range size that *succeeded* for a chain is cached
  (module-level, per `chainId`) and used as the starting size for subsequent
  batches, so once an RPC reveals a smaller real cap we don't re-fail from 10k
  each time. This is the lightweight form of "discover and act accordingly" —
  no upfront probing.

### Timestamps

- `useBlockTimestamps(chainId, blockNumbers)` dedupes the block numbers present in
  loaded events, fetches each via `getBlock`, and caches the result per
  `(chainId, blockNumber)` so the Dashboard card, Activity page, and account tab
  never refetch the same block.
- Rendered with a new `formatRelativeTime(unixSeconds)` helper in `lib/format.ts`.

## Hooks

- `useActivity({ sources, chainId })` — core infinite-query hook. `sources` is a
  list of `{ address, events }` (one per distinct contract). Handles batching,
  adaptive sizing, decode, merge, sort. Returns `{ events, fetchNextPage,
  hasNextPage, isLoading, isFetchingNextPage, error, oldestBlockLoaded }`.
- `useEcosystemActivity()` — composes `useActivity` over manager + bookingToken
  with their ecosystem catalog subsets.
- `useAccountActivity(address)` — composes `useActivity` over the single account
  address with the full account catalog subset.
- `useBlockTimestamps(chainId, blockNumbers)` — as above.

## UI surfaces

- **Dashboard** — new `RecentActivityCard` using `useEcosystemActivity`, rendering
  the latest **5** events only (no "Load older" here — always 1 batch, stays cheap),
  with a "View all →" link to `/activity`. Slots into the existing card grid.
- **Activity page** — new `pages/Activity.tsx` at route `/activity` (registered in
  `App.tsx`). Reverse-chron `ActivityList` + category filter chips + "Load older"
  button showing the loaded range (e.g. "showing last ~30,000 blocks").
- **Account detail** — new `pages/tabs/ActivityTab.tsx` added to `AccountWorkspace`
  tabs, using `useAccountActivity(address)` + "Load older".
- **Shared components** — `components/activity/ActivityList.tsx` and
  `ActivityRow.tsx` (icon · sentence · relative time · tx link via existing
  `explorerTxUrl`). Addresses rendered with existing `AddressDisplay`.

## Error & empty states

- A `getLogs` batch that fails even at the `500n` floor renders an inline
  "Couldn't load activity from this RPC" with a retry action. It must **not**
  crash the Dashboard card — the card degrades to the error state in place.
- Empty result for the loaded range → "No activity in the last 10,000 blocks."
  (the account tab uses "No activity for this account in the last 10,000 blocks.").
  The "Load older" control stays visible below so the user can page further back,
  making the rolling-window nature explicit.

## Testing (Vitest, existing patterns)

- Catalog: each entry's `render(args)` produces the expected sentence.
- `useActivity`: merge/sort ordering and pagination cursor math with a mocked
  `publicClient`; adaptive-halving fallback (mock `getLogs` to reject at 10k,
  succeed at 5k) and session-size memory.
- `useBlockTimestamps`: dedupe + per-block cache (no duplicate `getBlock`).
- Filter-chip logic over a fixed event set.
- `ActivityRow` rendering. Components rendered with `QueryClientProvider`
  (and mocked wagmi) per existing `*.test.tsx`.

## Files

New:

- `ui/src/lib/activity/catalog.ts`
- `ui/src/lib/activity/types.ts` (the `ActivityEvent` / category types)
- `ui/src/hooks/useActivity.ts`
- `ui/src/hooks/useBlockTimestamps.ts`
- `ui/src/pages/Activity.tsx`
- `ui/src/pages/tabs/ActivityTab.tsx`
- `ui/src/components/activity/ActivityList.tsx`
- `ui/src/components/activity/ActivityRow.tsx`
- Tests alongside each per existing convention.

Changed:

- `ui/src/App.tsx` — register `/activity` route.
- `ui/src/pages/Dashboard.tsx` — add `RecentActivityCard`.
- `ui/src/pages/AccountWorkspace.tsx` — add Activity tab.
- `ui/src/config/chains.ts` (or a new `config/activity.ts`) —
  `ACTIVITY_BATCH_BLOCKS` per chain.
- `ui/src/lib/format.ts` — add `formatRelativeTime`.
- Navigation (`components/Layout` or nav) — link to `/activity`.
