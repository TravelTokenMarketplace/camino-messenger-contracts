# Camino Messenger Contracts — Management UI Design

**Date:** 2026-06-23
**Status:** Approved (design)
**Author:** Ekrem Seren

## Overview

A browser-based, wallet-connected UI that lets developers and partners in the
Camino Messenger ecosystem create and manage their on-chain contract state — a
graphical equivalent of the existing Hardhat CLI tasks (`tasks/account.js`,
`tasks/manager.js`).

The app lives in a new `ui/` subfolder of this repo and is published as a static
site to GitHub Pages via a GitHub Actions workflow.

### Constraints

- **Static hosting only.** GitHub Pages has no backend, so the app is a fully
  client-side SPA. All chain interaction happens directly from the browser via
  RPC and the user's wallet.
- The app must stay in sync with deployed contracts without manual copying of
  addresses/ABIs.

## Goals (v1)

Core partner workflows, gated by on-chain roles:

- Connect a wallet, browse a read-only dashboard without one.
- Create a CM Account.
- Manage a CM Account: bots, payment tokens, supported services, wanted
  services, roles, encryption pubkeys, withdrawals, gas-money limits.

## Non-Goals (deferred to v2)

- Manager admin operations: pause/unpause, set-implementation, `btoken:set`,
  service registry register/unregister.
- CMAccount upgrade flow.
- BookingToken mint/cancel operations.
- Off-chain cheque (EIP-712) signing.

## Tech Stack

- **React + Vite + TypeScript** SPA.
- **wagmi + viem** for chain interaction; **RainbowKit** for wallet connection.
- **Tailwind CSS** for styling. Visual direction to be set during implementation
  using the `frontend-design` skill (avoid generic/templated AI aesthetics).

## Networks

| Network            | chainId | Status in UI                                  |
| ------------------ | ------- | --------------------------------------------- |
| Camino mainnet     | 500     | Enabled                                        |
| Base Sepolia       | 84532   | Enabled                                        |
| Base mainnet       | 8453    | Enabled when its address file exists           |
| Columbus testnet   | 501     | Defined but `enabled: false` (being phased out)|

Columbus stays in code as a disabled network so it can be re-enabled trivially,
but is not selectable/operational in the UI.

## Architecture

### Repo layout

```
ui/
  package.json              # vite, react, wagmi, viem, rainbowkit, tailwind
  vite.config.ts            # base path for GitHub Pages
  scripts/sync-contracts.ts # generates src/contracts/ from repo abi/ + ignition/
  src/
    contracts/              # GENERATED: resolved addresses + minimal ABIs per chain (git-ignored)
    config/chains.ts        # the networks + RPC endpoints
    wallet/                 # wagmi + rainbowkit setup
    pages/                  # Dashboard, Account workspace, etc.
    features/               # one module per workflow (bots, roles, services, tokens…)
    components/             # shared UI (tx button, address input, role badge…)
```

The `ui/` package is isolated from the Hardhat root `package.json` (its own
dependency tree).

### Contract sync script

`ui/scripts/sync-contracts.ts`:

- Reads `ignition/deployments/chain-*/deployed_addresses.json` and
  `abi/contracts/**`.
- Resolves the real proxy/contract addresses per chain, handling the
  `Module#Name` key format and ignoring historical/test modules. Of note, the
  Columbus (`chain-501`) file contains many legacy modules; only the canonical
  `CaminoMessengerModule#*Proxy` entries are used.
- Extracts only the ABIs the UI needs (CMAccountManager, CMAccount,
  BookingToken, plus minimal IERC20/IERC721) and writes typed output to
  `ui/src/contracts/`.
- Runs as a `predev`/`prebuild` step so generated config is always fresh. The
  generated directory is git-ignored; the Hardhat artifacts remain the source of
  truth.
- Base mainnet (and any future network) appears automatically once its address
  file exists.

### Wallet, networks & contract interaction

Two transports, by design:

1. **Wallet provider (writes + connected account).** Transactions are signed and
   broadcast through the user's wallet via its own configured RPC. Handled by the
   wagmi injected/RainbowKit connector.
2. **App-owned public RPC (reads).** All read calls use the app's configured
   viem `http()` transport, not the wallet's. Rationale: reads must work without
   a connected wallet (browsable dashboard); the app controls method support
   (`eth_getLogs` for role/event enumeration), reliability, and performance
   (direct HTTP is far faster than routing reads through the extension).

Other rules:

- **Single source of truth for chain = the wallet's connected chain.** When the
  wallet is on an unsupported chain (or Columbus), show a "switch network"
  prompt. Reads without a wallet use a default selected enabled network.
- **Default public RPCs** baked into `config/chains.ts`, with an optional custom
  RPC override (public endpoints rate-limit). The override affects only the read
  transport; writes always go through the wallet.
- **Shared transaction component** for every write: simulate → send → pending →
  confirmed/failed, with block-explorer links.
- **Role-aware UI.** Before showing a write action, check whether the connected
  address holds the required role (e.g. `BOT_ADMIN_ROLE` for bot management) and
  disable/explain unavailable actions instead of submitting reverting txs.

## v1 Feature Breakdown

### Dashboard (read-only, no wallet required)

- Network status card: resolved contract addresses, manager paused state,
  current CMAccount implementation version, BookingToken status.
- "My accounts": scan the manager for CM Accounts where the connected address
  holds any role (mirrors `account:find`).

### Create CM Account

- Form → create tx (via manager) → on success, deep-link into the new account's
  workspace.

### CM Account workspace (tabs, each gated by the relevant role)

- **Overview** — address, roles summary, native + ERC20 balances.
- **Bots** — list / add / remove messenger bots; set gas-money limit & period;
  withdraw gas for a bot.
- **Payment tokens** — list / add / remove supported payment tokens; toggle
  off-chain payment support.
- **Services (supported)** — list / add / remove; set restricted rate; set
  capabilities. Service names resolved against the manager's service registry.
- **Wanted services** — list / add / remove.
- **Roles** — view members per role; grant / revoke (gated by
  `DEFAULT_ADMIN_ROLE`).
- **Encryption pubkeys** — list / add / remove off-chain encryption public keys.
- **Withdrawals** — withdraw native, ERC20, ERC721 (gated by `WITHDRAWER_ROLE`).

Each write reuses the shared transaction component and role-gating.

## Error Handling

- Read failures (RPC down / rate-limited): surface a non-blocking error with a
  retry and a hint to set a custom RPC.
- Write failures: the transaction component reports revert reasons where
  available and links to the explorer.
- Wrong/unsupported chain: blocking "switch network" prompt before any write.
- Missing role: action disabled with an explanation of the required role.

## Testing

- Unit-test the sync script's address/ABI resolution (including the messy
  Columbus file and the missing-Base-mainnet case).
- Component/integration tests for role-gating logic and the transaction
  component states.
- Manual verification against Base Sepolia testnet before first deploy.

## Build & Deployment

### Contract sync

Runs automatically as `predev`/`prebuild`; output is git-ignored.

### GitHub Pages workflow (`.github/workflows/deploy-ui.yml`)

- **Trigger:** push to `dev` affecting `ui/**`, `abi/**`, or
  `ignition/deployments/**`; plus manual `workflow_dispatch`.
- **Steps:** checkout → setup Node → install deps (root + `ui/`) → run sync →
  `vite build` (with `base` set to the repo name) → upload artifact → deploy via
  `actions/deploy-pages`.
- **SPA fallback:** copy `index.html` → `404.html` so client-side routes resolve.

### Defaults

- **Deploy branch:** `dev`.
- **Vite `base`:** `/camino-messenger-contracts/` (project Pages URL). Becomes
  `/` if a custom domain is later configured.
