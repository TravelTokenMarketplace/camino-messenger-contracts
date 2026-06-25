# ERC20 balances in the Active account panel

**Date:** 2026-06-25
**Area:** `ui/` (React management UI)

## Problem

When a CM Account has no funds, the messenger bot's booking-token buy reverts
("execution reverted"). Today the **Active account** panel (`AccountSummary`,
shown on the left of the account workspace) only displays the account's **native**
balance, so an operator cannot tell at a glance whether the account is funded with
the ERC20 token it pays with.

We want the panel to also show selected **ERC20 token balances** for the CM Account,
and to warn the operator when any balance (native or ERC20) is zero — because a zero
balance means the account cannot buy booking tokens with that coin/token.

## Goals

- Show ERC20 token balances for the active CM Account in the Active account panel.
- The set of tokens is configurable and lives in one organized place.
- Symbols and balances are fetched on-chain (no hand-maintained metadata).
- When a balance is zero (native **or** any ERC20), warn the operator that the
  account cannot buy booking tokens with that token/coin.

## Non-goals

- No runtime add/remove of display tokens (no localStorage editor). Configuration
  is a committed code change.
- No transfers / funding flows from this panel — display + warning only.
- No changes to Solidity contracts or ABIs.

## Token source (decided)

The displayed token set is the **merge** of two sources, deduped case-insensitively
on address:

1. **Curated static config** — `src/config/tokens.ts`, keyed by chainId (mirrors
   `src/config/chains.ts`). Holds only addresses; metadata is read on-chain.
2. **The account's on-chain supported payment tokens** — `getSupportedTokens()` on
   the CMAccount. These are exactly the ERC20s the account can pay with, so they are
   the most directly relevant to "can it buy booking tokens".

The static config starts **empty** (per-chain comments as placeholders); the operator
fills in curated addresses later. The on-chain supported tokens populate the panel
regardless.

## Components

### `src/config/tokens.ts`

```ts
import { type Address } from "viem";

// Curated ERC20 token addresses to always display balances for, per chainId.
// Symbol/decimals/balance are read on-chain — only addresses live here.
export const EXTRA_TOKENS: Record<number, Address[]> = {
  500: [],    // Camino
  8453: [],   // Base
  84532: [],  // Base Sepolia
};
```

### `src/lib/erc20.ts`

A minimal ERC20 ABI fragment (`balanceOf`, `symbol`, `decimals`) — the project keeps
ABIs lean; the full ERC20 ABI is not needed here.

### `src/hooks/useErc20Balances.ts`

`useErc20Balances(account: Address)` returns
`{ tokens: TokenBalance[]; isLoading: boolean }` where

```ts
interface TokenBalance {
  address: Address;
  symbol: string;
  decimals: number;
  balance: bigint;
  formatted: string;
  isZero: boolean;
}
```

Behavior:

- Token set = `EXTRA_TOKENS[activeChainId]` merged with the account's
  `getSupportedTokens()` (read via the existing `useContractList` pattern / a
  `useReadContract`), deduped case-insensitively (compare lowercased addresses).
- One `useReadContracts({ allowFailure: true, chainId: activeChainId })` multicall
  reading `symbol`, `decimals`, `balanceOf(account)` for each address.
- Per-token result handling:
  - **`balanceOf` or `decimals` fails** → not a safely usable ERC20 → **drop**
    from the result, and `console.warn` in dev (`import.meta.env.DEV`) so a bad
    address in `tokens.ts` is catchable. (`decimals` is required: without a
    trustworthy scale the balance would be misformatted — e.g. a 6-decimal token
    rendered as 18 — so we drop rather than guess.)
  - **`balanceOf` + `decimals` ok but `symbol` fails** → keep the token with a
    cosmetic fallback: `symbol = shortAddress(address)`.
- Reads are keyed by `activeChainId` per the project's read/write convention.

### `src/components/AccountSummary.tsx` (modified)

- Keep the existing native balance row.
- Add a **Token balances** section below it listing each `TokenBalance` as
  `formatted symbol`, with `AddressDisplay` for the contract address. While loading,
  show a muted placeholder; if the merged set is empty, render nothing extra.
- **Zero-balance warnings** (reuse the amber `AlertTriangle` styling from `BotsTab`):
  - Native balance is `0` → amber warning: the account has no native coin for gas and
    cannot pay to buy booking tokens.
  - Each ERC20 with `isZero` → amber/info line:
    `"0 SYMBOL — can't buy booking tokens paid in SYMBOL."`

## Data flow

```
AccountSummary(account)
  ├─ useBalance(account)                      → native balance (existing)
  └─ useErc20Balances(account)
       ├─ EXTRA_TOKENS[chainId]               (static config)
       ├─ account.getSupportedTokens()        (on-chain)
       ├─ merge + dedupe (lowercased)
       └─ multicall symbol/decimals/balanceOf → TokenBalance[]
```

## Error handling

- Multicall uses `allowFailure: true`; failures are handled per-token (drop or
  fallback) as above — a single bad address never breaks the panel.
- Empty token set → panel shows only the native balance, exactly as today.

## Testing

A render test of `AccountSummary` (mocked wagmi: `useBalance` + `useReadContracts`,
plus a `QueryClientProvider` per the project's test convention) covering:

- Normal balances: ERC20 rows render with symbol + formatted amount.
- Zero ERC20 balance → the per-token warning is shown.
- Zero native balance → the native warning is shown.
- Dedup: a config token that is also a supported token appears once.
- Non-ERC20 address (`balanceOf` failure) → dropped, not rendered.

## Out of scope / follow-ups

- Pre-populating `EXTRA_TOKENS` with known addresses (e.g. USDC on Base) can be a
  later trivial config edit.
