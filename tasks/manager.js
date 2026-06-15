require("@nomicfoundation/hardhat-toolbox");

const MANAGER_SCOPE = scope("manager", "CM Account Manager Tasks");
const BT_SCOPE = scope("btoken", "Booking Token Tasks");

// TODO: Handle transaction failures

const ROLES = [
    "DEFAULT_ADMIN_ROLE",
    "PAUSER_ROLE",
    "UPGRADER_ROLE",
    "VERSIONER_ROLE",
    "SERVICE_REGISTRY_ADMIN_ROLE",
    //"CMACCOUNT_ROLE", // Disabled because it slows down the role:all output a lot. Use `account:list` instead
];

function bold(text) {
    const boldCode = "\x1b[1m";
    const resetCode = "\x1b[0m";
    return `${boldCode}${text}${resetCode}`;
}

function getAddressesForNetwork(hre) {
    let addresses;

    if (hre.network.name === "columbus") {
        console.log("Running on columbus");
        addresses = require("../ignition/deployments/chain-501/deployed_addresses.json");
    } else if (hre.network.name === "camino") {
        console.log("Running on camino");
        addresses = require("../ignition/deployments/chain-500/deployed_addresses.json");
    } else if (hre.network.name === "localhost") {
        console.log("Running on localhost");
        addresses = require("../ignition/deployments/chain-31337/deployed_addresses.json");
    } else if (hre.network.name === "base_sepolia") {
        console.log("Running on base_sepolia");
        addresses = require("../ignition/deployments/chain-84532/deployed_addresses.json");
    } else if (hre.network.name === "base") {
        console.log("Running on base");
        addresses = require("../ignition/deployments/chain-8453/deployed_addresses.json");
    } else {
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    return addresses;
}

async function getManager(hre) {
    const addresses = getAddressesForNetwork(hre);
    return await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);
}

async function getBookingToken(hre) {
    const addresses = getAddressesForNetwork(hre);
    return await ethers.getContractAt("BookingToken", addresses["CaminoMessengerModule#BookingTokenProxy"]);
}

// ServiceFeeToken helper removed as service fees are deprecated

async function handleRoles(taskArgs, hre, action, contractName) {
    let contract;

    if (contractName === "btoken") {
        contract = await getBookingToken(hre);
    } else if (contractName === "manager") {
        contract = await getManager(hre);
    } else {
        throw new Error(`Unsupported contract: ${contractName}`);
    }

    console.log(
        `${action === "grantRole" ? "Granting" : "Revoking"} role ${taskArgs.role} for address ${taskArgs.address}...`,
    );

    const role = await contract[taskArgs.role]();
    const tx = await contract[action](role, taskArgs.address);
    const txReceipt = await tx.wait();
    console.log("Tx:", txReceipt.hash);
}

function handleTransactionError(error, contract) {
    console.error("❌ Transaction failed!");

    if (error.data.data && contract) {
        const decodedError = contract.interface.parseError(error.data.data);
        console.error("Message:", error.message);
        console.error(`Reason: ${decodedError?.name} (${decodedError?.args})`);
    } else if (error.data?.message) {
        console.error(`Reason: ${error.data.message}`);
    } else if (error.message?.includes("[taskArgs.role] is not a function")) {
        console.error("Reason: CMAccount does not have this role.");
    } else if (error.message) {
        console.error("Message:", error.message);
        console.error("Error:", error);
    } else {
        // General error logging
        console.error("An unexpected error occurred.");
        console.error("Error:", error);
    }
}

async function handleServices(taskArgs, hre, action) {
    if (taskArgs.service && taskArgs.json) {
        throw new Error("Cannot provide both --service and --json parameters.");
    }

    let services = [];
    if (taskArgs.service) {
        services = [taskArgs.service];
    } else if (taskArgs.json) {
        const parsed = require(taskArgs.json);
        if (!Array.isArray(parsed) || parsed.length === 0 || !parsed.every((s) => typeof s === "string")) {
            throw new Error("JSON file must be a non-empty array of strings.");
        }
        services = parsed;
    } else {
        throw new Error("You must provide either --service or --json parameter.");
    }

    const manager = await getManager(hre);

    console.log(`${action === "register" ? "Registering" : "Unregistering"} services...`);

    for (const service of services) {
        console.log(`⏳ ${action === "register" ? "Registering" : "Unregistering"} Service:`, service);
        try {
            const tx = await manager[`${action}Service`](service);
            const txReceipt = await tx.wait();
            console.log("✅ Service:", service, "Tx:", txReceipt.hash);
        } catch (error) {
            handleTransactionError(error, manager);
        }
        console.log("-----------------------------------------------------------");
    }
}

MANAGER_SCOPE.task("status", "Print status of deployed contracts").setAction(async (taskArgs, hre) => {
    const addresses = getAddressesForNetwork(hre);

    const managerImplementation = await ethers.getContractAt(
        "CMAccountManager",
        addresses["CaminoMessengerModule#CMAccountManager"],
    );

    const manager = await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);

    const cmAccount = await ethers.getContractAt("CMAccount", addresses["CaminoMessengerModule#CMAccount"]);

    const bookingTokenImplementation = await ethers.getContractAt(
        "BookingToken",
        addresses["CaminoMessengerModule#BookingToken"],
    );

    const bookingToken = await ethers.getContractAt(
        "BookingToken",
        addresses["CaminoMessengerModule#BookingTokenProxy"],
    );

    console.log("================ MANAGER =======================================");
    console.log(`Proxy: ${await manager.getAddress()}`);
    console.log(`Implementation: ${await managerImplementation.getAddress()}`);

    console.log();
    console.log("================ CM ACCOUNT ====================================");
    console.log(`Implementation: ${await cmAccount.getAddress()}`);

    console.log();
    console.log("================ BOOKING TOKEN =================================");
    console.log(`Proxy: ${await bookingToken.getAddress()}`);
    console.log(`Implementation: ${await bookingTokenImplementation.getAddress()}`);

    console.log();
    console.log("================ CONFIGURATION on MANAGER ======================");
    console.log(`CM Account Impl: ${await manager.getAccountImplementation()}`);
});

MANAGER_SCOPE.task("services:register", "Register services")
    .addOptionalParam("json", "Full path to the services json file")
    .addOptionalParam("service", "Service name to register")
    .setAction(async (taskArgs, hre) => {
        await handleServices(taskArgs, hre, "register");
    });

MANAGER_SCOPE.task("services:unregister", "Unregister services")
    .addOptionalParam("json", "Full path to the services json file")
    .addOptionalParam("service", "Service name to unregister")
    .setAction(async (taskArgs, hre) => {
        await handleServices(taskArgs, hre, "unregister");
    });

MANAGER_SCOPE.task("services:list", "List registered services").setAction(async (taskArgs, hre) => {
    const addresses = getAddressesForNetwork(hre);
    const manager = await ethers.getContractAt("CMAccountManager", addresses["CaminoMessengerModule#ManagerProxy"]);
    console.log("Getting all registered services...");
    const services = await manager.getAllRegisteredServiceNames();
    console.log(services);
});

MANAGER_SCOPE.task("role:grant", "Grant role")
    .addParam("role", "Role to grant. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .addParam("address", "Address to grant role to")
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "grantRole", "manager");
    });

MANAGER_SCOPE.task("role:revoke", "Revoke role")
    .addParam("role", "Role to grant. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .addParam("address", "Address to revoke role to")
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "revokeRole", "manager");
    });

MANAGER_SCOPE.task("role:has", "Check if address has role")
    .addParam("role", "Role to check. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .addParam("address", "Address to check")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const role = await manager[taskArgs.role]();
        const hasRole = await manager.hasRole(role, taskArgs.address);
        console.log(`${taskArgs.address} ${hasRole ? "has" : "does not have"} role ${taskArgs.role}`);
        console.log(`${hasRole ? "🟢" : "🔴"}`, hasRole);
    });

MANAGER_SCOPE.task("role:members", "List role members")
    .addParam("role", "Role to list. Ex: SERVICE_REGISTRY_ADMIN_ROLE")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const role = await manager[taskArgs.role]();
        const count = await manager.getRoleMemberCount(role);
        console.log("Role:", taskArgs.role);
        console.log("Total Members:", count);

        // Iterate over the members of the role
        const members = [];
        for (let i = 0; i < count; i++) {
            const member = await manager.getRoleMember(role, i);
            members.push(member);
        }
        console.log(members);
    });

MANAGER_SCOPE.task("role:all", "List all roles").setAction(async (taskArgs, hre) => {
    const manager = await getManager(hre);
    for (const role of ROLES) {
        console.log(`🛡️  ${bold(role)}`);
        console.log(`${bold("=".repeat(48))}`);
        await hre.run({ scope: "manager", task: "role:members" }, { role });
        console.log();
    }
});

MANAGER_SCOPE.task("account:list", "List CM Accounts").setAction(async (taskArgs, hre) => {
    await hre.run({ scope: "manager", task: "role:members" }, { role: "CMACCOUNT_ROLE" });
});

MANAGER_SCOPE.task("account:set-implementation", "Set CMAccount implementation address")
    .addParam("address", "Implementation address to set as the new CMAccount impl")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const tx = await manager.setAccountImplementation(taskArgs.address);
        const txReceipt = await tx.wait();
        console.log("Tx:", txReceipt.hash);
    });

// Obsolete developer fee and prefund set tasks removed

MANAGER_SCOPE.task("btoken:set", "Set Booking Token address on the manager contract")
    .addParam("address", "Booking Token address")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        console.log(`Setting Booking Token Address to ${taskArgs.address}...`);
        const tx = await manager.setBookingTokenAddress(taskArgs.address);
        const txReceipt = await tx.wait();
        console.log("Tx:", txReceipt.hash);
    });

MANAGER_SCOPE.task("pause", "Pause CM Account Manager (stops account creation)").setAction(async (taskArgs, hre) => {
    const manager = await getManager(hre);
    console.log("Pausing CMAccountManager...");
    const tx = await manager.pause();
    const txReceipt = await tx.wait();
    console.log("Tx:", txReceipt.hash);
});

MANAGER_SCOPE.task("unpause", "Unpause CM Account Manager").setAction(async (taskArgs, hre) => {
    const manager = await getManager(hre);
    console.log("Unpausing CMAccountManager...");
    const tx = await manager.unpause();
    const txReceipt = await tx.wait();
    console.log("Tx:", txReceipt.hash);
});

BT_SCOPE.task("role:grant", "Grant role")
    .addParam("role", "Role to grant. Ex: MIN_EXPIRATION_ADMIN_ROLE")
    .addParam("address", "Address to grant role to")
    .setAction(async (taskArgs, hre) => {
        console.log(`📅 ${bold("BookingToken")}`);
        await handleRoles(taskArgs, hre, "grantRole", "btoken");
    });

BT_SCOPE.task("role:revoke", "Revoke role")
    .addParam("role", "Role to grant. Ex: MIN_EXPIRATION_ADMIN_ROLE")
    .addParam("address", "Address to revoke role to")
    .setAction(async (taskArgs, hre) => {
        console.log(`📅 ${bold("BookingToken")}`);
        await handleRoles(taskArgs, hre, "revokeRole", "btoken");
    });

BT_SCOPE.task("role:has", "Check if address has role")
    .addParam("role", "Role to check. Ex: MIN_EXPIRATION_ADMIN_ROLE")
    .addParam("address", "Address to check")
    .setAction(async (taskArgs, hre) => {
        const btoken = await getBookingToken(hre);
        console.log(`📅 ${bold("BookingToken")}`);
        const role = await btoken[taskArgs.role]();
        const hasRole = await btoken.hasRole(role, taskArgs.address);
        console.log(`${taskArgs.address} ${hasRole ? "has" : "does not have"} role ${taskArgs.role}`);
        console.log(`${hasRole ? "🟢" : "🔴"}`, hasRole);
    });

BT_SCOPE.task("status", "Show BookingToken status").setAction(async (taskArgs, hre) => {
    const btoken = await getBookingToken(hre);
    console.log(`📅 ${bold("BookingToken")}`);
    const name = await btoken.name();
    const symbol = await btoken.symbol();
    console.log(`${bold("Address")}: ${await btoken.getAddress()}`);
    console.log(`${bold("Name")}: ${name}`);
    console.log(`${bold("Symbol")}: ${symbol}`);
});

module.exports = {};
