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

type KeyPair struct {
	PrivKey *PrivateKey
	PubKey  *G1Point
}

func NewKeyPair(privKey *PrivateKey) *KeyPair {
	pubKey := MulByGeneratorG1(privKey) // pubkey = privateKey * G1
	return &KeyPair{privKey, &G1Point{pubKey}}
}

// IsImage use `bn254.PairingCheck` to check if:
//
// e(P1,g1) = e(g2,P2), i.e. pubkeyG1 <-> pubkeyG2
func (p1 *G1Point) IsImage(p2 *G2Point) (bool, error) {
	return bn254.PairingCheck(
		[]bn254.G1Affine{*p1.G1Affine, *new(bn254.G1Affine).Neg(GetG1Generator())},
		[]bn254.G2Affine{*GetG2Generator(), *p2.G2Affine},
	)
}

func GenRandomBlsKeys() (*KeyPair, error) {
	r, _ := new(big.Int).SetString(fr.Modulus().String(), 10)

	// Generate cryptographically strong pseudo-random between 0 - r (order of Fr)
	n, err := rand.Int(rand.Reader, r)
	if err != nil {
		return nil, err
	}

	sk := new(PrivateKey).SetBigInt(n)
	return NewKeyPair(sk), nil
}

type Signature struct {
	*G1Point `json:"g1_point"`
}

func NewZeroSignature() *Signature {
	return &Signature{NewZeroG1Point()}
}

func (s *Signature) Add(otherSignature *Signature) *Signature {
	s.G1Point.Add(otherSignature.G1Point)
	return s
}

// AggregateSignatures aggregates multiple signatures into one
func AggregateSignatures(sigs []*Signature) *Signature {
	aggSigs := NewZeroSignature()
	for _, sig := range sigs {
		aggSigs.Add(sig)
	}
	return aggSigs
}

// AggregateSignatures aggregates multiple signatures into one
func AggregateG2PubKey(G2pubKeys []*G2Point) *G2Point {
	aggG2PubKeys := NewZeroG2Point()
	for _, key := range G2pubKeys {
		aggG2PubKeys.Add(key)
	}
	return aggG2PubKeys
}

// This signs a hashed message (e.g. keccak256(msg)) on G1,
//
// Require a G2Pubkey for verification
func (k *KeyPair) SignMessage(digestHash [32]byte) *Signature {
	H := HashToG1(digestHash)
	sig := new(bn254.G1Affine).ScalarMultiplication(H, k.PrivKey.BigInt(new(big.Int)))
	return &Signature{&G1Point{sig}}
}

func (k *KeyPair) GetPubKeyG2() *G2Point {
	return &G2Point{MulByGeneratorG2(k.PrivKey)}
}

func (k *KeyPair) GetPubKeyG1() *G1Point {
	return k.PubKey
}

// verifySig internal methods to verify bls signature
//
// Check pairing function according to BLS scheme:
// e(Hash(message), signature) = e(pubkeyG2, g2)
func verifySig(sig *bn254.G1Affine, pubkey *bn254.G2Affine, msgBytes [32]byte) (bool, error) {
	g2Gen := GetG2Generator()

	msgPoint := HashToG1(msgBytes)

	var negSig bn254.G1Affine
	negSig.Neg((*bn254.G1Affine)(sig))

	P := [2]bn254.G1Affine{*msgPoint, negSig}
	Q := [2]bn254.G2Affine{*pubkey, *g2Gen}

	ok, err := bn254.PairingCheck(P[:], Q[:])
	if err != nil {
		return false, nil
	}
	return ok, nil
}

// VerifySignature verify a signed bls signature against the message and public key
func (s *Signature) VerifySignature(pubkey *G2Point, message [32]byte) (bool, error) {
	ok, err := verifySig(s.G1Affine, pubkey.G2Affine, message)
	if err != nil {
		return false, err
	}
	return ok, nil
}
