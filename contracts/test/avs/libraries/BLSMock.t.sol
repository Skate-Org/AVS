// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

import {Test, console2} from "forge-std/Test.sol";
import {BN254} from "../../../src/avs/libraries/BN254.sol";
import {BLS} from "../../../src/avs/libraries/BLS.sol";
import {BLSMock} from "../../../src/avs/mock/BLSMock.sol";

// NOTE: BLS publickey will be generated off-chain, as for any G2 operations
contract BLSTest is Test {
    using BN254 for BN254.G1Point;

    BLSMock public blsMock;

    function setUp() public {
        blsMock = new BLSMock();
    }

    // Private key for 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 - anvil test wallet 0
    uint256 blsPrivateKey0 = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 % BN254.R;
    BN254.G2Point pubKey0 =
        BN254.G2Point(
            [
                11829690919498240492717914986943835720956660440597679675829055482637656108571,
                2332637959155764237114944033008161729347763881858201243151775165598748610464
            ],
            [
                16844159549615964720138902510035506195836911611113351665735906210620230569546,
                16479896805003260112387785137189286363984636687122540735037453711640344896777
            ]
        );

    // Private key for 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 - anvil test wallet 1
    uint256 blsPrivateKey1 = 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d % BN254.R;
    BN254.G2Point pubKey1 =
        BN254.G2Point(
            [
                16553675810662402913332642109180320946410252990138787612031442987029708414951,
                3950887887877429316505810968405340209351770580495254664272056647280823110872
            ],
            [
                19767199458582240114868020780170509926238616450007142624546444508397469746103,
                17996234567193988288700669608184288873172459020548483946003084471765254786022
            ]
        );

    // Private key for 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC - anvil test wallet 1
    uint256 blsPrivateKey2 = 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a % BN254.R;
    BN254.G2Point pubKey2 =
        BN254.G2Point(
            [
                9792872845262080300649079268741298251040032822358403315276048280727343356779,
                15847210887164936986146813501197001106156225637036600642599857938920002570208
            ],
            [
                14792701865640318097083912335496875808840104722124871661180771905152637430638,
                15078585932495272831110558597484648487865512330810900597399063864832124095287
            ]
        );

    function testVerifyBatch() public view {
        bytes32 message = bytes32("BLS_signature");

        BN254.G1Point memory aggSignature = blsMock.signMessage(blsPrivateKey0, message);
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey1, message));
        aggSignature = BN254.addG1(aggSignature, blsMock.signMessage(blsPrivateKey2, message));


        BN254.G2Point[] memory pubKeys = new BN254.G2Point[](3);
        pubKeys[0] = pubKey0;
        pubKeys[1] = pubKey1;
        pubKeys[2] = pubKey2;

        bool valid = blsMock.verifyBatchPubKey(aggSignature, pubKeys, message);

        assert(valid);
    }

    function testVerifySingle() public view {
        bytes32 message = bytes32("BLS_signature");

        BN254.G1Point memory signature = blsMock.signMessage(blsPrivateKey0, message);

        bool valid = blsMock.verifySinglePubKey(signature, pubKey0, message);

        assert(valid);
    }

    function testSignMessage() public view {
        bytes32 message = bytes32(keccak256("BLS_signature"));

        BN254.G1Point memory signature = blsMock.signMessage(blsPrivateKey0, message);

        // deterministic message
        uint256 signatureX = 12285434762164843075105827537965160525425933292328646296935078631677402976423;
        uint256 signatureY = 10844361192912410820286160611829057254662817948362745840937011919688719785846;
        bool correctSignature = signature.X == signatureX && signature.Y == signatureY;

        assert(correctSignature);
    }
}
