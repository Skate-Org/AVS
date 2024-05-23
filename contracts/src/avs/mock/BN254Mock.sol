// SPDX-License-Identifier: MIT

pragma solidity ^0.8.12;

import {BN254} from "../libraries/BN254.sol";
import {BN254G2} from "../libraries/BN254G2.sol";

contract BN254Mock {
    function addG1(BN254.G1Point memory p0, BN254.G1Point memory p1) public view returns (BN254.G1Point memory) {
        return BN254.addG1(p0, p1);
    }

    function scalarMulG1(BN254.G1Point memory p, uint256 k) public view returns (BN254.G1Point memory) {
        return BN254.scalarMulG1(p, k);
    }

    function addG2(BN254G2.G2Jacobian memory p0, BN254G2.G2Jacobian memory p1) public pure returns (BN254G2.G2Jacobian memory) {
        return BN254G2.addG2(p0, p1);
    }

    function scalarMulG2(BN254G2.G2Jacobian memory p, uint256 k) public pure returns (BN254G2.G2Jacobian memory) {
        return BN254G2.scalarMulG2(p, k);
    }

    function toAffine(BN254G2.G2Jacobian memory p0) public view returns (BN254G2.G2Point memory) {
        return BN254G2.toAffine(p0);
    }

    function toJacobian(BN254G2.G2Point memory p0) public pure returns (BN254G2.G2Jacobian memory) {
        return BN254G2.toJacobian(p0);
    }
}
