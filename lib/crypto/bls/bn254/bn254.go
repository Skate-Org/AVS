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

// NOTE: BN254 curve (same as ETH core implementation)
// seed x₀=4965661367192848881
// Frobenius Trace: t = 6x₀² + 1
// [r = p - t + 1]
// 𝔽r: r=21888242871839275222246405745257275088548364400416034343698204186575808495617 (36x₀⁴+36x₀³+18x₀²+6x₀+1)
// 𝔽p: p=21888242871839275222246405745257275088696311157297823662689037894645226208583 (36x₀⁴+36x₀³+24x₀²+6x₀+1)
// (E/𝔽p): Y²=X³+3
// (Eₜ/𝔽p²): Y² = X³+3/(u+9) (D-type twist)

////////////////////////////////////////////
////////////////// G1 //////////////////////

// Define on the Elliptic curve over field 𝔽p, p = 2^254 - 127
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

// Hash the message to a field element
func hashToField(msgDigest [32]byte) *big.Int {
	z := new(big.Int).SetBytes(msgDigest[:])
	z.Mod(z, fp.Modulus())
	return z
}

// HashToG1 implements the simple hash-and-check (also sometimes try-and-increment) algorithm
// Reference: https://hackmd.io/@benjaminion/bls12-381#Hashing-to-the-curve
func HashToG1(digest [32]byte) *bn254.G1Affine {
	x := hashToField(digest)

	one := new(big.Int).SetUint64(1)
	three := new(big.Int).SetUint64(3)
	for {
		// Curve: y^2 = x^3 + 3
		xCube := new(big.Int).Exp(x, three, fp.Modulus())
		y := new(big.Int).Add(xCube, three)
		y.Mod(y, fp.Modulus())

		//  true sqrt means on curve, returns the G1 point
		if y.ModSqrt(y, fp.Modulus()) != nil {
			var fpX, fpY fp.Element
			fpX.SetBigInt(x)
			fpY.SetBigInt(y)

			return &bn254.G1Affine{
				X: fpX,
				Y: fpY,
			}
		}

		x.Add(x, one) // already mod when cubing
	}
}

/////////////////////////////////////////////
////////////////// G2 //////////////////////

// Define on the Elliptic curve over extension field 𝔽p², p = 2^254 - 127
//
// A point on G2 is (X{A0, A1}, Y{A0, A1})
type G2Point struct {
	*bn254.G2Affine
}

func NewG2Point(X, Y [2]*big.Int) *G2Point {
	// NOTE: swapped for convention with ETH implementation (P_g2 = X0 * i + X1 = A0 + A1 * i)
	return &G2Point{
		&bn254.G2Affine{
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

var (
	// NOTE: G2 sextic twist parameters
	// y^2 = x^3 + b0 + b1 * i
	// |E'(𝔽p²)| = p^2+1+2p-t^2 = (p-t+1)*(p+t-1) = |E(𝔽p)| * (p+t-1) = r * (p+t-1) = |G2| * (p+t-1)
	//
	// the co-factor of G2: cG2 = p+t-1
	G2Cofactor, _ = new(big.Int).SetString("21888242871839275222246405745257275088844257914179612981679871602714643921549", 10)

	// Equivalent to 3/(9+i) in 𝔽p²
	G2b0, _ = new(big.Int).SetString("19485874751759354771024239261021720505790618469301721065564631296452457478373", 10)
	G2b1, _ = new(big.Int).SetString("266929791119991161246907387137283842545076965332900288569378510910307636690", 10)
)

// WARNING: Must be same with smart contracts implementation
//
// HashToG2 take a digest message and map to a G2 Point
func HashToG2(digest [32]byte) *bn254.G2Affine {
	// return hashToG2SVDW(digest)
	return hasToG2TryAndIncrement(digest)
}

// hashToG2SVDW use the built-in SVDW map, more efficient but maybe more gas intensive than try and increment
func hashToG2SVDW(digest [32]byte) *bn254.G2Affine {
	x := hashToField(digest)
	var u bn254.E2
	u.A0.SetBigInt(x)

	g2point := bn254.MapToG2(u)
	return &g2point
}

// hasToG2TryAndIncrement implements the simple hash-and-check (try-and-increment) algorithm
func hasToG2TryAndIncrement(digest [32]byte) *bn254.G2Affine {
	x := hashToField(digest)

	one := new(big.Int).SetUint64(1)
	three := new(big.Int).SetUint64(3)
	for {
		var y2E2 bn254.E2
		xCube := new(big.Int).Exp(x, three, fp.Modulus())
		y0 := new(big.Int).Add(xCube, G2b0)
		y0.Mod(y0, fp.Modulus())

		y2E2.A0.SetBigInt(y0)
		y2E2.A1.SetBigInt(G2b1)

		// found a point on 𝔽p², cofactor clearing to G2
		if y2E2.Legendre() == 1 {
			var xE2, yE2 bn254.E2
			xE2.A0.SetBigInt(x)
			yE2.Sqrt(&y2E2)

			// Multiply the curve2Point by the cofactor to ensure it is in the correct subgroup
			curve2Point := &bn254.G2Affine{
				X: xE2,
				Y: yE2,
			}

			// TODO: why does multiplying with co-factor of G2 not work?
			// g2Point := new(bn254.G2Affine).ScalarMultiplication(curve2Point, G2Cofactor)
			// if !new(bn254.G2Affine).ScalarMultiplication(g2Point, fr.Modulus()).IsInfinity() {
			//   panic("not in subgroup")
			// }
			// groupOrder := new(big.Int).Mul(fr.Modulus(), G2Cofactor)
			// log.Printf("\n\nE(𝔽p²) order: %v\n\n", groupOrder)
			// log.Printf("\n\nG2 cofactor: %v\n\n", G2Cofactor)
			// if !new(bn254.G2Affine).ScalarMultiplication(curve2Point, groupOrder).IsInfinity() {
			//   panic("not in subgroup")
			// }

			g2Point := curve2Point.ClearCofactor(curve2Point)

			return g2Point
		}

		x.Add(x, one)
	}
}
