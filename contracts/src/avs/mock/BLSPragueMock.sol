// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import {BLS12_381} from "../libraries/BLS12_381.sol";
import {BLS_Prague} from "../libraries/BLS_Prague.sol";

contract BLSPragueMock {
    function verifySinglePubKey(
        BLS12_381.G1Point memory pubKey,
        BLS12_381.G2Point memory signature,
        bytes memory message
    ) public view returns (bool) {
        return BLS_Prague.verifySinglePubKey(pubKey, signature, message);
    }

    function verifyBatchPubKey(
        BLS12_381.G1Point[] memory pubKeys,
        BLS12_381.G2Point memory aggSignature,
        bytes memory message
    ) public view returns (bool) {
        return BLS_Prague.verifyBatchPubKey(pubKeys, aggSignature, message);
    }

    function hashMessageToG2(bytes memory message) public view returns (BLS12_381.G2Point memory) {
        return BLS12_381.hashMessageToG2(message);
    }

    function mockVerifySinglePubKey(
        BLS12_381.G1Point memory pubKey,
        BLS12_381.G2Point memory signature,
        bytes memory message
    ) public returns (bool) {
        return BLS_Prague.verifySinglePubKey(pubKey, signature, message);
    }

    function mockVerifyBatchPubKey(
        BLS12_381.G1Point[] memory pubKeys,
        BLS12_381.G2Point memory aggSignature,
        bytes memory message
    ) public returns (bool) {
        return BLS_Prague.verifyBatchPubKey(pubKeys, aggSignature, message);
    }
}
