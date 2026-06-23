# Contract Management UI Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a wallet-connected React SPA in `ui/` that lets partners create and manage their Camino Messenger CM Accounts, deployed to GitHub Pages.

**Architecture:** Client-side-only SPA (no backend). Reads go through an app-owned viem HTTP transport; writes go through the user's wallet via wagmi/RainbowKit. A build-time sync script resolves contract addresses + ABIs from the repo's `ignition/` and `abi/` outputs into generated TypeScript. Most CMAccount tabs share one generic, config-driven list-manager component.

**Tech Stack:** React 18, Vite, TypeScript, wagmi v2, viem v2, RainbowKit v2, TanStack Query (wagmi dep), Tailwind CSS v3, Vitest + @testing-library/react.

## Global Constraints

- Node `>=20` (dev env is v22), package manager `yarn` v1 inside `ui/` (its own lockfile, isolated from the Hardhat root).
- Networks: Camino mainnet (500, enabled), Base Sepolia (84532, enabled), Base mainnet (8453, enabled only when its address file exists), Columbus (501, defined but `enabled: false`).
- Reads use app-owned RPC transport; writes use the wallet connector. Active chain = the wallet's connected chain.
- Generated contract output (`ui/src/contracts/generated/`) is git-ignored; the Hardhat `abi/` and `ignition/deployments/` files are the source of truth.
- Vite `base`: `/camino-messenger-contracts/`.
- Every write goes through the shared `TxButton` component; every write action is role-gated.
- Repo root contains `abi/` and `ignition/deployments/chain-*/deployed_addresses.json`. The `ui/` folder is one level below root, so the sync script reads `../abi` and `../ignition` relative to `ui/`.

---

## File Structure

```
ui/
  package.json                         # isolated deps + scripts
  tsconfig.json, tsconfig.node.json
  vite.config.ts                       # base path, vitest config
  index.html
  tailwind.config.js, postcss.config.js
  .gitignore                           # ignores src/contracts/generated
  src/
    main.tsx                           # React entry
    App.tsx                            # router + providers
    index.css                          # tailwind directives
    contracts/
      generated/                       # GENERATED (git-ignored): addresses.ts, abis.ts
      index.ts                         # re-exports + getContractsForChain()
    config/
      chains.ts                        # network definitions (enabled flags, RPC, explorer)
    wallet/
      wagmi.ts                         # wagmi config (transports, connectors)
      Providers.tsx                    # WagmiProvider + QueryClient + RainbowKit
    lib/
      roles.ts                         # role name -> keccak hash map + helpers
      format.ts                        # address/amount formatting
    hooks/
      useActiveContracts.ts            # resolve addresses+abi for connected chain
      useHasRole.ts                    # role check for connected account
      useRoleMembers.ts                # read getRoleMembers(roleHash) -> address[]
    components/
      TxButton.tsx                     # simulate->send->pending->confirmed/failed
      RoleGate.tsx                     # wrapper: disable/explain if missing role
      AddressInput.tsx, Card.tsx, NetworkBadge.tsx
      ListManager.tsx                  # generic list + add + remove (getter or event source)
    pages/
      Dashboard.tsx                    # network status + my accounts (read-only)
      CreateAccount.tsx
      AccountWorkspace.tsx             # tab shell + routing
      tabs/
        OverviewTab.tsx
        BotsTab.tsx
        PaymentTokensTab.tsx
        ServicesTab.tsx                # supported + wanted
        RolesTab.tsx
        PubkeysTab.tsx
        WithdrawalsTab.tsx
  scripts/
    sync-contracts.ts                  # generates src/contracts/generated/*
    sync-contracts.test.ts
.github/workflows/deploy-ui.yml
```

---

## Task 1: Scaffold the `ui/` app (runnable shell)

**Files:**
- Create: `ui/package.json`, `ui/tsconfig.json`, `ui/tsconfig.node.json`, `ui/vite.config.ts`, `ui/index.html`, `ui/tailwind.config.js`, `ui/postcss.config.js`, `ui/.gitignore`, `ui/src/main.tsx`, `ui/src/App.tsx`, `ui/src/index.css`
- Test: `ui/src/App.test.tsx`

**Interfaces:**
- Produces: a Vite app that builds and runs; `App` component rendering a placeholder heading "Camino Messenger Contracts".

- [ ] **Step 1: Create `ui/package.json`**

```json
{
  "name": "camino-messenger-ui",
  "private": true,
  "version": "0.1.0",
  "type": "module",
  "scripts": {
    "sync": "tsx scripts/sync-contracts.ts",
    "predev": "yarn sync",
    "dev": "vite",
    "prebuild": "yarn sync",
    "build": "tsc -b && vite build",
    "preview": "vite preview",
    "test": "vitest run",
    "test:watch": "vitest"
  },
  "dependencies": {
    "@rainbow-me/rainbowkit": "^2.1.0",
    "@tanstack/react-query": "^5.51.0",
    "react": "^18.3.1",
    "react-dom": "^18.3.1",
    "react-router-dom": "^6.26.0",
    "viem": "^2.21.0",
    "wagmi": "^2.12.0"
  },
  "devDependencies": {
    "@testing-library/jest-dom": "^6.4.0",
    "@testing-library/react": "^16.0.0",
    "@types/react": "^18.3.0",
    "@types/react-dom": "^18.3.0",
    "@vitejs/plugin-react": "^4.3.0",
    "autoprefixer": "^10.4.0",
    "jsdom": "^24.1.0",
    "postcss": "^8.4.0",
    "tailwindcss": "^3.4.0",
    "tsx": "^4.16.0",
    "typescript": "^5.5.0",
    "vite": "^5.4.0",
    "vitest": "^2.0.0"
  }
}
```

- [ ] **Step 2: Create `ui/tsconfig.json`**

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx",
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "types": ["vitest/globals", "@testing-library/jest-dom"]
  },
  "include": ["src", "scripts"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

- [ ] **Step 3: Create `ui/tsconfig.node.json`**

```json
{
  "compilerOptions": {
    "composite": true,
    "skipLibCheck": true,
    "module": "ESNext",
    "moduleResolution": "bundler",
    "allowSyntheticDefaultImports": true,
    "strict": true
  },
  "include": ["vite.config.ts"]
}
```

- [ ] **Step 4: Create `ui/vite.config.ts`**

```ts
/// <reference types="vitest" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  base: "/camino-messenger-contracts/",
  test: {
    globals: true,
    environment: "jsdom",
    setupFiles: ["./src/test-setup.ts"],
  },
});
```

- [ ] **Step 5: Create `ui/src/test-setup.ts`**

```ts
import "@testing-library/jest-dom/vitest";
```

- [ ] **Step 6: Create `ui/index.html`**

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Camino Messenger Contracts</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
```

- [ ] **Step 7: Create `ui/tailwind.config.js`, `ui/postcss.config.js`, `ui/src/index.css`**

`tailwind.config.js`:
```js
export default {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: { extend: {} },
  plugins: [],
};
```

`postcss.config.js`:
```js
export default { plugins: { tailwindcss: {}, autoprefixer: {} } };
```

`src/index.css`:
```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

- [ ] **Step 8: Create `ui/.gitignore`**

```
node_modules
dist
src/contracts/generated
```

- [ ] **Step 9: Create `ui/src/App.tsx` and `ui/src/main.tsx`**

`App.tsx`:
```tsx
export default function App() {
  return (
    <main className="p-8">
      <h1 className="text-2xl font-bold">Camino Messenger Contracts</h1>
    </main>
  );
}
```

`main.tsx`:
```tsx
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
```

- [ ] **Step 10: Write the failing test `ui/src/App.test.tsx`**

```tsx
import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders app heading", () => {
  render(<App />);
  expect(
    screen.getByRole("heading", { name: /camino messenger contracts/i }),
  ).toBeInTheDocument();
});
```

- [ ] **Step 11: Install deps and run the test**

Run: `cd ui && yarn install && yarn test`
Expected: PASS (1 test). If install pulls slightly different patch versions, that is fine.

- [ ] **Step 12: Verify build works**

Run: `cd ui && yarn build`
Expected: build succeeds, `ui/dist/` produced. (Sync runs via `prebuild`; Task 2 creates it. Until then, temporarily run `yarn vite build` to confirm the Vite side. After Task 2 the full `yarn build` works.)

- [ ] **Step 13: Commit**

```bash
git add ui/
git commit -m "feat(ui): scaffold React + Vite + Tailwind app shell"
```

---

## Task 2: Contract sync script

**Files:**
- Create: `ui/scripts/sync-contracts.ts`
- Test: `ui/scripts/sync-contracts.test.ts`
- Create (generated at runtime, git-ignored): `ui/src/contracts/generated/addresses.ts`, `ui/src/contracts/generated/abis.ts`

**Interfaces:**
- Produces:
  - `resolveAddresses(deployments: Record<number, Record<string,string>>): Record<number, ResolvedAddresses>` where
    `ResolvedAddresses = { manager: string; cmAccountImpl: string; bookingToken: string }`.
  - The script's `main()` writes `addresses.ts` exporting `export const ADDRESSES: Record<number, ResolvedAddresses>` and `abis.ts` exporting `MANAGER_ABI`, `CMACCOUNT_ABI`, `BOOKINGTOKEN_ABI` (typed `as const`).
- Resolution rules: prefer keys ending `#ManagerProxy` for `manager`, `#BookingTokenProxy` for `bookingToken`, and the canonical `CaminoMessengerModule#CMAccount` for `cmAccountImpl`. Ignore any module prefix other than `CaminoMessengerModule` (handles the messy Columbus file). A chain with no `deployed_addresses.json` is omitted entirely.

- [ ] **Step 1: Write the failing test `ui/scripts/sync-contracts.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { resolveAddresses } from "./sync-contracts";

describe("resolveAddresses", () => {
  it("resolves canonical proxy addresses per chain", () => {
    const out = resolveAddresses({
      84532: {
        "CaminoMessengerModule#ManagerProxy": "0xMANAGER",
        "CaminoMessengerModule#BookingTokenProxy": "0xBT",
        "CaminoMessengerModule#CMAccount": "0xIMPL",
        "CaminoMessengerModule#ManagerERC1967Proxy": "0xMANAGER",
      },
    });
    expect(out[84532]).toEqual({
      manager: "0xMANAGER",
      bookingToken: "0xBT",
      cmAccountImpl: "0xIMPL",
    });
  });

  it("ignores non-canonical historical modules (Columbus mess)", () => {
    const out = resolveAddresses({
      501: {
        "CaminoMessengerModule#ManagerProxy": "0xGOOD",
        "CaminoMessengerModule#BookingTokenProxy": "0xGOODBT",
        "CaminoMessengerModule#CMAccount": "0xGOODIMPL",
        "RefactorCancellationModule#ManagerProxy": "0xBAD",
        "ERC20ServiceFeeModule#CMAccount": "0xBAD2",
      },
    });
    expect(out[501].manager).toBe("0xGOOD");
    expect(out[501].cmAccountImpl).toBe("0xGOODIMPL");
  });

  it("omits chains missing required keys", () => {
    const out = resolveAddresses({ 8453: { "Other#Thing": "0x0" } });
    expect(out[8453]).toBeUndefined();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run scripts/sync-contracts.test.ts`
Expected: FAIL ("resolveAddresses is not a function" / module not found).

- [ ] **Step 3: Implement `ui/scripts/sync-contracts.ts`**

```ts
import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

export interface ResolvedAddresses {
  manager: string;
  bookingToken: string;
  cmAccountImpl: string;
}

const CANONICAL = "CaminoMessengerModule";

export function resolveAddresses(
  deployments: Record<number, Record<string, string>>,
): Record<number, ResolvedAddresses> {
  const result: Record<number, ResolvedAddresses> = {};
  for (const [chainIdStr, keys] of Object.entries(deployments)) {
    const chainId = Number(chainIdStr);
    const manager = keys[`${CANONICAL}#ManagerProxy`];
    const bookingToken = keys[`${CANONICAL}#BookingTokenProxy`];
    const cmAccountImpl = keys[`${CANONICAL}#CMAccount`];
    if (!manager || !bookingToken || !cmAccountImpl) continue;
    result[chainId] = { manager, bookingToken, cmAccountImpl };
  }
  return result;
}

const __dirname = dirname(fileURLToPath(import.meta.url));
const REPO_ROOT = join(__dirname, "..", "..");
const DEPLOY_DIR = join(REPO_ROOT, "ignition", "deployments");
const ABI_DIR = join(REPO_ROOT, "abi", "contracts");
const OUT_DIR = join(__dirname, "..", "src", "contracts", "generated");
const CHAIN_IDS = [500, 501, 8453, 84532];

function loadDeployments(): Record<number, Record<string, string>> {
  const out: Record<number, Record<string, string>> = {};
  for (const id of CHAIN_IDS) {
    const file = join(DEPLOY_DIR, `chain-${id}`, "deployed_addresses.json");
    if (existsSync(file)) out[id] = JSON.parse(readFileSync(file, "utf8"));
  }
  return out;
}

function loadAbi(relPath: string): unknown {
  return JSON.parse(readFileSync(join(ABI_DIR, relPath), "utf8"));
}

function main(): void {
  const resolved = resolveAddresses(loadDeployments());
  mkdirSync(OUT_DIR, { recursive: true });

  writeFileSync(
    join(OUT_DIR, "addresses.ts"),
    `// AUTO-GENERATED by scripts/sync-contracts.ts. Do not edit.
export interface ResolvedAddresses { manager: string; bookingToken: string; cmAccountImpl: string; }
export const ADDRESSES: Record<number, ResolvedAddresses> = ${JSON.stringify(resolved, null, 2)};
`,
  );

  const manager = loadAbi("manager/CMAccountManager.sol/CMAccountManager.json");
  const cmAccount = loadAbi("account/CMAccount.sol/CMAccount.json");
  const bookingToken = loadAbi("booking-token/BookingToken.sol/BookingToken.json");

  writeFileSync(
    join(OUT_DIR, "abis.ts"),
    `// AUTO-GENERATED by scripts/sync-contracts.ts. Do not edit.
export const MANAGER_ABI = ${JSON.stringify(manager)} as const;
export const CMACCOUNT_ABI = ${JSON.stringify(cmAccount)} as const;
export const BOOKINGTOKEN_ABI = ${JSON.stringify(bookingToken)} as const;
`,
  );

  console.log(`Synced ${Object.keys(resolved).length} chains to ${OUT_DIR}`);
}

// Run main only when executed directly (not when imported by tests).
if (process.argv[1] && process.argv[1].endsWith("sync-contracts.ts")) {
  main();
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn vitest run scripts/sync-contracts.test.ts`
Expected: PASS (3 tests).

- [ ] **Step 5: Run the actual sync and confirm output**

Run: `cd ui && yarn sync`
Expected: prints `Synced 3 chains` (500, 501, 84532; 8453 omitted until deployed) and creates `ui/src/contracts/generated/addresses.ts` + `abis.ts`.

- [ ] **Step 6: Commit**

```bash
git add ui/scripts/
git commit -m "feat(ui): add contract sync script with address resolution"
```

---

## Task 3: Network config + contracts index

**Files:**
- Create: `ui/src/config/chains.ts`, `ui/src/contracts/index.ts`
- Test: `ui/src/config/chains.test.ts`

**Interfaces:**
- Consumes: `ADDRESSES` from `contracts/generated/addresses.ts`.
- Produces:
  - `chains.ts`: `export interface AppChain { id: number; name: string; enabled: boolean; rpcUrl: string; explorerUrl: string; nativeCurrency: { name: string; symbol: string; decimals: number }; }` and `export const APP_CHAINS: AppChain[]`, plus `export const ENABLED_CHAINS: AppChain[]` (enabled AND has addresses).
  - `contracts/index.ts`: `export function getContractsForChain(chainId: number): ResolvedAddresses | undefined` and re-exports the ABIs.

- [ ] **Step 1: Write the failing test `ui/src/config/chains.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { APP_CHAINS, ENABLED_CHAINS } from "./chains";

describe("chain config", () => {
  it("defines all four networks", () => {
    expect(APP_CHAINS.map((c) => c.id).sort()).toEqual([500, 501, 8453, 84532]);
  });

  it("marks Columbus (501) as disabled", () => {
    expect(APP_CHAINS.find((c) => c.id === 501)!.enabled).toBe(false);
  });

  it("excludes Columbus from enabled chains", () => {
    expect(ENABLED_CHAINS.some((c) => c.id === 501)).toBe(false);
  });

  it("only enables chains that have deployed addresses", () => {
    for (const c of ENABLED_CHAINS) {
      expect(c.enabled).toBe(true);
    }
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/config/chains.test.ts`
Expected: FAIL (module not found). Ensure `yarn sync` has been run so `generated/` exists.

- [ ] **Step 3: Implement `ui/src/contracts/index.ts`**

```ts
import { ADDRESSES, type ResolvedAddresses } from "./generated/addresses";
export { MANAGER_ABI, CMACCOUNT_ABI, BOOKINGTOKEN_ABI } from "./generated/abis";
export type { ResolvedAddresses };

export function getContractsForChain(chainId: number): ResolvedAddresses | undefined {
  return ADDRESSES[chainId];
}

export function hasContracts(chainId: number): boolean {
  return ADDRESSES[chainId] !== undefined;
}
```

- [ ] **Step 4: Implement `ui/src/config/chains.ts`**

```ts
import { hasContracts } from "../contracts";

export interface AppChain {
  id: number;
  name: string;
  enabled: boolean;
  rpcUrl: string;
  explorerUrl: string;
  nativeCurrency: { name: string; symbol: string; decimals: number };
}

export const APP_CHAINS: AppChain[] = [
  {
    id: 500,
    name: "Camino",
    enabled: true,
    rpcUrl: "https://api.camino.network/ext/bc/C/rpc",
    explorerUrl: "https://caminoscan.com",
    nativeCurrency: { name: "Camino", symbol: "CAM", decimals: 18 },
  },
  {
    id: 501,
    name: "Columbus (deprecated)",
    enabled: false,
    rpcUrl: "https://columbus.camino.network/ext/bc/C/rpc",
    explorerUrl: "https://columbus.caminoscan.com",
    nativeCurrency: { name: "Camino", symbol: "CAM", decimals: 18 },
  },
  {
    id: 8453,
    name: "Base",
    enabled: true,
    rpcUrl: "https://mainnet.base.org",
    explorerUrl: "https://basescan.org",
    nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  },
  {
    id: 84532,
    name: "Base Sepolia",
    enabled: true,
    rpcUrl: "https://base-sepolia.drpc.org",
    explorerUrl: "https://sepolia.basescan.org",
    nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  },
];

// A chain is usable only if marked enabled AND its contracts are deployed.
export const ENABLED_CHAINS: AppChain[] = APP_CHAINS.filter(
  (c) => c.enabled && hasContracts(c.id),
);
```

- [ ] **Step 5: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/config/chains.test.ts`
Expected: PASS (4 tests).

- [ ] **Step 6: Commit**

```bash
git add ui/src/config ui/src/contracts/index.ts
git commit -m "feat(ui): add network config and contracts resolver"
```

---

## Task 4: Wallet providers (wagmi + RainbowKit) and read transport

**Files:**
- Create: `ui/src/wallet/wagmi.ts`, `ui/src/wallet/Providers.tsx`
- Modify: `ui/src/main.tsx` (wrap App in Providers and Router)
- Test: `ui/src/wallet/wagmi.test.ts`

**Interfaces:**
- Consumes: `APP_CHAINS`, `ENABLED_CHAINS` from config.
- Produces: `export const wagmiConfig` built with viem `http(rpcUrl)` transports per enabled chain (app-owned read transport) and an injected/RainbowKit connector for writes; `export function Providers({ children })`.

- [ ] **Step 1: Write the failing test `ui/src/wallet/wagmi.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { wagmiConfig } from "./wagmi";
import { ENABLED_CHAINS } from "../config/chains";

describe("wagmiConfig", () => {
  it("registers a transport for every enabled chain", () => {
    for (const c of ENABLED_CHAINS) {
      expect(wagmiConfig.chains.some((wc) => wc.id === c.id)).toBe(true);
    }
  });

  it("does not register the disabled Columbus chain", () => {
    expect(wagmiConfig.chains.some((wc) => wc.id === 501)).toBe(false);
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/wallet/wagmi.test.ts`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/wallet/wagmi.ts`**

```ts
import { http } from "viem";
import { type Chain } from "viem";
import { createConfig } from "wagmi";
import { injected } from "wagmi/connectors";
import { ENABLED_CHAINS, type AppChain } from "../config/chains";

function toViemChain(c: AppChain): Chain {
  return {
    id: c.id,
    name: c.name,
    nativeCurrency: c.nativeCurrency,
    rpcUrls: { default: { http: [c.rpcUrl] } },
    blockExplorers: { default: { name: c.name, url: c.explorerUrl } },
  };
}

const viemChains = ENABLED_CHAINS.map(toViemChain);

export const wagmiConfig = createConfig({
  chains: viemChains as [Chain, ...Chain[]],
  connectors: [injected()],
  transports: Object.fromEntries(
    ENABLED_CHAINS.map((c) => [c.id, http(c.rpcUrl)]),
  ),
});
```

> Note: RainbowKit's `getDefaultConfig` can replace this later for richer wallet support; `injected()` keeps the first version dependency-light and testable. If using RainbowKit's connectors, set a WalletConnect `projectId` via `import.meta.env.VITE_WC_PROJECT_ID`.

- [ ] **Step 4: Implement `ui/src/wallet/Providers.tsx`**

```tsx
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { WagmiProvider } from "wagmi";
import { wagmiConfig } from "./wagmi";

const queryClient = new QueryClient();

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    </WagmiProvider>
  );
}
```

- [ ] **Step 5: Update `ui/src/main.tsx`**

```tsx
import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import { Providers } from "./wallet/Providers";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Providers>
      <BrowserRouter basename="/camino-messenger-contracts">
        <App />
      </BrowserRouter>
    </Providers>
  </React.StrictMode>,
);
```

- [ ] **Step 6: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/wallet/wagmi.test.ts`
Expected: PASS (2 tests).

- [ ] **Step 7: Commit**

```bash
git add ui/src/wallet ui/src/main.tsx
git commit -m "feat(ui): add wagmi config and wallet providers"
```

---

## Task 5: Role definitions and `useHasRole`

**Files:**
- Create: `ui/src/lib/roles.ts`, `ui/src/hooks/useHasRole.ts`
- Test: `ui/src/lib/roles.test.ts`

**Interfaces:**
- Produces:
  - `roles.ts`: `export const ROLE_HASHES: Record<RoleName, \`0x${string}\`>` where `RoleName` is a union of the account + manager role strings; `DEFAULT_ADMIN_ROLE` is the zero bytes32; all others are `keccak256(toBytes(name))`. Also `export function roleHash(name: RoleName)`.
  - `useHasRole.ts`: `export function useHasRole(contractAddress, abi, role: RoleName): { hasRole: boolean; isLoading: boolean }` using wagmi `useReadContract` calling `hasRole(role, connectedAddress)`.

- [ ] **Step 1: Write the failing test `ui/src/lib/roles.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { keccak256, toBytes } from "viem";
import { ROLE_HASHES } from "./roles";

describe("ROLE_HASHES", () => {
  it("uses zero bytes32 for DEFAULT_ADMIN_ROLE", () => {
    expect(ROLE_HASHES.DEFAULT_ADMIN_ROLE).toBe(`0x${"0".repeat(64)}`);
  });

  it("hashes named roles with keccak256", () => {
    expect(ROLE_HASHES.BOT_ADMIN_ROLE).toBe(keccak256(toBytes("BOT_ADMIN_ROLE")));
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/lib/roles.test.ts`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/lib/roles.ts`**

```ts
import { keccak256, toBytes, type Hex } from "viem";

export const ACCOUNT_ROLES = [
  "DEFAULT_ADMIN_ROLE",
  "UPGRADER_ROLE",
  "BOT_ADMIN_ROLE",
  "MESSENGER_BOT_ROLE",
  "GAS_WITHDRAWER_ROLE",
  "WITHDRAWER_ROLE",
  "BOOKING_OPERATOR_ROLE",
  "SERVICE_ADMIN_ROLE",
] as const;

export const MANAGER_ROLES = [
  "PAUSER_ROLE",
  "VERSIONER_ROLE",
  "SERVICE_REGISTRY_ADMIN_ROLE",
] as const;

export type RoleName = (typeof ACCOUNT_ROLES)[number] | (typeof MANAGER_ROLES)[number];

const ZERO_BYTES32 = `0x${"0".repeat(64)}` as Hex;

function compute(name: string): Hex {
  return name === "DEFAULT_ADMIN_ROLE" ? ZERO_BYTES32 : keccak256(toBytes(name));
}

export const ROLE_HASHES = Object.fromEntries(
  [...ACCOUNT_ROLES, ...MANAGER_ROLES].map((r) => [r, compute(r)]),
) as Record<RoleName, Hex>;

export function roleHash(name: RoleName): Hex {
  return ROLE_HASHES[name];
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/lib/roles.test.ts`
Expected: PASS (2 tests).

- [ ] **Step 5: Implement `ui/src/hooks/useHasRole.ts`** (no separate test; exercised via component tests)

```ts
import { type Abi, type Address } from "viem";
import { useAccount, useReadContract } from "wagmi";
import { roleHash, type RoleName } from "../lib/roles";

export function useHasRole(contractAddress: Address | undefined, abi: Abi, role: RoleName) {
  const { address } = useAccount();
  const { data, isLoading } = useReadContract({
    address: contractAddress,
    abi,
    functionName: "hasRole",
    args: address ? [roleHash(role), address] : undefined,
    query: { enabled: Boolean(contractAddress && address) },
  });
  return { hasRole: Boolean(data), isLoading };
}
```

- [ ] **Step 6: Commit**

```bash
git add ui/src/lib/roles.ts ui/src/lib/roles.test.ts ui/src/hooks/useHasRole.ts
git commit -m "feat(ui): add role hashes and useHasRole hook"
```

---

## Task 6: `useActiveContracts` hook

**Files:**
- Create: `ui/src/hooks/useActiveContracts.ts`
- Test: `ui/src/hooks/useActiveContracts.test.tsx`

**Interfaces:**
- Consumes: `getContractsForChain`, ABIs, wagmi `useChainId`/`useAccount`.
- Produces: `export function useActiveContracts(): { chainId: number | undefined; supported: boolean; manager?: Address; bookingToken?: Address; cmAccountImpl?: Address; managerAbi; cmAccountAbi; bookingTokenAbi }`. `supported` is true when the connected chain has resolved contracts.

- [ ] **Step 1: Write the failing test `ui/src/hooks/useActiveContracts.test.tsx`**

```tsx
import { describe, expect, it, vi } from "vitest";
import { renderHook } from "@testing-library/react";
import { useActiveContracts } from "./useActiveContracts";

vi.mock("wagmi", () => ({ useChainId: () => 84532 }));

describe("useActiveContracts", () => {
  it("reports supported=true for a chain with contracts", () => {
    const { result } = renderHook(() => useActiveContracts());
    expect(result.current.chainId).toBe(84532);
    expect(result.current.supported).toBe(true);
    expect(result.current.manager).toBeTruthy();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/hooks/useActiveContracts.test.tsx`
Expected: FAIL (module not found). Requires `yarn sync` so 84532 addresses exist.

- [ ] **Step 3: Implement `ui/src/hooks/useActiveContracts.ts`**

```ts
import { type Address } from "viem";
import { useChainId } from "wagmi";
import {
  BOOKINGTOKEN_ABI,
  CMACCOUNT_ABI,
  MANAGER_ABI,
  getContractsForChain,
} from "../contracts";

export function useActiveContracts() {
  const chainId = useChainId();
  const resolved = chainId ? getContractsForChain(chainId) : undefined;
  return {
    chainId,
    supported: Boolean(resolved),
    manager: resolved?.manager as Address | undefined,
    bookingToken: resolved?.bookingToken as Address | undefined,
    cmAccountImpl: resolved?.cmAccountImpl as Address | undefined,
    managerAbi: MANAGER_ABI,
    cmAccountAbi: CMACCOUNT_ABI,
    bookingTokenAbi: BOOKINGTOKEN_ABI,
  };
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/hooks/useActiveContracts.test.tsx`
Expected: PASS (1 test).

- [ ] **Step 5: Commit**

```bash
git add ui/src/hooks/useActiveContracts.ts ui/src/hooks/useActiveContracts.test.tsx
git commit -m "feat(ui): add useActiveContracts hook"
```

---

## Task 7: Shared `TxButton` component

**Files:**
- Create: `ui/src/components/TxButton.tsx`, `ui/src/lib/format.ts`
- Test: `ui/src/components/TxButton.test.tsx`

**Interfaces:**
- Consumes: wagmi `useWriteContract`, `useWaitForTransactionReceipt`.
- Produces:
  - `format.ts`: `export function shortAddress(a: string): string` (e.g. `0x1234…abcd`); `export function explorerTxUrl(base: string, hash: string)`.
  - `TxButton`: props `{ label: string; disabled?: boolean; write: () => Promise<\`0x${string}\`>; onConfirmed?: () => void; explorerBase?: string }`. Renders a button; on click calls `write()`, shows states idle → pending ("Confirming…") → success ("Confirmed") or error (message). Exposes an explorer link when a hash exists.

- [ ] **Step 1: Write the failing test `ui/src/lib/format.ts` test inline + component test `ui/src/components/TxButton.test.tsx`**

`ui/src/lib/format.test.ts`:
```ts
import { describe, expect, it } from "vitest";
import { shortAddress } from "./format";

describe("shortAddress", () => {
  it("truncates the middle", () => {
    expect(shortAddress("0x1234567890abcdef1234567890abcdef12345678")).toBe(
      "0x1234…5678",
    );
  });
});
```

`ui/src/components/TxButton.test.tsx`:
```tsx
import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { TxButton } from "./TxButton";

describe("TxButton", () => {
  it("calls write and shows confirmed on success", async () => {
    const write = vi.fn().mockResolvedValue("0xhash");
    const onConfirmed = vi.fn();
    render(<TxButton label="Do it" write={write} onConfirmed={onConfirmed} />);
    fireEvent.click(screen.getByRole("button", { name: /do it/i }));
    await waitFor(() => expect(write).toHaveBeenCalled());
    await waitFor(() => expect(onConfirmed).toHaveBeenCalled());
  });

  it("shows an error message on failure", async () => {
    const write = vi.fn().mockRejectedValue(new Error("user rejected"));
    render(<TxButton label="Do it" write={write} />);
    fireEvent.click(screen.getByRole("button", { name: /do it/i }));
    await waitFor(() =>
      expect(screen.getByText(/user rejected/i)).toBeInTheDocument(),
    );
  });
});
```

> Note: this `TxButton` takes a `write()` callback rather than calling wagmi directly, so it is unit-testable without mocking the whole wagmi stack. Callers build `write` from `useWriteContract` + `useWaitForTransactionReceipt` (shown in Task 9 and tabs). The button manages its own pending/confirmed/error UI from the promise.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `cd ui && yarn vitest run src/components/TxButton.test.tsx src/lib/format.test.ts`
Expected: FAIL (modules not found).

- [ ] **Step 3: Implement `ui/src/lib/format.ts`**

```ts
export function shortAddress(a: string): string {
  if (a.length < 10) return a;
  return `${a.slice(0, 6)}…${a.slice(-4)}`;
}

export function explorerTxUrl(base: string, hash: string): string {
  return `${base}/tx/${hash}`;
}

export function explorerAddrUrl(base: string, addr: string): string {
  return `${base}/address/${addr}`;
}
```

- [ ] **Step 4: Implement `ui/src/components/TxButton.tsx`**

```tsx
import { useState } from "react";
import { explorerTxUrl } from "../lib/format";

type Status = "idle" | "pending" | "success" | "error";

interface TxButtonProps {
  label: string;
  disabled?: boolean;
  write: () => Promise<`0x${string}`>;
  onConfirmed?: () => void;
  explorerBase?: string;
}

export function TxButton({ label, disabled, write, onConfirmed, explorerBase }: TxButtonProps) {
  const [status, setStatus] = useState<Status>("idle");
  const [hash, setHash] = useState<string>();
  const [error, setError] = useState<string>();

  async function handleClick() {
    setStatus("pending");
    setError(undefined);
    try {
      const h = await write();
      setHash(h);
      setStatus("success");
      onConfirmed?.();
    } catch (e) {
      setError(e instanceof Error ? e.message : String(e));
      setStatus("error");
    }
  }

  return (
    <div className="flex flex-col gap-1">
      <button
        type="button"
        disabled={disabled || status === "pending"}
        onClick={handleClick}
        className="rounded bg-indigo-600 px-3 py-1.5 text-white disabled:opacity-50"
      >
        {status === "pending" ? "Confirming…" : status === "success" ? "Confirmed" : label}
      </button>
      {hash && explorerBase && (
        <a className="text-xs text-indigo-500 underline" href={explorerTxUrl(explorerBase, hash)} target="_blank" rel="noreferrer">
          View transaction
        </a>
      )}
      {error && <span className="text-xs text-red-600">{error}</span>}
    </div>
  );
}
```

- [ ] **Step 5: Run the tests to verify they pass**

Run: `cd ui && yarn vitest run src/components/TxButton.test.tsx src/lib/format.test.ts`
Expected: PASS (3 tests total).

- [ ] **Step 6: Commit**

```bash
git add ui/src/components/TxButton.tsx ui/src/lib/format.ts ui/src/lib/format.test.ts ui/src/components/TxButton.test.tsx
git commit -m "feat(ui): add TxButton and formatting helpers"
```

---

## Task 8: `RoleGate` component

**Files:**
- Create: `ui/src/components/RoleGate.tsx`
- Test: `ui/src/components/RoleGate.test.tsx`

**Interfaces:**
- Produces: `RoleGate` with props `{ hasRole: boolean; isLoading?: boolean; roleName: string; children: ReactNode }`. When `hasRole` renders children; otherwise renders a disabled/explanatory message naming the required role. (Takes `hasRole` as a prop so it is decoupled from the chain; callers pass the result of `useHasRole`.)

- [ ] **Step 1: Write the failing test `ui/src/components/RoleGate.test.tsx`**

```tsx
import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { RoleGate } from "./RoleGate";

describe("RoleGate", () => {
  it("renders children when role present", () => {
    render(<RoleGate hasRole roleName="BOT_ADMIN_ROLE"><button>Add bot</button></RoleGate>);
    expect(screen.getByRole("button", { name: /add bot/i })).toBeInTheDocument();
  });

  it("explains the missing role otherwise", () => {
    render(<RoleGate hasRole={false} roleName="BOT_ADMIN_ROLE"><button>Add bot</button></RoleGate>);
    expect(screen.queryByRole("button", { name: /add bot/i })).toBeNull();
    expect(screen.getByText(/BOT_ADMIN_ROLE/)).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/components/RoleGate.test.tsx`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/components/RoleGate.tsx`**

```tsx
import { type ReactNode } from "react";

interface RoleGateProps {
  hasRole: boolean;
  isLoading?: boolean;
  roleName: string;
  children: ReactNode;
}

export function RoleGate({ hasRole, isLoading, roleName, children }: RoleGateProps) {
  if (isLoading) return <span className="text-xs text-gray-400">Checking permissions…</span>;
  if (!hasRole)
    return (
      <p className="text-xs text-amber-600">
        Requires <code>{roleName}</code> on the connected account.
      </p>
    );
  return <>{children}</>;
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/components/RoleGate.test.tsx`
Expected: PASS (2 tests).

- [ ] **Step 5: Commit**

```bash
git add ui/src/components/RoleGate.tsx ui/src/components/RoleGate.test.tsx
git commit -m "feat(ui): add RoleGate component"
```

---

## Task 9: `useRoleMembers` hook

A "bot" on a CMAccount is not tracked by add/remove events — it is **an address granted special roles** (`MESSENGER_BOT_ROLE`, `BOOKING_OPERATOR_ROLE`, `GAS_WITHDRAWER_ROLE`), exactly as the `account bot:list` CLI task lists them via `role:members` (`tasks/account.js:445`). So bot listing — and the Roles tab — both read `getRoleMembers(roleHash)`. This task builds the shared hook for that.

**Files:**
- Create: `ui/src/hooks/useRoleMembers.ts`
- Test: `ui/src/hooks/useRoleMembers.test.ts`

**Interfaces:**
- Consumes: wagmi `useReadContract`, `roleHash` from `lib/roles`.
- Produces:
  - `export function toMemberList(data: unknown): string[]` — pure helper normalizing the `address[]` return of `getRoleMembers` to a string array (returns `[]` when undefined).
  - `export function useRoleMembers(account: Address | undefined, abi: Abi, role: RoleName): { members: string[]; isLoading: boolean; refetch: () => void }` — wraps `useReadContract({ functionName: "getRoleMembers", args: [roleHash(role)] })`.

- [ ] **Step 1: Write the failing test `ui/src/hooks/useRoleMembers.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { toMemberList } from "./useRoleMembers";

describe("toMemberList", () => {
  it("returns [] for undefined", () => {
    expect(toMemberList(undefined)).toEqual([]);
  });

  it("stringifies address array entries", () => {
    expect(toMemberList(["0xAAA", "0xBBB"])).toEqual(["0xAAA", "0xBBB"]);
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/hooks/useRoleMembers.test.ts`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/hooks/useRoleMembers.ts`**

```ts
import { type Abi, type Address } from "viem";
import { useReadContract } from "wagmi";
import { roleHash, type RoleName } from "../lib/roles";

export function toMemberList(data: unknown): string[] {
  return ((data as unknown[]) ?? []).map((x) => String(x));
}

export function useRoleMembers(account: Address | undefined, abi: Abi, role: RoleName) {
  const { data, isLoading, refetch } = useReadContract({
    address: account,
    abi,
    functionName: "getRoleMembers",
    args: [roleHash(role)],
    query: { enabled: Boolean(account) },
  });
  return { members: toMemberList(data), isLoading, refetch: () => void refetch() };
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/hooks/useRoleMembers.test.ts`
Expected: PASS (2 tests).

- [ ] **Step 5: Commit**

```bash
git add ui/src/hooks/useRoleMembers.ts ui/src/hooks/useRoleMembers.test.ts
git commit -m "feat(ui): add useRoleMembers hook"
```

---

## Task 10: App layout, routing, and ConnectButton

**Files:**
- Modify: `ui/src/App.tsx`
- Create: `ui/src/components/Layout.tsx`, `ui/src/components/ConnectButton.tsx`, `ui/src/components/Card.tsx`, `ui/src/components/NetworkBadge.tsx`
- Test: `ui/src/App.test.tsx` (update)

**Interfaces:**
- Consumes: react-router, wagmi `useAccount`/`useChainId`/`useConnect`/`useDisconnect`, `useActiveContracts`.
- Produces: routes `/` → `Dashboard`, `/create` → `CreateAccount`, `/account/:address` → `AccountWorkspace`. `Layout` renders header (title, nav links, ConnectButton, NetworkBadge) + `<Outlet/>`. `Card` is a styled container. `NetworkBadge` shows the active chain name or "Unsupported network".

- [ ] **Step 1: Update the failing test `ui/src/App.test.tsx`**

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import App from "./App";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined, isConnected: false }),
  useChainId: () => 84532,
  useConnect: () => ({ connect: vi.fn(), connectors: [] }),
  useDisconnect: () => ({ disconnect: vi.fn() }),
  useReadContract: () => ({ data: undefined, isLoading: false }),
  usePublicClient: () => undefined,
}));

describe("App", () => {
  it("renders the header title and connect button", () => {
    render(
      <MemoryRouter initialEntries={["/"]}>
        <App />
      </MemoryRouter>,
    );
    expect(screen.getByText(/camino messenger/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /connect/i })).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/App.test.tsx`
Expected: FAIL (no connect button / components missing).

- [ ] **Step 3: Implement `ui/src/components/Card.tsx`**

```tsx
import { type ReactNode } from "react";
export function Card({ title, children }: { title?: string; children: ReactNode }) {
  return (
    <section className="rounded-lg border border-gray-200 bg-white p-4 shadow-sm">
      {title && <h2 className="mb-3 text-lg font-semibold">{title}</h2>}
      {children}
    </section>
  );
}
```

- [ ] **Step 4: Implement `ui/src/components/ConnectButton.tsx`**

```tsx
import { useAccount, useConnect, useDisconnect } from "wagmi";
import { shortAddress } from "../lib/format";

export function ConnectButton() {
  const { address, isConnected } = useAccount();
  const { connect, connectors } = useConnect();
  const { disconnect } = useDisconnect();

  if (isConnected && address)
    return (
      <button className="rounded border px-3 py-1.5" onClick={() => disconnect()}>
        {shortAddress(address)} · Disconnect
      </button>
    );

  return (
    <button
      className="rounded bg-indigo-600 px-3 py-1.5 text-white"
      onClick={() => connect({ connector: connectors[0] })}
    >
      Connect Wallet
    </button>
  );
}
```

- [ ] **Step 5: Implement `ui/src/components/NetworkBadge.tsx`**

```tsx
import { useActiveContracts } from "../hooks/useActiveContracts";
import { APP_CHAINS } from "../config/chains";

export function NetworkBadge() {
  const { chainId, supported } = useActiveContracts();
  const name = APP_CHAINS.find((c) => c.id === chainId)?.name;
  return (
    <span className={`rounded px-2 py-1 text-xs ${supported ? "bg-green-100 text-green-800" : "bg-red-100 text-red-800"}`}>
      {supported ? name : "Unsupported network"}
    </span>
  );
}
```

- [ ] **Step 6: Implement `ui/src/components/Layout.tsx`**

```tsx
import { Link, Outlet } from "react-router-dom";
import { ConnectButton } from "./ConnectButton";
import { NetworkBadge } from "./NetworkBadge";

export function Layout() {
  return (
    <div className="min-h-screen bg-gray-50">
      <header className="flex items-center justify-between border-b bg-white px-6 py-3">
        <div className="flex items-center gap-4">
          <Link to="/" className="font-bold">Camino Messenger</Link>
          <Link to="/" className="text-sm text-gray-600">Dashboard</Link>
          <Link to="/create" className="text-sm text-gray-600">Create Account</Link>
        </div>
        <div className="flex items-center gap-3">
          <NetworkBadge />
          <ConnectButton />
        </div>
      </header>
      <main className="mx-auto max-w-5xl p-6">
        <Outlet />
      </main>
    </div>
  );
}
```

- [ ] **Step 7: Implement `ui/src/App.tsx`**

```tsx
import { Route, Routes } from "react-router-dom";
import { Layout } from "./components/Layout";
import { Dashboard } from "./pages/Dashboard";
import { CreateAccount } from "./pages/CreateAccount";
import { AccountWorkspace } from "./pages/AccountWorkspace";

export default function App() {
  return (
    <Routes>
      <Route element={<Layout />}>
        <Route index element={<Dashboard />} />
        <Route path="create" element={<CreateAccount />} />
        <Route path="account/:address" element={<AccountWorkspace />} />
      </Route>
    </Routes>
  );
}
```

- [ ] **Step 8: Create minimal page stubs so the app compiles**

`ui/src/pages/Dashboard.tsx`:
```tsx
import { Card } from "../components/Card";
export function Dashboard() {
  return <Card title="Dashboard">Network status and your accounts will appear here.</Card>;
}
```

`ui/src/pages/CreateAccount.tsx`:
```tsx
import { Card } from "../components/Card";
export function CreateAccount() {
  return <Card title="Create CM Account">Form coming next.</Card>;
}
```

`ui/src/pages/AccountWorkspace.tsx`:
```tsx
import { useParams } from "react-router-dom";
import { Card } from "../components/Card";
export function AccountWorkspace() {
  const { address } = useParams();
  return <Card title="Account Workspace">{address}</Card>;
}
```

- [ ] **Step 9: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/App.test.tsx`
Expected: PASS.

- [ ] **Step 10: Verify dev server renders**

Run: `cd ui && yarn build`
Expected: full `yarn build` (sync + tsc + vite) succeeds.

- [ ] **Step 11: Commit**

```bash
git add ui/src
git commit -m "feat(ui): add layout, routing, wallet connect button, page stubs"
```

---

## Task 11: Dashboard — network status + my accounts

**Files:**
- Modify: `ui/src/pages/Dashboard.tsx`
- Create: `ui/src/hooks/useMyAccounts.ts`
- Test: `ui/src/hooks/useMyAccounts.test.ts`

**Interfaces:**
- Consumes: `useActiveContracts`, viem `getLogs` for `CMAccountCreated(address indexed account)`, `useReadContract` for manager `paused()` and `getAccountImplementation()`.
- Produces:
  - `useMyAccounts.ts`: `export function filterAccountsByRole(...)` pure helper is not needed; instead export `useMyAccounts(): { accounts: Address[]; isLoading: boolean }` which lists all `CMAccountCreated` accounts (role filtering deferred to per-account display to keep RPC load low). Export a pure helper `export function uniqueAddresses(addrs: string[]): string[]` for the test.
  - Dashboard renders a network status `Card` (paused state, implementation address) and a "Created accounts" list linking to `/account/:address`.

- [ ] **Step 1: Write the failing test `ui/src/hooks/useMyAccounts.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { uniqueAddresses } from "./useMyAccounts";

describe("uniqueAddresses", () => {
  it("dedupes case-insensitively, preserving checksum of first seen", () => {
    expect(uniqueAddresses(["0xABC", "0xabc", "0xDEF"])).toEqual(["0xABC", "0xDEF"]);
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/hooks/useMyAccounts.test.ts`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/hooks/useMyAccounts.ts`**

```ts
import { type Address } from "viem";
import { usePublicClient } from "wagmi";
import { useEffect, useState } from "react";
import { useActiveContracts } from "./useActiveContracts";

export function uniqueAddresses(addrs: string[]): string[] {
  const seen = new Set<string>();
  const out: string[] = [];
  for (const a of addrs) {
    const k = a.toLowerCase();
    if (!seen.has(k)) {
      seen.add(k);
      out.push(a);
    }
  }
  return out;
}

export function useMyAccounts() {
  const client = usePublicClient();
  const { manager, managerAbi } = useActiveContracts();
  const [accounts, setAccounts] = useState<Address[]>([]);
  const [isLoading, setLoading] = useState(false);

  useEffect(() => {
    if (!client || !manager) return;
    let cancelled = false;
    (async () => {
      setLoading(true);
      const event = (managerAbi as any).find(
        (x: any) => x.type === "event" && x.name === "CMAccountCreated",
      );
      const logs = await client.getLogs({ address: manager, event, fromBlock: 0n, toBlock: "latest" });
      const addrs = logs.map((l) => String((l as any).args.account));
      if (!cancelled) {
        setAccounts(uniqueAddresses(addrs) as Address[]);
        setLoading(false);
      }
    })();
    return () => { cancelled = true; };
  }, [client, manager, managerAbi]);

  return { accounts, isLoading };
}
```

- [ ] **Step 4: Implement `ui/src/pages/Dashboard.tsx`**

```tsx
import { Link } from "react-router-dom";
import { useReadContract } from "wagmi";
import { Card } from "../components/Card";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { useMyAccounts } from "../hooks/useMyAccounts";
import { shortAddress } from "../lib/format";

export function Dashboard() {
  const { manager, managerAbi, supported } = useActiveContracts();
  const { accounts, isLoading } = useMyAccounts();
  const { data: paused } = useReadContract({ address: manager, abi: managerAbi, functionName: "paused" });
  const { data: impl } = useReadContract({ address: manager, abi: managerAbi, functionName: "getAccountImplementation" });

  if (!supported) return <Card title="Dashboard">Connect to a supported network.</Card>;

  return (
    <div className="grid gap-4">
      <Card title="Network status">
        <dl className="grid grid-cols-2 gap-2 text-sm">
          <dt className="text-gray-500">Manager</dt><dd>{manager}</dd>
          <dt className="text-gray-500">Paused</dt><dd>{paused ? "Yes" : "No"}</dd>
          <dt className="text-gray-500">Account implementation</dt><dd>{impl as string}</dd>
        </dl>
      </Card>
      <Card title="Created accounts">
        {isLoading ? "Loading…" : (
          <ul className="divide-y">
            {accounts.map((a) => (
              <li key={a} className="py-2">
                <Link className="text-indigo-600 underline" to={`/account/${a}`}>{shortAddress(a)}</Link>
              </li>
            ))}
          </ul>
        )}
      </Card>
    </div>
  );
}
```

- [ ] **Step 5: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/hooks/useMyAccounts.test.ts`
Expected: PASS (1 test).

- [ ] **Step 6: Commit**

```bash
git add ui/src/pages/Dashboard.tsx ui/src/hooks/useMyAccounts.ts ui/src/hooks/useMyAccounts.test.ts
git commit -m "feat(ui): dashboard with network status and account list"
```

---

## Task 12: Create CM Account flow

**Files:**
- Modify: `ui/src/pages/CreateAccount.tsx`
- Create: `ui/src/lib/receipt.ts`
- Test: `ui/src/lib/receipt.test.ts`

**Interfaces:**
- Consumes: `useActiveContracts`, `useWriteContract`, `useAccount`, `usePublicClient`, viem `decodeEventLog`, react-router `useNavigate`.
- Produces:
  - `receipt.ts`: `export function findCreatedAccount(logs, abi): Address | undefined` — decodes the `CMAccountCreated` event from a receipt's logs and returns the `account` address.
  - `CreateAccount`: form with `admin` and `upgrader` address inputs (defaulting to connected address), a `TxButton` whose `write()` calls `createCMAccount(admin, upgrader)`, waits for the receipt, extracts the new address, and navigates to `/account/:address`.

- [ ] **Step 1: Write the failing test `ui/src/lib/receipt.test.ts`**

```ts
import { describe, expect, it } from "vitest";
import { encodeEventLog, parseAbi } from "viem";
import { findCreatedAccount } from "./receipt";

const abi = parseAbi(["event CMAccountCreated(address indexed account)"]);

describe("findCreatedAccount", () => {
  it("extracts the account address from logs", () => {
    const addr = "0x1111111111111111111111111111111111111111";
    const log = encodeEventLog({ abi, eventName: "CMAccountCreated", args: { account: addr } });
    const found = findCreatedAccount(
      [{ data: log.data, topics: log.topics }] as any,
      abi as any,
    );
    expect(found?.toLowerCase()).toBe(addr);
  });

  it("returns undefined when no matching event", () => {
    expect(findCreatedAccount([{ data: "0x", topics: ["0xdead"] }] as any, abi as any)).toBeUndefined();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/lib/receipt.test.ts`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/lib/receipt.ts`**

```ts
import { decodeEventLog, type Abi, type Address, type Log } from "viem";

export function findCreatedAccount(
  logs: Pick<Log, "data" | "topics">[],
  abi: Abi,
): Address | undefined {
  for (const log of logs) {
    try {
      const decoded = decodeEventLog({ abi, data: log.data, topics: log.topics as any });
      if (decoded.eventName === "CMAccountCreated") {
        return (decoded.args as { account: Address }).account;
      }
    } catch {
      // not our event, skip
    }
  }
  return undefined;
}
```

- [ ] **Step 4: Implement `ui/src/pages/CreateAccount.tsx`**

```tsx
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { type Address } from "viem";
import { useAccount, usePublicClient, useWriteContract } from "wagmi";
import { Card } from "../components/Card";
import { TxButton } from "../components/TxButton";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { findCreatedAccount } from "../lib/receipt";

export function CreateAccount() {
  const { address } = useAccount();
  const { manager, managerAbi, cmAccountAbi, supported } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const client = usePublicClient();
  const navigate = useNavigate();
  const [admin, setAdmin] = useState("");
  const [upgrader, setUpgrader] = useState("");

  const adminVal = (admin || address || "") as Address;
  const upgraderVal = (upgrader || address || "") as Address;

  async function write() {
    const hash = await writeContractAsync({
      address: manager!,
      abi: managerAbi,
      functionName: "createCMAccount",
      args: [adminVal, upgraderVal],
    });
    const receipt = await client!.waitForTransactionReceipt({ hash });
    const created = findCreatedAccount(receipt.logs, cmAccountAbi as any);
    if (created) navigate(`/account/${created}`);
    return hash;
  }

  if (!supported) return <Card title="Create CM Account">Connect to a supported network.</Card>;

  return (
    <Card title="Create CM Account">
      <div className="grid max-w-md gap-3">
        <label className="text-sm">Admin address
          <input className="mt-1 w-full rounded border px-2 py-1" placeholder={address} value={admin} onChange={(e) => setAdmin(e.target.value)} />
        </label>
        <label className="text-sm">Upgrader address
          <input className="mt-1 w-full rounded border px-2 py-1" placeholder={address} value={upgrader} onChange={(e) => setUpgrader(e.target.value)} />
        </label>
        <TxButton label="Create account" disabled={!address} write={write} />
      </div>
    </Card>
  );
}
```

- [ ] **Step 5: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/lib/receipt.test.ts`
Expected: PASS (2 tests).

- [ ] **Step 6: Commit**

```bash
git add ui/src/pages/CreateAccount.tsx ui/src/lib/receipt.ts ui/src/lib/receipt.test.ts
git commit -m "feat(ui): create CM Account flow with receipt parsing"
```

---

## Task 13: Account workspace shell with tabs

**Files:**
- Modify: `ui/src/pages/AccountWorkspace.tsx`
- Create: `ui/src/pages/tabs/OverviewTab.tsx`
- Test: `ui/src/pages/AccountWorkspace.test.tsx`

**Interfaces:**
- Consumes: react-router `useParams`, `useSearchParams`; `useActiveContracts`.
- Produces: `AccountWorkspace` reads `:address` param, renders a tab bar (Overview, Bots, Payment Tokens, Services, Roles, Pubkeys, Withdrawals) controlled by `?tab=` search param, and renders the active tab passing `account: Address` + `cmAccountAbi` as props. Each tab component signature: `({ account: Address }) => JSX`. `OverviewTab` shows the account address and native balance (`useBalance`).

- [ ] **Step 1: Write the failing test `ui/src/pages/AccountWorkspace.test.tsx`**

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { AccountWorkspace } from "./AccountWorkspace";

vi.mock("wagmi", () => ({
  useChainId: () => 84532,
  useBalance: () => ({ data: { formatted: "1.0", symbol: "ETH" } }),
  useAccount: () => ({ address: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false }),
  usePublicClient: () => undefined,
}));

describe("AccountWorkspace", () => {
  it("renders the tab bar with the account address", () => {
    const addr = "0x1111111111111111111111111111111111111111";
    render(
      <MemoryRouter initialEntries={[`/account/${addr}`]}>
        <Routes><Route path="account/:address" element={<AccountWorkspace />} /></Routes>
      </MemoryRouter>,
    );
    expect(screen.getByRole("link", { name: /bots/i })).toBeInTheDocument();
    expect(screen.getByText(new RegExp(addr.slice(0, 6), "i"))).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/pages/AccountWorkspace.test.tsx`
Expected: FAIL.

- [ ] **Step 3: Implement `ui/src/pages/tabs/OverviewTab.tsx`**

```tsx
import { type Address } from "viem";
import { useBalance } from "wagmi";
import { Card } from "../../components/Card";

export function OverviewTab({ account }: { account: Address }) {
  const { data } = useBalance({ address: account });
  return (
    <Card title="Overview">
      <dl className="grid grid-cols-2 gap-2 text-sm">
        <dt className="text-gray-500">Address</dt><dd className="break-all">{account}</dd>
        <dt className="text-gray-500">Native balance</dt><dd>{data ? `${data.formatted} ${data.symbol}` : "—"}</dd>
      </dl>
    </Card>
  );
}
```

- [ ] **Step 4: Implement `ui/src/pages/AccountWorkspace.tsx`**

```tsx
import { Link, useParams, useSearchParams } from "react-router-dom";
import { type Address } from "viem";
import { Card } from "../components/Card";
import { shortAddress } from "../lib/format";
import { OverviewTab } from "./tabs/OverviewTab";
import { BotsTab } from "./tabs/BotsTab";
import { PaymentTokensTab } from "./tabs/PaymentTokensTab";
import { ServicesTab } from "./tabs/ServicesTab";
import { RolesTab } from "./tabs/RolesTab";
import { PubkeysTab } from "./tabs/PubkeysTab";
import { WithdrawalsTab } from "./tabs/WithdrawalsTab";

const TABS = [
  { id: "overview", label: "Overview", Component: OverviewTab },
  { id: "bots", label: "Bots", Component: BotsTab },
  { id: "tokens", label: "Payment Tokens", Component: PaymentTokensTab },
  { id: "services", label: "Services", Component: ServicesTab },
  { id: "roles", label: "Roles", Component: RolesTab },
  { id: "pubkeys", label: "Pubkeys", Component: PubkeysTab },
  { id: "withdrawals", label: "Withdrawals", Component: WithdrawalsTab },
] as const;

export function AccountWorkspace() {
  const { address } = useParams();
  const [params] = useSearchParams();
  const active = params.get("tab") ?? "overview";
  const account = address as Address;
  const Active = (TABS.find((t) => t.id === active) ?? TABS[0]).Component;

  return (
    <div className="grid gap-4">
      <Card><h1 className="font-mono">{shortAddress(account)}</h1></Card>
      <nav className="flex gap-3 border-b text-sm">
        {TABS.map((t) => (
          <Link key={t.id} to={`?tab=${t.id}`} className={`pb-2 ${active === t.id ? "border-b-2 border-indigo-600 font-medium" : "text-gray-500"}`}>
            {t.label}
          </Link>
        ))}
      </nav>
      <Active account={account} />
    </div>
  );
}
```

- [ ] **Step 5: Create placeholder tab files so it compiles**

Create each of `ui/src/pages/tabs/BotsTab.tsx`, `PaymentTokensTab.tsx`, `ServicesTab.tsx`, `RolesTab.tsx`, `PubkeysTab.tsx`, `WithdrawalsTab.tsx` with this body (replace `NAME`):

```tsx
import { type Address } from "viem";
import { Card } from "../../components/Card";
export function NAME({ account }: { account: Address }) {
  return <Card title="NAME">Coming soon for {account}.</Card>;
}
```

Use exact export names: `BotsTab`, `PaymentTokensTab`, `ServicesTab`, `RolesTab`, `PubkeysTab`, `WithdrawalsTab`.

- [ ] **Step 6: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/pages/AccountWorkspace.test.tsx`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add ui/src/pages/AccountWorkspace.tsx ui/src/pages/tabs ui/src/pages/AccountWorkspace.test.tsx
git commit -m "feat(ui): account workspace shell with tab routing"
```

---

## Task 14: Generic `ListManager` component

**Files:**
- Create: `ui/src/components/ListManager.tsx`
- Test: `ui/src/components/ListManager.test.tsx`

**Interfaces:**
- Produces: `ListManager` with props:
  ```ts
  interface ListManagerProps {
    title: string;
    items: string[];
    isLoading: boolean;
    roleName: string;
    hasRole: boolean;
    addLabel: string;
    addPlaceholder: string;
    onAdd: (value: string) => Promise<\`0x${string}\`>;
    onRemove: (value: string) => Promise<\`0x${string}\`>;
    onChanged?: () => void;
    explorerBase?: string;
    renderItem?: (value: string) => React.ReactNode;
  }
  ```
  Renders: a list of items each with a remove `TxButton` (gated by `RoleGate`), plus an input + add `TxButton` (also gated). Pure presentational; data fetching lives in each tab. This is the shared engine for tokens/bots/pubkeys/services/roles tabs.

- [ ] **Step 1: Write the failing test `ui/src/components/ListManager.test.tsx`**

```tsx
import { describe, expect, it, vi } from "vitest";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { ListManager } from "./ListManager";

const base = {
  title: "Payment Tokens",
  isLoading: false,
  roleName: "DEFAULT_ADMIN_ROLE",
  addLabel: "Add token",
  addPlaceholder: "0x…",
  onRemove: vi.fn().mockResolvedValue("0xhash"),
};

describe("ListManager", () => {
  it("renders items and calls onAdd with the input value", async () => {
    const onAdd = vi.fn().mockResolvedValue("0xhash");
    render(<ListManager {...base} hasRole items={["0xAAA"]} onAdd={onAdd} />);
    expect(screen.getByText("0xAAA")).toBeInTheDocument();
    fireEvent.change(screen.getByPlaceholderText("0x…"), { target: { value: "0xBBB" } });
    fireEvent.click(screen.getByRole("button", { name: /add token/i }));
    await waitFor(() => expect(onAdd).toHaveBeenCalledWith("0xBBB"));
  });

  it("hides add/remove controls without the role", () => {
    render(<ListManager {...base} hasRole={false} items={["0xAAA"]} onAdd={vi.fn()} />);
    expect(screen.queryByRole("button", { name: /add token/i })).toBeNull();
    expect(screen.getByText(/DEFAULT_ADMIN_ROLE/)).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd ui && yarn vitest run src/components/ListManager.test.tsx`
Expected: FAIL (module not found).

- [ ] **Step 3: Implement `ui/src/components/ListManager.tsx`**

```tsx
import { type ReactNode, useState } from "react";
import { Card } from "./Card";
import { RoleGate } from "./RoleGate";
import { TxButton } from "./TxButton";

interface ListManagerProps {
  title: string;
  items: string[];
  isLoading: boolean;
  roleName: string;
  hasRole: boolean;
  addLabel: string;
  addPlaceholder: string;
  onAdd: (value: string) => Promise<`0x${string}`>;
  onRemove: (value: string) => Promise<`0x${string}`>;
  onChanged?: () => void;
  explorerBase?: string;
  renderItem?: (value: string) => ReactNode;
}

export function ListManager(props: ListManagerProps) {
  const { title, items, isLoading, roleName, hasRole, addLabel, addPlaceholder } = props;
  const [value, setValue] = useState("");

  return (
    <Card title={title}>
      {isLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y">
          {items.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {items.map((item) => (
            <li key={item} className="flex items-center justify-between py-2">
              <span className="font-mono text-sm">{props.renderItem ? props.renderItem(item) : item}</span>
              <RoleGate hasRole={hasRole} roleName={roleName}>
                <TxButton label="Remove" write={() => props.onRemove(item)} onConfirmed={props.onChanged} explorerBase={props.explorerBase} />
              </RoleGate>
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName={roleName}>
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder={addPlaceholder} value={value} onChange={(e) => setValue(e.target.value)} />
          <TxButton label={addLabel} disabled={!value} write={() => props.onAdd(value)} onConfirmed={() => { setValue(""); props.onChanged?.(); }} explorerBase={props.explorerBase} />
        </div>
      </RoleGate>
    </Card>
  );
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn vitest run src/components/ListManager.test.tsx`
Expected: PASS (2 tests).

- [ ] **Step 5: Commit**

```bash
git add ui/src/components/ListManager.tsx ui/src/components/ListManager.test.tsx
git commit -m "feat(ui): add generic ListManager component"
```

---

## Task 15: Payment Tokens tab (getter-based ListManager)

**Files:**
- Modify: `ui/src/pages/tabs/PaymentTokensTab.tsx`
- Create: `ui/src/hooks/useContractList.ts`
- Test: none new (logic covered by ListManager + useContractList is a thin wrapper; verified via build)

**Interfaces:**
- Consumes: `useReadContract`, `useWriteContract`, `useHasRole`, `useActiveContracts`, `ListManager`.
- Produces: `useContractList.ts`: `export function useContractList(account, abi, functionName): { items: string[]; isLoading: boolean; refetch: () => void }` — reads an `address[]`/`string[]` getter and returns string items.

- [ ] **Step 1: Implement `ui/src/hooks/useContractList.ts`**

```ts
import { type Abi, type Address } from "viem";
import { useReadContract } from "wagmi";

export function useContractList(account: Address, abi: Abi, functionName: string) {
  const { data, isLoading, refetch } = useReadContract({ address: account, abi, functionName });
  const items = ((data as unknown[]) ?? []).map((x) => String(x));
  return { items, isLoading, refetch: () => void refetch() };
}
```

- [ ] **Step 2: Implement `ui/src/pages/tabs/PaymentTokensTab.tsx`**

```tsx
import { type Address } from "viem";
import { useWriteContract } from "wagmi";
import { ListManager } from "../../components/ListManager";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

export function PaymentTokensTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const { items, isLoading, refetch } = useContractList(account, cmAccountAbi as any, "getSupportedTokens");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, cmAccountAbi as any, "DEFAULT_ADMIN_ROLE");

  return (
    <ListManager
      title="Payment Tokens"
      items={items}
      isLoading={isLoading || roleLoading}
      roleName="DEFAULT_ADMIN_ROLE"
      hasRole={hasRole}
      addLabel="Add token"
      addPlaceholder="Token address 0x…"
      onAdd={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "addSupportedToken", args: [v as Address] })}
      onRemove={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "removeSupportedToken", args: [v as Address] })}
      onChanged={refetch}
    />
  );
}
```

> Role note: confirm the exact role guarding `addSupportedToken` in `contracts/account/CMAccount.sol`. If it is not `DEFAULT_ADMIN_ROLE`, update `roleName` and the `useHasRole` arg accordingly. (The implementer must read the contract to verify; do not guess.)

- [ ] **Step 3: Verify build**

Run: `cd ui && yarn build`
Expected: succeeds.

- [ ] **Step 4: Commit**

```bash
git add ui/src/pages/tabs/PaymentTokensTab.tsx ui/src/hooks/useContractList.ts
git commit -m "feat(ui): payment tokens tab"
```

---

## Task 16: Bots tab (role-member based + gas money)

A bot is an address holding `MESSENGER_BOT_ROLE` (and typically `BOOKING_OPERATOR_ROLE` / `GAS_WITHDRAWER_ROLE`). `addMessengerBot(bot, gasMoney)` grants those roles and transfers `gasMoney` CAM from the account to the bot; `removeMessengerBot(bot)` revokes them. So the bot list = members of `MESSENGER_BOT_ROLE` (read via `useRoleMembers`), matching the `bot:list` CLI task.

**Files:**
- Modify: `ui/src/pages/tabs/BotsTab.tsx`
- Test: none new (`useRoleMembers` and `TxButton`/`RoleGate` already tested)

**Interfaces:**
- Consumes: `useRoleMembers(account, abi, "MESSENGER_BOT_ROLE")`, `useWriteContract`, `useHasRole` with `BOT_ADMIN_ROLE`, `Card`/`RoleGate`/`TxButton`. Add takes bot address + gas money (CAM), so it uses a custom two-field form rather than `ListManager`.
- Produces: Bots tab listing `MESSENGER_BOT_ROLE` members with add/remove.

- [ ] **Step 1: Implement `ui/src/pages/tabs/BotsTab.tsx`**

```tsx
import { useState } from "react";
import { type Address, parseEther } from "viem";
import { useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { useHasRole } from "../../hooks/useHasRole";

export function BotsTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, cmAccountAbi as any, "MESSENGER_BOT_ROLE");
  const { hasRole, isLoading: roleLoading } = useHasRole(account, cmAccountAbi as any, "BOT_ADMIN_ROLE");
  const [bot, setBot] = useState("");
  const [gas, setGas] = useState("0");

  return (
    <Card title="Messenger Bots">
      <p className="mb-3 text-xs text-gray-500">A bot is an address granted MESSENGER_BOT_ROLE (plus booking/gas roles) on this account.</p>
      {isLoading || roleLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y">
          {members.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {members.map((b) => (
            <li key={b} className="flex items-center justify-between py-2">
              <span className="font-mono text-sm">{b}</span>
              <RoleGate hasRole={hasRole} roleName="BOT_ADMIN_ROLE">
                <TxButton label="Remove" write={() => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "removeMessengerBot", args: [b as Address] })} onConfirmed={refetch} />
              </RoleGate>
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="BOT_ADMIN_ROLE">
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Bot address 0x…" value={bot} onChange={(e) => setBot(e.target.value)} />
          <input className="w-32 rounded border px-2 py-1" placeholder="Gas money (CAM)" value={gas} onChange={(e) => setGas(e.target.value)} />
          <TxButton label="Add bot" disabled={!bot} write={() => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "addMessengerBot", args: [bot as Address, parseEther(gas || "0")] })} onConfirmed={() => { setBot(""); setGas("0"); refetch(); }} />
        </div>
      </RoleGate>
    </Card>
  );
}
```

> Note: `addMessengerBot(address bot, uint256 gasMoney)` — `gasMoney` is the CAM amount transferred from the account to the bot (CLI `bot:add` accepts e.g. `1` or `0.1` CAM and the contract works in wei), so the form takes CAM and converts via `parseEther`.

- [ ] **Step 2: Verify build**

Run: `cd ui && yarn build`
Expected: succeeds.

- [ ] **Step 3: Commit**

```bash
git add ui/src/pages/tabs/BotsTab.tsx
git commit -m "feat(ui): bots tab listing MESSENGER_BOT_ROLE members with gas money"
```

---

## Task 17: Services tab (supported + wanted)

**Files:**
- Modify: `ui/src/pages/tabs/ServicesTab.tsx`
- Test: none new

**Interfaces:**
- Consumes: `useContractList` against `getSupportedServices` and `getWantedServices` (verify return types/getter names against ABI — `getSupportedServices`/`getWantedServices` exist), `useWriteContract` for `addService`/`removeService`/`addWantedServices`/`removeWantedServices`, `useHasRole` with `SERVICE_ADMIN_ROLE`, two `ListManager`s.
- Produces: Services tab with two sections.

- [ ] **Step 1: Verify getter return shapes and add/remove arg types**

Run: `cd /hgst/work/github.com/TravelTokenMarketplace/camino-messenger/camino-messenger-contracts && node -e 'const a=require("./abi/contracts/account/CMAccount.sol/CMAccount.json");["getSupportedServices","getWantedServices","addService","removeService","addWantedServices","removeWantedServices"].forEach(n=>{const f=a.find(x=>x.type==="function"&&x.name===n);console.log(n,"in:",JSON.stringify(f.inputs.map(i=>i.type)),"out:",JSON.stringify((f.outputs||[]).map(o=>o.type)))})'`
Expected: prints signatures. If `addService` takes a struct or multiple args (not a single string), the simple ListManager input is insufficient — in that case implement a custom form like BotsTab and note it. Decide based on actual output; the code below assumes string service names. **If signatures differ, adapt before writing.**

- [ ] **Step 2: Implement `ui/src/pages/tabs/ServicesTab.tsx`**

```tsx
import { type Address } from "viem";
import { useWriteContract } from "wagmi";
import { ListManager } from "../../components/ListManager";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

export function ServicesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const supported = useContractList(account, cmAccountAbi as any, "getSupportedServices");
  const wanted = useContractList(account, cmAccountAbi as any, "getWantedServices");
  const { hasRole } = useHasRole(account, cmAccountAbi as any, "SERVICE_ADMIN_ROLE");

  return (
    <div className="grid gap-4">
      <ListManager
        title="Supported Services"
        items={supported.items}
        isLoading={supported.isLoading}
        roleName="SERVICE_ADMIN_ROLE"
        hasRole={hasRole}
        addLabel="Add service"
        addPlaceholder="Service name"
        onAdd={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "addService", args: [v] })}
        onRemove={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "removeService", args: [v] })}
        onChanged={supported.refetch}
      />
      <ListManager
        title="Wanted Services"
        items={wanted.items}
        isLoading={wanted.isLoading}
        roleName="SERVICE_ADMIN_ROLE"
        hasRole={hasRole}
        addLabel="Add wanted"
        addPlaceholder="Service name"
        onAdd={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "addWantedServices", args: [[v]] })}
        onRemove={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "removeWantedServices", args: [[v]] })}
        onChanged={wanted.refetch}
      />
    </div>
  );
}
```

> Note: `addWantedServices`/`removeWantedServices` are plural — likely accept an array; the code wraps the single value in `[v]`. Confirm in Step 1 and adjust. `setServiceRestrictedRate`/`setServiceCapabilities` are deferred (advanced) — not in v1 per spec scope.

- [ ] **Step 3: Verify build and commit**

Run: `cd ui && yarn build`
```bash
git add ui/src/pages/tabs/ServicesTab.tsx
git commit -m "feat(ui): services tab (supported + wanted)"
```

---

## Task 18: Roles tab

**Files:**
- Modify: `ui/src/pages/tabs/RolesTab.tsx`
- Test: none new

**Interfaces:**
- Consumes: `useRoleMembers` per account role, `useWriteContract` for `grantRole`/`revokeRole`, `useHasRole` with `DEFAULT_ADMIN_ROLE`, `ACCOUNT_ROLES`, `roleHash`, `ListManager`.
- Produces: Roles tab rendering one `ListManager` per account role, listing members and granting/revoking by address.

- [ ] **Step 1: Implement `ui/src/pages/tabs/RolesTab.tsx`**

```tsx
import { type Address } from "viem";
import { useWriteContract } from "wagmi";
import { ListManager } from "../../components/ListManager";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { useRoleMembers } from "../../hooks/useRoleMembers";
import { ACCOUNT_ROLES, roleHash, type RoleName } from "../../lib/roles";

function RoleSection({ account, role, hasAdmin }: { account: Address; role: RoleName; hasAdmin: boolean }) {
  const { cmAccountAbi } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, cmAccountAbi as any, role);
  return (
    <ListManager
      title={role}
      items={members}
      isLoading={isLoading}
      roleName="DEFAULT_ADMIN_ROLE"
      hasRole={hasAdmin}
      addLabel="Grant"
      addPlaceholder="Account address 0x…"
      onAdd={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "grantRole", args: [roleHash(role), v as Address] })}
      onRemove={(v) => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "revokeRole", args: [roleHash(role), v as Address] })}
      onChanged={refetch}
    />
  );
}

export function RolesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { hasRole } = useHasRole(account, cmAccountAbi as any, "DEFAULT_ADMIN_ROLE");
  return (
    <div className="grid gap-4">
      {ACCOUNT_ROLES.map((r) => (
        <RoleSection key={r} account={account} role={r} hasAdmin={hasRole} />
      ))}
    </div>
  );
}
```

- [ ] **Step 2: Verify build and commit**

Run: `cd ui && yarn build`
```bash
git add ui/src/pages/tabs/RolesTab.tsx
git commit -m "feat(ui): roles tab with grant/revoke per role"
```

---

## Task 19: Pubkeys and Withdrawals tabs

**Files:**
- Modify: `ui/src/pages/tabs/PubkeysTab.tsx`, `ui/src/pages/tabs/WithdrawalsTab.tsx`
- Test: none new

**Interfaces:**
- Pubkeys: `useContractList` against `getPublicKeysAddresses`; add via `addPublicKey(address, bytes)`, remove via `removePublicKey(address)`. Add needs two fields (address + data bytes), so use a custom form like BotsTab.
- Withdrawals: native `withdraw(recipient, amount)`, `withdrawERC20`/`withdrawERC721` (verify exact function names from ABI), gated by `WITHDRAWER_ROLE`.

- [ ] **Step 1: Verify withdrawal + pubkey function signatures**

Run: `cd /hgst/work/github.com/TravelTokenMarketplace/camino-messenger/camino-messenger-contracts && node -e 'const a=require("./abi/contracts/account/CMAccount.sol/CMAccount.json");a.filter(x=>x.type==="function"&&/withdraw|PublicKey/i.test(x.name)).forEach(f=>console.log(f.name+"("+f.inputs.map(i=>i.type+" "+i.name).join(", ")+")"))'`
Expected: prints `withdraw`, ERC20/ERC721 withdraw names, `addPublicKey`, `removePublicKey`, `getPublicKeysAddresses`. Use exact names below.

- [ ] **Step 2: Implement `ui/src/pages/tabs/PubkeysTab.tsx`**

```tsx
import { useState } from "react";
import { type Address, type Hex } from "viem";
import { useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useContractList } from "../../hooks/useContractList";
import { useHasRole } from "../../hooks/useHasRole";

export function PubkeysTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const { items, isLoading, refetch } = useContractList(account, cmAccountAbi as any, "getPublicKeysAddresses");
  const { hasRole } = useHasRole(account, cmAccountAbi as any, "DEFAULT_ADMIN_ROLE");
  const [addr, setAddr] = useState("");
  const [data, setData] = useState("");

  return (
    <Card title="Encryption Public Keys">
      {isLoading ? <p>Loading…</p> : (
        <ul className="mb-4 divide-y">
          {items.length === 0 && <li className="py-2 text-sm text-gray-400">None</li>}
          {items.map((k) => (
            <li key={k} className="flex items-center justify-between py-2">
              <span className="font-mono text-sm">{k}</span>
              <RoleGate hasRole={hasRole} roleName="DEFAULT_ADMIN_ROLE">
                <TxButton label="Remove" write={() => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "removePublicKey", args: [k as Address] })} onConfirmed={refetch} />
              </RoleGate>
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName="DEFAULT_ADMIN_ROLE">
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Address 0x…" value={addr} onChange={(e) => setAddr(e.target.value)} />
          <input className="flex-1 rounded border px-2 py-1" placeholder="Pubkey data (hex 0x…)" value={data} onChange={(e) => setData(e.target.value)} />
          <TxButton label="Add" disabled={!addr || !data} write={() => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "addPublicKey", args: [addr as Address, data as Hex] })} onConfirmed={refetch} />
        </div>
      </RoleGate>
    </Card>
  );
}
```

> Role note: confirm the role guarding `addPublicKey` in the contract; adjust `roleName`/`useHasRole` if not `DEFAULT_ADMIN_ROLE`.

- [ ] **Step 3: Implement `ui/src/pages/tabs/WithdrawalsTab.tsx`**

```tsx
import { useState } from "react";
import { type Address, parseEther } from "viem";
import { useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";

export function WithdrawalsTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  const { writeContractAsync } = useWriteContract();
  const { hasRole } = useHasRole(account, cmAccountAbi as any, "WITHDRAWER_ROLE");
  const [recipient, setRecipient] = useState("");
  const [amount, setAmount] = useState("0");

  return (
    <Card title="Withdraw native">
      <RoleGate hasRole={hasRole} roleName="WITHDRAWER_ROLE">
        <div className="flex items-end gap-2">
          <input className="flex-1 rounded border px-2 py-1" placeholder="Recipient 0x…" value={recipient} onChange={(e) => setRecipient(e.target.value)} />
          <input className="w-32 rounded border px-2 py-1" placeholder="Amount (ETH)" value={amount} onChange={(e) => setAmount(e.target.value)} />
          <TxButton label="Withdraw" disabled={!recipient} write={() => writeContractAsync({ address: account, abi: cmAccountAbi as any, functionName: "withdraw", args: [recipient as Address, parseEther(amount || "0")] })} />
        </div>
      </RoleGate>
    </Card>
  );
}
```

> Note: ERC20/ERC721 withdraw forms can be added once Step 1 confirms their exact function names and arg lists; v1 ships native withdraw here, and a `// TODO(v2): erc20/erc721` is acceptable as a follow-up issue, not a code placeholder. If the names are confirmed and simple, add two more small forms mirroring the native one.

- [ ] **Step 4: Verify build and commit**

Run: `cd ui && yarn build && yarn test`
Expected: build succeeds; all tests pass.
```bash
git add ui/src/pages/tabs/PubkeysTab.tsx ui/src/pages/tabs/WithdrawalsTab.tsx
git commit -m "feat(ui): pubkeys and withdrawals tabs"
```

---

## Task 20: GitHub Pages deploy workflow + docs

**Files:**
- Create: `.github/workflows/deploy-ui.yml`, `ui/README.md`
- Modify: `ui/vite.config.ts` (add a `copy index.html to 404.html` build step via a small plugin) OR add a post-build script

**Interfaces:**
- Produces: a workflow that builds `ui/` and deploys `ui/dist` to GitHub Pages on push to `dev` touching `ui/**`, `abi/**`, `ignition/deployments/**`, plus `workflow_dispatch`.

- [ ] **Step 1: Add SPA 404 fallback to the build**

Append a small plugin to `ui/vite.config.ts` `plugins` array:

```ts
import { copyFileSync } from "node:fs";
import { resolve } from "node:path";

const spaFallback = {
  name: "spa-404-fallback",
  closeBundle() {
    copyFileSync(resolve("dist/index.html"), resolve("dist/404.html"));
  },
};
```

Add `spaFallback` to the `plugins` array after `react()`.

- [ ] **Step 2: Create `.github/workflows/deploy-ui.yml`**

```yaml
name: Deploy UI

on:
  push:
    branches: [dev]
    paths:
      - "ui/**"
      - "abi/**"
      - "ignition/deployments/**"
      - ".github/workflows/deploy-ui.yml"
  workflow_dispatch:

permissions:
  contents: read
  pages: write
  id-token: write

concurrency:
  group: pages
  cancel-in-progress: true

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 20
      - name: Install UI deps
        working-directory: ui
        run: yarn install --frozen-lockfile
      - name: Build UI
        working-directory: ui
        run: yarn build
      - uses: actions/upload-pages-artifact@v3
        with:
          path: ui/dist

  deploy:
    needs: build
    runs-on: ubuntu-latest
    environment:
      name: github-pages
      url: ${{ steps.deployment.outputs.page_url }}
    steps:
      - id: deployment
        uses: actions/deploy-pages@v4
```

- [ ] **Step 3: Write `ui/README.md`**

Include: how to run locally (`yarn install && yarn dev`), that `yarn sync` regenerates contract data from `../abi` and `../ignition`, the network table, the custom-RPC note (public RPCs rate-limit `eth_getLogs`; bots/account listing rely on logs), and how the Pages deploy works (push to `dev`). Note that GitHub Pages must be set to "GitHub Actions" source in repo settings (manual one-time step).

- [ ] **Step 4: Verify the full build once more**

Run: `cd ui && yarn install --frozen-lockfile && yarn build`
Expected: succeeds; `ui/dist/index.html` and `ui/dist/404.html` exist.

- [ ] **Step 5: Commit**

```bash
git add .github/workflows/deploy-ui.yml ui/README.md ui/vite.config.ts ui/yarn.lock
git commit -m "ci(ui): add GitHub Pages deploy workflow and README"
```

---

## Self-Review Notes (coverage vs spec)

- **Networks (4, Columbus disabled):** Tasks 2–4. ✓
- **App-owned read transport / wallet for writes / follow wallet chain:** Tasks 4, 6, 10. ✓
- **Role-aware UI + shared tx component:** Tasks 5, 7, 8, 14. ✓
- **Sync script handling messy Columbus + missing Base mainnet:** Task 2. ✓
- **Dashboard (status + accounts):** Task 11. ✓
- **Create account:** Task 12. ✓
- **Workspace tabs — overview, bots, payment tokens, services (supported+wanted), roles, pubkeys, withdrawals:** Tasks 13, 15–19. ✓
- **Deferred v2 (manager admin, upgrades, booking token, restricted-rate/capabilities, cheque signing):** correctly omitted. ✓
- **GitHub Pages workflow + SPA fallback + base path:** Tasks 1, 20. ✓

**Open verification items the implementer MUST resolve by reading the contracts/ABIs (flagged inline, not guessed):** exact roles guarding `addSupportedToken`, `addPublicKey`; exact signatures of `addService`/`addWantedServices`/withdrawal functions. (Bots are confirmed role-based: `addMessengerBot`/`removeMessengerBot` manage `MESSENGER_BOT_ROLE` members, gas money in CAM — see Task 16.)
