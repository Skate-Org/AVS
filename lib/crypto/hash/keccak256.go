package hash

import (
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func Keccak256(data ...[]byte) []byte {
	return ethcrypto.Keccak256(data...)
}

func Keccak256Message(data ...[]byte) []byte {
	digest := ethcrypto.Keccak256(data...)
	prefix := []byte("\x19Ethereum Signed Message:\n32")
	return Keccak256(prefix, digest)
}
