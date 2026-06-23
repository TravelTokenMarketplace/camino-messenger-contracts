# Camino Messenger Contracts

Solidity smart contracts for the Camino Messenger ecosystem, built with Hardhat. A wallet-connected management UI lives in `ui/` (see `ui/CLAUDE.md`).

## Layout

- `contracts/` — Solidity sources (Solidity 0.8.24, optimizer `runs: 1`, `evmVersion: paris`)
    - `manager/` — `CMAccountManager` (factory + registry + roles), `ICMAccountManager`
    - `account/` — `CMAccount` (per-partner account: bots, tokens, services, pubkeys, withdrawals), `GasMoneyManager`, `ICMAccount`
    - `booking-token/` — `BookingToken` (ERC-721) + cancellable/operator extensions
    - `partner/` — shared base contracts: `PartnerConfiguration`, `ServiceRegistry`
    - `test/` — mocks/helpers used by tests
- `tasks/` — Hardhat tasks for management: `manager.js`, `account.js` (registered in `hardhat.config.js`). The UI mirrors these.
- `ignition/` — Hardhat Ignition deployment: `modules/messenger.js`, per-network `*_parameters.json`, and recorded deployments in `deployments/chain-<id>/`
- `abi/` — exported ABIs (see ABIs note below)
- `scripts/`, `services/`, `utils/`, `examples/`, `go/` — supporting code
- `docs/` — generated contract docs + design specs/plans under `docs/superpowers/`
- `ui/` — the React management UI (separate toolchain; has its own `CLAUDE.md`)

## Architecture notes

- Contracts are **UUPS upgradeable** (OpenZeppelin upgradeable + `@openzeppelin/hardhat-upgrades`). `CMAccountManager.createCMAccount` deploys an `ERC1967Proxy` per account and grants it `CMACCOUNT_ROLE`.
- **Access control** is role-based (`AccessControlEnumerable`). `DEFAULT_ADMIN_ROLE` is `0x00…`; other roles are `keccak256(name)`. A "messenger bot" is an address granted three roles at once: `MESSENGER_BOT_ROLE`, `BOOKING_OPERATOR_ROLE`, `GAS_WITHDRAWER_ROLE`.
- Services are referenced by name (`cmp.services.<package>.<version>.<Name>`) and must be registered in the manager's `ServiceRegistry` before an account can support/want them.

## Commands

- `yarn compile` — compile contracts (also runs `contract-sizer`)
- `yarn test` — run the Hardhat test suite (`REPORT_GAS=true`)
- `yarn lint` — Prettier + ESLint + Solhint; `yarn format` to auto-fix with Prettier
- `yarn docgen` — regenerate `docs/`
- `yarn hardhat export-abi` — regenerate `abi/` (see below)
- Deploy with Hardhat Ignition, e.g. `yarn hardhat ignition deploy ignition/modules/messenger.js --network <name> --parameters ignition/<network>_parameters.json`

## Networks

Defined in `hardhat.config.js`: `camino` (500), `columbus` (501, deprecated), `base` (8453), `base_sepolia` (84532), plus `localhost`. RPC URLs and deployer keys come from Hardhat **configuration variables** (`vars.get(...)`), e.g. `CAMINO_DEPLOYER_PRIVATE_KEY`, `BASE_SEPOLIA_URL`, `BASESCAN_API_KEY` — set them with `yarn hardhat vars set <NAME>`.

## ABIs (important)

There are two ABI locations and they serve different purposes:

- `artifacts/` — Hardhat's compiler output. **Hardhat and tests use this.**
- `abi/` — a flattened export produced by `hardhat-abi-exporter` via `yarn hardhat export-abi`. **The UI consumes this** (`ui/scripts/sync-contracts.ts` reads it).

After changing contracts, run `yarn compile` and then `yarn hardhat export-abi` so `abi/` (and therefore the UI) stays in sync. `abi/` is committed; don't hand-edit it.
