// SPDX-License-Identifier: MIT

pragma solidity ^0.8.12;

import {BN254} from "../libraries/BN254.sol";
import {BLS} from "../libraries/BLS.sol";

contract BLSMock {
    function verifySinglePubKey(BN254.G1Point memory signature, BN254.G2Point memory pubKeyG2, bytes32 message) public view returns (bool) {
        return BLS.verifySinglePubKey(signature, pubKeyG2, message);
    }

    function verifyBatchPubKey(BN254.G1Point memory signature, BN254.G2Point[] memory pubKeysG2, bytes32 message) public view returns (bool) {
        return BLS.verifyBatchPubKey(signature, pubKeysG2, message);
    }

    function signMessage(uint256 blsPrivateKey, bytes32 message) public view returns (BN254.G1Point memory signature) {
        return BLS.signMessage(blsPrivateKey, message);
    }
}
