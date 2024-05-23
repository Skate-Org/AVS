package bn254

import (
	"crypto/rand"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

type PrivateKey = fr.Element

// NewPrivateKey create a scalar private key on Fr
//
// Use `(bn254/fr).Element.SetString`, octal or hexadecimal string must be prefixed with "0o" or "0x", respectively
func NewPrivateKey(stringKey string) (*PrivateKey, error) {
	privateKey, err := new(fr.Element).SetString(stringKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

type BLSKey struct {
	PrivKey *PrivateKey
	PubKey  *G2Point
}

func NewKeyPair(privKey *PrivateKey) *BLSKey {
	pubKey := MulByGeneratorG2(privKey) // pubkey = privateKey * G1
	return &BLSKey{
		PrivKey: privKey,
		PubKey:  &G2Point{pubKey},
	}
}

func GenRandomBlsKeys() (*BLSKey, error) {
	r, _ := new(big.Int).SetString(fr.Modulus().String(), 10)

	// Generate cryptographically strong pseudo-random element of group Fr
	n, err := rand.Int(rand.Reader, r)
	if err != nil {
		return nil, err
	}

	sk := new(PrivateKey).SetBigInt(n)
	return NewKeyPair(sk), nil
}

type Signature struct {
	*G1Point
}

func NewZeroSignature() *Signature {
	return &Signature{NewZeroG1Point()}
}

// Should rarely be used, use Jacobian instead
func (s *Signature) Add(otherSignature *Signature) *Signature {
	s.G1Point.Add(otherSignature.G1Point)
	return s
}

// AggregateSignatures aggregates multiple signatures
// NOTE: for G1 signature
func AggregateSignatures(sigs []*Signature) *Signature {
	aggSigJac := bn254.G1Jac{}
	for _, sig := range sigs {
		sigJac := new(bn254.G1Jac).FromAffine(sig.G1Affine)
		aggSigJac.AddAssign(sigJac)
	}
	return &Signature{
		&G1Point{
			new(bn254.G1Affine).FromJacobian(&aggSigJac),
		},
	}
}

// AggregateSignatures aggregates multiple publickey
// NOTE: for G2 pubkey
func AggregatePubKey(pubKeys []*G2Point) *G2Point {
	aggPubkeyJac := bn254.G2Jac{}
	for _, key := range pubKeys {
		keyJac := new(bn254.G2Jac).FromAffine(key.G2Affine)
		aggPubkeyJac.AddAssign(keyJac)
	}
	return &G2Point{
		new(bn254.G2Affine).FromJacobian(&aggPubkeyJac),
	}
}

// This signs a hashed message (e.g. keccak256(msg)) on G1,
//
// Require a G2Pubkey for verification
func (k *BLSKey) SignMessage(msgDigest [32]byte) *Signature {
	H := HashToG1(msgDigest)
	sig := new(bn254.G1Affine).ScalarMultiplication(H, k.PrivKey.BigInt(new(big.Int)))
	return &Signature{&G1Point{sig}}
}

// verifySignature internal methods to verify bls signature
//
// Check pairing function according to BLS scheme:
// e(pubKeyG1, Hash(message)G2) = e(g1, signatureG2)
func verifySignature(signature *bn254.G1Affine, pubkey *bn254.G2Affine, msgDigest [32]byte) (bool, error) {
	negateG2 := new(bn254.G2Affine).Neg(GetG2Generator())

	messageG1 := HashToG1(msgDigest)

	P := [2]bn254.G1Affine{*messageG1, *signature}
	Q := [2]bn254.G2Affine{*pubkey, *negateG2}

	ok, err := bn254.PairingCheck(P[:], Q[:])
	if err != nil {
		return false, err
	}
	return ok, nil
}

// Verify verify a signed bls signature against the message and public key
func (s *Signature) Verify(pubkey *G2Point, msgDigest [32]byte) (bool, error) {
	ok, err := verifySignature(s.G1Affine, pubkey.G2Affine, msgDigest)
	if err != nil {
		return false, err
	}
	return ok, nil
}
