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

    // Verify the aggregated BLS signature:
    //     e(Hash(message), pubKeyG2) = e(signature, g2)
    // <=> e(Hash(message), pubKeyG2) * e(signature, -g2) = 1
    function verifySinglePubKey(BN254.G1Point memory signature, BN254.G2Point memory pubKey, bytes32 message) internal view returns (bool) {
        BN254.G1Point memory messageG1 = BN254.hashToG1(message);
        bool valid = BN254.pairing(messageG1, pubKey, signature, BN254.negGeneratorG2());
        return valid;
    }

    // NOTE: on-chain point addition for G2 cost ~23k comparing to 37k gas using this pre-compiled.
    // => Fix BN254G2 to further optimize (later)
    function verifyBatchPubKey(BN254.G1Point memory aggSignature, BN254.G2Point[] memory pubKeys, bytes32 message) internal view returns (bool) {
        BN254.G1Point memory messageG1 = BN254.hashToG1(message);
        uint256 inputSize = 6 + 6 * pubKeys.length;

        uint256[] memory input = new uint256[](inputSize);
        input[0] = aggSignature.X;
        input[1] = aggSignature.Y;
        input[2] = BN254.G2x1;
        input[3] = BN254.G2x0;
        input[4] = BN254.nG2y1;
        input[5] = BN254.nG2y0;

        for (uint i = 1; i <= pubKeys.length; i++) {
            input[i * 6] = messageG1.X;
            input[i * 6 + 1] = messageG1.Y;
            input[i * 6 + 2] = pubKeys[i - 1].X[1];
            input[i * 6 + 3] = pubKeys[i - 1].X[0];
            input[i * 6 + 4] = pubKeys[i - 1].Y[1];
            input[i * 6 + 5] = pubKeys[i - 1].Y[0];
        }

        return BN254._ecPairing(input);
    }

    // NOTE: This function should only be used as an utility for testing
    function signMessage(uint256 blsPrivateKey, bytes32 message) internal view returns (BN254.G1Point memory signature) {
        if (blsPrivateKey >= BN254.R) {
            revert("BLS.signMessage: Key must in Fr group");
        }

        BN254.G1Point memory messageG1 = BN254.hashToG1(message);
        signature = messageG1.scalarMulG1(blsPrivateKey);
    }
}
