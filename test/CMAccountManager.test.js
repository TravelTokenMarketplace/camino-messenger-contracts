/**
 * @dev CMAccountManager tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

const { setCode } = require("@nomicfoundation/hardhat-network-helpers");

// Fixtures
const {
    setupSigners,
    deployNullUSDFixture,
    deployCMAccountManagerFixture,
    deployCMAccountImplFixture,
    deployCMAccountManagerWithCMAccountImplFixture,
    deployAndConfigureAllFixture,
} = require("./utils/fixtures");

describe("CMAccountManager", function () {
    describe("Main", function () {
        it("should deploy correctly with the right state", async function () {
            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const DEFAULT_ADMIN_ROLE = await cmAccountManager.DEFAULT_ADMIN_ROLE();
            const PAUSER_ROLE = await cmAccountManager.PAUSER_ROLE();
            const UPGRADER_ROLE = await cmAccountManager.UPGRADER_ROLE();
            const VERSIONER_ROLE = await cmAccountManager.VERSIONER_ROLE();

            // Check roles
            expect(await cmAccountManager.hasRole(DEFAULT_ADMIN_ROLE, signers.managerAdmin.address)).to.be.true;
            expect(await cmAccountManager.hasRole(PAUSER_ROLE, signers.managerPauser.address)).to.be.true;
            expect(await cmAccountManager.hasRole(UPGRADER_ROLE, signers.managerUpgrader.address)).to.be.true;
            expect(await cmAccountManager.hasRole(VERSIONER_ROLE, signers.managerVersioner.address)).to.be.true;

            // Check state
            expect(await cmAccountManager.paused()).to.be.false;
        });

        it("should get role counts correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const DEFAULT_ADMIN_ROLE = await cmAccountManager.DEFAULT_ADMIN_ROLE();
            const PAUSER_ROLE = await cmAccountManager.PAUSER_ROLE();
            const UPGRADER_ROLE = await cmAccountManager.UPGRADER_ROLE();
            const VERSIONER_ROLE = await cmAccountManager.VERSIONER_ROLE();

            expect(await cmAccountManager.getRoleMemberCount(DEFAULT_ADMIN_ROLE)).to.be.equal(1);
            expect(await cmAccountManager.getRoleMemberCount(PAUSER_ROLE)).to.be.equal(1);
            expect(await cmAccountManager.getRoleMemberCount(UPGRADER_ROLE)).to.be.equal(1);
            expect(await cmAccountManager.getRoleMemberCount(VERSIONER_ROLE)).to.be.equal(1);
        });

        it("should set and get booking token addr correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Get booking token address
            expect(await cmAccountManager.getBookingTokenAddress()).to.be.not.reverted;

            // Try to set booking token addr with non-auth address
            await expect(
                cmAccountManager.connect(signers.otherAccount3).setBookingTokenAddress(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");

            // Try to set booking token to invalid addresses
            await expect(cmAccountManager.connect(signers.managerVersioner).setBookingTokenAddress(ethers.ZeroAddress))
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidBookingTokenAddress")
                .withArgs(ethers.ZeroAddress);

            await expect(
                cmAccountManager
                    .connect(signers.managerVersioner)
                    .setBookingTokenAddress(signers.otherAccount1.address),
            )
                .to.be.revertedWithCustomError(cmAccountManager, "InvalidBookingTokenAddress")
                .withArgs(signers.otherAccount1.address);
        });
    });

    describe("Upgrades", function () {
        it("should upgrade correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const CMAccountManagerTest = await ethers.getContractFactory(
                "CMAccountManagerTest",
                signers.managerUpgrader,
            );
            const cmAccountManagerTest = await upgrades.upgradeProxy(cmAccountManager, CMAccountManagerTest);

            await expect(await cmAccountManagerTest.getVersion()).to.be.equal("TESTING");
        });

        it("should not upgrade if the caller does not have the upgrader role", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            const CMAccountManagerTest = await ethers.getContractFactory("CMAccountManagerTest", signers.managerPauser);

            await expect(upgrades.upgradeProxy(cmAccountManager, CMAccountManagerTest)).to.be.revertedWithCustomError(
                CMAccountManagerTest,
                "AccessControlUnauthorizedAccount",
            );
        });
    });

    describe("CMAccount Implementation", function () {
        it("should set CMAccount implementation correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);
            const { cmAccountImpl } = await loadFixture(deployCMAccountImplFixture);

            const cmAccountImplAddress = await cmAccountImpl.getAddress();

            await expect(
                await cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(cmAccountImplAddress),
            )
                .to.emit(cmAccountManager, "CMAccountImplementationUpdated")
                .withArgs(ethers.ZeroAddress, cmAccountImplAddress);
        });

        it("should get CMAccount implementation correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
                deployCMAccountManagerWithCMAccountImplFixture,
            );

            await expect(await cmAccountManager.getAccountImplementation()).to.be.equal(cmAccountImplAddress);
        });

        it("should revert if the implementation is zero code length address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
                deployCMAccountManagerWithCMAccountImplFixture,
            );

            await expect(
                cmAccountManager.connect(signers.managerVersioner).setAccountImplementation(ethers.ZeroAddress),
            )
                .to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidImplementation")
                .withArgs(ethers.ZeroAddress);
        });

        it("should revert if the caller does not have the versioner role", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccountImplAddress } = await loadFixture(
                deployCMAccountManagerWithCMAccountImplFixture,
            );

            await expect(
                cmAccountManager.connect(signers.otherAccount1).setAccountImplementation(cmAccountImplAddress),
            ).to.be.revertedWithCustomError(cmAccountManager, "AccessControlUnauthorizedAccount");
        });
    });

    describe("Pausable", function () {
        it("should pause and unpause the contract", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            await cmAccountManager.connect(signers.managerPauser).pause();
            await expect(await cmAccountManager.paused()).to.be.true;

            await cmAccountManager.connect(signers.managerPauser).unpause();
            await expect(await cmAccountManager.paused()).to.be.false;
        });

        it("should not allow non-pauser to pause", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployCMAccountManagerFixture);

            await expect(cmAccountManager.connect(signers.otherAccount1).pause()).to.be.revertedWithCustomError(
                cmAccountManager,
                "AccessControlUnauthorizedAccount",
            );

            await expect(cmAccountManager.connect(signers.otherAccount1).unpause()).to.be.revertedWithCustomError(
                cmAccountManager,
                "AccessControlUnauthorizedAccount",
            );
        });
    });

    describe("CMAccount", function () {
        it("should create CMAccount correctly", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, cmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const cmAccountManagerAddress = await cmAccountManager.getAddress();
            const cmAccountAddress = await cmAccount.getAddress();

            expect(await cmAccountManager.isCMAccount(cmAccountAddress)).to.be.true;
            expect(await cmAccountManager.isCMAccount(signers.otherAccount1.address)).to.be.false;
            expect(await cmAccountManager.isCMAccount(ethers.ZeroAddress)).to.be.false;
            expect(await cmAccount.getManagerAddress()).to.be.equal(cmAccountManagerAddress);

            // Check balance for native deposit
            expect(await ethers.provider.getBalance(cmAccountAddress)).to.be.equal(ethers.parseEther("100"));
        });

        it("should revert creating for zero code btoken and impl addr", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager, bookingToken } = await loadFixture(deployAndConfigureAllFixture);

            // Get account impl
            const CMAccountImplAddr = await cmAccountManager.getAccountImplementation();

            // Set booking token code to zero
            await network.provider.send("hardhat_setCode", [await bookingToken.getAddress(), "0x"]);

            // Create CMAccount
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader.address, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "InvalidBookingTokenAddress");

            // Set acct impl code to zero
            await network.provider.send("hardhat_setCode", [CMAccountImplAddr, "0x"]);

            // Create CMAccount
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader.address, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidImplementation");
        });

        it("should fail if admin is zero address", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            await expect(
                cmAccountManager.createCMAccount(ethers.ZeroAddress, signers.cmAccountUpgrader, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "CMAccountInvalidAdmin");
        });

        it("should fail if the manager is paused", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            await cmAccountManager.connect(signers.managerPauser).pause();
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(cmAccountManager, "EnforcedPause");
        });

        it("should not fail for any msg.value", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Zero msg.value
            await expect(
                cmAccountManager.createCMAccount(signers.cmAccountAdmin.address, signers.cmAccountUpgrader, {
                    value: 0n,
                }),
            ).to.be.not.reverted;

            const nonZeroMsgValue = ethers.parseEther("300");

            // Non-zero msg.value
            const nonZeroMsgValueTx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader,
                {
                    value: nonZeroMsgValue,
                },
            );

            await expect(nonZeroMsgValueTx).to.be.not.reverted;

            const receipt = await nonZeroMsgValueTx.wait();

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const cmAccountAddress = parsedEvent.args.account;

            // Check balances
            await expect(nonZeroMsgValueTx).to.changeEtherBalances(
                [signers.managerAdmin, cmAccountAddress],
                [-nonZeroMsgValue, nonZeroMsgValue],
            );
        });

        it("should set and get correct account creator", async function () {
            // Set up signers
            await setupSigners();

            const { cmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Create distributor CMAccount
            // This is called with managerAdmin as the signer
            const tx = await cmAccountManager.createCMAccount(
                signers.cmAccountAdmin.address,
                signers.cmAccountUpgrader.address,
                { value: ethers.parseEther("100") },
            );

            const receipt = await tx.wait();

            // Parse event to get the CMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return cmAccountManager.interface.parseLog(log).name === "CMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = cmAccountManager.interface.parseLog(event);
            const newCMAccountAddress = parsedEvent.args.account;

            expect(await cmAccountManager.getCMAccountCreator(newCMAccountAddress)).to.be.equal(
                signers.managerAdmin.address,
            );
        });
    });
});
