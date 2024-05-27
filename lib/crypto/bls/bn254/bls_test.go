package bn254_test

import (
	"log"
	"testing"

	"github.com/Skate-Org/AVS/lib/crypto/bls/bn254"
	libHash "github.com/Skate-Org/AVS/lib/crypto/hash"
	gnarkBn254 "github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	// 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 - anvil test wallet 0
	privKey, err := bn254.NewPrivateKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	// 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 - anvil test wallet 1
	// privKey, err := bn254.NewPrivateKey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	// 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC - anvil test wallet 1
	// privKey, err := bn254.NewPrivateKey("0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a")
	assert.Nil(t, err)
	// log.Printf("\n\nPrivate key: %v\n\n", privKey)

	keyPair := bn254.NewKeyPair(privKey)
	log.Printf("\nExample pubkey G2: \nx=%v\ny=%v\n\n", keyPair.PubKey.X.String(), keyPair.PubKey.Y.String())
	z2 := bn254.NewZeroG2Point()
	log.Printf("\nG2 x inverse: %v\n\n", z2.X.Inverse(&keyPair.PubKey.X))
	// pubKeyG1 := bn254.MulByGeneratorG1(privKey)
	// log.Printf("\n\nPublic key G1: %v\n\n", pubKeyG1)

	// // Message to sign
	// message := "BLS_signature"
	// msgHash := [32]byte(libHash.Keccak256([]byte(message)))
	//
	// // Sign the message
	// signature := keyPair.SignMessage(msgHash)
	//
	// log.Printf("\n\nSignature: %v\n\n", signature)
	//
	// g1 := bn254.GetG1Generator()
	// result_g1 := g1.Add(g1, pubKeyG1)
	// log.Printf("\n\nAdd g1 result: %v\n\n", result_g1.String())
	//
	//
	g2Jac := new(gnarkBn254.G2Jac).FromAffine(bn254.GetG2Generator())
	pubKeyG2Jac := new(gnarkBn254.G2Jac).FromAffine(keyPair.PubKey.G2Affine)
	resultG2Jac := g2Jac.AddAssign(pubKeyG2Jac)
	log.Printf(
		"\n(g2+pubkeyG2) Jacobian result [for contracts testing]: \nx=%v\ny=%v\nz=%v\n\n",
		resultG2Jac.X.String(), resultG2Jac.Y.String(), resultG2Jac.Z.String(),
	)

	resultG2 := bn254.NewZeroG2Point().FromJacobian(resultG2Jac)
	log.Printf("\n(g2+pubkeyG2) result: \nx=%v\ny=%v\n\n", resultG2.X.String(), resultG2.Y.String())
}

func TestSingleSignature(t *testing.T) {
	// Generate key pair
	keyPair, err := bn254.GenRandomBlsKeys()
	assert.Nil(t, err)

	// Message to sign
	message := "Hello, BLS!"
	msgHash := [32]byte(libHash.Keccak256([]byte(message)))

	// Sign the message
	signature := keyPair.SignMessage(msgHash)

	// Verify the signature
	pubkeyG1 := keyPair.PubKey
	valid, err := signature.Verify(pubkeyG1, msgHash)
	assert.Nil(t, err)
	assert.True(t, valid, "The signature should be valid")

	// Tamper with the message and verify again
	tamperedMessage := "Hello, BLS! (tampered)"
	tamperedHash := [32]byte(libHash.Keccak256([]byte(tamperedMessage)))
	valid, err = signature.Verify(pubkeyG1, tamperedHash)
	assert.Nil(t, err)
	assert.False(t, valid, "The signature should be invalid")
}

func TestAggregatedSignature(t *testing.T) {
	QUORUM_SIZE := 10
	// Message to sign
	message := "Aggregate BLS!"
	msgHash := [32]byte(libHash.Keccak256([]byte(message)))

	// Tamper with the message and verify again
	tamperedMessage := "Hello, BLS! (tampered)"
	tamperedHash := [32]byte(libHash.Keccak256([]byte(tamperedMessage)))

	sigs := make([]*bn254.Signature, QUORUM_SIZE)
	pubKeys := make([]*bn254.G2Point, QUORUM_SIZE)
	invalidSigs := make([]*bn254.Signature, QUORUM_SIZE)

	for i := 0; i < QUORUM_SIZE; i++ {
		key, err := bn254.GenRandomBlsKeys()
		assert.Nil(t, err)
		pubKeys[i] = key.PubKey

		sig := key.SignMessage(msgHash)
		sigs[i] = sig

		invalidSig := key.SignMessage(tamperedHash)
		invalidSigs[i] = invalidSig
	}

	aggSig := bn254.AggregateSignatures(sigs)
	aggPubKeys := bn254.AggregatePubKey(pubKeys)

	// NOTE: Should be able to verify QUORUM_SIZE of valid signatures

	// Scenario 0: sigs are normal
	valid, err := aggSig.Verify(aggPubKeys, msgHash)
	assert.Nil(t, err)
	assert.True(t, valid, "The aggregated signature should be valid for the aggregated key")

	// Scenario 1: sigs order is mixed
	reorderedSigs := append(sigs[1:], sigs[0])
	aggSig2 := bn254.AggregateSignatures(reorderedSigs)
	valid2, err := aggSig2.Verify(aggPubKeys, msgHash)
	assert.Nil(t, err)
	assert.True(t, valid2, "The aggregated signature should be valid when sigs are re-ordered")

	// NOTE: Should detects invalid signatures

	// Scenario 0: all signatures are invalid
	invalidAggSig0 := bn254.AggregateSignatures(invalidSigs)
	invalid0, err := invalidAggSig0.Verify(aggPubKeys, msgHash)
	assert.Nil(t, err)
	assert.False(t, invalid0, "The aggregated signature should be invalid when all sigs are invalid")

	// Scenario 1: 1 signatures is invalid
	invalidSigsArray1 := append(sigs[:QUORUM_SIZE-1], invalidSigs[0])
	invalidAggSig1 := bn254.AggregateSignatures(invalidSigsArray1)
	invalid1, err := invalidAggSig1.Verify(aggPubKeys, msgHash)
	assert.Nil(t, err)
	assert.False(t, invalid1, "The aggregated signature should be invalid when 1 sig is invalid")
}
