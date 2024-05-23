// SPDX-License-Identifier: MIT
//
// G1 operations adapted from:
// + HarryR: https://github.com/HarryR/solcrypto/blob/master/contracts/altbn128.sol (MIT license):
// + LayrLabs Inc.: https://github.com/Layr-Labs/eigenlayer-middleware/blob/fb313de7ccfa8b4391e51a3c20b213aab2d035b7/src/libraries/BN254.sol (MIT license)

pragma solidity ^0.8.12;

import {Mod} from "./Mod.sol";

/**
 *
 * @notice add, mul in G1, and pairing is natively supported.
 *ETH natively supports alt_bn128 in
 *  [EIP-196](https://eips.ethereum.org/EIPS/eip-196) and [EIP-197](https://eips.ethereum.org/EIPS/eip-197)
 *  These refers to the same curve with different naming.
 *
 * @title Library for operations on the BN254 elliptic curve
 * @author Skate Organization
 *
 * @dev Operations on 𝔽p² curve, or in G2 subgroup, e.g. add, mul, etc. should also done off-chain
 */
library BN254 {
    // modulus for the underlying field 𝔽p of the elliptic curve
    uint256 internal constant P = 21888242871839275222246405745257275088696311157297823662689037894645226208583;
    // modulus for the underlying field 𝔽r of the elliptic curve
    uint256 internal constant R = 21888242871839275222246405745257275088548364400416034343698204186575808495617;

    // points on elliptic curve over field 𝔽p
    struct G1Point {
        uint256 X;
        uint256 Y;
    }

    /// points on elliptic curve over extension field 𝔽p²
    /// Encoding of field elements is: X[0] + X[1] * i
    /// WARNING: this is reverse of the encoding EIP-197 conventinon (which doesn't follow normal convention).
    /// https://eips.ethereum.org/EIPS/eip-197#encoding, specifically:
    /// "Elements a * i + b of F_p^2 are encoded as two elements of F_p, (a, b)."
    struct G2Point {
        uint256[2] X;
        uint256[2] Y;
    }

    function pairing(G1Point memory a1, G2Point memory a2, G1Point memory b1, G2Point memory b2) internal view returns (bool) {
        uint256[] memory input = new uint256[](12);
        // uint256[12] memory input;

        input[0] = a1.X;
        input[1] = a1.Y;

        // NOTE: mind the G2 order, as discussed above
        input[2] = a2.X[1];
        input[3] = a2.X[0];
        input[4] = a2.Y[1];
        input[5] = a2.Y[0];

        input[6] = b1.X;
        input[7] = b1.Y;

        // NOTE: mind the G2 order, as discussed above
        input[8] = b2.X[1];
        input[9] = b2.X[0];
        input[10] = b2.Y[1];
        input[11] = b2.Y[0];

        return _ecPairing(input);
    }

    /**
     *  @return The result of computing the pairing check
     *         e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1
     *  @dev Gas cost is cap at gas() - 2000
     */
    function _ecPairing(uint256[] memory input) internal view returns (bool) {
        uint256[1] memory out;
        bool success;

        uint256 inputSize = input.length;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x08, add(input, 0x20), mul(inputSize, 0x20), out, 0x20)
        }

        require(success, "ec-pairing-failed");

        return out[0] != 0;
    }

    //////////////////////////////////////////////////////////////////////////////////////////
    //////////////////////////////// BEGIN G1 operations /////////////////////////////////////

    /// common generators is (1,2)
    function generatorG1() internal pure returns (G1Point memory) {
        return G1Point(1, 2);
    }

    function negGeneratorG1() internal pure returns (G1Point memory) {
        return G1Point(1, P - 2);
    }

    /**
     * @return The negation of `p`, i.e. p.plus(p.negate()) should be zero.
     */
    function negateG1(G1Point memory p) internal pure returns (G1Point memory) {
        return G1Point(p.X, (P - p.Y) % P);
    }

    /**
     * Plus to G1 points.
     *
     * @return r the sum of two points of G1
     * @dev NOTE: only G1 operations are precompiled
     */
    function addG1(G1Point memory p1, G1Point memory p2) internal view returns (G1Point memory r) {
        uint256[4] memory input;
        input[0] = p1.X;
        input[1] = p1.Y;
        input[2] = p2.X;
        input[3] = p2.Y;

        uint256[2] memory result = _ecadd(input);

        r.X = result[0];
        r.Y = result[1];
    }

    /**
     * Plus P, Q on elliptic curve to a result point Z
     *
     * @param input P.X, P.Y, Q.X, Q.Y
     * @return r Z.X, Z.Y
     */
    function _ecadd(uint256[4] memory input) private view returns (uint256[2] memory r) {
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        // 500 gas as per EIP-196
        assembly {
            success := staticcall(sub(gas(), 2000), 0x06, input, 0x80, r, 0x40)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }

        require(success, "ec-add-failed");
    }

    /**
     * @return r the product of a point on G1 by a scalar.
     *
     * @dev NOTE: only G1 operations are precompiled
     */
    function scalarMulG1(G1Point memory p, uint256 s) internal view returns (G1Point memory r) {
        uint256[3] memory input;
        input[0] = p.X;
        input[1] = p.Y;
        input[2] = s;
        uint256[2] memory result = _ecmul(input);
        r.X = result[0];
        r.Y = result[1];
    }

    function _ecmul(uint256[3] memory input) internal view returns (uint256[2] memory r) {
        bool success;
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(sub(gas(), 2000), 0x07, input, 0x60, r, 0x40)
            // Use "invalid" to make gas estimation work
            switch success
            case 0 {
                invalid()
            }
        }
        require(success, "ec-mul-failed");
    }

    /// @return hashedG1 the keccak256 hash of the G1 Point
    /// @dev used for BLS signatures
    function hashG1Point(G1Point memory pk) internal pure returns (bytes32 hashedG1) {
        // Equivalent to:
        // hashedG1 = keccak256(abi.encodePacked(pk.X, pk.Y));
        assembly {
            mstore(0, mload(pk))
            mstore(0x20, mload(add(0x20, pk)))
            hashedG1 := keccak256(0, 0x40)
        }
    }

    /// @return hashedG2 the keccak256 hash of the G2 Point
    /// @dev used for BLS signatures
    function hashG2Point(G2Point memory pk) internal pure returns (bytes32 hashedG2) {
        // Equivalent to:
        // hashedG2 = keccak256(abi.encodePacked(pk.X[0], pk.X[1], pk.Y[0], pk.Y[1]));
        assembly {
            mstore(0, mload(pk))
            mstore(0x20, mload(add(0x20, pk)))
            mstore(0x40, mload(add(0x40, pk)))
            mstore(0x60, mload(add(0x60, pk)))
            hashedG2 := keccak256(0, 0x80)
        }
    }

    // since p == 3 (mod 4), uses the identity
    // a^(2 * (p+1)/4) == a^(p+1) == a^2  (mod p)
    //
    // => y = sqrt(beta) = beta^((p + 1)/4) [IF EXIST]
    uint256 internal constant EXP_SQRT = 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52;

    /**
     * hashToG1 implements the try and increment strategy, non-constant time but cheaper on-chain
     */
    function hashToG1(bytes32 digestHash) internal view returns (G1Point memory) {
        uint256 x = uint256(digestHash) % P;

        while (true) {
            // RHS = x^3 + 3
            // G1 curve: y^2 = x^3 + 3
            uint256 RHS = addmod(mulmod(mulmod(x, x, P), x, P), 3, P);
            uint256 y = Mod.modExp(RHS, EXP_SQRT, P);

            if (RHS == mulmod(y, y, P)) {
                return G1Point(x, y);
            }

            x = addmod(x, 1, P);
        }

        return G1Point(0, 0);
    }

    ////////////////////////////////// END G1 operations /////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////////////////////

    // generator of group G2
    /// @dev Generator point in F_p^2 is of the form: (x0 + ix1, y0 + iy1).
    uint256 internal constant G2x0 = 10857046999023057135944570762232829481370756359578518086990519993285655852781;
    uint256 internal constant G2x1 = 11559732032986387107991004021392285783925812861821192530917403151452391805634;
    uint256 internal constant G2y0 = 8495653923123431417604973247489272438418190587263600148770280649306958101930;
    uint256 internal constant G2y1 = 4082367875863433681332203403145435568316851327593401208105741076214120093531;

    // Negation of the generator of group G2
    uint256 internal constant nG2y0 = 13392588948715843804641432497768002650278120570034223513918757245338268106653;
    uint256 internal constant nG2y1 = 17805874995975841540914202342111839520379459829704422454583296818431106115052;

    /// @notice returns the G2 generator
    function generatorG2() internal pure returns (G2Point memory) {
        return G2Point([G2x0, G2x1], [G2y0, G2y1]);
    }

    function negGeneratorG2() internal pure returns (G2Point memory) {
        return G2Point([G2x0, G2x1], [nG2y0, nG2y1]);
    }
}
