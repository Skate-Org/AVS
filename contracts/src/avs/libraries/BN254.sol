// SPDX-License-Identifier: MIT
// 
// G1 operations adapted from:
// + HarryR: https://github.com/HarryR/solcrypto/blob/master/contracts/altbn128.sol (MIT license):
// + LayrLabs Inc.: https://github.com/Layr-Labs/eigenlayer-middleware/blob/fb313de7ccfa8b4391e51a3c20b213aab2d035b7/src/libraries/BN254.sol (MIT license)

pragma solidity ^0.8.12;

import {Mod} from "./Mod.sol";

/**
 *
 * @notice add, mul in G1, and pairing is natively supported as per:
 *  [EIP-196](https://eips.ethereum.org/EIPS/eip-196) and [EIP-197](https://eips.ethereum.org/EIPS/eip-197)
 *  Note that "alt_bn128" refers to the same curve with different naming.
 *
 * @title Library for operations on the BN254 elliptic curve
 * @author Skate Organization
 *
 * @dev Operations on 𝔽p² curve, or in G2 subgroup, e.g. add, mul, etc. should also done off-chain
 */
library BN254 {
    using BN254 for BN254.G1Point;

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

    //////////////////////////////////////////////////////////////////////////////////////////
    //////////////////////////////// BEGIN G2 operations /////////////////////////////////////

    using BN254 for uint256[2]; // E2 elements
    using BN254 for BN254.G2Point;
    using BN254 for BN254.G2Jacobian;

    struct G2Jacobian {
        uint256[2] X;
        uint256[2] Y;
        uint256[2] Z;
    }

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

    //////////////////////////////////////////////////////////////////////////////////////////
    //////////////////////////////// BEGIN G2 operations /////////////////////////////////////

    function toAffine(G2Jacobian memory pJac) internal view returns (G2Point memory pAff) {
        if (!pJac.Z.isZero()) {
            uint256[2] memory zInverse = pJac.Z.inverse();
            uint256[2] memory zInverseSquare = zInverse.square();

            pAff.X = pJac.X.mul(zInverseSquare);
            pAff.Y = pJac.Y.mul(zInverseSquare).mul(zInverse);
        }
    }

    function toJacobian(G2Point memory pAff) internal pure returns (G2Jacobian memory pJac) {
        // infinity
        if (pAff.X.isZero() && pAff.Y.isZero()) {
            pJac.X = [uint256(1), 0];
            pJac.Y = [uint256(1), 0];
            pJac.Z = [uint256(0), 0];
        } else {
            pJac.X = pAff.X;
            pJac.Y = pAff.Y;
            pJac.Z = [uint256(1), 0];
        }
    }

    /**
     * scalarMulG2 performs scalar multiplication of a G2 Jacobian point
     * NOTE: ensure k < |E'(𝔽p²)|, else this would perform redundant calculations
     */
    function scalarMulG2(G2Jacobian memory point, uint256 k) internal pure returns (G2Jacobian memory result) {
        G2Jacobian memory temp = point; // Temporary variable to hold intermediate results
        while (k != 0) {
            if ((k & 1) != 0) {
                result = addG2(result, temp);
            }
            temp = temp.doubleG2();
            k >>= 1;
        }
    }

    /**
     * addG2 performs curve addition of two G2 Jacobian points
     *
     * @dev See https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#addition-add-2007-bl
     */
    function addG2(G2Jacobian memory p, G2Jacobian memory q) internal pure returns (G2Jacobian memory result) {
        if (p.Z.isZero()) {
            return q;
        }
        if (q.Z.isZero()) {
            return p;
        }

        // pZZ = p.Z^2 || qZZ = q.Z^2
        uint256[2] memory pZZ = p.Z.square();
        uint256[2] memory qZZ = q.Z.square();

        // u2 = pX * qZZ || u1 = qX * pZZ
        uint256[2] memory u2 = mul(p.X, qZZ);
        uint256[2] memory u1 = mul(q.X, pZZ);

        // s2 = pY * qZ * qZZ || s1 = pY * pZ * pZZ
        uint256[2] memory s2 = p.Y.mul(q.Z).mul(qZZ);
        uint256[2] memory s1 = q.Y.mul(p.Z).mul(pZZ);

        // NOTE: disable this optimization save 1.1k gas (out of 23.7K -> 4.5%). In addition, this optimization unlikely to happen since pubkey should be different.
        // if (u1.eq(u2) && s1.eq(s2)) {
        //     // equivalent to p == q
        //     return p.doubleG2();
        // }

        // H = u2 - u1 = qX * pZZ - u1
        uint256[2] memory H = u2.sub(u1);
        // I = (2H)^2
        uint256[2] memory I = square(H.add(H));
        // J = H * I
        uint256[2] memory J = mul(H, I);
        // V = u1 * I
        uint256[2] memory V = mul(u1, I);
        // r = 2 * (s2 - s1) = 2(qY * pZ * pZZ - s1)
        uint256[2] memory r = s2.sub(s1).scalar_mul(2);

        // result.X = r^2 - J - 2V
        result.X = r.square().sub(J).sub(V).sub(V);
        // result.Y = r(V - result.X) - 2(s1 * J)
        result.Y = sub(V, result.X).mul(r);
        result.Y = result.Y.sub(mul(s1, J).scalar_mul(2));
        // result.Z = [(pZ+qZ)^2 - pZZ - qZZ] * H
        result.Z = add(p.Z, q.Z).square();
        result.Z = result.Z.sub(pZZ).sub(qZZ);
        result.Z = result.Z.mul(H);
    }

    function subG2(G2Jacobian memory p, G2Jacobian memory q) internal pure returns (G2Jacobian memory result) {
        q.Y = q.Y.neg();
        return addG2(p, q);
    }

    /**
     * doubleG2 performs doubling of an G2 Jacobian point
     *
     * @dev See https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#doubling-dbl-2007-bl
     */
    function doubleG2(G2Jacobian memory p) internal pure returns (G2Jacobian memory result) {
        // x^2
        uint256[2] memory xx = square(p.X);
        // y^2, y^4
        uint256[2] memory yy = square(p.Y);
        uint256[2] memory yyyy = yy.square();
        // z^2
        uint256[2] memory zz = p.Z.square();

        // S = 2 * [(p.X+yy)^2 - xx - yyyy]
        uint256[2] memory S = add(p.X, yy).square().sub(xx).sub(yyyy).scalar_mul(2);
        // M = 3xx
        uint256[2] memory M = scalar_mul(xx, 3);
        uint256[2] memory T = sub(M.square(), S.scalar_mul(2));

        result.X = T;
        result.Y = sub(M.mul(S.sub(T)), yyyy.scalar_mul(8));
        result.Z = add(p.Y, p.Z).square().sub(yy).sub(zz);
    }

    uint256 internal constant bTWIST0 = 19485874751759354771024239261021720505790618469301721065564631296452457478373;
    uint256 internal constant bTWIST1 = 266929791119991161246907387137283842545076965332900288569378510910307636690;

    // ψ o π o ψ⁻¹, where ψ:E → E' is the degree 6 isomorphism defined over 𝔽p¹²
    uint256 internal constant ENDO_X0 = 21575463638280843010398324269430826099269044274347216827212613867836435027261;
    uint256 internal constant ENDO_X1 = 10307601595873709700152284273816112264069230130616436755625194854815875713954;
    uint256 internal constant ENDO_Y0 = 2821565182194536844548159561693502659359617185244120367078079554186484126554;
    uint256 internal constant ENDO_Y1 = 3505843767911556378687030309984248845540243509899259641013678093033130930403;

    // Seed x₀ of the curve
    uint256 internal constant X_SEED = 4965661367192848881;

    /**
     * psi implements the psi function
     *
     * ψ(p) = u o π o u⁻¹ where u: E'→E isomorphism from the twist to E
     */
    function psi(G2Jacobian memory point) internal pure returns (G2Jacobian memory result) {
        // X = conjugate(x) * endo.X
        result.X = mul([point.X[0], P - point.X[1]], [ENDO_X0, ENDO_X1]);
        // Y = conjugate(y) * endo.Y
        result.Y = mul([point.Y[0], P - point.Y[1]], [ENDO_Y0, ENDO_Y1]);
        // Z = conjugate(Z)
        result.Z[1] = P - point.Z[1];
        result.Z[0] = point.Z[0];
    }

    /**
     * @notice Checks if a G2Affine point is on 𝔽p² curve
     *
     * Satisfy equation: Y² = X³ + 3/(u+9) (D-type twist, d=6)
     */
    function isOnCurve(G2Point memory p) internal pure returns (bool) {
        uint256[2] memory lhs = square(p.Y);
        uint256[2] memory rhs = square(p.X).mul(p.X).add([bTWIST0, bTWIST1]);
        return sub(lhs, rhs).isZero();
    }

    /**
     * @notice Checks if a G2Jacobian point is on 𝔽p² curve
     *
     * Satisfy equation: Y² = X³ + Z^6 * 3/(u+9) (D-type twist, d=6)
     */
    function isOnCurve(G2Jacobian memory p) internal pure returns (bool) {
        uint256[2] memory lhs = square(p.Y);
        uint256[2] memory z6 = (p.Z.square().mul(p.Z)).square();
        uint256[2] memory rhs = square(p.X).mul(p.X).add(z6.mul([bTWIST0, bTWIST1]));
        return sub(lhs, rhs).isZero();
    }

    /**
     * _clearCofactor maps a point on 𝔽p² curve to the G2 subgroup (r-torsion)
     *
     * @notice implements https://github.com/Consensys/gnark-crypto/tree/564b6f724c3beac52d805e6e600d0a1fda9770b5/ecc/bn254/g2.go#L558
     * cf http://cacr.uwaterloo.ca/techreports/2011/cacr2011-26.pdf, 6.1
     */
    function clearCofactor(G2Jacobian memory p) internal pure returns (G2Jacobian memory g2Point) {
        // p0 = p * x
        // NOTE: expensive operation 62 doubles + ~20 adds
        G2Jacobian memory point0 = scalarMulG2(p, X_SEED);
        // p1 = ψ(3p * x)
        G2Jacobian memory point1 = psi(addG2(point0, doubleG2(point0)));
        // p2= ψ²(p * x)
        G2Jacobian memory point2 = psi(psi(point0));
        // p3= ψ3(p)
        G2Jacobian memory point3 = psi(psi(p));

        return addG2(point0, addG2(point1, addG2(point2, point3)));
    }

    /**
     * @notice Checks if a G2Jacobian point is in the correct G2 subgroup (generator specified in BN254)
     *
     * https://eprint.iacr.org/2022/348.pdf, sec. 3 and 5.1
     * [r]P == 0 <==> [x₀+1]P + ψ([x₀]P) + ψ²([x₀]P) = ψ³([2x₀]P)
     *
     * @dev https://github.com/Consensys/gnark-crypto/blob/564b6f724c3beac52d805e6e600d0a1fda9770b5/ecc/bn254/g2.go#L410
     */
    function isInSubGroup(G2Jacobian memory p) internal pure returns (bool) {
        G2Jacobian memory x0P = scalarMulG2(p, X_SEED);
        G2Jacobian memory psi_x0P = psi(x0P);
        G2Jacobian memory psi2_x0P = psi(psi_x0P);

        G2Jacobian memory lhs = (x0P.addG2(p)).addG2(psi_x0P).addG2(psi2_x0P);
        G2Jacobian memory rhs = doubleG2(psi(psi2_x0P));

        G2Jacobian memory result = rhs.subG2(lhs);

        return result.isOnCurve() && result.Z.isZero();
    }

    /**
     * @notice Checks if a G2Affine point is in the correct G2 subgroup (generator specified in BN254)
     */
    function isInSubGroup(G2Point memory p) internal pure returns (bool) {
        return p.toJacobian().isInSubGroup();
    }

    /**
     * Remind curve equation: Y² = X³ + 3/(u+9) (D-type twist, d=6)
     */
    function hashToG2(bytes32 digest) internal view returns (G2Point memory point) {
        uint256 x = uint256(digest) % P;

        while (true) {
            uint256[2] memory rhs = add([x, 0].square().mul([x, 0]), [bTWIST0, bTWIST1]);

            if (rhs.legendre() == 1) {
                point.X = [x, 0];
                point.Y = rhs.sqrt();
                point = point.toJacobian().clearCofactor().toAffine();
            }

            x += 1;
        }

        return G2Point([uint256(0), 0], [uint256(0), 0]);
    }

    ////////////////////////////////// END G2 operations /////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////////////////////

    //////////////////////////////////////////////////////////////////////////////////////////
    ////////////////////////////// BEGIN E2 arithmetics //////////////////////////////////////
    // E2 are elements of 𝔽p², represented as uint256[2]
    // WARNING: must use with elements in 𝔽p (e.g. 0<=val<P), other wise revert with overflow

    // SQRT_EXP_G2 = (P-3)/4, use for efficient sqrt algorithm
    uint256 internal constant SQRT_EXP_G2 = 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f51;
    // LEGENDRE_EXP = (P-1)/2
    uint256 internal constant LEGENDRE_EXP = 0x183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea3;

    function _submodP(uint256 a, uint256 b) private pure returns (uint256) {
        return addmod(a, P - b, P);
    }

    /**
     * @notice E2 multiplication
     */
    function mul(uint256[2] memory x, uint256[2] memory y) internal pure returns (uint256[2] memory) {
        return [_submodP(mulmod(x[0], y[0], P), mulmod(x[1], y[1], P)), addmod(mulmod(x[0], y[1], P), mulmod(x[1], y[0], P), P)];
    }

    /**
     * @notice E2 squaring
     */
    function square(uint256[2] memory x) internal pure returns (uint256[2] memory) {
        return [_submodP(mulmod(x[0], x[0], P), mulmod(x[1], x[1], P)), mulmod(mulmod(x[0], x[1], P), 2, P)];
    }

    /**
     * @notice E2 is zero if (0, 0)
     */
    function isZero(uint256[2] memory x) internal pure returns (bool) {
        return x[0] == 0 && x[1] == 0;
    }

    /**
     * @notice E2 is one if (1, 0)
     */
    function isOne(uint256[2] memory x) internal pure returns (bool) {
        return x[0] == 0 && x[1] == 0;
    }

    /**
     * @notice E2 scalar multiplication
     */
    function scalar_mul(uint256[2] memory x, uint256 k) internal pure returns (uint256[2] memory) {
        return [mulmod(x[0], k, P), mulmod(x[1], k, P)];
    }

    /**
     * @notice E2 addition
     */
    function add(uint256[2] memory x, uint256[2] memory y) internal pure returns (uint256[2] memory) {
        return [addmod(x[0], y[0], P), addmod(x[1], y[1], P)];
    }

    /**
     * @notice E2 substraction
     */
    function sub(uint256[2] memory x, uint256[2] memory y) internal pure returns (uint256[2] memory) {
        return [_submodP(x[0], y[0]), _submodP(x[1], y[1])];
    }

    function eq(uint256[2] memory x, uint256[2] memory y) internal pure returns (bool) {
        return sub(x, y).isZero();
    }

    /**
     * @notice E2 negation
     */
    function neg(uint256[2] memory x) internal pure returns (uint256[2] memory) {
        return [(P - x[0]) % P, (P - x[1]) % P];
    }

    function norm(uint256[2] memory x) internal pure returns (uint256) {
        return addmod(mulmod(x[0], x[0], P), mulmod(x[1], x[1], P), P);
    }

    // Algorithm 8 from https://eprint.iacr.org/2010/354.pdf
    function inverse(uint256[2] memory x) internal view returns (uint256[2] memory) {
        // NOTE: this can be optimized: https://github.com/pornin/bingcd/blob/main/doc/bingcd.pdf
        uint256 normInverse = Mod.modInverse(norm(x), P);
        return [mulmod(x[0], normInverse, P), P - mulmod(x[1], normInverse, P)];
    }

    function legendre(uint256[2] memory x) internal view returns (uint256) {
        return Mod.modExp(norm(x), LEGENDRE_EXP, P);
    }

    // NOTE: only takes positive exponent
    function exp(uint256[2] memory x, uint256 _exp) internal pure returns (uint256[2] memory result) {
        result = [uint256(1), 0];
        if (_exp == 0) {
            return result;
        }

        while (_exp != 0) {
            if ((_exp & 1) != 0) {
                result = result.square();
            }
            result = result.mul(x);
            _exp >>= 1;
        }
    }

    // cf https://eprint.iacr.org/2012/685.pdf (algo 9)
    // @dev doesn't check if x is quadratic residue
    function sqrt(uint256[2] memory x) internal pure returns (uint256[2] memory) {
        uint256[2] memory a1 = x.exp(SQRT_EXP_G2);
        uint256[2] memory alpha = a1.square().mul(x);
        uint256[2] memory x0 = a1.mul(x);

        // NOTE: skip legendre check
        if (alpha[0] == P - 1 && alpha[1] == P - 1) {
            // i * x0
            return [P - x0[1], x[0]];
        }

        // x0 * (1 + alpha)^LEGENDRE_EXP
        return mul(alpha.add([uint256(1), 0]).exp(LEGENDRE_EXP), x0);
    }

    ///////////////////////////////// END E2 arithmetics /////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////////////////////
}
