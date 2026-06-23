# Camino Messenger Contract Management UI

Wallet-connected web app for creating and managing the contracts in this repo (`../contracts`). A read-write GUI mirroring the Hardhat tasks in `../tasks`, deployable to GitHub Pages. This is a **separate toolchain** from the root Hardhat project — run all commands below from `ui/`.

## Stack

React 18 + Vite + TypeScript · Tailwind (class-based dark mode) · wagmi v2 + viem v2 · TanStack Query · `@headlessui/react` + `fuse.js` (combobox/menus) · lucide-react (icons) · react-jazzicon (identicons) · react-router-dom · Vitest + Testing Library.

## Commands (run from `ui/`)

- `yarn dev` — Vite dev server (runs `yarn sync` first via `predev`)
- `yarn build` — `tsc -b && vite build` (runs `yarn sync` first via `prebuild`)
- `yarn sync` — regenerate `src/contracts/generated/` from on-chain deployment data + ABIs
- `yarn test` — run Vitest once. **Run from `ui/`** (root `yarn test` runs Hardhat). If `yarn test` resolves to the root, call `./node_modules/.bin/vitest run <path>` directly.

## Generated contract data (build-time sync)

`scripts/sync-contracts.ts` resolves deployed addresses from `../ignition/deployments/chain-<id>/` and ABIs from `../abi/` into `src/contracts/generated/{addresses.ts,abis.ts}`. That directory is **git-ignored and must never be hand-edited** — re-run `yarn sync` after a new deployment or ABI export. Consume it only through `src/contracts/index.ts` (`getContractsForChain`, `hasContracts`, `MANAGER_ABI`, `CMACCOUNT_ABI`, `BOOKINGTOKEN_ABI`).

If service reads break with decode errors, the usual cause is a stale `../abi/` — regenerate it from the repo root with `yarn hardhat export-abi`.

## Structure

- `config/chains.ts` — `APP_CHAINS` (Columbus 501 defined but `enabled: false`) and `ENABLED_CHAINS` (enabled **and** has deployed contracts).
- `wallet/` — `wagmi.ts` (config), `Providers.tsx` (Wagmi + Query + ActiveChain + Tx providers), `activeChain.tsx` (`useActiveChain`: active chain follows the wallet when connected, free selection when not).
- `tx/TxProvider.tsx` — global transaction tracking (`useTx().track`).
- `hooks/` — `useActiveContracts`, `useHasRole`, `useRoleMembers`, `useContractList`, `useMyAccounts`.
- `components/` — shared UI: `TxButton`, `TxPanel`, `RoleGate`, `PermissionHint`, `RowAction`, `Autocomplete`, `Checkbox`, `Input`, `Tooltip`, `AddressDisplay`, `CopyButton`, `RefreshButton`, `Card`, `NetworkSelector`, `ConnectButton`, etc.
- `pages/` — `Dashboard`, `CreateAccount`, `AccountWorkspace` + `tabs/` (Bots, PaymentTokens, Services, Roles, Pubkeys, Withdrawals).
- `lib/` — `format`, `roles` (role names + `roleHash`), `serviceName` (parse `cmp.services.<pkg>.<version>.<Name>`, group by package), `receipt`.

## Conventions

- **Reads vs writes:** reads use the app-owned RPC keyed by `useActiveChain().activeChainId`; writes go through the wallet. Always pass `chainId: activeChainId` to read hooks.
- **All user-facing writes go through `TxButton`** (which calls `useTx().track`) — don't call `writeContractAsync` directly for actions without tracking. `track` waits for the receipt, then `onConfirmed` fires and **all queries are invalidated** so the whole view refreshes. Use `onConfirmed` for side effects like clearing inputs; do **not** add manual refetch timers.
- **Permission-gated actions** are wrapped in `RoleGate` with a human `action` label (renders a `PermissionHint` like "Can't add bot" when the wallet lacks the role).
- **Inputs** use the shared `Input` component / `inputClass`; selects and menus use HeadlessUI (`NetworkSelector`, `ConnectButton`, `Checkbox`, `Autocomplete`).
- **Dark mode** is Tailwind `class` strategy; every color needs a `dark:` variant.

## Gotchas

- **Don't use `eth_getLogs` over large ranges** (free-tier RPCs reject it). Enumerate accounts via `CMACCOUNT_ROLE` members and lists via role members / array getters, not events.
- **viem overload ambiguity:** `getServiceRestrictedRate` / `getServiceCapabilities` are overloaded by `(string)` and `(bytes32)`. Use the explicit single-overload `bytes32` ABI fragments in `ServicesTab.tsx`.
- **`getSupportedServices` tuple** doesn't decode reliably — list `getAllServiceHashes()` and resolve names via the manager + per-hash getters instead.
- **Tests:** components using `useQueryClient`/`useTx`/wagmi need a `QueryClientProvider` (and usually mocked wagmi) in the test render — see existing `*.test.tsx`.

## Deploy

`../.github/workflows/deploy-ui.yml` builds and publishes to GitHub Pages on pushes to `dev` touching `ui/`, `abi/`, or `ignition/deployments/` (also `workflow_dispatch`). Vite/router base path is `/camino-messenger-contracts/`.
