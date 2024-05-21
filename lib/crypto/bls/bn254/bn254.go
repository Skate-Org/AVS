package bn254

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

func newFpElement(x *big.Int) fp.Element {
	var p fp.Element
	p.SetBigInt(x)
	return p
}

////////////////////////////////////////////
////////////////// G1 //////////////////////

// Define on the Elliptic curve over field Fp, p = 2^254 - 127
//
// A point on G1 is (X, Y)
type G1Point struct {
	*bn254.G1Affine
}

func NewG1Point(x, y *big.Int) *G1Point {
	return &G1Point{
		&bn254.G1Affine{
			X: newFpElement(x),
			Y: newFpElement(y),
		},
	}
}

func NewZeroG1Point() *G1Point {
	return NewG1Point(big.NewInt(0), big.NewInt(0))
}

// Add another G1 point to this one
func (p *G1Point) Add(p2 *G1Point) *G1Point {
	p.G1Affine.Add(p.G1Affine, p2.G1Affine)
	return p
}

// Sub another G1 point from this one
func (p *G1Point) Sub(p2 *G1Point) *G1Point {
	p.G1Affine.Sub(p.G1Affine, p2.G1Affine)
	return p
}

func SerializeG1(p *bn254.G1Affine) []byte {
	b := make([]byte, 0)

	BytesX := p.X.Bytes()
	b = append(b, BytesX[:]...)
	BytesY := p.Y.Bytes()
	b = append(b, BytesY[:]...)

	return b
}

func (p *G1Point) Serialize() []byte {
	return SerializeG1(p.G1Affine)
}

func DeserializeG1(b []byte) *bn254.G1Affine {
	p := new(bn254.G1Affine)
	p.X.SetBytes(b[0:32])
	p.Y.SetBytes(b[32:64])
	return p
}

// HashToG1 implements the simple hash-and-check (also sometimes try-and-increment) algorithm
// see https://hackmd.io/@benjaminion/bls12-381#Hash-and-check
// Note that this function needs to be the same as the one used in the contract:
// https://github.com/Layr-Labs/eigenlayer-middleware/blob/1feb6ae7e12f33ce8eefb361edb69ee26c118b5d/src/libraries/BN254.sol#L292
// we don't use the newer constant time hash-to-curve algorithms as they are gas-expensive to compute onchain
func HashToG1(digest [32]byte) *bn254.G1Affine {
	one := new(big.Int).SetUint64(1)
	three := new(big.Int).SetUint64(3)
	x := new(big.Int)
	x.SetBytes(digest[:])
	for {
		// y = x^3 + 3
		xP3 := new(big.Int).Exp(x, big.NewInt(3), fp.Modulus())
		y := new(big.Int).Add(xP3, three)
		y.Mod(y, fp.Modulus())

		if y.ModSqrt(y, fp.Modulus()) == nil {
			x.Add(x, one).Mod(x, fp.Modulus())
		} else {
			var fpX, fpY fp.Element
			fpX.SetBigInt(x)
			fpY.SetBigInt(y)
			return &bn254.G1Affine{
				X: fpX,
				Y: fpY,
			}
		}
	}
}

func GetG1Generator() *bn254.G1Affine {
	g1Gen := new(bn254.G1Affine)
	_, err := g1Gen.X.SetString("1")
	if err != nil {
		return nil
	}
	_, err = g1Gen.Y.SetString("2")
	if err != nil {
		return nil
	}
	return g1Gen
}

func MulByGeneratorG1(a *fr.Element) *bn254.G1Affine {
	g1Gen := GetG1Generator()
	return new(bn254.G1Affine).ScalarMultiplication(g1Gen, a.BigInt(new(big.Int)))
}

/////////////////////////////////////////////
////////////////// G2 //////////////////////

// Define on the Elliptic curve over extension field Fp^2, p = 2^254 - 127
//
// A point on G2 is (X{A0, A1}, Y{A0, A1})
type G2Point struct {
	*bn254.G2Affine
}

func NewG2Point(X, Y [2]*big.Int) *G2Point {
	return &G2Point{
		&bn254.G2Affine{
			// TODO: why are 1 and 0 swapped here?
			X: struct{ A0, A1 fp.Element }{
				A0: newFpElement(X[1]),
				A1: newFpElement(X[0]),
			},
			Y: struct{ A0, A1 fp.Element }{
				A0: newFpElement(Y[1]),
				A1: newFpElement(Y[0]),
			},
		},
	}
}

// Add another G2 point to this one
func (p *G2Point) Add(p2 *G2Point) *G2Point {
	p.G2Affine.Add(p.G2Affine, p2.G2Affine)
	return p
}

// Sub another G2 point from this one
func (p *G2Point) Sub(p2 *G2Point) *G2Point {
	p.G2Affine.Sub(p.G2Affine, p2.G2Affine)
	return p
}

func NewZeroG2Point() *G2Point {
	return NewG2Point([2]*big.Int{big.NewInt(0), big.NewInt(0)}, [2]*big.Int{big.NewInt(0), big.NewInt(0)})
}

func SerializeG2(p *bn254.G2Affine) []byte {
	b := make([]byte, 0)

	BytesX0 := p.X.A0.Bytes()
	b = append(b, BytesX0[:]...)
	BytesX1 := p.X.A1.Bytes()
	b = append(b, BytesX1[:]...)

	BytesY0 := p.Y.A0.Bytes()
	b = append(b, BytesY0[:]...)
	BytesY1 := p.Y.A1.Bytes()
	b = append(b, BytesY1[:]...)

	return b
}

func DeserializeG2(b []byte) *bn254.G2Affine {
	p := new(bn254.G2Affine)
	p.X.A0.SetBytes(b[0:32])
	p.X.A1.SetBytes(b[32:64])
	p.Y.A0.SetBytes(b[64:96])
	p.Y.A1.SetBytes(b[96:128])
	return p
}

func (p *G2Point) Serialize() []byte {
	return SerializeG2(p.G2Affine)
}

func GetG2Generator() *bn254.G2Affine {
	g2Gen := new(bn254.G2Affine)
	g2Gen.X.SetString(
		"10857046999023057135944570762232829481370756359578518086990519993285655852781",
		"11559732032986387107991004021392285783925812861821192530917403151452391805634",
	)
	g2Gen.Y.SetString(
		"8495653923123431417604973247489272438418190587263600148770280649306958101930",
		"4082367875863433681332203403145435568316851327593401208105741076214120093531",
	)
	return g2Gen
}

func MulByGeneratorG2(a *fr.Element) *bn254.G2Affine {
	g2Gen := GetG2Generator()
	return new(bn254.G2Affine).ScalarMultiplication(g2Gen, a.BigInt(new(big.Int)))
}
