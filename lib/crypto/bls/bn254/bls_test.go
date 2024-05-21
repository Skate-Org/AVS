package bn254_test

import (
	"testing"

	"github.com/Skate-Org/AVS/lib/crypto/bls/bn254"
	libHash "github.com/Skate-Org/AVS/lib/crypto/hash"
	"github.com/stretchr/testify/assert"
)

func TestSignAndVerify(t *testing.T) {
	// Generate key pair
	keyPair, err := bn254.GenRandomBlsKeys()
	assert.Nil(t, err)

	// Message to sign
	message := "Hello, BLS!"
	hash := [32]byte(libHash.Keccak256([]byte(message)))

	// Sign the message
	signature := keyPair.SignMessage(hash)

	// Verify the signature
	pubKeyG2 := keyPair.GetPubKeyG2()
	valid, err := signature.VerifySignature(pubKeyG2, hash)
	assert.Nil(t, err)
	assert.True(t, valid, "The signature should be valid")

	// Tamper with the message and verify again
	tamperedMessage := "Hello, BLS! (tampered)"
	tamperedHash := [32]byte(libHash.Keccak256([]byte(tamperedMessage)))
	valid, err = signature.VerifySignature(pubKeyG2, tamperedHash)
	assert.Nil(t, err)
	assert.False(t, valid, "The signature should be invalid")
}

func TestAggregateSignatures(t *testing.T) {
	QUORUM_SIZE := 100
	// Message to sign
	message := "Aggregate BLS!"
	hash := [32]byte(libHash.Keccak256([]byte(message)))

	// Tamper with the message and verify again
	tamperedMessage := "Hello, BLS! (tampered)"
	tamperedHash := [32]byte(libHash.Keccak256([]byte(tamperedMessage)))

	sigs := make([]*bn254.Signature, QUORUM_SIZE)
	G2pubKeys := make([]*bn254.G2Point, QUORUM_SIZE)
	invalidSigs := make([]*bn254.Signature, QUORUM_SIZE)

	for i := 0; i < 100; i++ {
		key, err := bn254.GenRandomBlsKeys()
		assert.Nil(t, err)
		G2pubKeys[i] = key.GetPubKeyG2()

		sig := key.SignMessage(hash)
		sigs[i] = sig

		invalidSig := key.SignMessage(tamperedHash)
		invalidSigs[i] = invalidSig
	}

	aggSig := bn254.AggregateSignatures(sigs)
	aggPubKeys := bn254.AggregateG2PubKey(G2pubKeys)

	// NOTE: Should be able to verify QUORUM_SIZE of valid signatures
	valid, err := aggSig.VerifySignature(aggPubKeys, hash)
	assert.Nil(t, err)
	assert.True(t, valid, "The aggregated signature should be valid for the aggregated key")

	// NOTE: Should detects invalid signatures

	// Scenario 0: all signatures are invalid
	invalidAggSig0 := bn254.AggregateSignatures(invalidSigs)
	invalid0, err := invalidAggSig0.VerifySignature(aggPubKeys, hash)
	assert.Nil(t, err)
	assert.False(t, invalid0, "The aggregated signature should be invalid when all sigs are invalid")

	// Scenario 1: 1 signatures are invalid
	invalidSigsArray1 := append(sigs[:20], invalidSigs[30])
	invalidSigsArray1 = append(invalidSigsArray1, sigs[21:]...)
	invalidAggSig1 := bn254.AggregateSignatures(invalidSigsArray1)
	invalid1, err := invalidAggSig1.VerifySignature(aggPubKeys, hash)
	assert.Nil(t, err)
	assert.False(t, invalid1, "The aggregated signature should be invalid when all sigs are invalid")

	// Scenario 2: signature ordered mismatch
	invalidSigsArray2 := append(sigs[1:], sigs[0])
	invalidAggSig2 := bn254.AggregateSignatures(invalidSigsArray2)
	invalid2, err := invalidAggSig2.VerifySignature(aggPubKeys, hash)
	assert.Nil(t, err)
	assert.False(t, invalid2, "The aggregated signature should be invalid when all sigs are invalid")
}
