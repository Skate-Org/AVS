package ecdsa

import (
	"crypto/ecdsa"
	// "encoding/hex"
	"io"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type (
	PrivateKey = ecdsa.PrivateKey
	PublicKey  = ecdsa.PublicKey
)

// Sign returns a signature from input data.
//
// Follow EVM standard: [R (32b) || S (32b) || V (1b)] and (V={27, 28}).
// NOTE: digestHash must be 32 bytes
func Sign(digestHash []byte, key *PrivateKey) ([65]byte, error) {
	sig, err := ethcrypto.Sign(digestHash[:], key)
	if err != nil {
		return [65]byte(sig), err
	}
	// NOTE: COMPATIBILITY REQUIREMENTS: Adjust V from secp256k1 0/1 to Ethereum 27/28
	sig[64] += 27
	return [65]byte(sig), nil
}

// Recover the public key that sign a given hash
//
// NOTE: Equivalent to Openzeppelin's ECDSA.recover function:
// https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/cryptography/ECDSA.sol
//
// NOTE: digestHash must be 32 bytes
func EcRecover(digestHash []byte, sig [65]byte) (common.Address, error) {
	if v := sig[64]; v != 27 && v != 28 {
		return common.Address{}, errors.New("invalid recovery id (V) format, must be 27 or 28")
	}
	// NOTE: COMPATIBILITY REQUIREMENTS: Adjust V from Ethereum 27/28 to secp256k1 0/1
	sig[64] -= 27

	pubkey, err := ethcrypto.SigToPub(digestHash[:], sig[:])
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Recover public key")
	}

	return ethcrypto.PubkeyToAddress(*pubkey), nil
}

// Verify ethereum signed message
// NOTE: digestHash must be 32 bytes
func Verify(address common.Address, digestHash []byte, sig [65]byte) (bool, error) {
	actual, err := EcRecover(digestHash, sig)

	return actual == address, err
}

func PubkeyToAddress(publicKey PublicKey) common.Address {
	return ethcrypto.PubkeyToAddress(publicKey)
}

func Keccak256(data ...[]byte) []byte {
	return ethcrypto.Keccak256(data...)
}

func Keccak256Message(data ...[]byte) []byte {
	digest := ethcrypto.Keccak256(data...)
	prefix := []byte("\x19Ethereum Signed Message:\n32")
	return Keccak256(prefix, digest)
}

func S256() *secp256k1.BitCurve {
	return secp256k1.S256()
}

func FromECDSAPub(pubKey *ecdsa.PublicKey) []byte {
	return ethcrypto.FromECDSAPub(pubKey)
}

func KeyGen(curve *secp256k1.BitCurve, rand io.Reader) (*PrivateKey, error) {
	return ecdsa.GenerateKey(curve, rand)
}
