// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract CounterScript is Script {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    function setUp() public {}

    function run() public view {
      bytes32 hash = bytes32(0x03eba95f3aa153f41311f64559ee5862e4f7c2c5bb25fca537cbc53ac9c8f5ef);
      bytes memory signature = hex"1074d76350c1ff54ae960be64130cb7b9ef293b4617785fcbcd53f38c6be97123569075e83ba6aad1b52032b528367d84e669c1ea3753baa9b9b37ba4e49da161c";
      address recover = hash.recover(signature);
      console.log("Recovered address", recover);


      uint256 taskId = 40;
      bytes memory msgData = hex"2874d76350c1ff54ae960be64a";

      bytes32 ethSignedDigest = keccak256(abi.encodePacked(taskId, msgData)).toEthSignedMessageHash();
      console.logBytes32(ethSignedDigest);

      bytes32 signedDigest = keccak256(abi.encodePacked(taskId, msgData));
      console.logBytes32(signedDigest);
    }
}
