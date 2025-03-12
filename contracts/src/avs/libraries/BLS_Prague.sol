// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import {BLS12_381} from "./BLS12_381.sol";

/**
 * Follow the standard defined in [IETF BLS draft v4](https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-bls-signature-04#section-2.1)
 *
 * Curve BLS12-381 as per [EIP-2537](https://eips.ethereum.org/EIPS/eip-2537)
 *
 * NOTE: Use G1 for pubkey, G2 for signature.
 *
 * @author Skate Organization
 */
library BLS_Prague {
    using BLS12_381 for BLS12_381.G1Point;
    using BLS12_381 for BLS12_381.G2Point;

    /**
     * Verify a single BLS signature
     *     e(pubKeyG1, Hash(message)) = e(g1, signature)
     * <=> e(pubKeyG1, Hash(message)) * e(-g1, signature) = 1
     */
    function verifySinglePubKey(
        BLS12_381.G1Point memory pubKey,
        BLS12_381.G2Point memory signature,
        bytes memory message
    ) internal view returns (bool valid) {
        BLS12_381.G1Point[] memory g1s = new BLS12_381.G1Point[](2);
        g1s[0] = pubKey;
        g1s[1] = BLS12_381.negGeneratorG1();

        BLS12_381.G2Point[] memory g2s = new BLS12_381.G2Point[](2);
        g2s[0] = BLS12_381.hashMessageToG2(message);
        g2s[1] = signature;

        valid = BLS12_381.pairing(g1s, g2s);
    }

    /**
     * Verify aggregated single BLS signature:
     *        e(Σ_i(pubKeyG1_i), Hash(message)) = e(g1, aggSignature)
     * OR <=> ∏_i[e(pubKeyG1_i, Hash(message)] * e(-g1, aggSignature) = 1  [Method 1]
     * OR <=> e(Σ_i(pubKeyG1_i, Hash(message) * e(-g1, aggSignature) = 1   [Method 2]
     *
     * NOTE: Though it's much cheaper to aggregate the public key [method 1] than
     * verifying multiple pubKey pairs [method 2]. Cost analysis
     * 1/ Aggregate publicKey then ECPairing:
     *    a/ Pubkey aggregation cost: 375 * keySize
     *    b/ Pairing cost: 32,600 + 37,700
     * 2/ ECPairing batch:
     *    a/ Pairing cost: 32,600 * keySize + 37,700
     */
    function verifyBatchPubKey(
        BLS12_381.G1Point[] memory pubKeys,
        BLS12_381.G2Point memory aggSignature,
        bytes memory message
    ) internal view returns (bool valid) {
        require(pubKeys.length > 0, "Invalid pubKey set");

        BLS12_381.G1Point memory aggregatedPubKey = pubKeys[0];
        for (uint256 i = 1; i < pubKeys.length; i++) {
            aggregatedPubKey = aggregatedPubKey.addG1(pubKeys[i]);
        }
        BLS12_381.G1Point[] memory g1s = new BLS12_381.G1Point[](2);
        g1s[0] = aggregatedPubKey;
        g1s[1] = BLS12_381.negGeneratorG1();

        BLS12_381.G2Point[] memory g2s = new BLS12_381.G2Point[](2);
        g2s[0] = BLS12_381.hashMessageToG2(message);
        g2s[1] = aggSignature;

        valid = BLS12_381.pairing(g1s, g2s);
    }

    function signMessage(
        uint256 blsPrivateKey,
        bytes memory message
    ) internal view returns (BLS12_381.G2Point memory signature) {
        if (blsPrivateKey >= BLS12_381.R) {
            revert("Private Key must be within Fp group order R");
        }

        BLS12_381.G2Point memory messageG2 = BLS12_381.hashMessageToG2(message);
        signature = messageG2.scalarMulG2(blsPrivateKey);
    }

    function getPubKey(uint256 blsPrivateKey) internal view returns (BLS12_381.G1Point memory pubKey) {
        if (blsPrivateKey >= BLS12_381.R) {
            revert("Private Key must be within Fp group order R");
        }

        pubKey = BLS12_381.generatorG1().scalarMulG1(blsPrivateKey);
    }
}
