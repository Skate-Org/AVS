// SPDX-License-Identifier: MIT

pragma solidity ^0.8.12;

import {BN254} from "../libraries/BN254.sol";

contract BN254Mock {
    function addG1(BN254.G1Point memory p0, BN254.G1Point memory p1) public view returns (BN254.G1Point memory) {
        return BN254.addG1(p0, p1);
    }

    function scalarMulG1(BN254.G1Point memory p, uint256 k) public view returns (BN254.G1Point memory) {
        return BN254.scalarMulG1(p, k);
    }

    function addG2(BN254.G2Jacobian memory p0, BN254.G2Jacobian memory p1) public pure returns (BN254.G2Jacobian memory) {
        return BN254.addG2(p0, p1);
    }

    function doubleG2(BN254.G2Jacobian memory p0) public pure returns (BN254.G2Jacobian memory) {
        return BN254.doubleG2(p0);
    }

    function scalarMulG2(BN254.G2Jacobian memory p, uint256 k) public pure returns (BN254.G2Jacobian memory) {
        return BN254.scalarMulG2(p, k);
    }

    function toAffine(BN254.G2Jacobian memory p0) public view returns (BN254.G2Point memory) {
        return BN254.toAffine(p0);
    }

    function toJacobian(BN254.G2Point memory p0) public pure returns (BN254.G2Jacobian memory) {
        return BN254.toJacobian(p0);
    }

    function e2Inverse(uint256[2] memory e) public view returns (uint256[2] memory) {
        return BN254.inverse(e);
    }
}
