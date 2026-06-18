# Camino Messenger Contracts

[![CAMINO NETWORK](https://img.shields.io/badge/CAMINO-NETWORK-b440fc?style=for-the-badge&logoColor=white&labelColor=0085ff)](https://camino.network/)
[![CHAT WITH US](https://img.shields.io/badge/DISCORD-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.com/channels/949247897688494150/1182680860797960253)

[![CI](https://github.com/TravelTokenMarketplace/camino-messenger-contracts/actions/workflows/ci.yaml/badge.svg)](https://github.com/TravelTokenMarketplace/camino-messenger-contracts/actions/workflows/ci.yaml)

This repository contains the smart contracts for the [Camino
Messenger](https://camino.network/camino-messenger-sets-the-global-standard-in-travel-data-management-and-distribution/).

## Camino (mainnet) Deployed Contracts

Below is a table of deployed contracts and their addresses on Camino mainnet.

| Contract                       | Address                                                                                                                              |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| CMAccountManager               | [0xf9FE1eaAB73a2902136FE7A83E0703338D3b9F1e](https://caminoscan.com/address/0xf9FE1eaAB73a2902136FE7A83E0703338D3b9F1e?tab=contract) |
| BookingToken                   | [0xe2b8c92B6519d1A2020dA0A5fBbA99a43A2c0922](https://caminoscan.com/address/0xe2b8c92B6519d1A2020dA0A5fBbA99a43A2c0922?tab=contract) |
| BookingTokenOperator (Library) | [0x65C34Ca1FCdF46B60C2b9b8f81475f69086116dD](https://caminoscan.com/address/0x65C34Ca1FCdF46B60C2b9b8f81475f69086116dD?tab=contract) |
| CMAccount (Implementation)     | [0x52D94b6ccDa96BE4a99ED9C8D39682D6B4EE4702](https://caminoscan.com/address/0x52D94b6ccDa96BE4a99ED9C8D39682D6B4EE4702?tab=contract) |

## Base Sepolia (testnet) Deployed Contracts

Below is a table of deployed contracts and their addresses on Base Sepolia
testnet.

| Contract                       | Address                                                                                                                            |
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------- |
| CMAccountManager               | [0xEcf9b5ca23257969B4F9bb3Efca2d5bb850FAcE9](https://sepolia.basescan.org/address/0xEcf9b5ca23257969B4F9bb3Efca2d5bb850FAcE9#code) |
| BookingToken                   | [0x459EEdD4bE13bD7D1Af27DA5DdA6d69407118C83](https://sepolia.basescan.org/address/0x459EEdD4bE13bD7D1Af27DA5DdA6d69407118C83#code) |
| BookingTokenOperator (Library) | [0x579EF9939b884E2E9424736AfCcE6623FC728A66](https://sepolia.basescan.org/address/0x579EF9939b884E2E9424736AfCcE6623FC728A66)      |
| CMAccount (Implementation)     | [0x7AEFbc8FC7d103bDf79e14F7CC4F42d93B916b61](https://sepolia.basescan.org/address/0x7AEFbc8FC7d103bDf79e14F7CC4F42d93B916b61#code) |

## Chain4Travel Messenger Server

Chain4Travel is running the first and currently only messenger server.

| Camino Mainnet                       | Address                                      |
| ------------------------------------ | -------------------------------------------- |
| Messenger URL                        | `https://messenger.chain4travel.com`         |
| Messenger CM Account                 | `0x16DFfB3911BB0b1B53eF4d774804381f0B38B5d7` |
| Messenger Service Bot (`toBot`) Addr | `0xbeb027D2f439805E17EAA16Da26c1FCa68a30232` |

| Base Sepolia Testnet                 | Address                              |
| ------------------------------------ | ------------------------------------ |
| Messenger URL                        | `https://messenger.chain4travel.com` |
| Messenger CM Account                 | _TBD_                                |
| Messenger Service Bot (`toBot`) Addr | _TBD_                                |

## Quickstart

### Clone the repo and change directory into

```sh
git clone git@github.com:TravelTokenMarketplace/camino-messenger-contracts.git
cd camino-messenger-contracts
```

### Install packages

```sh
yarn install
```

### Run tests. This will compile the contracts and run the tests:

```sh
yarn test
```

### Setting Hardhat Vars

For Camino (mainnet), Base Sepolia (testnet), and Base (mainnet) networks, we
are using hardhat's vars tool to store private keys and URLs. To set these you
can use the commands below:

```
yarn hardhat vars set BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY
```

```
yarn hardhat vars set BASE_DEPLOYER_PRIVATE_KEY
```

```
yarn hardhat vars set CAMINO_DEPLOYER_PRIVATE_KEY
```

Optional settings for Base networks:

```
yarn hardhat vars set BASE_SEPOLIA_URL
```

(Defaults to `https://base-sepolia.drpc.org` if not set. **Note**: Public
endpoints are subject to rate limiting. For high-volume usage, users should use
private endpoints from third-party RPC providers such as QuickNode, Ankr,
Infura, or Alchemy.)

```
yarn hardhat vars set BASE_URL
```

(Defaults to `https://base.drpc.org` if not set)

```
yarn hardhat vars set BASESCAN_API_KEY
```

(Used for verifying contracts on Base and Base Sepolia)

These will also be used for `yarn hardhat manager` tasks. These variables are
stored in the `/home/$USER/.config/hardhat-nodejs/vars.json` file, so they are
not accidentally pushed to git.

## Contracts

### CMAccount

The `CMAccount` contract represents a Camino Messenger account. Currently, it
includes functionalities for the management of bots. More features will be
introduced in the future.

This contract works closely with the `CMAccountManager` to handle accounts.

### CMAccountManager

The `CMAccountManager` contract acts as a manager for `CMAccount` contracts. It
handles the creation, registration, verification, and management of accounts. It
also keeps records for the developer wallet, fees, and `CMAccount`
implementation address. Accounts can only be upgraded to the implementation
address that the manager holds.

### PartnerConfiguration

The `PartnerConfiguration` contract is used by the `CMAccount` and implements
features to register supported (supplier) and wanted (distributor) services,
register public keys that would be used to encrypt private data, off-chain
payment support, and on-chain supported payment token addresses.

### ServiceRegistry

The `ServiceRegistry` contract is used by the `CMAccountManager` contract and
implements a registry that is used to hash service names to keccak256 hashes and
store them in a mapping as `hash => service name`. `CMAccount` use these to
resolve hashes to service names and service names to hashes.

### BookingToken

The `BookingToken` contract is an ERC-721 NFT contract that is used by the
partners to mint and buy Booking Tokens. A Booking Token represents a booking
done on the Camino Messenger ecosystem.

Only the `CMAccount` contracts are allowed to mint and buy the tokens.

### Proxies

For `CMAccountManager` and `CMAccount` contracts, an `ERC1967Proxy` (UUPS) is
used.

The **`hardhat-ignition`** module deploys the `CMAccountManager` contract and
then deploys an `ERC1967Proxy` proxy, setting the implementation address to the
`CMAccountManager`'s address. We will call this proxy **managerProxy** or simply
**manager** in this document.

Then a `CMAccount` contract is deployed, and its address is set by calling
`managerProxy.setAccountImplementation(CMAccount.getAddress())`. After that, the
manager is ready to create CM accounts.

Calling `managerProxy.createCMAccount(...)` with the necessary arguments creates
an `ERC1967Proxy` and sets the implementation address to the recorded account
implementation address in the manager. After it is deployed, it is immediately
(same transaction) initialized with the given arguments.

## Deploy Contracts Locally

### Run local hardhat node

```
yarn hardhat node
```

### Deploy contracts using the ignition module

```
yarn hardhat ignition deploy ignition/modules/0_development.js --network localhost
```

### Output should be similar to this

```
yarn run v1.22.19
$ /hgst/work/github.com/TravelTokenMarketplace/camino-messenger-contracts/node_modules/.bin/hardhat ignition deploy ignition/modules/0_development.js --network localhost
Hardhat Ignition 🚀

Deploying [ CMAccountManagerModule ]

Batch #1
  Executed BookingTokenProxyModule#BookingToken
  Executed CMAccountManagerModule#BookingTokenOperator
  Executed ManagerProxyModule#CMAccountManager

Batch #2
  Executed BookingTokenProxyModule#ERC1967Proxy
  Executed CMAccountManagerModule#CMAccount
  Executed ManagerProxyModule#ERC1967Proxy

Batch #3
  Executed CMAccountManagerModule#BookingToken
  Executed CMAccountManagerModule#CMAccountManager

Batch #4
  Executed CMAccountManagerModule#BookingToken.initialize
  Executed CMAccountManagerModule#CMAccountManager.initialize
  Executed CMAccountManagerModule#CMAccountManager.setAccountImplementation
  Executed CMAccountManagerModule#CMAccountManager.setBookingTokenAddress

[ CMAccountManagerModule ] successfully deployed 🚀

Deployed Addresses

BookingTokenProxyModule#BookingToken - 0x5FbDB2315678afecb367f032d93F642f64180aa3
CMAccountManagerModule#BookingTokenOperator - 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
ManagerProxyModule#CMAccountManager - 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
BookingTokenProxyModule#ERC1967Proxy - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
CMAccountManagerModule#CMAccount - 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
ManagerProxyModule#ERC1967Proxy - 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
CMAccountManagerModule#BookingToken - 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
CMAccountManagerModule#CMAccountManager - 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
Done in 3.78s.
```

You can also see your deployed contract addresses in the
`ignition/deployments/<chainid>/deployed_addresses.json` file.

### Visualize the deployment

```
yarn hardhat ignition visualize ignition/modules/0_development.js
```

This will open a browswer tab with the deployment flow visualized.

## Camino Messenger Account Setup

> [!WARNING] This guide is for development purposes on Base Sepolia only. For
> officially registered CM Accounts on Base, please wait for the frontend GUI to
> be finished.

To set up your Camino Messenger Account (CM Account) for use with the Camino
Messenger Bot on Base Sepolia, you need to:

1. Create a CM Account
2. Register your bot on your CM Account
3. Register the services you provide on your CM Account

Follow the steps below to complete this process.

### Prerequisites

Before you begin, ensure you have completed the following steps:

- **Compile the Contracts:** Ensure all contracts are successfully compiled.
  (`yarn compile --force`)
- **KYC Verification:** is not required on Base Sepolia
- **Fund Your Wallet:** Use one of the faucets to obtain at least 0.1ETH. Most
  of them require real ETH on Mainnet or Mainnet activity. See:
  https://docs.base.org/base-chain/network-information/network-faucets. After
  the fee removal feature has been deployed, there is no requirement for 100CAM
  or 100 ERC-20 tokens.

### Creating a CM Account

To create your CM Account, run the following command:

```
yarn hardhat account create --private-key <PrivateKeyValue> --network base_sepolia
```

<details>
<summary>Example output:</summary>

```
yarn run v1.22.22
$ '.../camino-messenger-contracts/node_modules/.bin/hardhat' account create --private-key <private-key-deducted> --network base_sepolia
Running on base_sepolia
We need to approve the ServiceFeeToken for the manager to create the CMAccount.
Getting ServiceFeeToken...
Running on base_sepolia
ServiceFeeToken Address: 0xF8BD889B94142aae5d6ed585B16c77C8c378CAAD
ServiceFeeToken Name   : USD Service Fee Token
ServiceFeeToken Symbol : USD.test
Signer: 0x066c6B010358771af982D5BaEb0c9f46bf1bE762
Getting Prefund Amount...
Prefund Amount: 100.0
Approving the manager for 100.0 USD.test ...
Tx: 0x817ed0b31dfd63a4be0d34bb2adf00cc0123cd225926a4fc9a970e0852ade23f
Creating CMAccount... (Sending 0.0 CAM to the new CMAccount)
Tx: 0x3ea49567a6fbbfddb049b3588350542d6a4e2c23ae0e1e30d019d3f8a453621f
CMAccount Address: 0x5e1c75F35be2f3E0093525E1CcEE03C1A359D7e7
Done in 4.67s.
```

**The above example output was taken before the message fee deposit was
removed** **Check the created contract using BaseScan
(https://sepolia.basescan.org/)**

</details>

> [!TIP] Instead of specifying your private key and CM Account address from the
> CLI, you can export them as variables.
>
> Use these commands to set the variables, and then you can omit the
> `--private-key` and `--cm-account` arguments from the `yarn hardhat account`
> commands below:
>
> `export CMACCOUNT_PK=0x...`
>
> `export CMACCOUNT_ADDRESS=0x...`

#### Command Parameters

- **`--private-key`:** Enter the static private key of your wallet.
- **`--network`:** Specify the network where you wish to create your account (as
  configured in the `hardhat.config.js`). For development purposes, use
  `base_sepolia`.

The command output will provide a new CM Account address that you must save for
use in the following steps.

### Registering Your Bot

After creating your CM Account, you need to register the address of your bot on
the CM Account to authorize it. Execute the following command:

```
yarn hardhat account bot:add --cm-account <CMAccountAddress> --private-key <PrivateKeyValue> --bot <BotAddress> --network base_sepolia
```

<details>
<summary>Example output:</summary>

```
yarn run v1.22.22
$ /hgst/work/github.com/chain4travel/camino-messenger-contracts/node_modules/.bin/hardhat account bot:add --cm-account 0xe7C3AAd26fA667a2dc51A4B3c56f8919574b275A --private-key <private-key-deducted> --network base_sepolia --bot 0xd41786599F2B225A5A1eA35cDc4A2a6Fa9E92BeA
CMAccount: 0xe7C3AAd26fA667a2dc51A4B3c56f8919574b275A
Bot: 0xd41786599F2B225A5A1eA35cDc4A2a6Fa9E92BeA
Gas: 0 (This amount will be transferred from the CMAccount to the bot address)
Adding bot to CMAccount...
Signer: 0xFe77dcE375C3814F15F8035bCAC1A791D3dCdf21
Tx: 0x24421b26f36caadcdb007c1a5d010416b8889586c36471d399a3ab367ff967a4
Done in 2.16s.
```

</details>

#### Command Parameters

- **`--cm-account`:** The EVM contract address of your newly created CM Account.
- **`--private-key`:** The static private key of the wallet used for creating
  the CM Account.
- **`--bot`:** The address of the bot you are registering. Ensure that this
  address is your bot's wallet address and not the same as your CM Account
  wallet.
- **`--network`:** Specify the network (as configured in the
  `hardhat.config.js`).

### Registering Services

With your CM Account and bot registered, you can now add supported services. For
example, to register the Ping Service, use the following command:

```
yarn hardhat account service:add --cm-account  <CMAccountAddress> --private-key <PrivateKeyValue> --service-name cmp.services.ping.v1.PingService --fee 10 --network base_sepolia
```

<details>
<summary>Example output:</summary>

```
yarn run v1.22.22
$ /hgst/work/github.com/chain4travel/camino-messenger-contracts/node_modules/.bin/hardhat account service:add --service-name cmp.services.ping.v1.PingService --fee 10 --cm-account 0xe7C3AAd26fA667a2dc51A4B3c56f8919574b275A --private-key <private-key-deducted> --network base_sepolia
CMAccount: 0xe7C3AAd26fA667a2dc51A4B3c56f8919574b275A
Service Name: cmp.services.ping.v1.PingService
Fee: 10
Restricted Rate: false
Capabilities: undefined
Adding service to CMAccount...
Signer: 0xFe77dcE375C3814F15F8035bCAC1A791D3dCdf21
Tx: 0xe84c6e7e8bfa90f9a35a4534add7b182a5b2240c6b0c1f2433cfad5156b4628d
Done in 2.11s.
```

</details>

#### Command Parameters

- **`--cm-account`:** The EVM contract address of your CM Account.
- **`--private-key`:** The static private key of the wallet used to create your
  CM Account.
- **`--service-name`:** The full-service name to register. For a complete list
  of supported services, consult the [Camino Messenger Protocol documentation](https://buf.build/chain4travel/camino-messenger-protocol/docs).
  (You can also check the service names here in the [services](./services/)
  folder)
- **`--fee`:** The service fee associated with your service in `aCAM` (`wei` in
  EVM terms).
- **`--network`:** Specify the network (as configured in the
  `hardhat.config.js`).

### Summary

Following these steps, you will have:

1. Created your CM Account.
2. Registered your bot with the newly created CM Account.
3. Added and configured services to enhance your account's capabilities.

You can now add your CM Account address to the Camino Messenger Bot
configuration and start running the bot.

## License

The Camino Messenger Contracts are licensed under the terms of the [Camino Messenger License](LICENSE.md).

## Data Protection

Please take note of the [Camino Messenger Data Protection Guidelines](DATA_PROTECTION.md).
