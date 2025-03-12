// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

/**
 * @notice Natively supported after EVM Prague as per:
 *  [EIP-2537](https://eips.ethereum.org/EIPS/eip-2537)
 *
 * @title Library for operations on the BN254 elliptic curve
 * @author Skate Organization
 *
 * Curve: Y^2 = X^3 + AX + B (mod P)
 *  where P = 0x1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab
 *        B = 0x4
 *        A = 0x0
 */
library BLS12_381 {
    // Modulus for the underlying base field 𝔽p of the elliptic curve
    // P = [P0 || P1] as bytes64
    uint256 internal constant P0 = 0x1a0111ea397fe69a4b1ba7b6434bacd76;
    uint256 internal constant P1 = 0x4774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab;
    // Order of the base field 𝔽p
    uint256 internal constant R = 0x73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001;

    /// G1 point on Fp
    struct G1Point {
        uint256[2] X;
        uint256[2] Y;
    }

    /// G2 point on extension field 𝔽p²
    /// Encoding of field elements is: X = X[0] + X[1] * i => X[0] is first 2 bytes, while X[1] are the laters.
    struct G2Point {
        uint256[4] X;
        uint256[4] Y;
    }

    /**
     *  @return The result of computing the pairing check
     *         e(Pg1[0], Pg2[0]) *  .... * e(Pg1[n], Pg2[n]) == 1
     *
     *  @dev Cost (32600*k + 37700) gas, where k is the number of pairs
     */
    function pairing(G1Point[] memory P1s, G2Point[] memory P2s) internal view returns (bool) {
        uint256 pairCount = P1s.length;
        require(pairCount == P2s.length, "G1 and G2 points length mismatch");

        uint256 SLICE_SIZE = 12;
        uint256 inputSize = pairCount * SLICE_SIZE;
        uint256[] memory input = new uint256[](inputSize);

        for (uint256 i = 0; i < pairCount; i++) {
            G1Point memory p1i = P1s[i];
            G2Point memory p2i = P2s[i];

            uint256 offset = i * SLICE_SIZE;

            input[offset + 0] = p1i.X[0];
            input[offset + 1] = p1i.X[1];
            input[offset + 2] = p1i.Y[0];
            input[offset + 3] = p1i.Y[1];

            input[offset + 4] = p2i.X[0];
            input[offset + 5] = p2i.X[1];
            input[offset + 6] = p2i.X[2];
            input[offset + 7] = p2i.X[3];

            input[offset + 8] = p2i.Y[0];
            input[offset + 9] = p2i.Y[1];
            input[offset + 10] = p2i.Y[2];
            input[offset + 11] = p2i.Y[3];
        }

        uint256[1] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0f, add(input, 0x20), mul(inputSize, 0x20), out, 0x20)
        }

        require(success, "[precompiles 0x0f]: bls12_381-ec_pairing-failed");

        return out[0] == 0x1;
    }

    //////////////////////////////////////////////////////////////////////////////////////////
    //////////////////////////////// BEGIN G1 operations /////////////////////////////////////
    function generatorG1() internal pure returns (G1Point memory g1) {
        g1.X = [0x17f1d3a73197d7942695638c4fa9ac0f, 0xc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb];
        g1.Y = [0x08b3f481e3aaa0f1a09e30ed741d8ae4, 0xfcf5e095d5d00af600db18cb2c04b3edd03cc744a2888ae40caa232946c5e7e1];
    }

    function negGeneratorG1() internal pure returns (G1Point memory nG1) {
        nG1.X = [0x17f1d3a73197d7942695638c4fa9ac0f, 0xc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb];
        nG1.Y = [0x114d1d6855d545a8aa7d76c8cf2e21f2, 0x67816aef1db507c96655b9d5caac42364e6f38ba0ecb751bad54dcd6b939c2ca];
    }

    /**
     * @return result - the sum of two G1 points
     *
     * @dev Cost 375 gas
     */
    function addG1(G1Point memory p1, G1Point memory p2) internal view returns (G1Point memory result) {
        uint256[8] memory input;
        input[0] = p1.X[0];
        input[1] = p1.X[1];
        input[2] = p1.Y[0];
        input[3] = p1.Y[1];

        input[4] = p2.X[0];
        input[5] = p2.X[1];
        input[6] = p2.Y[0];
        input[7] = p2.Y[1];

        uint256[4] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0b, input, 0x100, out, 0x80)
        }
        require(success, "[precompiles 0x0b]::bls12_381-g1-add-failed");

        result.X = [out[0], out[1]];
        result.Y = [out[2], out[3]];
    }

    /**
     * @return result the product of p * s
     *
     * @dev Cost 12,000 gas
     */
    function scalarMulG1(G1Point memory p, uint256 s) internal view returns (G1Point memory result) {
        uint256[5] memory input;
        input[0] = p.X[0];
        input[1] = p.X[1];
        input[2] = p.Y[0];
        input[3] = p.Y[1];

        input[4] = s;

        uint256[4] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0c, input, 0xa0, out, 0x80)
        }
        require(success, "[precompiles 0x0c]::bls12_381-g1-msm-failed");

        result.X = [out[0], out[1]];
        result.Y = [out[2], out[3]];
    }

    /**
     * @return result the sum of product of Pi * Si
     *
     * @dev Cost (12,000 * length(P) * discount factor) gas. For discount factor, see [EIP-2537](https://eips.ethereum.org/EIPS/eip-2537)).
     * Require P.length == S.length
     */
    function multiScalarMulG1(G1Point[] memory P, uint256[] memory S) internal view returns (G1Point memory result) {
        uint256 pairCount = P.length;
        require(pairCount == S.length, "Length mismatch");

        uint256 SLICE_SIZE = 5;
        uint256 inputSize = pairCount * SLICE_SIZE;
        uint256[] memory input = new uint256[](inputSize);

        uint256[4] memory out;
        bool success;

        for (uint256 i = 0; i < pairCount; i++) {
            G1Point memory pi = P[i];
            uint256 si = S[i];
            uint256 offset = i * SLICE_SIZE;
            input[offset + 0] = pi.X[0];
            input[offset + 1] = pi.X[1];

            input[offset + 2] = pi.Y[0];
            input[offset + 3] = pi.Y[1];

            input[offset + 4] = si;
        }

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0c, add(input, 0x20), mul(inputSize, 0x20), out, 0x80)
        }
        require(success, "[precompiles 0x0c]::bls12_381-g1-msm-failed");

        result.X = [out[0], out[1]];
        result.Y = [out[2], out[3]];
    }

    /**
     * @return result G1 point from a base field element of 𝔽p
     *
     * @dev Cost 5,500 gas
     */
    function mapToG1(bytes32[2] memory Fp) internal view returns (G1Point memory result) {
        uint256[4] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x10, Fp, 0x40, out, 0x80)
        }
        require(success, "[precompiles 0x10]::bls12_381-g1-map-failed");

        result.X = [out[0], out[1]];
        result.Y = [out[2], out[3]];
    }

    /**
     * @return p1 - a G1 point
     *
     * @dev - TODO: Implements robust XOF instead of using below keccak256 trick
     */
    function hashMessageToG1(bytes memory message) internal view returns (G1Point memory p1) {
        bytes32 zero = 0x0;
        bytes32 part0 = keccak256(abi.encodePacked(message, uint256(0)));

        p1 = mapToG1([zero, part0]);
    }

    ////////////////////////////////// END G1 operations /////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////////////////////

    //////////////////////////////////////////////////////////////////////////////////////////
    //////////////////////////////// BEGIN G2 operations /////////////////////////////////////

    /**
     * @return result the product of p * s
     *
     * @dev Cost 22,500 gas
     */
    function scalarMulG2(G2Point memory p, uint256 s) internal view returns (G2Point memory result) {
        uint256[9] memory input;
        input[0] = p.X[0];
        input[1] = p.X[1];
        input[2] = p.X[2];
        input[3] = p.X[3];

        input[4] = p.Y[0];
        input[5] = p.Y[1];
        input[6] = p.Y[2];
        input[7] = p.Y[3];

        input[8] = s;

        uint256[8] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0e, input, 0x120, out, 0x100)
        }
        require(success, "[precompiles 0x0e]::bls12_381-g2-msm-failed");

        result.X = [out[0], out[1], out[2], out[3]];
        result.Y = [out[4], out[5], out[6], out[7]];
    }

    /**
     * @return result the sum of product of Pi * Si
     *
     * @dev Cost (22,500 * length(P) * discount factor) gas. For discount factor, see [EIP-2537](https://eips.ethereum.org/EIPS/eip-2537)).
     * Require P.length == S.length
     */
    function multiScalarMulG2(G2Point[] memory P, uint256[] memory S) internal view returns (G2Point memory result) {
        uint256 pairCount = P.length;
        require(pairCount == S.length, "Length mismatch");

        uint256 PAIR_SIZE = 9;
        uint256 inputSize = pairCount * PAIR_SIZE;
        uint256[] memory input = new uint256[](inputSize);

        uint256[8] memory out;
        bool success;

        for (uint256 i = 0; i < pairCount; i++) {
            G2Point memory pi = P[i];
            uint256 si = S[i];
            uint256 offset = i * PAIR_SIZE;
            input[offset + 0] = pi.X[0];
            input[offset + 1] = pi.X[1];
            input[offset + 2] = pi.X[2];
            input[offset + 3] = pi.X[3];

            input[offset + 4] = pi.Y[0];
            input[offset + 5] = pi.Y[1];
            input[offset + 6] = pi.Y[2];
            input[offset + 7] = pi.Y[3];

            input[offset + 8] = si;
        }

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0e, add(input, 0x20), mul(inputSize, 0x20), out, 0x100)
        }
        require(success, "[precompiles 0x0e]::bls12_381-g2-msm-failed");

        result.X = [out[0], out[1], out[2], out[3]];
        result.Y = [out[4], out[5], out[6], out[7]];
    }

    /**
     * @return result the sum of two G2 points
     *
     * @dev Cost 600 gas
     */
    function addG2(G2Point memory p1, G2Point memory p2) internal view returns (G2Point memory result) {
        uint256[16] memory input;
        input[0] = p1.X[0];
        input[1] = p1.X[1];
        input[2] = p1.X[2];
        input[3] = p1.X[3];
        input[4] = p1.Y[0];
        input[5] = p1.Y[1];
        input[6] = p1.Y[2];
        input[7] = p1.Y[3];

        input[8] = p2.X[0];
        input[9] = p2.X[1];
        input[10] = p2.X[2];
        input[11] = p2.X[3];
        input[12] = p2.Y[0];
        input[13] = p2.Y[1];
        input[14] = p2.Y[2];
        input[15] = p2.Y[3];

        uint256[8] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x0d, input, 0x200, out, 0x100)
        }
        require(success, "[precompiles 0x0d]::bls12_381-g2-add-failed");

        result.X = [out[0], out[1], out[2], out[3]];
        result.Y = [out[4], out[5], out[6], out[7]];
    }

    /**
     * @return result G2 point from a base field element of 𝔽p2
     *
     * @dev Cost 23,800 gas
     */
    function mapToG2(bytes32[4] memory Fp2) internal view returns (G2Point memory result) {
        uint256[8] memory out;
        bool success;

        // solium-disable-next-line security/no-inline-assembly
        assembly {
            success := staticcall(gas(), 0x11, Fp2, 0x80, out, 0x100)
        }
        require(success, "[precompiles 0x11]::bls12_381-g2-map-failed");

        result.X = [out[0], out[1], out[2], out[3]];
        result.Y = [out[4], out[5], out[6], out[7]];
    }

    /**
     * @return p2 - a G2 point
     *
     * @dev - TODO: Implements robust XOF instead of using below keccak256 trick
     */
    function hashMessageToG2(bytes memory message) internal view returns (G2Point memory p2) {
        bytes32 zero = 0x0;
        bytes32 part0 = keccak256(abi.encodePacked(message, uint256(0)));
        bytes32 part1 = keccak256(abi.encodePacked(message, uint256(1)));

        p2 = mapToG2([zero, part0, zero, part1]);
    }

    ////////////////////////////////// END G2 operations /////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////////////////////
}
