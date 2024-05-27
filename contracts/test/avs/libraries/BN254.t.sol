// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.20;

import {Test, console2} from "forge-std/Test.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import {IStrategyManager, IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";

import {BN254} from "../../../src/avs/libraries/BN254.sol";
import {BN254G2} from "../../../src/avs/libraries/BN254G2.sol";
import {BLS} from "../../../src/avs/libraries/BLS.sol";
import {BN254Mock} from "../../../src/avs/mock/BN254Mock.sol";

import {Vm} from "forge-std/Vm.sol";

contract BN254Test is Test {
    BN254Mock public bn254Mock;
    using BN254 for BN254.G1Point;
    using BN254G2 for BN254G2.G2Point;
    using BN254G2 for BN254G2.G2Jacobian;

    function setUp() public {
        bn254Mock = new BN254Mock();
    }

    // Private key for 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 - anvil test wallet 0
    uint256 blsPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 % BN254.R;

    BN254G2.G2Point pubKeyG2 =
        BN254G2.G2Point(
            [
                11829690919498240492717914986943835720956660440597679675829055482637656108571,
                2332637959155764237114944033008161729347763881858201243151775165598748610464
            ],
            [
                16844159549615964720138902510035506195836911611113351665735906210620230569546,
                16479896805003260112387785137189286363984636687122540735037453711640344896777
            ]
        );

    BN254.G1Point pubKeyG1 =
        BN254.G1Point(
            10675399722913712025295756746724046633692424641404814156513759249824090513920,
            17239416246131612069967213304966345009400148519694852386882551553158056846891
        );

    function testAddG1() public view {
        BN254.G1Point memory g1 = BN254.generatorG1();
        BN254.G1Point memory g1_ = BN254.negGeneratorG1();

        BN254.G1Point memory zero = bn254Mock.addG1(g1, g1_);
        assert(zero.X == 0 && zero.Y == 0);

        // BN254.G1Point memory p = BN254.negGeneratorG1();
        BN254.G1Point memory result = bn254Mock.addG1(g1, pubKeyG1);
        uint256 expectedX = 9567269949748979410036934805147496020977968906406702497620651285550190253776;
        uint256 expectedY = 11032316415206286220900079148529059454077136295866602837557047271282522280133;
        assertEq(result.X, expectedX);
        assertEq(result.Y, expectedY);
    }

    function testScalarMulG1() public view {
        BN254.G1Point memory g1 = BN254.generatorG1();

        // Multiply by group order
        BN254.G1Point memory zero = bn254Mock.scalarMulG1(g1, BN254.R);
        assert(zero.X == 0 && zero.Y == 0);

        BN254.G1Point memory result = bn254Mock.scalarMulG1(g1, blsPrivateKey);
        assertEq(result.X, pubKeyG1.X);
        assertEq(result.Y, pubKeyG1.Y);
    }

    // TODO: finish functionalities of BN254G2
    function testAddG2() public view {
        BN254G2.G2Point memory g2 = BN254G2.generatorG2();
        BN254G2.G2Point memory g2_ = BN254G2.negGeneratorG2();

        BN254G2.G2Point memory zero = bn254Mock.addG2(g2.toJacobian(), g2_.toJacobian()).toAffine();
        assert(zero.X[0] == 0 && zero.X[1] == 0 && zero.Y[0] == 0 && zero.Y[1] == 0);

        BN254G2.G2Point memory result = bn254Mock.addG2(g2.toJacobian(), pubKeyG2.toJacobian()).toAffine();
        uint256[2] memory expectedX = [
            5089185281081355630965216039114289061805807966378485977784466132038562598413,
            19185479619893335706385157721662219355785721478713265516367218835785551166058
        ];
        uint256[2] memory expectedY = [
            609266954259096775064279300199517870991685285436374777840595655246630800772,
            19352729685888935666605530468178106257173769676130225054715042437407307181534
        ];
        assertEq(result.X[0], expectedX[0]);
        assertEq(result.X[1], expectedX[1]);
        assertEq(result.Y[0], expectedY[0]);
        assertEq(result.Y[1], expectedY[1]);
    }

    function testScalarMulG2() public view {
        BN254G2.G2Point memory g2 = BN254G2.generatorG2();

        // Multiply by group order
        BN254G2.G2Point memory zero = bn254Mock.scalarMulG2(g2.toJacobian(), BN254.R).toAffine();
        assert(zero.X[0] == 0 && zero.X[1] == 0 && zero.Y[0] == 0 && zero.Y[1] == 0);

        BN254G2.G2Point memory result = bn254Mock.scalarMulG2(g2.toJacobian(), blsPrivateKey).toAffine();
        assertEq(result.X[0], pubKeyG2.X[0]);
        assertEq(result.X[1], pubKeyG2.X[1]);
        assertEq(result.Y[0], pubKeyG2.Y[0]);
        assertEq(result.Y[1], pubKeyG2.Y[1]);
    }

    function testToAffine() public view {
        BN254G2.G2Jacobian memory pubKeyG2FromAffine = bn254Mock.toJacobian(pubKeyG2);

        BN254G2.G2Point memory result0 = bn254Mock.toAffine(pubKeyG2FromAffine);
        assertEq(result0.X[0], pubKeyG2.X[0]);
        assertEq(result0.X[1], pubKeyG2.X[1]);
        assertEq(result0.Y[0], pubKeyG2.Y[0]);
        assertEq(result0.Y[1], pubKeyG2.Y[1]);

        BN254G2.G2Jacobian memory pubKeyG2Jac = BN254G2.G2Jacobian(
            [
                3329639083096165885731220919082359124592048864283657592811763527427566306241,
                8806852823229444710244328042456680910008558339241855251396229563951045655138
            ],
            [
                9313001925923743241224391674616694715610801962044656918391441663944101203615,
                6532828922907635986206876418879070648299755993683570467682966769978105914363
            ],
            [
                19942955030888908508699717295835262609524502995259500485011966915941225697003,
                18454188147661245741752119976768248109156097959925982575531255971707286390340
            ]
        );

        BN254G2.G2Point memory result1 = bn254Mock.toAffine(pubKeyG2Jac);
        assertEq(result1.X[0], pubKeyG2.X[0]);
        assertEq(result1.X[1], pubKeyG2.X[1]);
        assertEq(result1.Y[0], pubKeyG2.Y[0]);
        assertEq(result1.Y[1], pubKeyG2.Y[1]);
    }
}
