package ecdsa_test

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"skatechain.org/lib/crypto/ecdsa"
)

func TestECDSA(t *testing.T) {
	// Generate a random private key on K256 curve
	privateKey, err := ecdsa.KeyGen(ecdsa.S256(), rand.Reader)
	if err != nil {
		t.Fatal("Error generating private key:", err)
	}

	// Create a msgHash to sign
	message := []byte("Hello World")
	msgHash := ecdsa.Keccak256(message)

	// Sign the hash
	signature, err := ecdsa.Sign(msgHash[:], privateKey)
	if err != nil {
		t.Fatal("Error signing:", err)
	}

	// Verify the recovered address matches the original address
	originalAddress := ecdsa.PubkeyToAddress(privateKey.PublicKey)
	valid, err := ecdsa.Verify(originalAddress, msgHash[:], signature)
	if err != nil {
		t.Fatal("Error verifying signature:", err)
	}
	assert.True(t, valid, "Signature should be valid")

	// Recover public key from signature
	recoveredAddress, err := ecdsa.EcRecover(msgHash, signature)
	if err != nil {
		t.Fatal("Error recovering public key:", err)
	}
	assert.Equal(t, originalAddress, recoveredAddress, "Recovered address should match original address")
}

func TestKeccak256(t *testing.T) {
	bytesA := []byte("A")
	bytesB := []byte("B")
	hashA := ecdsa.Keccak256(bytesA, bytesB)
	hashB := ecdsa.Keccak256(append(bytesA, bytesB...))

	assert.Equal(t, hex.EncodeToString(hashA), hex.EncodeToString(hashB), "Expected: keccack256((A | B)) == keccak256(A, B)")

	dataDigest, _ := hex.DecodeString("2874d76350c1ff54ae960be64a")
	taskId := 40
	buf32 := make([]byte, 32) // taskId is uint256 in avs contract
	taskIdBytes := new(big.Int).SetUint64(uint64(taskId)).FillBytes(buf32)

	// equivalent to keccak256(abi.encodePacked(uint256, bytes))
	hash := hex.EncodeToString(ecdsa.Keccak256(taskIdBytes, dataDigest))
	expectedHash := "15191343bf37f6bff64be7d1bae5b940df32a0e2ba2dc2583a3127a72b948f6f"
	assert.Equal(t, hash, expectedHash, "Expected: return keccak256(abi.encodePacked(uint256, bytes))")

	// equivalent to keccak256(abi.encodePacked(uint256, bytes))
	ethHash := hex.EncodeToString(ecdsa.Keccak256Message(taskIdBytes, dataDigest))
	expectedEthHash := "3e942da6c4b6b2f60b7ff49540a59cd9fcf055e3385edf016596788984157f86"
	assert.Equal(t, ethHash, expectedEthHash, "Expected: return keccak256(abi.encodePacked(uint256, bytes)).toEthSignedMessageHash()")
}
