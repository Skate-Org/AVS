// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {BLSPragueMock} from "src/avs/mock/BLSPragueMock.sol";
import {BLSMock} from "src/avs/mock/BLSMock.sol";

contract DeployBLSMock is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        new BLSMock();
        new BLSPragueMock();

        vm.stopBroadcast();
    }
}
