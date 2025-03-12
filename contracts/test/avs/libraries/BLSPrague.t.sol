// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test, console2} from "forge-std/Test.sol";
import {BLS12_381} from "../../../src/avs/libraries/BLS12_381.sol";
import {BLSPragueMock} from "../../../src/avs/mock/BLSPragueMock.sol";
import {BLS_Prague} from "../../../src/avs/libraries/BLS_Prague.sol";

contract BLSPragueTest is Test {
    using BLS12_381 for BLS12_381.G1Point;
    using BLS12_381 for BLS12_381.G2Point;

    BLSPragueMock public blsPragueMock;

    // Default Anvil Test wallets with explicit private keys included
    // (0) Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
    uint256 blsPrivateKey0 = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 % BLS12_381.R;
    BLS12_381.G1Point pubKey0 = BLS_Prague.getPubKey(blsPrivateKey0);

    // (1) Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
    uint256 blsPrivateKey1 = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d % BLS12_381.R;
    BLS12_381.G1Point pubKey1 = BLS_Prague.getPubKey(blsPrivateKey1);

    // (2) Private Key: 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
    uint256 blsPrivateKey2 = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a % BLS12_381.R;
    BLS12_381.G1Point pubKey2 = BLS_Prague.getPubKey(blsPrivateKey2);

    // (3) Private Key: 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
    uint256 blsPrivateKey3 = 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6 % BLS12_381.R;
    BLS12_381.G1Point pubKey3 = BLS_Prague.getPubKey(blsPrivateKey3);

    // (4) Private Key: 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
    uint256 blsPrivateKey4 = 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a % BLS12_381.R;
    BLS12_381.G1Point pubKey4 = BLS_Prague.getPubKey(blsPrivateKey4);

    // (5) Private Key: 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
    uint256 blsPrivateKey5 = 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba % BLS12_381.R;
    BLS12_381.G1Point pubKey5 = BLS_Prague.getPubKey(blsPrivateKey5);

    // (6) Private Key: 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
    uint256 blsPrivateKey6 = 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e % BLS12_381.R;
    BLS12_381.G1Point pubKey6 = BLS_Prague.getPubKey(blsPrivateKey6);

    // (7) Private Key: 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
    uint256 blsPrivateKey7 = 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356 % BLS12_381.R;
    BLS12_381.G1Point pubKey7 = BLS_Prague.getPubKey(blsPrivateKey7);

    // (8) Private Key: 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
    uint256 blsPrivateKey8 = 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97 % BLS12_381.R;
    BLS12_381.G1Point pubKey8 = BLS_Prague.getPubKey(blsPrivateKey8);

    // (9) Private Key: 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
    uint256 blsPrivateKey9 = 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6 % BLS12_381.R;
    BLS12_381.G1Point pubKey9 = BLS_Prague.getPubKey(blsPrivateKey9);

    function setUp() public {
        blsPragueMock = new BLSPragueMock();
    }

    function testVerifyBatch() public view {
        bytes memory message = "A Skate data batch";

        // Aggregate signatures from all 10 keys
        BLS12_381.G2Point memory aggSignature = BLS_Prague.signMessage(blsPrivateKey0, message);
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey1, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey2, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey3, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey4, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey5, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey6, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey7, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey8, message));
        aggSignature = aggSignature.addG2(BLS_Prague.signMessage(blsPrivateKey9, message));

        // Create an array of 10 public keys
        BLS12_381.G1Point[] memory pubKeys = new BLS12_381.G1Point[](10);
        pubKeys[0] = pubKey0;
        pubKeys[1] = pubKey1;
        pubKeys[2] = pubKey2;
        pubKeys[3] = pubKey3;
        pubKeys[4] = pubKey4;
        pubKeys[5] = pubKey5;
        pubKeys[6] = pubKey6;
        pubKeys[7] = pubKey7;
        pubKeys[8] = pubKey8;
        pubKeys[9] = pubKey9;

        bool valid = blsPragueMock.verifyBatchPubKey(pubKeys, aggSignature, message);
        assert(valid);
    }

    function testVerifySingle() public view {
        bytes memory message = "A Skate data batch";

        BLS12_381.G2Point memory signature = BLS_Prague.signMessage(blsPrivateKey0, message);
        bool valid = blsPragueMock.verifySinglePubKey(pubKey0, signature, message);
        assert(valid);
    }

    function testHashMessage() public view {
        bytes memory message = "A Skate data batch";
        blsPragueMock.hashMessageToG2(message);
    }
}

