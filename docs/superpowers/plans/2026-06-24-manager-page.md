# Manager Page Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a wallet-connected **Manager** page to the UI exposing `CMAccountManager` and `BookingToken` ecosystem-management functions (config, service registry, roles).

**Architecture:** A new `/manager` route renders a tabbed `ManagerWorkspace` mirroring the existing `AccountWorkspace`. A generic `RolesPanel` is extracted from the account `RolesTab` and reused for account, manager, and booking-token role sections. Reads use the app RPC keyed by `activeChainId`; writes go through `TxButton` → `useTx().track`; write controls are gated per-role with `RoleGate`.

**Tech Stack:** React 18 + Vite + TypeScript, wagmi v2 + viem v2, TanStack Query, Tailwind (class dark mode), react-router-dom, Vitest + Testing Library.

## Global Constraints

- All commands run from `ui/`. Tests: `./node_modules/.bin/vitest run <path>` (root `yarn test` runs Hardhat).
- Reads must pass `chainId: activeChainId` (from `useActiveChain`/`useActiveContracts`).
- All user-facing writes go through `TxButton` (never call `writeContractAsync` directly without tracking); use `onConfirmed` for side effects, no manual refetch timers.
- Permission-gated writes wrapped in `RoleGate` with a human `action` label.
- Every Tailwind color needs a `dark:` variant.
- Do NOT edit `src/contracts/generated/` (git-ignored, auto-synced). No contract or ABI changes are needed — all functions already exist in the generated ABIs.
- No `Co-Authored-By` trailer in commits. Format with `yarn format` (from `ui/`) before committing UI changes.
- Contract facts: `CMAccountManager` is `AccessControlEnumerable` (has `getRoleMembers`); `BookingToken` is plain `AccessControl` (NO `getRoleMembers`/`getRoleMemberCount` — cannot enumerate members).

---

### Task 1: Extend role definitions in `lib/roles.ts`

**Files:**
- Modify: `ui/src/lib/roles.ts`
- Test: `ui/src/lib/roles.test.ts`

**Interfaces:**
- Produces: `MANAGER_ROLES` (now includes `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`), `BOOKINGTOKEN_ROLES: readonly RoleName[]`, and `MIN_EXPIRATION_ADMIN_ROLE` present in `ROLE_HASHES`/`RoleName`.

- [ ] **Step 1: Write the failing test**

Append to `ui/src/lib/roles.test.ts`:

```ts
import { BOOKINGTOKEN_ROLES, MANAGER_ROLES, ROLE_HASHES } from "./roles";

describe("manager & booking token roles", () => {
  it("manager roles include admin, pauser, upgrader, versioner, service registry admin", () => {
    expect([...MANAGER_ROLES]).toEqual(
      expect.arrayContaining([
        "DEFAULT_ADMIN_ROLE",
        "PAUSER_ROLE",
        "UPGRADER_ROLE",
        "VERSIONER_ROLE",
        "SERVICE_REGISTRY_ADMIN_ROLE",
      ]),
    );
  });

  it("booking token roles include admin, upgrader, min expiration admin", () => {
    expect([...BOOKINGTOKEN_ROLES]).toEqual([
      "DEFAULT_ADMIN_ROLE",
      "UPGRADER_ROLE",
      "MIN_EXPIRATION_ADMIN_ROLE",
    ]);
  });

  it("hashes MIN_EXPIRATION_ADMIN_ROLE with keccak256", () => {
    expect(ROLE_HASHES.MIN_EXPIRATION_ADMIN_ROLE).toBe(keccak256(toBytes("MIN_EXPIRATION_ADMIN_ROLE")));
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/lib/roles.test.ts`
Expected: FAIL (`BOOKINGTOKEN_ROLES` undefined / `MIN_EXPIRATION_ADMIN_ROLE` missing).

- [ ] **Step 3: Implement the role changes**

Replace the `MANAGER_ROLES`/`RoleName`/`ROLE_HASHES` section of `ui/src/lib/roles.ts` so the file reads:

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
  "DEFAULT_ADMIN_ROLE",
  "PAUSER_ROLE",
  "UPGRADER_ROLE",
  "VERSIONER_ROLE",
  "SERVICE_REGISTRY_ADMIN_ROLE",
] as const;

export const BOOKINGTOKEN_ROLES = [
  "DEFAULT_ADMIN_ROLE",
  "UPGRADER_ROLE",
  "MIN_EXPIRATION_ADMIN_ROLE",
] as const;

export type RoleName =
  | (typeof ACCOUNT_ROLES)[number]
  | (typeof MANAGER_ROLES)[number]
  | (typeof BOOKINGTOKEN_ROLES)[number];

const ZERO_BYTES32 = `0x${"0".repeat(64)}` as Hex;

function compute(name: string): Hex {
  return name === "DEFAULT_ADMIN_ROLE" ? ZERO_BYTES32 : keccak256(toBytes(name));
}

export const ROLE_HASHES = Object.fromEntries(
  [...ACCOUNT_ROLES, ...MANAGER_ROLES, ...BOOKINGTOKEN_ROLES].map((r) => [r, compute(r)]),
) as Record<RoleName, Hex>;

export function roleHash(name: RoleName): Hex {
  return ROLE_HASHES[name];
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `./node_modules/.bin/vitest run src/lib/roles.test.ts`
Expected: PASS.

- [ ] **Step 5: Commit**

```bash
yarn format
git add src/lib/roles.ts src/lib/roles.test.ts
git commit -m "feat(ui): add manager and booking token role definitions"
```

---

### Task 2: Extract generic `RolesPanel` and rewire account `RolesTab`

**Files:**
- Create: `ui/src/components/RolesPanel.tsx`
- Modify: `ui/src/pages/tabs/RolesTab.tsx`
- Test: `ui/src/components/RolesPanel.test.tsx`

**Interfaces:**
- Consumes: `useRoleMembers`, `useHasRole`, `roleHash`, `RoleName`, `shortRoleName`, `TxButton`, `RoleGate`, `RowAction`, `AddressDisplay`.
- Produces: `RolesPanel({ address, abi, roles, enumerable })` — `enumerable` true renders per-role member lists + grant/revoke; false renders grant/revoke-by-address forms + "you hold this role" indicator and no member list. Grant/revoke gated on `DEFAULT_ADMIN_ROLE` of `address`.

- [ ] **Step 1: Write the failing test**

Create `ui/src/components/RolesPanel.test.tsx`:

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RolesPanel } from "./RolesPanel";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

const addr = "0x1111111111111111111111111111111111111111" as const;

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("RolesPanel", () => {
  it("renders a row per role (enumerable)", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE", "PAUSER_ROLE"]} enumerable />);
    expect(screen.getByText("Admin")).toBeInTheDocument();
    expect(screen.getByText("Pauser")).toBeInTheDocument();
  });

  it("notes members are not listable in non-enumerable mode", () => {
    wrap(<RolesPanel address={addr} abi={[]} roles={["DEFAULT_ADMIN_ROLE"]} enumerable={false} />);
    expect(screen.getByText(/cannot list/i)).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/components/RolesPanel.test.tsx`
Expected: FAIL (`RolesPanel` does not exist).

- [ ] **Step 3: Create `RolesPanel.tsx`**

Create `ui/src/components/RolesPanel.tsx`:

```tsx
import { useState } from "react";
import { ChevronRight, ShieldCheck, ShieldPlus, Trash2 } from "lucide-react";
import { type Abi, type Address } from "viem";
import { useWriteContract } from "wagmi";
import { AddressDisplay } from "./AddressDisplay";
import { RoleGate } from "./RoleGate";
import { RowAction } from "./RowAction";
import { TxButton } from "./TxButton";
import { useHasRole } from "../hooks/useHasRole";
import { useRoleMembers } from "../hooks/useRoleMembers";
import { roleHash, type RoleName } from "../lib/roles";
import { shortRoleName } from "../lib/format";

const inputClass =
  "rounded border border-gray-300 bg-white px-2 py-1.5 text-sm focus:border-indigo-500 focus:outline-none dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100";

function RoleHeader({ label, role, open, badge }: { label: string; role: RoleName; open: boolean; badge: React.ReactNode }) {
  return (
    <>
      <ChevronRight className={`h-4 w-4 shrink-0 text-gray-400 transition-transform ${open ? "rotate-90" : ""}`} />
      <span className="flex-1 text-sm font-medium text-gray-800 dark:text-gray-100">{label}</span>
      <span className="font-mono text-[11px] text-gray-400">{role}</span>
      {badge}
    </>
  );
}

function GrantForm({ account, abi, role, label, onDone }: { account: Address; abi: Abi; role: RoleName; label: string; onDone: () => void }) {
  const { writeContractAsync } = useWriteContract();
  const [grantee, setGrantee] = useState("");
  return (
    <div className="flex items-end gap-2">
      <input className={`flex-1 ${inputClass}`} placeholder="Address 0x…" value={grantee} onChange={(e) => setGrantee(e.target.value)} />
      <TxButton
        label="Grant"
        icon={<ShieldPlus className="h-4 w-4" />}
        disabled={!grantee.trim()}
        tooltip={`Grants ${label} to this address — sends a transaction to your wallet.`}
        write={() => writeContractAsync({ address: account, abi, functionName: "grantRole", args: [roleHash(role), grantee.trim() as Address] })}
        onConfirmed={() => { setGrantee(""); onDone(); }}
      />
    </div>
  );
}

function EnumerableRoleRow({ account, abi, role, hasAdmin, open, onToggle }: {
  account: Address; abi: Abi; role: RoleName; hasAdmin: boolean; open: boolean; onToggle: () => void;
}) {
  const { writeContractAsync } = useWriteContract();
  const { members, isLoading, refetch } = useRoleMembers(account, abi, role);
  const label = shortRoleName(role);

  return (
    <li className="rounded-md border border-gray-100 dark:border-gray-700/60">
      <button type="button" onClick={onToggle} aria-expanded={open} className="flex w-full items-center gap-2 px-3 py-2 text-left">
        <RoleHeader label={label} role={role} open={open} badge={
          <span
            className={`rounded-full px-2 py-0.5 text-xs font-medium ${!isLoading && members.length > 0
              ? "bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300"
              : "bg-gray-100 text-gray-500 dark:bg-gray-700 dark:text-gray-400"}`}
            title={`${members.length} address${members.length === 1 ? "" : "es"} with this role`}
          >
            {isLoading ? "…" : members.length}
          </span>
        } />
      </button>
      {open && (
        <div className="space-y-3 border-t border-gray-100 px-3 py-3 dark:border-gray-700/60">
          <ul className="divide-y dark:divide-gray-700">
            {isLoading && <li className="py-2 text-sm text-gray-400">Loading…</li>}
            {!isLoading && members.length === 0 && <li className="py-2 text-sm text-gray-400">No members</li>}
            {members.map((m) => (
              <li key={m} className="group flex items-center justify-between gap-3 py-2">
                <AddressDisplay address={m} className="text-sm" />
                {hasAdmin && (
                  <RowAction>
                    <TxButton
                      label="Revoke" variant="danger" icon={<Trash2 className="h-4 w-4" />}
                      tooltip={`Revokes ${label} from this address — sends a transaction to your wallet.`}
                      write={() => writeContractAsync({ address: account, abi, functionName: "revokeRole", args: [roleHash(role), m as Address] })}
                      onConfirmed={refetch}
                    />
                  </RowAction>
                )}
              </li>
            ))}
          </ul>
          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`grant ${label}`}>
            <GrantForm account={account} abi={abi} role={role} label={label} onDone={refetch} />
          </RoleGate>
        </div>
      )}
    </li>
  );
}

function NonEnumerableRoleRow({ account, abi, role, hasAdmin, open, onToggle }: {
  account: Address; abi: Abi; role: RoleName; hasAdmin: boolean; open: boolean; onToggle: () => void;
}) {
  const { writeContractAsync } = useWriteContract();
  const { hasRole: youHold } = useHasRole(account, abi, role);
  const [revokee, setRevokee] = useState("");
  const label = shortRoleName(role);

  return (
    <li className="rounded-md border border-gray-100 dark:border-gray-700/60">
      <button type="button" onClick={onToggle} aria-expanded={open} className="flex w-full items-center gap-2 px-3 py-2 text-left">
        <RoleHeader label={label} role={role} open={open} badge={
          youHold ? (
            <span className="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2 py-0.5 text-xs font-medium text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300">
              <ShieldCheck className="h-3 w-3" /> You
            </span>
          ) : <span className="w-px" />
        } />
      </button>
      {open && (
        <div className="space-y-3 border-t border-gray-100 px-3 py-3 dark:border-gray-700/60">
          <p className="text-xs text-gray-400">This contract cannot list role members on-chain. Grant or revoke by address.</p>
          <RoleGate hasRole={hasAdmin} roleName="DEFAULT_ADMIN_ROLE" action={`manage ${label}`}>
            <div className="space-y-2">
              <GrantForm account={account} abi={abi} role={role} label={label} onDone={() => {}} />
              <div className="flex items-end gap-2">
                <input className={`flex-1 ${inputClass}`} placeholder="Address 0x… to revoke" value={revokee} onChange={(e) => setRevokee(e.target.value)} />
                <TxButton
                  label="Revoke" variant="danger" icon={<Trash2 className="h-4 w-4" />}
                  disabled={!revokee.trim()}
                  tooltip={`Revokes ${label} from this address — sends a transaction to your wallet.`}
                  write={() => writeContractAsync({ address: account, abi, functionName: "revokeRole", args: [roleHash(role), revokee.trim() as Address] })}
                  onConfirmed={() => setRevokee("")}
                />
              </div>
            </div>
          </RoleGate>
        </div>
      )}
    </li>
  );
}

export function RolesPanel({ address, abi, roles, enumerable }: {
  address: Address; abi: Abi; roles: readonly RoleName[]; enumerable: boolean;
}) {
  const { hasRole: hasAdmin } = useHasRole(address, abi, "DEFAULT_ADMIN_ROLE");
  const [openRole, setOpenRole] = useState<RoleName | null>(null);
  const Row = enumerable ? EnumerableRoleRow : NonEnumerableRoleRow;

  return (
    <ul className="space-y-2">
      {roles.map((r) => (
        <Row
          key={r}
          account={address}
          abi={abi}
          role={r}
          hasAdmin={hasAdmin}
          open={openRole === r}
          onToggle={() => setOpenRole((cur) => (cur === r ? null : r))}
        />
      ))}
    </ul>
  );
}
```

- [ ] **Step 4: Rewire account `RolesTab` to use `RolesPanel`**

Replace `ui/src/pages/tabs/RolesTab.tsx` entirely with:

```tsx
import { type Abi, type Address } from "viem";
import { Card } from "../../components/Card";
import { RolesPanel } from "../../components/RolesPanel";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { ACCOUNT_ROLES } from "../../lib/roles";

export function RolesTab({ account }: { account: Address }) {
  const { cmAccountAbi } = useActiveContracts();
  return (
    <Card title="Roles">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Expand a role to see its members and grant or revoke it. Only one role is open at a time.
      </p>
      <RolesPanel address={account} abi={cmAccountAbi as Abi} roles={ACCOUNT_ROLES} enumerable />
    </Card>
  );
}
```

- [ ] **Step 5: Run tests + typecheck**

Run: `./node_modules/.bin/vitest run src/components/RolesPanel.test.tsx && ./node_modules/.bin/tsc -b`
Expected: PASS, no type errors.

- [ ] **Step 6: Commit**

```bash
yarn format
git add src/components/RolesPanel.tsx src/components/RolesPanel.test.tsx src/pages/tabs/RolesTab.tsx
git commit -m "refactor(ui): extract reusable RolesPanel from account RolesTab"
```

---

### Task 3: `ManagerWorkspace` page, route, and nav link

**Files:**
- Create: `ui/src/pages/ManagerWorkspace.tsx`
- Create: `ui/src/components/ManagerSummary.tsx`
- Create: `ui/src/pages/ManagerWorkspace.test.tsx`
- Modify: `ui/src/App.tsx`
- Modify: `ui/src/components/Layout.tsx`

**Interfaces:**
- Consumes: `useActiveContracts`, `AddressDisplay`, `TxPanel`, `RefreshButton`.
- Produces: `ManagerWorkspace` page at `/manager` with tabs `config`, `services`, `roles`, `booking-token` selected via `?tab=`. Tab components (`ManagerConfigTab`, `ServiceRegistryTab`, `ManagerRolesTab`, `BookingTokenTab`) are imported from later tasks; this task creates placeholder tab components inline as separate files? No — tab files are created in Tasks 4-7. To keep this task self-contained, it ships the workspace shell with stub tab components defined locally, replaced by Tasks 4-7.

> Implementation note: create the four tab files as **minimal stubs** in this task (each a `Card` saying "…"), so the workspace compiles and routes. Tasks 4–7 replace each stub's body. This keeps every task independently testable.

- [ ] **Step 1: Write the failing test**

Create `ui/src/pages/ManagerWorkspace.test.tsx`:

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ManagerWorkspace } from "./ManagerWorkspace";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useBalance: () => ({ data: undefined }),
  useReadContract: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useReadContracts: () => ({ data: undefined, isLoading: false, refetch: vi.fn() }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn(), items: [] }) }));

describe("ManagerWorkspace", () => {
  it("renders the manager tab bar", () => {
    render(
      <QueryClientProvider client={new QueryClient()}>
        <MemoryRouter initialEntries={["/manager"]}>
          <Routes><Route path="manager" element={<ManagerWorkspace />} /></Routes>
        </MemoryRouter>
      </QueryClientProvider>,
    );
    expect(screen.getByRole("link", { name: /service registry/i })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /booking token/i })).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/pages/ManagerWorkspace.test.tsx`
Expected: FAIL (`ManagerWorkspace` does not exist).

- [ ] **Step 3: Create stub tab files**

Create each of these four files with a stub body (replaced in Tasks 4–7):

`ui/src/pages/tabs/ManagerConfigTab.tsx`:
```tsx
import { Card } from "../../components/Card";
export function ManagerConfigTab() {
  return <Card title="Manager Configuration">…</Card>;
}
```

`ui/src/pages/tabs/ServiceRegistryTab.tsx`:
```tsx
import { Card } from "../../components/Card";
export function ServiceRegistryTab() {
  return <Card title="Service Registry">…</Card>;
}
```

`ui/src/pages/tabs/ManagerRolesTab.tsx`:
```tsx
import { Card } from "../../components/Card";
export function ManagerRolesTab() {
  return <Card title="Manager Roles">…</Card>;
}
```

`ui/src/pages/tabs/BookingTokenTab.tsx`:
```tsx
import { Card } from "../../components/Card";
export function BookingTokenTab() {
  return <Card title="Booking Token">…</Card>;
}
```

- [ ] **Step 4: Create `ManagerSummary.tsx`**

Create `ui/src/components/ManagerSummary.tsx`:

```tsx
import { type Abi } from "viem";
import { useReadContract } from "wagmi";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { AddressDisplay } from "./AddressDisplay";

export function ManagerSummary() {
  const { manager, bookingToken, managerAbi, chainId, supported } = useActiveContracts();
  const { data: paused } = useReadContract({ chainId, address: manager, abi: managerAbi as Abi, functionName: "paused", query: { enabled: supported } });

  return (
    <aside className="h-fit rounded-lg border border-gray-200 bg-white p-4 shadow-sm dark:border-gray-800 dark:bg-gray-800">
      <h2 className="mb-3 text-xs font-semibold uppercase tracking-wide text-gray-400">Ecosystem</h2>
      <dl className="grid grid-cols-1 gap-3 text-sm">
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Manager</dt>
          <dd>{manager ? <AddressDisplay address={manager} /> : "—"}</dd>
        </div>
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Booking token</dt>
          <dd>{bookingToken ? <AddressDisplay address={bookingToken} /> : "—"}</dd>
        </div>
        <div>
          <dt className="text-gray-500 dark:text-gray-400">Account creation</dt>
          <dd>{paused ? "Paused" : "Active"}</dd>
        </div>
      </dl>
    </aside>
  );
}
```

- [ ] **Step 5: Create `ManagerWorkspace.tsx`**

Create `ui/src/pages/ManagerWorkspace.tsx`:

```tsx
import { Link, useSearchParams } from "react-router-dom";
import { KeyRound, Server, Settings, Ticket } from "lucide-react";
import { Card } from "../components/Card";
import { ManagerSummary } from "../components/ManagerSummary";
import { RefreshButton } from "../components/RefreshButton";
import { TxPanel } from "../components/TxPanel";
import { useActiveContracts } from "../hooks/useActiveContracts";
import { ManagerConfigTab } from "./tabs/ManagerConfigTab";
import { ServiceRegistryTab } from "./tabs/ServiceRegistryTab";
import { ManagerRolesTab } from "./tabs/ManagerRolesTab";
import { BookingTokenTab } from "./tabs/BookingTokenTab";

const TABS = [
  { id: "config", label: "Config", Icon: Settings, Component: ManagerConfigTab },
  { id: "services", label: "Service Registry", Icon: Server, Component: ServiceRegistryTab },
  { id: "roles", label: "Manager Roles", Icon: KeyRound, Component: ManagerRolesTab },
  { id: "booking-token", label: "Booking Token", Icon: Ticket, Component: BookingTokenTab },
] as const;

export function ManagerWorkspace() {
  const { supported } = useActiveContracts();
  const [params] = useSearchParams();
  const active = params.get("tab") ?? TABS[0].id;
  const Active = (TABS.find((t) => t.id === active) ?? TABS[0]).Component;

  if (!supported) return <Card title="Manager">Connect to a supported network.</Card>;

  return (
    <div className="grid items-start gap-6 md:grid-cols-[260px_1fr]">
      <div className="flex flex-col gap-4">
        <ManagerSummary />
        <TxPanel />
      </div>
      <div className="flex min-w-0 flex-col gap-4">
        <div className="flex items-end justify-between gap-3 border-b text-sm dark:border-gray-800">
          <nav className="flex flex-wrap gap-3">
            {TABS.map((t) => (
              <Link
                key={t.id}
                to={`?tab=${t.id}`}
                className={`-mb-px inline-flex items-center gap-1.5 border-b-2 pb-2 ${
                  active === t.id
                    ? "border-indigo-600 font-medium text-gray-900 dark:text-gray-100"
                    : "border-transparent text-gray-500 dark:text-gray-400"
                }`}
              >
                <t.Icon className="h-4 w-4" /> {t.label}
              </Link>
            ))}
          </nav>
          <div className="pb-1.5">
            <RefreshButton />
          </div>
        </div>
        <Active />
      </div>
    </div>
  );
}
```

- [ ] **Step 6: Register route in `App.tsx`**

In `ui/src/App.tsx`, add the import and route:

```tsx
import { ManagerWorkspace } from "./pages/ManagerWorkspace";
```
and inside `<Route element={<Layout />}>`, after the `create` route:
```tsx
        <Route path="manager" element={<ManagerWorkspace />} />
```

- [ ] **Step 7: Add nav link in `Layout.tsx`**

In `ui/src/components/Layout.tsx`, add between the Dashboard and Create Account `NavLink`s:

```tsx
          <NavLink to="/manager" className={navLinkClass}>Manager</NavLink>
```

- [ ] **Step 8: Run test + typecheck**

Run: `./node_modules/.bin/vitest run src/pages/ManagerWorkspace.test.tsx && ./node_modules/.bin/tsc -b`
Expected: PASS, no type errors.

- [ ] **Step 9: Commit**

```bash
yarn format
git add src/pages/ManagerWorkspace.tsx src/pages/ManagerWorkspace.test.tsx src/components/ManagerSummary.tsx src/pages/tabs/ManagerConfigTab.tsx src/pages/tabs/ServiceRegistryTab.tsx src/pages/tabs/ManagerRolesTab.tsx src/pages/tabs/BookingTokenTab.tsx src/App.tsx src/components/Layout.tsx
git commit -m "feat(ui): add Manager page shell, route, and nav link"
```

---

### Task 4: `ManagerConfigTab` — pause, implementation, booking token address

**Files:**
- Modify: `ui/src/pages/tabs/ManagerConfigTab.tsx`
- Test: `ui/src/pages/tabs/ManagerConfigTab.test.tsx`

**Interfaces:**
- Consumes: `useActiveContracts` (`manager`, `managerAbi`, `chainId`, `supported`), `useHasRole`, `useReadContract`, `useWriteContract`, `RoleGate`, `TxButton`, `AddressDisplay`, `Card`.
- Reads: `paused`, `getAccountImplementation`, `getBookingTokenAddress`. Writes: `pause`/`unpause` (`PAUSER_ROLE`), `setAccountImplementation` (`VERSIONER_ROLE`), `setBookingTokenAddress` (`VERSIONER_ROLE`).

- [ ] **Step 1: Write the failing test**

Create `ui/src/pages/tabs/ManagerConfigTab.test.tsx`:

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ManagerConfigTab } from "./ManagerConfigTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => ({
    data: functionName === "paused" ? false : undefined,
    isLoading: false,
    refetch: vi.fn(),
  }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("ManagerConfigTab", () => {
  it("shows pause status and config sections", () => {
    wrap(<ManagerConfigTab />);
    expect(screen.getByText(/account creation/i)).toBeInTheDocument();
    expect(screen.getByText(/account implementation/i)).toBeInTheDocument();
    expect(screen.getByText(/booking token address/i)).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/pages/tabs/ManagerConfigTab.test.tsx`
Expected: FAIL (stub renders only "…").

- [ ] **Step 3: Implement `ManagerConfigTab.tsx`**

Replace `ui/src/pages/tabs/ManagerConfigTab.tsx` with:

```tsx
import { useState } from "react";
import { Pause, Play, Save } from "lucide-react";
import { type Abi, type Address } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";

const inputClass =
  "rounded border border-gray-300 bg-white px-2 py-1.5 text-sm focus:border-indigo-500 focus:outline-none dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100";

function AddressSetting({ title, current, functionName, roleName, action, isLoading, refetch }: {
  title: string; current: Address | undefined; functionName: string; roleName: "VERSIONER_ROLE"; action: string;
  isLoading: boolean; refetch: () => void;
}) {
  const { manager, managerAbi } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { hasRole } = useHasRole(manager, abi, roleName);
  const { writeContractAsync } = useWriteContract();
  const [value, setValue] = useState("");

  return (
    <Card title={title}>
      <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
        <dt className="text-gray-500 dark:text-gray-400">Current</dt>
        <dd>{isLoading ? "…" : current ? <AddressDisplay address={current} /> : "—"}</dd>
      </dl>
      <RoleGate hasRole={hasRole} roleName={roleName} action={action}>
        <div className="flex items-end gap-2">
          <input className={`flex-1 ${inputClass}`} placeholder="0x…" value={value} onChange={(e) => setValue(e.target.value)} />
          <TxButton
            label="Save" icon={<Save className="h-4 w-4" />} disabled={!value.trim()}
            tooltip={`${action} — sends a transaction to your wallet.`}
            write={() => writeContractAsync({ address: manager!, abi, functionName, args: [value.trim() as Address] })}
            onConfirmed={() => { setValue(""); refetch(); }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}

export function ManagerConfigTab() {
  const { manager, managerAbi, chainId, supported } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const { hasRole: canPause } = useHasRole(manager, abi, "PAUSER_ROLE");

  const { data: paused, isLoading: pausedLoading, refetch: refetchPaused } =
    useReadContract({ chainId, address: manager, abi, functionName: "paused", query: { enabled: supported } });
  const { data: impl, isLoading: implLoading, refetch: refetchImpl } =
    useReadContract({ chainId, address: manager, abi, functionName: "getAccountImplementation", query: { enabled: supported } });
  const { data: btoken, isLoading: btokenLoading, refetch: refetchBtoken } =
    useReadContract({ chainId, address: manager, abi, functionName: "getBookingTokenAddress", query: { enabled: supported } });

  const isPaused = paused === true;

  return (
    <div className="grid gap-4">
      <Card title="Account creation">
        <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-gray-500 dark:text-gray-400">Status</dt>
          <dd>{pausedLoading ? "…" : isPaused ? "Paused" : "Active"}</dd>
        </dl>
        <RoleGate hasRole={canPause} roleName="PAUSER_ROLE" action={isPaused ? "unpause manager" : "pause manager"}>
          <TxButton
            label={isPaused ? "Unpause" : "Pause"}
            variant={isPaused ? "primary" : "danger"}
            icon={isPaused ? <Play className="h-4 w-4" /> : <Pause className="h-4 w-4" />}
            tooltip={isPaused
              ? "Resumes CM account creation — sends a transaction to your wallet."
              : "Pauses CM account creation — sends a transaction to your wallet."}
            write={() => writeContractAsync({ address: manager!, abi, functionName: isPaused ? "unpause" : "pause", args: [] })}
            onConfirmed={() => refetchPaused()}
          />
        </RoleGate>
      </Card>

      <AddressSetting
        title="Account implementation"
        current={impl as Address | undefined}
        functionName="setAccountImplementation"
        roleName="VERSIONER_ROLE"
        action="set account implementation"
        isLoading={implLoading}
        refetch={() => refetchImpl()}
      />

      <AddressSetting
        title="Booking token address"
        current={btoken as Address | undefined}
        functionName="setBookingTokenAddress"
        roleName="VERSIONER_ROLE"
        action="set booking token address"
        isLoading={btokenLoading}
        refetch={() => refetchBtoken()}
      />
    </div>
  );
}
```

- [ ] **Step 4: Run test + typecheck**

Run: `./node_modules/.bin/vitest run src/pages/tabs/ManagerConfigTab.test.tsx && ./node_modules/.bin/tsc -b`
Expected: PASS, no type errors.

- [ ] **Step 5: Commit**

```bash
yarn format
git add src/pages/tabs/ManagerConfigTab.tsx src/pages/tabs/ManagerConfigTab.test.tsx
git commit -m "feat(ui): manager config tab (pause, implementation, booking token)"
```

---

### Task 5: `ServiceRegistryTab` — register / unregister services

**Files:**
- Modify: `ui/src/pages/tabs/ServiceRegistryTab.tsx`
- Test: `ui/src/pages/tabs/ServiceRegistryTab.test.tsx`

**Interfaces:**
- Consumes: `useActiveContracts`, `useHasRole`, `useReadContract`, `useWriteContract`, `groupServicesByPackage`, `Card`, `RoleGate`, `TxButton`, `RowAction`, `CopyButton`.
- Reads: `getAllRegisteredServiceNames`. Writes: `registerService(string)` / `unregisterService(string)` (`SERVICE_REGISTRY_ADMIN_ROLE`).

- [ ] **Step 1: Write the failing test**

Create `ui/src/pages/tabs/ServiceRegistryTab.test.tsx`:

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ServiceRegistryTab } from "./ServiceRegistryTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => ({
    data: functionName === "getAllRegisteredServiceNames"
      ? ["cmp.services.accommodation.v1.AccommodationSearchService"]
      : undefined,
    isLoading: false,
    refetch: vi.fn(),
  }),
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("ServiceRegistryTab", () => {
  it("lists registered services", () => {
    wrap(<ServiceRegistryTab />);
    expect(screen.getByText("AccommodationSearchService")).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/pages/tabs/ServiceRegistryTab.test.tsx`
Expected: FAIL (stub renders "…").

- [ ] **Step 3: Implement `ServiceRegistryTab.tsx`**

Replace `ui/src/pages/tabs/ServiceRegistryTab.tsx` with:

```tsx
import { useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { type Abi } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { Card } from "../../components/Card";
import { CopyButton } from "../../components/CopyButton";
import { RoleGate } from "../../components/RoleGate";
import { RowAction } from "../../components/RowAction";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { groupServicesByPackage, parseServiceName } from "../../lib/serviceName";

const inputClass =
  "rounded border border-gray-300 bg-white px-2 py-1.5 text-sm focus:border-indigo-500 focus:outline-none dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100";

export function ServiceRegistryTab() {
  const { manager, managerAbi, chainId, supported } = useActiveContracts();
  const abi = managerAbi as Abi;
  const { hasRole } = useHasRole(manager, abi, "SERVICE_REGISTRY_ADMIN_ROLE");
  const { writeContractAsync } = useWriteContract();
  const { data, isLoading, refetch } = useReadContract({
    chainId, address: manager, abi, functionName: "getAllRegisteredServiceNames", query: { enabled: supported },
  });
  const names = (data as string[] | undefined) ?? [];
  const groups = groupServicesByPackage(names.map((n) => ({ name: n })));
  const [name, setName] = useState("");

  return (
    <Card title="Service Registry">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Services must be registered here before any CM Account can support or want them.
      </p>
      {isLoading ? <p className="py-2 text-sm text-gray-400">Loading…</p> : names.length === 0 ? (
        <p className="mb-4 py-2 text-sm text-gray-400">No services registered.</p>
      ) : (
        <div className="mb-4 space-y-5">
          {groups.map((g) => (
            <div key={g.pkg}>
              <h4 className="mb-1.5 text-xs font-semibold uppercase tracking-wide text-gray-500 dark:text-gray-400">{g.pkg} <span className="text-gray-400">{g.items.length}</span></h4>
              <ul className="space-y-2">
                {g.items.map((s) => {
                  const parsed = parseServiceName(s.name);
                  return (
                    <li key={s.name} className="group flex items-center justify-between gap-3 rounded-md border border-gray-100 px-3 py-2 dark:border-gray-700/60">
                      <span className="flex min-w-0 items-baseline gap-2">
                        {parsed.version && <span className="rounded bg-indigo-50 px-1.5 py-0.5 font-mono text-xs font-medium text-indigo-700 dark:bg-indigo-950 dark:text-indigo-300">{parsed.version}</span>}
                        <span className="break-all font-mono text-sm font-medium text-gray-900 dark:text-gray-100">{parsed.name}</span>
                        <CopyButton value={s.name} label="Copy full service name" />
                      </span>
                      {hasRole && (
                        <RowAction>
                          <TxButton
                            label="Unregister" variant="danger" icon={<Trash2 className="h-4 w-4" />}
                            tooltip="Unregisters this service from the manager — sends a transaction to your wallet."
                            write={() => writeContractAsync({ address: manager!, abi, functionName: "unregisterService", args: [s.name] })}
                            onConfirmed={() => refetch()}
                          />
                        </RowAction>
                      )}
                    </li>
                  );
                })}
              </ul>
            </div>
          ))}
        </div>
      )}
      <RoleGate hasRole={hasRole} roleName="SERVICE_REGISTRY_ADMIN_ROLE" action="Register service">
        <div className="flex items-end gap-2">
          <input
            className={`flex-1 ${inputClass}`}
            placeholder="cmp.services.<package>.<version>.<Name>"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <TxButton
            label="Register" icon={<Plus className="h-4 w-4" />} disabled={!name.trim()}
            tooltip="Registers a new service name in the manager — sends a transaction to your wallet."
            write={() => writeContractAsync({ address: manager!, abi, functionName: "registerService", args: [name.trim()] })}
            onConfirmed={() => { setName(""); refetch(); }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}
```

> Note: `parseServiceName` and `groupServicesByPackage` are both exported from `ui/src/lib/serviceName.ts` (verified). `groupServicesByPackage` returns groups with `.pkg` and `.items` (each `{ name }`), as used in `ServicesTab.tsx`.

- [ ] **Step 4: Run test + typecheck**

Run: `./node_modules/.bin/vitest run src/pages/tabs/ServiceRegistryTab.test.tsx && ./node_modules/.bin/tsc -b`
Expected: PASS, no type errors.

- [ ] **Step 5: Commit**

```bash
yarn format
git add src/pages/tabs/ServiceRegistryTab.tsx src/pages/tabs/ServiceRegistryTab.test.tsx
git commit -m "feat(ui): service registry tab (register/unregister services)"
```

---

### Task 6: `ManagerRolesTab` — manager role members + grant/revoke

**Files:**
- Modify: `ui/src/pages/tabs/ManagerRolesTab.tsx`

**Interfaces:**
- Consumes: `useActiveContracts` (`manager`, `managerAbi`), `RolesPanel`, `MANAGER_ROLES`, `Card`.

- [ ] **Step 1: Implement `ManagerRolesTab.tsx`**

Replace `ui/src/pages/tabs/ManagerRolesTab.tsx` with:

```tsx
import { type Abi } from "viem";
import { Card } from "../../components/Card";
import { RolesPanel } from "../../components/RolesPanel";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { MANAGER_ROLES } from "../../lib/roles";

export function ManagerRolesTab() {
  const { manager, managerAbi } = useActiveContracts();
  if (!manager) return <Card title="Manager Roles">Connect to a supported network.</Card>;
  return (
    <Card title="Manager Roles">
      <p className="mb-3 text-xs text-gray-500 dark:text-gray-400">
        Expand a role to see its members and grant or revoke it. Grant/revoke requires the Admin role.
      </p>
      <RolesPanel address={manager} abi={managerAbi as Abi} roles={MANAGER_ROLES} enumerable />
    </Card>
  );
}
```

- [ ] **Step 2: Verify it renders in the workspace test**

Run: `./node_modules/.bin/vitest run src/pages/ManagerWorkspace.test.tsx && ./node_modules/.bin/tsc -b`
Expected: PASS, no type errors.

- [ ] **Step 3: Commit**

```bash
yarn format
git add src/pages/tabs/ManagerRolesTab.tsx
git commit -m "feat(ui): manager roles tab"
```

---

### Task 7: `BookingTokenTab` — info, settings, roles

**Files:**
- Modify: `ui/src/pages/tabs/BookingTokenTab.tsx`
- Test: `ui/src/pages/tabs/BookingTokenTab.test.tsx`

**Interfaces:**
- Consumes: `useActiveContracts` (`bookingToken`, `bookingTokenAbi`, `chainId`, `supported`), `useHasRole`, `useReadContract`, `useWriteContract`, `RolesPanel`, `BOOKINGTOKEN_ROLES`, `RoleGate`, `TxButton`, `AddressDisplay`, `Card`.
- Reads: `name`, `symbol`, `version` (returns `[major, minor, patch]`), `getManagerAddress`, `getMinExpirationTimestampDiff`. Writes: `setManagerAddress` (`DEFAULT_ADMIN_ROLE`), `setMinExpirationTimestampDiff` (`MIN_EXPIRATION_ADMIN_ROLE`).

- [ ] **Step 1: Write the failing test**

Create `ui/src/pages/tabs/BookingTokenTab.test.tsx`:

```tsx
import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BookingTokenTab } from "./BookingTokenTab";

vi.mock("wagmi", () => ({
  useAccount: () => ({ address: undefined }),
  useReadContract: ({ functionName }: { functionName: string }) => {
    const map: Record<string, unknown> = {
      name: "BookingToken",
      symbol: "TRIP",
      version: [1n, 0n, 0n],
      getManagerAddress: "0x2222222222222222222222222222222222222222",
      getMinExpirationTimestampDiff: 60n,
    };
    return { data: map[functionName], isLoading: false, refetch: vi.fn() };
  },
  useWriteContract: () => ({ writeContractAsync: vi.fn() }),
}));
vi.mock("../../wallet/activeChain", () => ({
  useActiveChain: () => ({ activeChainId: 84532, setActiveChainId: vi.fn() }),
}));
vi.mock("../../tx/TxProvider", () => ({ useTx: () => ({ track: vi.fn() }) }));

function wrap(ui: React.ReactNode) {
  return render(<QueryClientProvider client={new QueryClient()}>{ui}</QueryClientProvider>);
}

describe("BookingTokenTab", () => {
  it("shows token info and settings", () => {
    wrap(<BookingTokenTab />);
    expect(screen.getByText("BookingToken")).toBeInTheDocument();
    expect(screen.getByText("TRIP")).toBeInTheDocument();
    expect(screen.getByText(/min expiration/i)).toBeInTheDocument();
  });
});
```

- [ ] **Step 2: Run test to verify it fails**

Run: `./node_modules/.bin/vitest run src/pages/tabs/BookingTokenTab.test.tsx`
Expected: FAIL (stub renders "…").

- [ ] **Step 3: Implement `BookingTokenTab.tsx`**

Replace `ui/src/pages/tabs/BookingTokenTab.tsx` with:

```tsx
import { useState } from "react";
import { Save } from "lucide-react";
import { type Abi, type Address } from "viem";
import { useReadContract, useWriteContract } from "wagmi";
import { AddressDisplay } from "../../components/AddressDisplay";
import { Card } from "../../components/Card";
import { RoleGate } from "../../components/RoleGate";
import { RolesPanel } from "../../components/RolesPanel";
import { TxButton } from "../../components/TxButton";
import { useActiveContracts } from "../../hooks/useActiveContracts";
import { useHasRole } from "../../hooks/useHasRole";
import { BOOKINGTOKEN_ROLES } from "../../lib/roles";

const inputClass =
  "rounded border border-gray-300 bg-white px-2 py-1.5 text-sm focus:border-indigo-500 focus:outline-none dark:border-gray-600 dark:bg-gray-900 dark:text-gray-100";

function formatVersion(v: unknown): string {
  if (!Array.isArray(v) || v.length < 3) return "—";
  return `${v[0]}.${v[1]}.${v[2]}`;
}

export function BookingTokenTab() {
  const { bookingToken, bookingTokenAbi, chainId, supported } = useActiveContracts();
  const abi = bookingTokenAbi as Abi;
  const { writeContractAsync } = useWriteContract();
  const enabled = { query: { enabled: supported } };

  const { data: name } = useReadContract({ chainId, address: bookingToken, abi, functionName: "name", ...enabled });
  const { data: symbol } = useReadContract({ chainId, address: bookingToken, abi, functionName: "symbol", ...enabled });
  const { data: version } = useReadContract({ chainId, address: bookingToken, abi, functionName: "version", ...enabled });
  const { data: managerAddr, refetch: refetchManager } = useReadContract({ chainId, address: bookingToken, abi, functionName: "getManagerAddress", ...enabled });
  const { data: minDiff, refetch: refetchMinDiff } = useReadContract({ chainId, address: bookingToken, abi, functionName: "getMinExpirationTimestampDiff", ...enabled });

  const { hasRole: isAdmin } = useHasRole(bookingToken, abi, "DEFAULT_ADMIN_ROLE");
  const { hasRole: canSetMin } = useHasRole(bookingToken, abi, "MIN_EXPIRATION_ADMIN_ROLE");

  const [newManager, setNewManager] = useState("");
  const [newMin, setNewMin] = useState("");

  if (!bookingToken) return <Card title="Booking Token">Connect to a supported network.</Card>;

  return (
    <div className="grid gap-4">
      <Card title="Booking Token">
        <dl className="grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-gray-500 dark:text-gray-400">Address</dt>
          <dd><AddressDisplay address={bookingToken} /></dd>
          <dt className="text-gray-500 dark:text-gray-400">Name</dt>
          <dd>{(name as string) ?? "—"}</dd>
          <dt className="text-gray-500 dark:text-gray-400">Symbol</dt>
          <dd>{(symbol as string) ?? "—"}</dd>
          <dt className="text-gray-500 dark:text-gray-400">Version</dt>
          <dd>{formatVersion(version)}</dd>
        </dl>
      </Card>

      <Card title="Manager address">
        <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-gray-500 dark:text-gray-400">Current</dt>
          <dd>{managerAddr ? <AddressDisplay address={managerAddr as Address} /> : "—"}</dd>
        </dl>
        <RoleGate hasRole={isAdmin} roleName="DEFAULT_ADMIN_ROLE" action="set manager address">
          <div className="flex items-end gap-2">
            <input className={`flex-1 ${inputClass}`} placeholder="0x…" value={newManager} onChange={(e) => setNewManager(e.target.value)} />
            <TxButton
              label="Save" icon={<Save className="h-4 w-4" />} disabled={!newManager.trim()}
              tooltip="Sets the manager address on the booking token — sends a transaction to your wallet."
              write={() => writeContractAsync({ address: bookingToken, abi, functionName: "setManagerAddress", args: [newManager.trim() as Address] })}
              onConfirmed={() => { setNewManager(""); refetchManager(); }}
            />
          </div>
        </RoleGate>
      </Card>

      <Card title="Min expiration timestamp diff">
        <dl className="mb-3 grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm">
          <dt className="text-gray-500 dark:text-gray-400">Current</dt>
          <dd>{minDiff !== undefined ? `${minDiff} seconds` : "—"}</dd>
        </dl>
        <RoleGate hasRole={canSetMin} roleName="MIN_EXPIRATION_ADMIN_ROLE" action="set min expiration diff">
          <div className="flex items-end gap-2">
            <input className={`w-40 ${inputClass}`} type="number" min="0" placeholder="seconds" value={newMin} onChange={(e) => setNewMin(e.target.value)} />
            <TxButton
              label="Save" icon={<Save className="h-4 w-4" />} disabled={!newMin.trim()}
              tooltip="Sets the minimum reservation expiration difference — sends a transaction to your wallet."
              write={() => writeContractAsync({ address: bookingToken, abi, functionName: "setMinExpirationTimestampDiff", args: [BigInt(newMin.trim())] })}
              onConfirmed={() => { setNewMin(""); refetchMinDiff(); }}
            />
          </div>
        </RoleGate>
      </Card>

      <Card title="Booking Token Roles">
        <RolesPanel address={bookingToken} abi={abi} roles={BOOKINGTOKEN_ROLES} enumerable={false} />
      </Card>
    </div>
  );
}
```

- [ ] **Step 4: Run test + typecheck**

Run: `./node_modules/.bin/vitest run src/pages/tabs/BookingTokenTab.test.tsx && ./node_modules/.bin/tsc -b`
Expected: PASS, no type errors.

- [ ] **Step 5: Commit**

```bash
yarn format
git add src/pages/tabs/BookingTokenTab.tsx src/pages/tabs/BookingTokenTab.test.tsx
git commit -m "feat(ui): booking token tab (info, settings, roles)"
```

---

### Task 8: Full suite, build, and final verification

**Files:** none (verification only).

- [ ] **Step 1: Run the full UI test suite**

Run: `./node_modules/.bin/vitest run`
Expected: PASS (all existing + new tests).

- [ ] **Step 2: Typecheck + production build**

Run: `yarn build`
Expected: `tsc -b` clean and Vite build succeeds. (Note: `prebuild` runs `yarn sync`, which reads `../abi/` and `../ignition/deployments/` — ensure those exist; if `yarn sync` fails locally due to env, run `./node_modules/.bin/tsc -b` and `./node_modules/.bin/vite build` is optional.)

- [ ] **Step 3: Lint**

Run: `yarn lint` (from repo root) or the UI's eslint if configured; fix any reported issues.

- [ ] **Step 4: Manual smoke (optional, if a dev server is available)**

Run: `yarn dev`, open the app, click **Manager** in the nav, and verify all four tabs render and reads populate on a supported network (e.g. Base Sepolia 84532).

- [ ] **Step 5: Final commit (if lint/format produced changes)**

```bash
yarn format
git add -A
git commit -m "chore(ui): formatting and lint fixes for manager page"
```

---

## Self-Review Notes

- **Spec coverage:** Config (pause/unpause, setAccountImplementation, setBookingTokenAddress) → Task 4. Service registry (register/unregister + list) → Task 5. Manager roles (enumerable grant/revoke) → Task 6. BookingToken (name/symbol/version/manager/minExpiration + setManagerAddress + setMinExpirationTimestampDiff + non-enumerable roles) → Task 7. Nav link + read-only-for-all + route → Task 3. RolesPanel refactor + roles.ts → Tasks 1–2. Excluded items (reinitializeV2, upgrades, onlyCMAccount fns) are intentionally absent.
- **Non-enumerable BookingToken roles** handled explicitly in Task 2 `NonEnumerableRoleRow` (no `getRoleMembers` call) and used in Task 7.
- **Type consistency:** `RolesPanel({ address, abi, roles, enumerable })` defined in Task 2 and called with those exact props in Tasks 2, 6, 7. `MANAGER_ROLES`/`BOOKINGTOKEN_ROLES` defined in Task 1 and consumed in Tasks 6/7. `roleHash`/`RoleName` unchanged signatures.
- **serviceName helper risk** flagged in Task 5 Step 4 with a concrete fallback (use `s.parsed`).
