// SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.24;

interface ICMAccountManager {
    function getAccountImplementation() external view returns (address);

    function isCMAccount(address account) external view returns (bool);

    function getRegisteredServiceHashByName(string memory serviceName) external view returns (bytes32 serviceHash);

    function getServiceHashByName(string memory serviceName) external view returns (bytes32 serviceHash);

    function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string memory serviceName);

    function getServiceNameByHash(bytes32 serviceHash) external view returns (string memory serviceName);
}
