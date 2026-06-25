# Human-readable Token Display & Numeric Typography Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Show ERC20 tokens by symbol + name (address compact/on-hover) in the Payment Tokens tab and the Active Account panel, and render fund values in a dedicated tabular monospace font.

**Architecture:** A reusable `useTokenMetadata` hook multicalls `symbol`/`name`/`decimals`; `useErc20Balances` builds on it. A shared `TokenDisplay` component renders symbol/name/compact-address consistently in both the tab and the left panel. JetBrains Mono (bundled via Fontsource) + `tabular-nums` style all balances, with a `formatAmount` helper trimming over-long values.

**Tech Stack:** React 18, TypeScript, wagmi v2 + viem v2, Tailwind (class dark mode), Vitest + Testing Library, `@fontsource/jetbrains-mono`.

## Global Constraints

- All commands run from `ui/`. Run tests with `./node_modules/.bin/vitest run <path>` (root `yarn test` runs Hardhat).
- Reads pass `chainId` from `useActiveContracts()`; never use `eth_getLogs`.
- Every color needs a `dark:` variant (Tailwind `class` dark mode).
- Addresses keep `font-mono`; only numeric/fund values use the new `font-num`.
- No CDN fonts — bundle via Fontsource.
- Commit messages: plain, no Co-Authored-By trailer.

---

### Task 1: `formatAmount` helper

**Files:**
- Modify: `ui/src/lib/format.ts`
- Test: `ui/src/lib/format.test.ts`

**Interfaces:**
- Produces: `formatAmount(value: string, sigFractionDigits?: number): { display: string; full: string }` — `display` groups the integer part with thousands separators and keeps up to `sigFractionDigits` (default 6) significant fractional digits with trailing zeros stripped; `full` is the input verbatim.

- [ ] **Step 1: Write the failing test**

Append to `ui/src/lib/format.test.ts`:

```ts
import { formatAmount } from "./format";

describe("formatAmount", () => {
  it("trims long fractional precision to significant digits", () => {
    const { display, full } = formatAmount("0.00200000000063198");
    expect(display).toBe("0.002");
    expect(full).toBe("0.00200000000063198");
  });

  it("groups the integer part with thousands separators", () => {
    expect(formatAmount("1850").display).toBe("1,850");
    expect(formatAmount("1234567.5").display).toBe("1,234,567.5");
  });

  it("leaves short values unchanged", () => {
    expect(formatAmount("1.5").display).toBe("1.5");
    expect(formatAmount("0").display).toBe("0");
  });

  it("keeps significant digits for very small numbers", () => {
    expect(formatAmount("0.000001234567").display).toBe("0.000001234567");
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/lib/format.test.ts`
Expected: FAIL — `formatAmount is not a function`.

- [ ] **Step 3: Add the implementation**

Append to `ui/src/lib/format.ts`:

```ts
/**
 * Format a decimal amount string for display: thousands separators on the
 * integer part, and the fractional part trimmed to `sigFractionDigits`
 * significant digits (skipping leading zeros) with trailing zeros stripped.
 * Returns the trimmed `display` and the untouched `full` value (for tooltips).
 */
export function formatAmount(value: string, sigFractionDigits = 6): { display: string; full: string } {
  const full = value;
  const neg = value.startsWith("-");
  const v = neg ? value.slice(1) : value;
  const [intRaw, fracRaw = ""] = v.split(".");
  const intGrouped = intRaw.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  let frac = fracRaw;
  if (frac) {
    let lead = 0;
    while (lead < frac.length && frac[lead] === "0") lead++;
    frac = frac.slice(0, lead + sigFractionDigits).replace(/0+$/, "");
  }
  const display = (neg ? "-" : "") + intGrouped + (frac ? "." + frac : "");
  return { display, full };
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `./node_modules/.bin/vitest run src/lib/format.test.ts`
Expected: PASS (all cases).

- [ ] **Step 5: Commit**

```bash
git add ui/src/lib/format.ts ui/src/lib/format.test.ts
git commit -m "feat(ui): add formatAmount helper for fund values"
```

---

### Task 2: `name` ERC20 fragment + `useTokenMetadata` hook

**Files:**
- Modify: `ui/src/lib/erc20.ts`
- Create: `ui/src/hooks/useTokenMetadata.ts`
- Test: `ui/src/hooks/useTokenMetadata.test.tsx`

**Interfaces:**
- Consumes: `ERC20_ABI` (gains a `name` view fragment), `useActiveContracts().chainId`.
- Produces:
  - `interface TokenMeta { address: Address; symbol?: string; name?: string; decimals?: number }`
  - `useTokenMetadata(addresses: Address[]): { meta: Map<string, TokenMeta>; isLoading: boolean }` — `meta` keyed by lowercase address; addresses deduped case-insensitively; each field is `undefined` when its on-chain read fails.

- [ ] **Step 1: Add the `name` fragment to `ERC20_ABI`**

In `ui/src/lib/erc20.ts`, add after the `symbol` fragment:

```ts
  {
    type: "function",
    name: "name",
    stateMutability: "view",
    inputs: [],
    outputs: [{ name: "", type: "string" }],
  },
```

- [ ] **Step 2: Write the failing test**

Create `ui/src/hooks/useTokenMetadata.test.tsx`:

```tsx
import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import type { Address } from "viem";

let mockMulticall: { data: unknown; isLoading: boolean };

vi.mock("wagmi", () => ({ useReadContracts: () => mockMulticall }));
vi.mock("./useActiveContracts", () => ({ useActiveContracts: () => ({ chainId: 84532 }) }));

import { useTokenMetadata } from "./useTokenMetadata";

const A = "0xAAaA000000000000000000000000000000000001" as Address;
// 3 reads per token: symbol, name, decimals
const ok = (s: string, n: string, d: number) => [
  { status: "success", result: s },
  { status: "success", result: n },
  { status: "success", result: d },
];

afterEach(() => {
  mockMulticall = { data: undefined, isLoading: false };
});

describe("useTokenMetadata", () => {
  it("maps symbol/name/decimals keyed by lowercase address", () => {
    mockMulticall = { data: ok("EURe", "Monerium EUR emoney", 18), isLoading: false };
    const { result } = renderHook(() => useTokenMetadata([A]));
    expect(result.current.meta.get(A.toLowerCase())).toMatchObject({
      symbol: "EURe",
      name: "Monerium EUR emoney",
      decimals: 18,
    });
  });

  it("dedupes addresses case-insensitively", () => {
    mockMulticall = { data: ok("EURe", "Monerium", 18), isLoading: false };
    const { result } = renderHook(() => useTokenMetadata([A, A.toLowerCase() as Address]));
    expect(result.current.meta.size).toBe(1);
  });

  it("leaves a field undefined when its read fails", () => {
    mockMulticall = {
      data: [
        { status: "failure", error: new Error("x") }, // symbol
        { status: "success", result: "Token" }, // name
        { status: "success", result: 6 }, // decimals
      ],
      isLoading: false,
    };
    const { result } = renderHook(() => useTokenMetadata([A]));
    const m = result.current.meta.get(A.toLowerCase());
    expect(m?.symbol).toBeUndefined();
    expect(m?.name).toBe("Token");
    expect(m?.decimals).toBe(6);
  });
});
```

- [ ] **Step 3: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/hooks/useTokenMetadata.test.tsx`
Expected: FAIL — cannot resolve `./useTokenMetadata`.

- [ ] **Step 4: Implement the hook**

Create `ui/src/hooks/useTokenMetadata.ts`:

```ts
import { type Address } from "viem";
import { useReadContracts } from "wagmi";
import { ERC20_ABI } from "../lib/erc20";
import { useActiveContracts } from "./useActiveContracts";

export interface TokenMeta {
  address: Address;
  symbol?: string;
  name?: string;
  decimals?: number;
}

/**
 * Resolve ERC20 symbol/name/decimals for a list of token addresses via a single
 * multicall. Returns a map keyed by lowercase address; any field whose on-chain
 * read fails is left undefined (consumers fall back to the address).
 */
export function useTokenMetadata(addresses: Address[]): { meta: Map<string, TokenMeta>; isLoading: boolean } {
  const { chainId } = useActiveContracts();

  const seen = new Set<string>();
  const list: Address[] = [];
  for (const a of addresses) {
    const key = a.toLowerCase();
    if (seen.has(key)) continue;
    seen.add(key);
    list.push(a);
  }

  const { data, isLoading } = useReadContracts({
    allowFailure: true,
    contracts: list.flatMap((address) => [
      { chainId, address, abi: ERC20_ABI, functionName: "symbol" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "name" } as const,
      { chainId, address, abi: ERC20_ABI, functionName: "decimals" } as const,
    ]),
    query: { enabled: list.length > 0 },
  });

  const meta = new Map<string, TokenMeta>();
  if (data) {
    list.forEach((address, i) => {
      const s = data[i * 3];
      const n = data[i * 3 + 1];
      const d = data[i * 3 + 2];
      meta.set(address.toLowerCase(), {
        address,
        symbol: s?.status === "success" ? (s.result as string) : undefined,
        name: n?.status === "success" ? (n.result as string) : undefined,
        decimals: d?.status === "success" ? Number(d.result) : undefined,
      });
    });
  }

  return { meta, isLoading };
}
```

- [ ] **Step 5: Run test to verify it passes**

Run: `./node_modules/.bin/vitest run src/hooks/useTokenMetadata.test.tsx`
Expected: PASS.

- [ ] **Step 6: Commit**

```bash
git add ui/src/lib/erc20.ts ui/src/hooks/useTokenMetadata.ts ui/src/hooks/useTokenMetadata.test.tsx
git commit -m "feat(ui): add useTokenMetadata hook and ERC20 name fragment"
```

---

### Task 3: Refactor `useErc20Balances` onto `useTokenMetadata` (+ `name`)

**Files:**
- Modify: `ui/src/hooks/useErc20Balances.ts`
- Test: `ui/src/hooks/useErc20Balances.test.tsx` (rewrite)

**Interfaces:**
- Consumes: `useTokenMetadata` (Task 2), `ERC20_ABI`, `EXTRA_TOKENS`, `useActiveContracts`.
- Produces: `TokenBalance` gains `name?: string`. Drop rule unchanged: a token is dropped unless `meta.decimals` is defined **and** its `balanceOf` read succeeds. `symbol` falls back to `shortAddress(address)` when metadata lacks it.

- [ ] **Step 1: Rewrite the test**

Replace the whole contents of `ui/src/hooks/useErc20Balances.test.tsx`:

```tsx
import { afterEach, describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import type { Address } from "viem";
import type { TokenMeta } from "./useTokenMetadata";

let mockSupported: unknown;
let mockBalances: { data: unknown; isLoading: boolean };
let mockMeta: Map<string, TokenMeta>;
const warnSpy = vi.fn();

vi.mock("wagmi", () => ({
  useReadContract: () => ({ data: mockSupported, isLoading: false }),
  useReadContracts: () => mockBalances,
}));
vi.mock("./useTokenMetadata", () => ({ useTokenMetadata: () => ({ meta: mockMeta, isLoading: false }) }));
vi.mock("./useActiveContracts", () => ({ useActiveContracts: () => ({ chainId: 84532, cmAccountAbi: [] }) }));
vi.mock("../config/tokens", () => ({ EXTRA_TOKENS: { 84532: ["0xAAaA000000000000000000000000000000000001"] } }));

import { useErc20Balances } from "./useErc20Balances";

const account = "0x1111111111111111111111111111111111111111" as const;
const A = "0xAAaA000000000000000000000000000000000001";
const metaOf = (over: Partial<TokenMeta>): Map<string, TokenMeta> =>
  new Map([[A.toLowerCase(), { address: A as Address, symbol: "USDC", name: "USD Coin", decimals: 6, ...over }]]);
const bal = (b: bigint) => ({ data: [{ status: "success", result: b }], isLoading: false });

afterEach(() => {
  mockSupported = [];
  mockBalances = { data: undefined, isLoading: false };
  mockMeta = metaOf({});
  warnSpy.mockClear();
  vi.unstubAllGlobals();
});

describe("useErc20Balances", () => {
  it("returns formatted balances, name and isZero flags", () => {
    mockSupported = [];
    mockMeta = metaOf({});
    mockBalances = bal(1500000n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
    expect(result.current.tokens[0]).toMatchObject({ symbol: "USDC", name: "USD Coin", decimals: 6, formatted: "1.5", isZero: false });
  });

  it("flags zero balances", () => {
    mockBalances = bal(0n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens[0].isZero).toBe(true);
  });

  it("dedupes a config token that is also a supported token", () => {
    mockSupported = ["0xaaaa000000000000000000000000000000000001"];
    mockBalances = bal(1n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(1);
  });

  it("drops a token whose balanceOf read fails and warns in dev", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockBalances = { data: [{ status: "failure", error: new Error("x") }], isLoading: false };
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
    expect(warnSpy).toHaveBeenCalled();
  });

  it("falls back to a short address when symbol metadata is missing", () => {
    mockMeta = metaOf({ symbol: undefined });
    mockBalances = bal(5000000n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens[0].symbol).toContain("…");
  });

  it("drops a token when decimals metadata is missing (avoids misformatting)", () => {
    vi.stubGlobal("console", { ...console, warn: warnSpy });
    mockMeta = metaOf({ decimals: undefined });
    mockBalances = bal(1500000n);
    const { result } = renderHook(() => useErc20Balances(account));
    expect(result.current.tokens).toHaveLength(0);
    expect(warnSpy).toHaveBeenCalled();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/hooks/useErc20Balances.test.tsx`
Expected: FAIL — current hook does its own symbol/decimals multicall and has no `name`.

- [ ] **Step 3: Rewrite the hook**

Replace the whole contents of `ui/src/hooks/useErc20Balances.ts`:

```ts
import { type Address, formatUnits } from "viem";
import { useReadContract, useReadContracts } from "wagmi";
import { ERC20_ABI } from "../lib/erc20";
import { shortAddress } from "../lib/format";
import { EXTRA_TOKENS } from "../config/tokens";
import { useActiveContracts } from "./useActiveContracts";
import { useTokenMetadata } from "./useTokenMetadata";

export interface TokenBalance {
  address: Address;
  symbol: string;
  name?: string;
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
  const supported = (supportedRaw as Address[] | undefined) ?? [];
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

  const { meta, isLoading: metaLoading } = useTokenMetadata(addresses);

  const { data: balances, isLoading: balLoading } = useReadContracts({
    allowFailure: true,
    contracts: addresses.map(
      (address) => ({ chainId, address, abi: ERC20_ABI, functionName: "balanceOf", args: [account] }) as const,
    ),
    query: { enabled: addresses.length > 0 },
  });

  const tokens: TokenBalance[] = [];
  if (balances) {
    addresses.forEach((address, i) => {
      const m = meta.get(address.toLowerCase());
      const balanceRes = balances[i];

      // Require balanceOf AND a trustworthy decimals: without decimals we'd
      // misformat the balance (e.g. a 6-decimal token shown as 18), so drop
      // rather than guess. symbol/name are cosmetic and may fall back.
      if (balanceRes?.status !== "success" || m?.decimals === undefined) {
        if (import.meta.env.DEV) {
          // eslint-disable-next-line no-console
          console.warn(`[useErc20Balances] ${address} is not a usable ERC20 on chain ${chainId} — dropping.`);
        }
        return;
      }

      const balance = balanceRes.result as bigint;
      tokens.push({
        address,
        symbol: m.symbol ?? shortAddress(address),
        name: m.name,
        decimals: m.decimals,
        balance,
        formatted: formatUnits(balance, m.decimals),
        isZero: balance === 0n,
      });
    });
  }

  return { tokens, isLoading: supportedLoading || metaLoading || balLoading };
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `./node_modules/.bin/vitest run src/hooks/useErc20Balances.test.tsx`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add ui/src/hooks/useErc20Balances.ts ui/src/hooks/useErc20Balances.test.tsx
git commit -m "refactor(ui): build useErc20Balances on useTokenMetadata, add token name"
```

---

### Task 4: `TokenDisplay` component

**Files:**
- Create: `ui/src/components/TokenDisplay.tsx`
- Test: `ui/src/components/TokenDisplay.test.tsx`

**Interfaces:**
- Consumes: `Identicon`, `Tooltip`, `CopyButton`, `shortAddress`.
- Produces: `TokenDisplay(props: { address: string; symbol?: string; name?: string; className?: string })` — renders identicon, symbol (or compact address when no symbol) as headline, name as muted secondary, and the compact address (full on hover) with a copy button.

- [ ] **Step 1: Write the failing test**

Create `ui/src/components/TokenDisplay.test.tsx`:

```tsx
import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { TokenDisplay } from "./TokenDisplay";

const A = "0x29F37F6adCa168B79B8d9567eab9BE3fBF21db85";

describe("TokenDisplay", () => {
  it("shows symbol headline, name, and a compact address", () => {
    render(<TokenDisplay address={A} symbol="EURe" name="Monerium EUR emoney" />);
    expect(screen.getByText("EURe")).toBeInTheDocument();
    expect(screen.getByText("Monerium EUR emoney")).toBeInTheDocument();
    expect(screen.getByText("0x29F3…db85")).toBeInTheDocument();
  });

  it("falls back to a compact address headline when symbol is missing", () => {
    render(<TokenDisplay address={A} />);
    // Both headline and the secondary line render the compact form.
    expect(screen.getAllByText("0x29F3…db85").length).toBeGreaterThanOrEqual(1);
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/components/TokenDisplay.test.tsx`
Expected: FAIL — cannot resolve `./TokenDisplay`.

- [ ] **Step 3: Implement the component**

Create `ui/src/components/TokenDisplay.tsx`:

```tsx
import { shortAddress } from "../lib/format";
import { CopyButton } from "./CopyButton";
import { Identicon } from "./Identicon";
import { Tooltip } from "./Tooltip";

/**
 * Render an ERC20 token in a human-readable way: identicon + symbol headline
 * (or compact address when the symbol is unknown), the full name as muted
 * secondary text, and the compacted address (full on hover) with a copy button.
 */
export function TokenDisplay({
  address,
  symbol,
  name,
  className = "",
}: {
  address: string;
  symbol?: string;
  name?: string;
  className?: string;
}) {
  const short = shortAddress(address);
  return (
    <span className={`inline-flex min-w-0 items-center gap-2 ${className}`}>
      <Identicon address={address} size={20} />
      <span className="min-w-0">
        <span className="flex min-w-0 items-baseline gap-2">
          <span className="truncate font-medium">{symbol ?? short}</span>
          {symbol && name && <span className="truncate text-xs text-gray-500 dark:text-gray-400">{name}</span>}
        </span>
        <span className="flex items-center gap-1">
          <Tooltip content={address}>
            <span className="font-mono text-xs text-gray-500 dark:text-gray-400">{short}</span>
          </Tooltip>
          <CopyButton value={address} label="Copy address" />
        </span>
      </span>
    </span>
  );
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `./node_modules/.bin/vitest run src/components/TokenDisplay.test.tsx`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add ui/src/components/TokenDisplay.tsx ui/src/components/TokenDisplay.test.tsx
git commit -m "feat(ui): add TokenDisplay component (symbol + name + compact address)"
```

---

### Task 5: JetBrains Mono numeric font

**Files:**
- Modify: `ui/package.json` (dependency), `ui/src/main.tsx`, `ui/tailwind.config.js`

**Interfaces:**
- Produces: a `font-num` Tailwind utility (JetBrains Mono) for numeric/fund values. Used in Tasks 6–7 as `className="font-num tabular-nums"`.

- [ ] **Step 1: Add the font dependency**

Run: `yarn add @fontsource/jetbrains-mono`
Expected: `package.json` gains `@fontsource/jetbrains-mono` under dependencies.

- [ ] **Step 2: Import the font CSS**

In `ui/src/main.tsx`, add below `import "./index.css";`:

```ts
import "@fontsource/jetbrains-mono/400.css";
import "@fontsource/jetbrains-mono/500.css";
```

- [ ] **Step 3: Register the Tailwind family**

In `ui/tailwind.config.js`, replace `theme: { extend: {} },` with:

```js
  theme: {
    extend: {
      fontFamily: {
        num: ['"JetBrains Mono"', "ui-monospace", "monospace"],
      },
    },
  },
```

- [ ] **Step 4: Verify the build picks it up**

Run: `yarn build`
Expected: build succeeds (`tsc -b && vite build`), no errors.

- [ ] **Step 5: Commit**

```bash
git add ui/package.json ui/yarn.lock ui/src/main.tsx ui/tailwind.config.js
git commit -m "feat(ui): bundle JetBrains Mono as the font-num numeric family"
```

---

### Task 6: Wire `TokenDisplay` into the Payment Tokens tab

**Files:**
- Modify: `ui/src/pages/tabs/PaymentTokensTab.tsx`

**Interfaces:**
- Consumes: `useTokenMetadata` (Task 2), `TokenDisplay` (Task 4).

- [ ] **Step 1: Resolve metadata and render with `TokenDisplay`**

Replace the whole contents of `ui/src/pages/tabs/PaymentTokensTab.tsx`:

```tsx
import { type Abi, type Address } from "viem";
import { useWriteContract } from "wagmi";
import { ListManager } from "../../components/ListManager";
import { TokenDisplay } from "../../components/TokenDisplay";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";
import { useTokenMetadata } from "../../hooks/useTokenMetadata";

export function PaymentTokensTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const abi = cmAccountAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { items, isLoading, refetch } = useContractList(account, abi, "getSupportedTokens");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, abi, "SERVICE_ADMIN_ROLE");
  const { meta } = useTokenMetadata(items as Address[]);

  return (
    <ListManager
      title="Payment Tokens"
      items={items}
      isLoading={isLoading || roleLoading}
      roleName="SERVICE_ADMIN_ROLE"
      hasRole={hasRole}
      addLabel="Add token"
      addPlaceholder="Token address 0x…"
      onAdd={(v) => writeContractAsync({ address: account, abi, functionName: "addSupportedToken", args: [v as Address] })}
      onRemove={(v) => writeContractAsync({ address: account, abi, functionName: "removeSupportedToken", args: [v as Address] })}
      onChanged={refetch}
      renderItem={(v) => {
        const m = meta.get(v.toLowerCase());
        return <TokenDisplay address={v} symbol={m?.symbol} name={m?.name} />;
      }}
    />
  );
}
```

- [ ] **Step 2: Verify type-check and existing tab tests pass**

Run: `yarn build`
Expected: PASS (no TypeScript errors).

- [ ] **Step 3: Commit**

```bash
git add ui/src/pages/tabs/PaymentTokensTab.tsx
git commit -m "feat(ui): show symbol + name in Payment Tokens tab"
```

---

### Task 7: Apply `TokenDisplay`, `formatAmount` and `font-num` in the Active Account panel

**Files:**
- Modify: `ui/src/components/AccountSummary.tsx`
- Test: `ui/src/components/AccountSummary.test.tsx` (verify still passes)

**Interfaces:**
- Consumes: `TokenDisplay` (Task 4), `formatAmount` (Task 1), `useErc20Balances` `name` field (Task 3), `font-num` utility (Task 5), `Tooltip`.

- [ ] **Step 1: Update native balance and token rows**

In `ui/src/components/AccountSummary.tsx`:

Add imports near the existing component imports:

```ts
import { formatAmount } from "../lib/format";
import { Tooltip } from "./Tooltip";
import { TokenDisplay } from "./TokenDisplay";
```

Replace the native-balance `<dd>` (`<dd>{bal ? `${bal.formatted} ${bal.symbol}` : "—"}</dd>`) with:

```tsx
          <dd className="font-num tabular-nums">
            {bal ? (
              <Tooltip content={`${bal.formatted} ${bal.symbol}`}>
                <span>
                  {formatAmount(bal.formatted).display} {bal.symbol}
                </span>
              </Tooltip>
            ) : (
              "—"
            )}
          </dd>
```

Replace the token row (the `<div key={t.address}>` block's inner `flex items-center justify-between` line) so it uses `TokenDisplay` and the numeric font:

```tsx
                  <div className="flex items-center justify-between gap-2">
                    <TokenDisplay address={t.address} symbol={t.symbol} name={t.name} className="min-w-0 text-sm" />
                    <Tooltip content={`${t.formatted} ${t.symbol}`}>
                      <span className={`font-num tabular-nums whitespace-nowrap ${t.isZero ? "text-amber-700 dark:text-amber-300" : ""}`}>
                        {formatAmount(t.formatted).display} {t.symbol}
                      </span>
                    </Tooltip>
                  </div>
```

(Leave the existing zero-balance warning `<p>` below it unchanged.)

- [ ] **Step 2: Run the panel tests**

Run: `./node_modules/.bin/vitest run src/components/AccountSummary.test.tsx`
Expected: PASS — the amount text still renders as `1.5 USDC` (formatAmount("1.5") → "1.5"), and the zero/native warnings are unchanged.

- [ ] **Step 3: Run the full UI test suite**

Run: `./node_modules/.bin/vitest run`
Expected: PASS (all suites).

- [ ] **Step 4: Lint/format**

Run: `yarn format` (from `ui/`)
Expected: files formatted, no errors.

- [ ] **Step 5: Commit**

```bash
git add ui/src/components/AccountSummary.tsx
git commit -m "feat(ui): readable token rows and tabular fund values in account panel"
```

---

## Self-Review

**Spec coverage:**
- Symbol + name + compact/hover address → Tasks 4 (component), 6 (tab), 7 (panel). ✓
- Reusable metadata hook + `useErc20Balances` reuse → Tasks 2, 3. ✓
- JetBrains Mono + tabular figures on fund values → Tasks 5, 7. ✓
- `formatAmount` trim with full-on-hover → Tasks 1, 7. ✓
- Consistency ripple (Withdrawals/BookingToken): on inspection those render inputs and the token's symbol string, not computed fund *values*, so there is no displayed amount to restyle — no task needed; documented here so it isn't mistaken for a gap.

**Placeholder scan:** none — every code/test step contains full content.

**Type consistency:** `TokenMeta` (`symbol?`/`name?`/`decimals?`) defined in Task 2 and consumed in Task 3 test + hook; `TokenBalance.name?` added in Task 3 and consumed in Task 7; `formatAmount` `{ display, full }` defined Task 1, consumed Task 7; `TokenDisplay` props consistent across Tasks 4/6/7; `font-num` defined Task 5, used Tasks 7. ✓
