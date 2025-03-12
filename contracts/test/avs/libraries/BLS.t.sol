// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Test, console2} from "forge-std/Test.sol";
import {BN254} from "src/avs/libraries/BN254.sol";
import {BLS} from "src/avs/libraries/BLS.sol";
import {BLSMock} from "src/avs/mock/BLSMock.sol";

// NOTE: BLS public key will be generated off-chain.
contract BLSTest is Test {
    using BN254 for BN254.G1Point;
    using BN254 for BN254.G2Point;
    using BN254 for BN254.G2Jacobian;

    BLSMock public blsMock;

    function setUp() public {
        blsMock = new BLSMock();
    }

    // Default Anvil Test wallets with explicit private keys
    // (0) Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
    uint256 blsPrivateKey0 = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 % BN254.R;
    BN254.G2Point pubKey0 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey0).toAffine();

    // (1) Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
    uint256 blsPrivateKey1 = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d % BN254.R;
    BN254.G2Point pubKey1 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey1).toAffine();

    // (2) Private Key: 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
    uint256 blsPrivateKey2 = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a % BN254.R;
    BN254.G2Point pubKey2 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey2).toAffine();

    // (3) Private Key: 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
    uint256 blsPrivateKey3 = 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6 % BN254.R;
    BN254.G2Point pubKey3 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey3).toAffine();

    // (4) Private Key: 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
    uint256 blsPrivateKey4 = 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a % BN254.R;
    BN254.G2Point pubKey4 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey4).toAffine();

    // (5) Private Key: 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
    uint256 blsPrivateKey5 = 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba % BN254.R;
    BN254.G2Point pubKey5 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey5).toAffine();

    // (6) Private Key: 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
    uint256 blsPrivateKey6 = 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e % BN254.R;
    BN254.G2Point pubKey6 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey6).toAffine();

    // (7) Private Key: 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
    uint256 blsPrivateKey7 = 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356 % BN254.R;
    BN254.G2Point pubKey7 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey7).toAffine();

    // (8) Private Key: 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
    uint256 blsPrivateKey8 = 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97 % BN254.R;
    BN254.G2Point pubKey8 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey8).toAffine();

    // (9) Private Key: 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
    uint256 blsPrivateKey9 = 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6 % BN254.R;
    BN254.G2Point pubKey9 = BN254.generatorG2().toJacobian().scalarMulG2(blsPrivateKey9).toAffine();

    function testVerifyBatch() public view {
        bytes32 message = bytes32("BLS_signature");

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

        bool valid = blsMock.verifyBatchPubKey(aggSignature, pubKeys, message);
        assert(valid);
    }

    function testVerifySingle() public view {
        bytes32 message = bytes32("BLS_signature");

        // Aggregate individual signatures from all 10 keys
        BN254.G1Point memory signature = blsMock.signMessage(blsPrivateKey0, message);
        bool valid = blsMock.verifySinglePubKey(signature, pubKey0, message);
        assert(valid);
    }

    function testSignMessage() public view {
        bytes32 message = bytes32(keccak256("BLS_signature"));

        BN254.G1Point memory signature = blsMock.signMessage(blsPrivateKey0, message);

        // Deterministic expected signature for verification (example values)
        uint256 signatureX = 12285434762164843075105827537965160525425933292328646296935078631677402976423;
        uint256 signatureY = 10844361192912410820286160611829057254662817948362745840937011919688719785846;
        bool correctSignature = signature.X == signatureX && signature.Y == signatureY;

        assert(correctSignature);
    }
}
