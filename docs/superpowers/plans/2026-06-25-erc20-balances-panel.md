# ERC20 Balances Panel Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Show configurable ERC20 token balances for the active CM Account in the Active account panel, with zero-balance warnings for native and ERC20 tokens.

**Architecture:** A curated per-chain address config (`config/tokens.ts`) is merged with the account's on-chain `getSupportedTokens()`, deduped, and resolved on-chain via a single multicall (`symbol`/`decimals`/`balanceOf`) inside a new `useErc20Balances` hook. `AccountSummary` consumes the hook to render balance rows and amber zero-balance warnings.

**Tech Stack:** React 18 + TypeScript, wagmi v2 (`useReadContracts`, `useReadContract`, `useBalance`), viem v2 (`formatUnits`), Vitest + Testing Library. All commands run from `ui/`.

## Global Constraints

- All UI work happens in `ui/`; run commands from there. If `yarn test` resolves to the root Hardhat project, call `./node_modules/.bin/vitest run <path>` directly.
- Reads must pass `chainId: activeChainId` (from `useActiveChain`/`useActiveContracts`), per the read/write convention.
- Dark mode is Tailwind `class` strategy — every color needs a `dark:` variant.
- Don't hand-edit `src/contracts/generated/`. Don't add new dependencies.
- Reuse existing components: `AddressDisplay`, and the amber `AlertTriangle` warning styling already used in `BotsTab`.
- Tests for components using wagmi/query need a `QueryClientProvider` wrapper and mocked `wagmi` + `../wallet/activeChain` (see `RolesPanel.test.tsx`).

---

### Task 1: ERC20 ABI fragment + token config

**Files:**
- Create: `ui/src/lib/erc20.ts`
- Create: `ui/src/config/tokens.ts`

**Interfaces:**
- Produces: `ERC20_ABI` (viem `Abi` with `balanceOf(address)→uint256`, `symbol()→string`, `decimals()→uint8`); `EXTRA_TOKENS: Record<number, Address[]>`.

- [ ] **Step 1: Write `ui/src/lib/erc20.ts`**

```ts
import { type Abi } from "viem";

/** Minimal ERC20 fragments — only what the balances panel reads. */
export const ERC20_ABI = [
  {
    type: "function",
    name: "balanceOf",
    stateMutability: "view",
    inputs: [{ name: "account", type: "address" }],
    outputs: [{ name: "", type: "uint256" }],
  },
  {
    type: "function",
    name: "symbol",
    stateMutability: "view",
    inputs: [],
    outputs: [{ name: "", type: "string" }],
  },
  {
    type: "function",
    name: "decimals",
    stateMutability: "view",
    inputs: [],
    outputs: [{ name: "", type: "uint8" }],
  },
] as const satisfies Abi;
```

- [ ] **Step 2: Write `ui/src/config/tokens.ts`**

```ts
import { type Address } from "viem";

// Curated ERC20 token addresses to always display balances for, per chainId.
// Only addresses live here — symbol/decimals/balance are read on-chain.
// Merged with each account's on-chain getSupportedTokens() at render time.
export const EXTRA_TOKENS: Record<number, Address[]> = {
  500: [], // Camino
  8453: [], // Base
  84532: [], // Base Sepolia
};
```

- [ ] **Step 3: Typecheck**

Run: `cd ui && yarn tsc -b --noEmit` (or `./node_modules/.bin/tsc -b --noEmit`)
Expected: no errors.

- [ ] **Step 4: Commit**

```bash
git add ui/src/lib/erc20.ts ui/src/config/tokens.ts
git commit -m "feat(ui): add ERC20 ABI fragment and curated token config"
```

---

### Task 2: `useErc20Balances` hook

**Files:**
- Create: `ui/src/hooks/useErc20Balances.ts`
- Test: `ui/src/hooks/useErc20Balances.test.tsx`

**Interfaces:**
- Consumes: `ERC20_ABI`, `EXTRA_TOKENS` (Task 1); `useActiveContracts` (`chainId`, `cmAccountAbi`); wagmi `useReadContract`, `useReadContracts`.
- Produces:
  ```ts
  interface TokenBalance {
    address: Address;
    symbol: string;
    decimals: number;
    balance: bigint;
    formatted: string;
    isZero: boolean;
  }
  function useErc20Balances(account: Address): { tokens: TokenBalance[]; isLoading: boolean };
  ```

- [ ] **Step 1: Write the failing test**

```tsx
import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";

let mockSupported: unknown;
let mockMulticall: { data: unknown; isLoading: boolean };
const warnSpy = vi.fn();

vi.mock("wagmi", () => ({
  useReadContract: () => ({ data: mockSupported, isLoading: false }),
  useReadContracts: () => mockMulticall,
}));
vi.mock("../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({ chainId: 84532, cmAccountAbi: [] }),
}));
vi.mock("../config/tokens", () => ({ EXTRA_TOKENS: { 84532: ["0xAAaA000000000000000000000000000000000001"] } }));

import { useErc20Balances } from "./useErc20Balances";

const account = "0x1111111111111111111111111111111111111111" as const;
// 3 calls per token: symbol, decimals, balanceOf
const ok = (s: string, d: number, b: bigint) => [
  { status: "success", result: s },
  { status: "success", result: d },
  { status: "success", result: b },
];

afterEach(() => {
  mockSupported = undefined;
  mockMulticall = { data: undefined, isLoading: false };
  warnSpy.mockClear();
});

describe("useErc20Balances", () => {
  it("returns formatted balances and isZero flags", () => {
    mockSupported = [];
    mockMulticall = { data: ok("USDC", 6, 1500000n), isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
    expect(result.current.tokens[0]).toMatchObject({ symbol: "USDC", decimals: 6, formatted: "1.5", isZero: false });
  });

  it("flags zero balances", () => {
    mockSupported = [];
    mockMulticall = { data: ok("USDC", 6, 0n), isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens[0].isZero).toBe(true);
  });

  it("dedupes a config token that is also a supported token", () => {
    mockSupported = ["0xaaaa000000000000000000000000000000000001"]; // same as config, different case
    mockMulticall = { data: ok("USDC", 6, 1n), isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
  });

  it("drops a non-ERC20 (balanceOf failure) and warns in dev", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockSupported = [];
    mockMulticall = {
      data: [
        { status: "failure", error: new Error("x") },
        { status: "failure", error: new Error("x") },
        { status: "failure", error: new Error("x") },
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
  });

  it("keeps token with symbol/decimals fallback when only those fail", () => {
    mockSupported = [];
    mockMulticall = {
      data: [
        { status: "failure", error: new Error("x") }, // symbol
        { status: "failure", error: new Error("x") }, // decimals
        { status: "success", result: 5000000000000000000n }, // balanceOf
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
    expect(result.current.tokens[0]).toMatchObject({ decimals: 18, formatted: "5" });
    // symbol falls back to the shortened address (contains the ellipsis)
    expect(result.current.tokens[0].symbol).toContain("…");
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd ui && ./node_modules/.bin/vitest run src/hooks/useErc20Balances.test.tsx`
Expected: FAIL — `useErc20Balances` not found / module missing.

- [ ] **Step 3: Write minimal implementation**

```ts
import { type Address, formatUnits } from "viem";
import { useReadContract, useReadContracts } from "wagmi";
import { ERC20_ABI } from "../lib/erc20";
import { shortAddress } from "../lib/format";
import { EXTRA_TOKENS } from "../config/tokens";
import { useActiveContracts } from "./useActiveContracts";

export interface TokenBalance {
  address: Address;
  symbol: string;
  decimals: number;
  balance: bigint;
  formatted: string;
  isZero: boolean;
}

export function useErc20Balances(account: Address): { tokens: TokenBalance[]; isLoading: boolean } {
  const { chainId, cmAccountAbi } = useActiveContracts();

  // On-chain payment tokens the account supports.
  const { data: supportedRaw, isLoading: supportedLoading } = useReadContract({
    chainId,
    address: account,
    abi: cmAccountAbi,
    functionName: "getSupportedTokens",
  });
  const supported = ((supportedRaw as Address[] | undefined) ?? []).map((a) => a as Address);
  const configured = EXTRA_TOKENS[chainId] ?? [];

  // Merge + dedupe case-insensitively, keeping the first-seen casing.
  const seen = new Set<string>();
  const addresses: Address[] = [];
  for (const a of [...configured, ...supported]) {
    const key = a.toLowerCase();
    if (seen.has(key)) continue;
    seen.add(key);
    addresses.push(a);
  }

  const { data: multi, isLoading: multiLoading } = useReadContracts({
    allowFailure: true,
    contracts: addresses.flatMap((address) => [
      { chainId, address, abi: ERC20_ABI, functionName: "symbol" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "decimals" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "balanceOf", args: [account] } as const,
    ]),
    query: { enabled: addresses.length > 0 },
  });

  const tokens: TokenBalance[] = [];
  if (multi) {
    addresses.forEach((address, i) => {
      const symbolRes = multi[i * 3];
      const decimalsRes = multi[i * 3 + 1];
      const balanceRes = multi[i * 3 + 2];

      if (balanceRes?.status !== "success") {
        if (import.meta.env.DEV) {
          // eslint-disable-next-line no-console
          console.warn(`[useErc20Balances] ${address} is not a usable ERC20 on chain ${chainId} — dropping.`);
        }
        return;
      }

      const symbol = symbolRes?.status === "success" ? (symbolRes.result as string) : shortAddress(address);
      const decimals = decimalsRes?.status === "success" ? Number(decimalsRes.result) : 18;
      const balance = balanceRes.result as bigint;
      tokens.push({
        address,
        symbol,
        decimals,
        balance,
        formatted: formatUnits(balance, decimals),
        isZero: balance === 0n,
      });
    });
  }

  return { tokens, isLoading: supportedLoading || multiLoading };
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd ui && ./node_modules/.bin/vitest run src/hooks/useErc20Balances.test.tsx`
Expected: PASS (5 tests).

- [ ] **Step 5: Commit**

```bash
git add ui/src/hooks/useErc20Balances.ts ui/src/hooks/useErc20Balances.test.tsx
git commit -m "feat(ui): add useErc20Balances hook (merged config + on-chain tokens)"
```

---

### Task 3: Render balances + warnings in AccountSummary

**Files:**
- Modify: `ui/src/components/AccountSummary.tsx`
- Test: `ui/src/components/AccountSummary.test.tsx` (create)

**Interfaces:**
- Consumes: `useErc20Balances` → `{ tokens: TokenBalance[]; isLoading }` (Task 2); existing `useBalance`.

- [ ] **Step 1: Write the failing test**

```tsx
import { afterEach, describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { AccountSummary } from "./AccountSummary";
import type { TokenBalance } from "../hooks/useErc20Balances";

let mockNative: { formatted: string; symbol: string } | undefined;
let mockTokens: TokenBalance[];

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useBalance: () => ({ data: mockNative }),
}));
vi.mock("../hooks/useActiveContracts", () => ({
  useActiveContracts: () => ({ chainId: 84532 }),
}));
vi.mock("../hooks/useMyAccounts", () => ({ useAccountRolesFor: () => [] }));
vi.mock("../hooks/useErc20Balances", () => ({ useErc20Balances: () => ({ tokens: mockTokens, isLoading: false }) }));

const account = "0x1111111111111111111111111111111111111111" as const;
const tok = (over: Partial<TokenBalance>): TokenBalance => ({
  address: "0xAAaA000000000000000000000000000000000001",
  symbol: "USDC",
  decimals: 6,
  balance: 1500000n,
  formatted: "1.5",
  isZero: false,
  ...over,
});

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

afterEach(() => {
  mockNative = { formatted: "1.0", symbol: "ETH" };
  mockTokens = [];
});

describe("AccountSummary ERC20 balances", () => {
  it("renders a row per token with formatted amount and symbol", () => {
    mockNative = { formatted: "1.0", symbol: "ETH" };
    mockTokens = [tok({})];
    wrap(<AccountSummary account={account} />);
    expect(screen.getByText(/1\.5 USDC/)).toBeInTheDocument();
  });

  it("warns when an ERC20 balance is zero", () => {
    mockNative = { formatted: "1.0", symbol: "ETH" };
    mockTokens = [tok({ balance: 0n, formatted: "0", isZero: true })];
    wrap(<AccountSummary account={account} />);
    expect(screen.getByText(/can't buy booking tokens paid in USDC/i)).toBeInTheDocument();
  });

  it("warns when native balance is zero", () => {
    mockNative = { formatted: "0", symbol: "ETH" };
    mockTokens = [];
    wrap(<AccountSummary account={account} />);
    expect(screen.getByText(/no .* for gas|can't pay gas/i)).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd ui && ./node_modules/.bin/vitest run src/components/AccountSummary.test.tsx`
Expected: FAIL — token text / warnings not rendered.

- [ ] **Step 3: Modify `AccountSummary.tsx`**

Add imports at the top:

```tsx
import { AlertTriangle } from "lucide-react";
import { useErc20Balances } from "../hooks/useErc20Balances";
```

Inside the component, after the existing `bal` line, add:

```tsx
  const { tokens } = useErc20Balances(account);
  const nativeZero = bal ? Number(bal.value) === 0 : false;
```

Replace the existing native balance `<div>` and add the token section. The native
balance block becomes:

```tsx
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Native balance</dt>
          <dd>{bal ? `${bal.formatted} ${bal.symbol}` : "—"}</dd>
          {nativeZero && (
            <p className="mt-1 flex items-start gap-1 text-xs text-amber-700 dark:text-amber-300">
              <AlertTriangle className="mt-0.5 h-3 w-3 shrink-0" />
              No {bal?.symbol} for gas — this account can't pay gas to buy booking tokens.
            </p>
          )}
        </div>
        {tokens.length > 0 && (
          <div>
            <dt className="mb-1 text-gray-500 dark:text-gray-400">Token balances</dt>
            <dd className="flex flex-col gap-2">
              {tokens.map((t) => (
                <div key={t.address}>
                  <div className="flex items-center justify-between gap-2">
                    <AddressDisplay address={t.address} className="text-xs" />
                    <span className={t.isZero ? "text-amber-700 dark:text-amber-300" : ""}>
                      {t.formatted} {t.symbol}
                    </span>
                  </div>
                  {t.isZero && (
                    <p className="mt-0.5 flex items-start gap-1 text-xs text-amber-700 dark:text-amber-300">
                      <AlertTriangle className="mt-0.5 h-3 w-3 shrink-0" />
                      0 {t.symbol} — can't buy booking tokens paid in {t.symbol}.
                    </p>
                  )}
                </div>
              ))}
            </dd>
          </div>
        )}
```

Note: `useBalance` returns `value` (bigint) and `formatted`. If `AddressDisplay`
does not accept a `className` prop, drop the prop (check `AddressDisplay.tsx`).

- [ ] **Step 4: Run test to verify it passes**

Run: `cd ui && ./node_modules/.bin/vitest run src/components/AccountSummary.test.tsx`
Expected: PASS (3 tests).

- [ ] **Step 5: Lint + typecheck**

Run: `cd ui && yarn lint` and `./node_modules/.bin/tsc -b --noEmit`
Expected: no errors. If `yarn format` is needed, run it (per project convention, format with `yarn format`, not raw prettier).

- [ ] **Step 6: Commit**

```bash
git add ui/src/components/AccountSummary.tsx ui/src/components/AccountSummary.test.tsx
git commit -m "feat(ui): show ERC20 balances and zero-balance warnings in account panel"
```

---

### Task 4: Full suite verification

**Files:** none (verification only)

- [ ] **Step 1: Run the full UI test suite**

Run: `cd ui && ./node_modules/.bin/vitest run`
Expected: all tests pass, including the new hook + component tests.

- [ ] **Step 2: Build check**

Run: `cd ui && yarn build`
Expected: `tsc -b` + `vite build` succeed (sync runs via `prebuild`).

- [ ] **Step 3: Manual smoke (optional, if a dev wallet is available)**

Run: `cd ui && yarn dev`, open an account workspace, confirm the Active account
panel shows token balances and a warning appears for any zero balance.

---

## Self-Review Notes

- **Spec coverage:** config file (Task 1), ERC20 ABI (Task 1), merged+deduped source & multicall & drop/fallback handling (Task 2), panel rows & native/ERC20 zero warnings (Task 3), tests (Tasks 2–3), suite verification (Task 4). All spec sections covered.
- **Types:** `TokenBalance` and `useErc20Balances` signatures are consistent across Tasks 2 and 3.
- **No new deps:** uses existing wagmi/viem/lucide-react.
- **Open config item:** `EXTRA_TOKENS` ships empty per the spec; pre-populating known addresses (e.g. USDC) is a later one-line edit.
