// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {SkateGateway} from "src/gateway/SkateGateway.sol";
import {SkateApp} from "src/skate/SkateApp.sol";

contract DeployGateway is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        address relayer = 0x786775c9ecB916bd7f5a59c150491871fCfCEe86;
        SkateGateway gateway = new SkateGateway(relayer);
        gateway.registerRelayer(relayer);

        vm.stopBroadcast();
    }
}

contract DeploySkateApp is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        address relayer = 0x786775c9ecB916bd7f5a59c150491871fCfCEe86;
        new SkateApp(relayer);

        vm.stopBroadcast();
    }
}
