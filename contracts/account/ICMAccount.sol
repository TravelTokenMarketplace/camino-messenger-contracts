// SPDX-License-Identifier: LGPL-3.0-or-later

pragma solidity 0.8.24;

// `uint256 prefundAmount` is removed as it is no longer used in the contract @2025-08-28
interface ICMAccount {
    function initialize(address manager, address bookingToken, address owner, address upgrader) external;
}
