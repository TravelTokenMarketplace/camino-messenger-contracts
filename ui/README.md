# Camino Messenger Contracts — Management UI

A wallet-connected, client-side web app for creating and managing Camino
Messenger CM Accounts. It is a graphical equivalent of the Hardhat CLI tasks in
`../tasks/`, and is deployed as a static site to GitHub Pages.

## Quick start

```sh
cd ui
yarn install
yarn dev
```

`yarn dev` (and `yarn build`) first run `yarn sync`, which regenerates contract
data from the repo's `../abi` and `../ignition/deployments` into
`src/contracts/generated/` (git-ignored). Re-run `yarn sync` after contracts are
redeployed.

## Scripts

| Script        | Description                                              |
| ------------- | ------------------------------------------------------- |
| `yarn dev`    | Sync contracts, then start the Vite dev server.          |
| `yarn build`  | Sync contracts, typecheck, and build to `dist/`.         |
| `yarn preview`| Preview the production build locally.                    |
| `yarn test`   | Run the Vitest suite.                                    |
| `yarn sync`   | Regenerate `src/contracts/generated/` from repo outputs. |

## Networks

| Network        | Chain ID | Status in UI                                |
| -------------- | -------- | ------------------------------------------- |
| Camino         | 500      | Enabled                                     |
| Base           | 8453     | Enabled once its address file exists         |
| Base Sepolia   | 84532    | Enabled                                     |
| Columbus       | 501      | Defined but disabled (being phased out)      |

A network is only selectable if it is enabled **and** has deployed contracts in
`../ignition/deployments/chain-<id>/deployed_addresses.json`.

## How it works

- **Reads** use an app-owned viem HTTP transport (the RPC URLs in
  `src/config/chains.ts`), so the dashboard is browsable without a wallet.
- **Writes** go through the connected wallet (MetaMask / injected) on whatever
  chain the wallet is connected to. The UI follows the wallet's chain and prompts
  to switch when it is on an unsupported network.
- **Role-aware:** write actions are gated by the on-chain role they require
  (e.g. `SERVICE_ADMIN_ROLE` for tokens/services/pubkeys, `BOT_ADMIN_ROLE` for
  bots, `WITHDRAWER_ROLE` for withdrawals, `DEFAULT_ADMIN_ROLE` for role
  management).

> **Note on public RPCs:** the default RPC endpoints are public and may
> rate-limit. The "Created accounts" list relies on `eth_getLogs` over the
> manager's `CMAccountCreated` events; if a public endpoint caps log ranges,
> point `src/config/chains.ts` at a private RPC (QuickNode, Ankr, Alchemy, …).

## Deployment (GitHub Pages)

`.github/workflows/deploy-ui.yml` builds `ui/` and deploys `ui/dist` to GitHub
Pages on every push to `dev` that touches `ui/**`, `abi/**`, or
`ignition/deployments/**` (and on manual `workflow_dispatch`).

One-time setup: in the repository settings, set **Pages → Build and deployment →
Source** to **GitHub Actions**.

The Vite `base` is `/camino-messenger-contracts/` (project Pages URL). If a
custom domain is configured, change `base` to `/` in `vite.config.ts`.
