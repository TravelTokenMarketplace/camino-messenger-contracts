# Human-readable token display & numeric typography

Date: 2026-06-25
Scope: `ui/` only

## Problem

The Payment Tokens tab and the Active Account left panel surface ERC20 tokens by
**raw contract address only**, which is not human-readable. The left panel does
show a symbol for tokens it has balances for, but the address dominates the row.
Fund/balance numbers use the default proportional font and show excessive
precision (e.g. `0.00200000000063198 ETH`).

## Goals

1. Show each token as **symbol (bold headline) + name (muted secondary)**, with
   the address compacted (`0x29F3…db85`), smaller/muted, copyable, and full on hover.
2. Apply a dedicated fixed-width font (**JetBrains Mono**, bundled via Fontsource)
   with tabular figures to all fund/balance values.
3. Trim over-long amounts to ~6 significant decimals for display, full precision on hover.

These apply to Payment Tokens + the left panel, and (approved) ripple to the
other amount displays (Withdrawals, BookingToken) for consistency.

## Design

### 1. `useTokenMetadata(addresses)` hook — `ui/src/hooks/useTokenMetadata.ts`

- Input: `Address[]` (deduped case-insensitively).
- Multicalls `symbol` / `name` / `decimals` (existing `ERC20_ABI` gains a `name`
  fragment) via `useReadContracts({ allowFailure: true })`, keyed by `chainId`.
- Returns `{ meta: Map<lowercaseAddress, { address, symbol?, name?, decimals? }>, isLoading }`.
- Per-token graceful fallback: a failed `symbol`/`name` is simply absent; consumers
  fall back to the compact address as headline.

### 2. `useErc20Balances` refactor — reuse the metadata hook

- Delegate `symbol`/`name`/`decimals` resolution to `useTokenMetadata`; this hook
  keeps the `getSupportedTokens` + `EXTRA_TOKENS` merge and layers `balanceOf` on top.
- `TokenBalance` gains an optional `name?: string`.
- Existing drop rule unchanged: require trustworthy `decimals` + `balanceOf`, else drop.

### 3. `TokenDisplay` component — `ui/src/components/TokenDisplay.tsx`

Props: `{ address, symbol?, name?, className? }`. Layout:

```
[identicon] EURe                 ← symbol, bold (or compact address if no symbol)
            Monerium EUR emoney  ← name, text-xs muted (omitted if absent)
            0x29F3…db85  ⧉       ← shortAddress, text-xs muted, copy button
```

- Full address shown via the existing `Tooltip` wrapping the compact address.
- Copy button reuses the existing `CopyButton` (or AddressDisplay's copy logic).
- Replaces `<AddressDisplay>` in `PaymentTokensTab` rows and the token rows in `AccountSummary`.

### 4. Numeric typography — JetBrains Mono

- Add dependency `@fontsource/jetbrains-mono`; import its CSS in `src/main.tsx`.
- Tailwind `theme.extend.fontFamily.num = ['"JetBrains Mono"', 'monospace']`.
- New utility/class applies `font-num` + `tabular-nums`. Apply to: native balance,
  token balances (left panel), payment-token list is labels not numbers (no change),
  Withdrawals amounts, BookingToken amounts.
- Addresses keep `font-mono` (unchanged).

### 5. Amount formatting — `ui/src/lib/format.ts`

- `formatAmount(value: string, opts?)`: trims a decimal string to ~6 significant
  fractional digits, adds thousands separators to the integer part; returns
  `{ display, full }` (or two helpers). Full precision used in a hover tooltip.
- Used by native balance and token balances; the raw `formatUnits` string is the `full`.

## Components touched

- New: `hooks/useTokenMetadata.ts`, `components/TokenDisplay.tsx`.
- Edit: `lib/erc20.ts` (+`name`), `hooks/useErc20Balances.ts`, `lib/format.ts`
  (+`formatAmount`), `components/AccountSummary.tsx`, `pages/tabs/PaymentTokensTab.tsx`,
  `tailwind.config.js`, `src/main.tsx`, `package.json`.
- Consistency ripple: `tabs/WithdrawalsTab.tsx`, `tabs/BookingTokenTab.tsx` (number font + `formatAmount`).

## Testing

- `useTokenMetadata`: dedupe + fallback on failed calls (Vitest + mocked wagmi).
- `formatAmount`: trimming, separators, full-precision passthrough (pure unit test).
- `TokenDisplay`: renders symbol/name/compact address; address-only fallback.
- Update existing `AccountSummary.test.tsx` for the new row structure.

## Out of scope

- Token logos/registry beyond on-chain `symbol`/`name`.
- Changing which tokens are listed (`EXTRA_TOKENS` / `getSupportedTokens` logic).
