// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {BN254} from "src/avs/libraries/BN254.sol";
import {BLS} from "src/avs/libraries/BLS.sol";
import {BLSMock} from "src/avs/mock/BLSMock.sol";
import {BLS12_381} from "src/avs/libraries/BLS12_381.sol";
import {BLSPragueMock} from "src/avs/mock/BLSPragueMock.sol";
import {BLS_Prague} from "src/avs/libraries/BLS_Prague.sol";

contract BLSBenchmark is Script {
    using BLS12_381 for BLS12_381.G1Point;
    using BLS12_381 for BLS12_381.G2Point;

    using BN254 for BN254.G1Point;
    using BN254 for BN254.G2Point;
    using BN254 for BN254.G2Jacobian;

    // Default Anvil Test wallets with explicit private keys
    // (0) Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
    uint256 blsPrivateKey0 = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 % BN254.R;
    BN254.G2Point bn254PubKey0 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey0).toAffine();
    BLS12_381.G1Point bls12381PubKey0 = BLS_Prague.getPubKey(blsPrivateKey0);

    // (1) Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
    uint256 blsPrivateKey1 = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d % BN254.R;
    BN254.G2Point bn254PubKey1 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey1).toAffine();
    BLS12_381.G1Point bls12381PubKey1 = BLS_Prague.getPubKey(blsPrivateKey1);

    // (2) Private Key: 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
    uint256 blsPrivateKey2 = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a % BN254.R;
    BN254.G2Point bn254PubKey2 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey2).toAffine();
    BLS12_381.G1Point bls12381PubKey2 = BLS_Prague.getPubKey(blsPrivateKey2);

    // (3) Private Key: 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
    uint256 blsPrivateKey3 = 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6 % BN254.R;
    BN254.G2Point bn254PubKey3 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey3).toAffine();
    BLS12_381.G1Point bls12381PubKey3 = BLS_Prague.getPubKey(blsPrivateKey3);

    // (4) Private Key: 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
    uint256 blsPrivateKey4 = 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a % BN254.R;
    BN254.G2Point bn254PubKey4 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey4).toAffine();
    BLS12_381.G1Point bls12381PubKey4 = BLS_Prague.getPubKey(blsPrivateKey4);

    // (5) Private Key: 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
    uint256 blsPrivateKey5 = 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba % BN254.R;
    BN254.G2Point bn254PubKey5 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey5).toAffine();
    BLS12_381.G1Point bls12381PubKey5 = BLS_Prague.getPubKey(blsPrivateKey5);

    // (6) Private Key: 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
    uint256 blsPrivateKey6 = 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e % BN254.R;
    BN254.G2Point bn254PubKey6 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey6).toAffine();
    BLS12_381.G1Point bls12381PubKey6 = BLS_Prague.getPubKey(blsPrivateKey6);

    // (7) Private Key: 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
    uint256 blsPrivateKey7 = 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356 % BN254.R;
    BN254.G2Point bn254PubKey7 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey7).toAffine();
    BLS12_381.G1Point bls12381PubKey7 = BLS_Prague.getPubKey(blsPrivateKey7);

    // (8) Private Key: 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
    uint256 blsPrivateKey8 = 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97 % BN254.R;
    BN254.G2Point bn254PubKey8 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey8).toAffine();
    BLS12_381.G1Point bls12381PubKey8 = BLS_Prague.getPubKey(blsPrivateKey8);

    // (9) Private Key: 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
    uint256 blsPrivateKey9 = 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6 % BN254.R;
    BN254.G2Point bn254PubKey9 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey9).toAffine();
    BLS12_381.G1Point bls12381PubKey9 = BLS_Prague.getPubKey(blsPrivateKey9);

    BLSMock blsMock = BLSMock(0xc4A3A9e357A59c1Ca894B3Ebbdd00d88B5bB3A5c);
    BLSPragueMock blsPragueMock = BLSPragueMock(0x8c6E8631AF90312d88Ca68738A340f3101937235);

    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        ///////////////////////////////////////////////////////////////////////////
        /////////////////////////////// Cancun versions ///////////////////////////
        bytes32 message = bytes32("BLS_signature");
        BN254.G1Point memory signature = blsMock.signMessage(blsPrivateKey0, message);
        blsMock.mockVerifySinglePubKey(signature, bn254PubKey0, message);

        // Aggregate signatures from all 10 keys
        BN254.G1Point memory aggSignature = blsMock.signMessage(blsPrivateKey0, message);
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey1, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey2, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey3, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey4, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey5, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey6, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey7, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey8, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey9, message));
        // Create an array of 10 public keys
        BN254.G2Point[] memory pubKeys = new BN254.G2Point[](10);
        pubKeys[0] = bn254PubKey0;
        pubKeys[1] = bn254PubKey1;
        pubKeys[2] = bn254PubKey2;
        pubKeys[3] = bn254PubKey3;
        pubKeys[4] = bn254PubKey4;
        pubKeys[5] = bn254PubKey5;
        pubKeys[6] = bn254PubKey6;
        pubKeys[7] = bn254PubKey7;
        pubKeys[8] = bn254PubKey8;
        pubKeys[9] = bn254PubKey9;

        blsMock.mockVerifyBatchPubKey(aggSignature, pubKeys, message);

        ///////////////////////////////////////////////////////////////////////////
        /////////////////////////////// Prague versions ///////////////////////////
        bytes memory message2 = "BLS_Signature";
        BLS12_381.G2Point memory signature2 = BLS_Prague.signMessage(blsPrivateKey0, message2);
        blsPragueMock.mockVerifySinglePubKey(bls12381PubKey0, signature2, message2);

        // Aggregate signatures from all 10 keys
        BLS12_381.G2Point memory aggSignature2 = BLS_Prague.signMessage(blsPrivateKey0, message2);
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey1, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey2, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey3, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey4, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey5, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey6, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey7, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey8, message2));
        aggSignature2 = aggSignature2.addG2(BLS_Prague.signMessage(blsPrivateKey9, message2));

        // Create an array of 10 public keys
        BLS12_381.G1Point[] memory pubKeys2 = new BLS12_381.G1Point[](10);
        pubKeys2[0] = bls12381PubKey0;
        pubKeys2[1] = bls12381PubKey1;
        pubKeys2[2] = bls12381PubKey2;
        pubKeys2[3] = bls12381PubKey3;
        pubKeys2[4] = bls12381PubKey4;
        pubKeys2[5] = bls12381PubKey5;
        pubKeys2[6] = bls12381PubKey6;
        pubKeys2[7] = bls12381PubKey7;
        pubKeys2[8] = bls12381PubKey8;
        pubKeys2[9] = bls12381PubKey9;

        blsPragueMock.mockVerifyBatchPubKey(pubKeys2, aggSignature2, message2);

        vm.stopBroadcast();
    }
}
