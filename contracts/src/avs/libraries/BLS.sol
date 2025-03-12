// SPDX-License-Identifier: MIT

pragma solidity ^0.8.12;

import {BN254} from "./BN254.sol";

/**
 * Follow the standard defined in [IETF BLS draft v4](https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-bls-signature-04#section-2.1)
 *
 * Since G2 operations are not natively supported, point multiplication is prohibitively expensive.
 * We use publicKey on G2, and signature on G1, i.e. minimal-signature-size.
 *
 * Though [ETH2.0 specs](https://github.com/ethereum/consensus-specs/blob/f968d6245919c54f3135d1625d1c10c8a68588df/specs/phase0/beacon-chain.md#crypto)
 * use minimal-pubkey-size scheme, i.e. G1 for pubkey, G2 for signature.
 *
 * @author Skate Organization
 */
library BLS {
    using BN254 for BN254.G1Point;
    using BN254 for BN254.G2Point;
    using BN254 for BN254.G2Jacobian;

    // Verify the aggregated BLS signature:
    //     e(Hash(message), pubKeyG2) = e(signature, g2)
    // <=> e(Hash(message), pubKeyG2) * e(signature, -g2) = 1
    function verifySinglePubKey(
        BN254.G1Point memory signature,
        BN254.G2Point memory pubKey,
        bytes32 message
    ) internal view returns (bool) {
        BN254.G1Point memory messageG1 = BN254.hashToG1(message);
        bool valid = BN254.pairing(messageG1, pubKey, signature, BN254.negGeneratorG2());
        return valid;
    }

    // NOTE: This function is inefficient, DO NOT USE. For testing reference only.
    // on-chain point addition for G2 cost ~23.6k + 3k (Jacobian transform) comparing to 37k gas using this pre-compiled.
    // => using G2 addition will save 10k gas per operators + gas for payload size
    function verifyBatchPubKey(
        BN254.G1Point memory aggSignature,
        BN254.G2Point[] memory pubKeys,
        bytes32 message
    ) internal view returns (bool) {
        BN254.G1Point memory messageG1 = BN254.hashToG1(message);

        require(pubKeys.length > 0, "Invalid pubKey set");
        BN254.G2Jacobian memory aggregatedPubKey = pubKeys[0].toJacobian();

        for (uint256 i = 1; i < pubKeys.length; i++) {
            aggregatedPubKey = aggregatedPubKey.addG2(pubKeys[i].toJacobian());
        }

        bool valid = BN254.pairing(messageG1, aggregatedPubKey.toAffine(), aggSignature, BN254.negGeneratorG2());
        return valid;
    }

    // NOTE: This function should only be used as an utility for testing
    function signMessage(
        uint256 blsPrivateKey,
        bytes32 message
    ) internal view returns (BN254.G1Point memory signature) {
        if (blsPrivateKey >= BN254.R) {
            revert("Private Key must be within Fp group order R");
        }

        BN254.G1Point memory messageG1 = BN254.hashToG1(message);
        signature = messageG1.scalarMulG1(blsPrivateKey);
    }
}
