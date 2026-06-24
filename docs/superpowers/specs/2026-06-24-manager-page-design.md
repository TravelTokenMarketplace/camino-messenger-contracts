# Manager Page — Ecosystem Management UI

**Date:** 2026-06-24
**Status:** Approved (design)

## Goal

Add a wallet-connected **Manager** page to the UI (`ui/`) that exposes the
ecosystem-management functions of the `CMAccountManager` and `BookingToken`
contracts. Today the UI only surfaces a read-only network status on the
Dashboard and per-account management in `AccountWorkspace`. This adds the
missing manager/booking-token administration surface.

## Scope

### Included (CMAccountManager)
- `pause()` / `unpause()` — `PAUSER_ROLE`
- `setAccountImplementation(address)` — `VERSIONER_ROLE`
- `setBookingTokenAddress(address)` — `VERSIONER_ROLE`
- `registerService(string)` / `unregisterService(string)` — `SERVICE_REGISTRY_ADMIN_ROLE`
- Role grant/revoke (`grantRole`/`revokeRole`) — `DEFAULT_ADMIN_ROLE`
- Reads: `paused`, `getAccountImplementation`, `getBookingTokenAddress`,
  `getAllRegisteredServiceNames`, role members.

### Included (BookingToken)
- `setManagerAddress(address)` — `DEFAULT_ADMIN_ROLE`
- `setMinExpirationTimestampDiff(uint256)` — `MIN_EXPIRATION_ADMIN_ROLE`
- Role grant/revoke — `DEFAULT_ADMIN_ROLE`
- Reads: `name`, `symbol`, `version`, `getManagerAddress`,
  `getMinExpirationTimestampDiff`, `hasRole`.

### Excluded
- `BookingToken.reinitializeV2(name, symbol)` — risky/one-time; deliberately
  omitted.
- Proxy upgrade execution (`upgradeToAndCall`) — performed via Hardhat/Ignition,
  not a UI form. `UPGRADER_ROLE` is still shown in role lists.
- All `onlyCMAccount` booking-token functions (`safeMintWithReservation`,
  `buyReservedToken`, cancellation flow, `recordExpiration`) — these are
  account-level actions, not ecosystem management.

## Key contract constraint

- **`CMAccountManager` is `AccessControlEnumerableUpgradeable`** — exposes
  `getRoleMembers(role)`. Manager roles can list members (reuse existing
  `useRoleMembers`).
- **`BookingToken` is plain `AccessControlUpgradeable`** — NO
  `getRoleMembers`/`getRoleMemberCount`. BookingToken roles cannot enumerate
  members. The UI offers grant/revoke-by-address plus a "you hold this role"
  indicator only, and states in the UI that members can't be listed.

## Visibility

Per decision: the **Manager** nav link is **always visible** (no role gating on
the link). The page is read-only for everyone; each write control is wrapped in
`RoleGate` so users without the relevant role see a `PermissionHint` instead of
an active button. This matches the account-tab UX.

## Architecture

### Routing & navigation
- `App.tsx`: add `<Route path="manager" element={<ManagerWorkspace />} />`.
- `Layout.tsx`: add a `Manager` `NavLink` between *Dashboard* and *Create
  Account*.

### `ManagerWorkspace` page
Mirrors `AccountWorkspace` layout: a two-column grid with a left column
(summary card + `TxPanel`) and a right column (tab nav + active tab). Tabs are
selected via `?tab=` search param. Tabs:

1. **Config** (`ManagerConfigTab`) — `CMAccountManager` config:
   - Manager proxy address; `paused` status + Pause/Unpause (gated `PAUSER_ROLE`).
   - Account implementation: read `getAccountImplementation` + set form (gated
     `VERSIONER_ROLE`).
   - Booking token address: read `getBookingTokenAddress` + set form (gated
     `VERSIONER_ROLE`).
2. **Service Registry** (`ServiceRegistryTab`) — list
   `getAllRegisteredServiceNames` grouped by package (reuse
   `groupServicesByPackage` + `ServiceLabel`). Register (input) / unregister
   (row action), gated `SERVICE_REGISTRY_ADMIN_ROLE`.
3. **Roles** (`ManagerRolesTab`) — `RolesPanel` (enumerable) over the manager
   contract for `DEFAULT_ADMIN_ROLE`, `PAUSER_ROLE`, `UPGRADER_ROLE`,
   `VERSIONER_ROLE`, `SERVICE_REGISTRY_ADMIN_ROLE`.
4. **Booking Token** (`BookingTokenTab`):
   - Info card: `name`, `symbol`, `version`, manager address.
   - Manager address: `setManagerAddress` (gated `DEFAULT_ADMIN_ROLE`).
   - Min expiration diff: read `getMinExpirationTimestampDiff` + set (gated
     `MIN_EXPIRATION_ADMIN_ROLE`).
   - Roles: `RolesPanel` (non-enumerable) over the booking token for
     `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `MIN_EXPIRATION_ADMIN_ROLE`.

### Shared `RolesPanel` refactor
Extract a generic `RolesPanel` component from the current account `RolesTab`:
- Props: `address: Address`, `abi: Abi`, `roles: readonly RoleName[]`,
  `enumerable: boolean`.
- `enumerable={true}`: member list via `useRoleMembers` + grant/revoke (used by
  account `RolesTab` and `ManagerRolesTab`).
- `enumerable={false}`: grant/revoke forms + "you hold this role" indicator via
  `useHasRole`, no member list (used by BookingToken roles).
- Grant/revoke gated on `DEFAULT_ADMIN_ROLE` of the target contract.
- `RolesTab` is rewired to render `<RolesPanel enumerable />` — no behavior
  change to the account page.

### `lib/roles.ts` changes
- Extend `MANAGER_ROLES` to include `DEFAULT_ADMIN_ROLE` and `UPGRADER_ROLE`
  (currently `PAUSER_ROLE`, `VERSIONER_ROLE`, `SERVICE_REGISTRY_ADMIN_ROLE`).
- Add `BOOKINGTOKEN_ROLES = [DEFAULT_ADMIN_ROLE, UPGRADER_ROLE,
  MIN_EXPIRATION_ADMIN_ROLE]`.
- Ensure `MIN_EXPIRATION_ADMIN_ROLE` is in the hash map and `RoleName` union.

### Reused infrastructure
- `useActiveContracts` already exposes `manager`, `bookingToken`, `managerAbi`,
  `bookingTokenAbi`.
- `useHasRole`, `useRoleMembers`, `RoleGate`, `PermissionHint`, `TxButton`,
  `RowAction`, `AddressDisplay`, `Card`, `Autocomplete`, `RefreshButton` are
  reused as-is.
- All writes go through `TxButton` → `useTx().track`, per UI conventions.

## Error handling / edge cases
- Read hooks always pass `chainId: activeChainId`; unsupported network renders a
  "connect to a supported network" notice (as Dashboard does).
- Set-address forms validate non-empty input before enabling the button
  (contract reverts on zero-code addresses; surfaced via the tx panel/wallet).
- BookingToken roles tab clearly states members cannot be enumerated.

## Testing
- Vitest unit tests following existing conventions (QueryClientProvider +
  mocked wagmi):
  - `RolesPanel` in both enumerable and non-enumerable modes.
  - `ManagerWorkspace` tab routing render.
  - Smoke render of each new tab.
- `roles.test.ts`: assert new role hashes and lists.
- Run `yarn test` from `ui/` (or `./node_modules/.bin/vitest run`).

## Out of scope / follow-ups
- No contract changes; ABIs already contain all needed functions (verified
  against `ui/src/contracts/generated/abis.ts`). No `yarn hardhat export-abi`
  needed.
