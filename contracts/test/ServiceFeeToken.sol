// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

/**
 * @title ServiceFeeToken
 * @notice This contract is deprecated and removed as part of Milestone 1 service fee removal.
 */
contract ServiceFeeToken {
    function name() public pure returns (string memory) {
        return "USD Service Fee Token";
    }

    function symbol() public pure returns (string memory) {
        return "USD.test";
    }

    function decimals() public pure returns (uint8) {
        return 18;
    }

    function mint(address, uint256) public {}
}
