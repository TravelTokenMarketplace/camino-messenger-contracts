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

## Design system ("transit-board terminal")

Identity is a departures board crossed with a financial terminal — dense, legible, operator-trustworthy. All tokens live in `tailwind.config.js`; **use the semantic tokens, never raw `gray-*`/`indigo-*`**.

- **Color:** `tarmac` (blue-black ink / dark surfaces, full 50–950 scale), `paper` + `paper-raised` (cool light bg, deliberately not cream), `camino` (brand teal — means **active / confirmed / "go"**), `departure` (amber — means **pending / in-transit**), `signal` (danger; only `DEFAULT`/`fg`/`dark`, no numeric scale — keep error reds as plain `red-*`). The teal↔amber pairing is semantic, not decorative.
- **Type:** `font-display` (Space Grotesk — titles, brand, big numerals; used sparingly), `font-sans` (IBM Plex Sans — body/UI, the default), `font-mono`/`font-num` (IBM Plex Mono — all on-chain data: addresses, hashes, amounts). Fonts are self-hosted via `@fontsource/*`, imported in `main.tsx`.
- **Component utilities** (in `index.css`): `.board` (hairline-ruled panel surface — `Card`/`TxPanel` use it), `.eyebrow` (small uppercase mono section label — Card titles render as this), `.board-grid` (faint grid texture for the Dashboard manifest hero).
- **Signature:** `TxPanel` renders a split-flap status chip — pending tx flips on amber (`animate-flap`) and resolves to teal `CONFIRMED`. Keep this the one loud element; everything else stays quiet. Respect `prefers-reduced-motion` (handled globally in `index.css`).
- The activity category palette in `lib/activity/style.ts` is an intentionally varied categorical system (each event type a distinct hue) — don't flatten it to brand colors.

**Building new UI (stay on-style):**

- **Reuse primitives, don't hand-roll.** Panels → `<Card>` (board surface + `.eyebrow` title). Tx actions → `TxButton` (never raw `writeContractAsync` + a custom button). Utility/secondary buttons follow `RefreshButton`/`ThemeToggle`: `border-tarmac-300 … hover:bg-tarmac-50`, squared. On-chain data → `AddressDisplay`/`TokenDisplay`; toggles → `Switch`; rich hover → `Tooltip`; icons → `lucide-react` (h-3.5–h-4, `opacity-70` when secondary).
- **Badges/chips** copy `RoleBadge`/`NetworkBadge`: `rounded-[3px]`, hairline border, `font-mono text-[0.625rem] uppercase tracking-[0.08em]`. Show state with a dot: camino = live/ok, `departure` + `animate-lamp` = pending, `signal`/`red` = error.
- **Labels & muted text:** section labels = `.eyebrow`; field labels = `text-tarmac-500 dark:text-tarmac-400`; faint/zero/disabled = `text-tarmac-300 dark:text-tarmac-600`. Numbers use `font-num tabular-nums`.
- **Lists/tables** follow the Dashboard CM Accounts board: one **fixed-width** grid template string shared by the header row and every data row — avoid `auto` columns, since variable content (e.g. a chip vs a dash) misaligns the columns. Use `.eyebrow` column headers and mono numerals.
- **Surfaces are flat:** hairline borders (`border-tarmac-200 dark:border-tarmac-800`), small radius, `shadow-board` only, `divide-tarmac-200/60` dividers. No heavy shadows or large radii.
- **Focus & motion:** a global teal focus-visible ring lives in `index.css` — keep elements focusable, don't override it. Keep motion minimal (the split-flap is the only loud animation; reduced-motion is handled globally).

## Gotchas

- **Don't use `eth_getLogs` over large ranges** (free-tier RPCs reject it). Enumerate accounts via `CMACCOUNT_ROLE` members and lists via role members / array getters, not events.
- **viem overload ambiguity:** `getServiceRestrictedRate` / `getServiceCapabilities` are overloaded by `(string)` and `(bytes32)`. Use the explicit single-overload `bytes32` ABI fragments in `ServicesTab.tsx`.
- **`getSupportedServices` tuple** doesn't decode reliably — list `getAllServiceHashes()` and resolve names via the manager + per-hash getters instead.
- **Tests:** components using `useQueryClient`/`useTx`/wagmi need a `QueryClientProvider` (and usually mocked wagmi) in the test render — see existing `*.test.tsx`.

## Deploy

`../.github/workflows/deploy-ui.yml` builds and publishes to GitHub Pages on pushes to `dev` touching `ui/`, `abi/`, or `ignition/deployments/` (also `workflow_dispatch`). Vite/router base path is `/camino-messenger-contracts/`.
