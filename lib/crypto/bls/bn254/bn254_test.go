package bn254_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/Skate-Org/AVS/lib/crypto/bls/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/stretchr/testify/assert"
)

func TestG1Point(t *testing.T) {
	x := big.NewInt(1)
	y := big.NewInt(2)

	// Test NewG1Point
	g1Point := bn254.NewG1Point(x, y)
	assert.NotNil(t, g1Point, "G1 point should not be nil")
	assert.Equal(t, x, g1Point.X.BigInt(new(big.Int)), "G1 point X coordinate mismatch")
	assert.Equal(t, y, g1Point.Y.BigInt(new(big.Int)), "G1 point Y coordinate mismatch")

	// Test Add
	zero0 := bn254.NewZeroG1Point()
	zero0.Add(g1Point)
	assert.Equal(t, x, zero0.X.BigInt(new(big.Int)), "G1 point X coordinate mismatch after addition")
	assert.Equal(t, y, zero0.Y.BigInt(new(big.Int)), "G1 point Y coordinate mismatch after addition")

	// Test Sub
	zero1 := bn254.NewZeroG1Point()

	g1Point.Sub(zero1)
	assert.Equal(t, x, g1Point.X.BigInt(new(big.Int)), "G1 point X coordinate mismatch after subtraction")
	assert.Equal(t, y, g1Point.Y.BigInt(new(big.Int)), "G1 point Y coordinate mismatch after subtraction")

	zero1.Sub(g1Point)
	negY := new(big.Int).Mod(new(big.Int).Neg(y), fp.Modulus())
	assert.Equal(t, x, zero1.X.BigInt(new(big.Int)), "G1 point X coordinate mismatch after subtraction")
	assert.Equal(t, negY, zero1.Y.BigInt(new(big.Int)), "G1 point Y coordinate mismatch after subtraction")

	// Test Serialize and Deserialize
	serialized := g1Point.Serialize()
	deserialized := bn254.DeserializeG1(serialized)
	assert.Equal(t, g1Point.X, deserialized.X, "G1 point X coordinate mismatch after deserialization")
	assert.Equal(t, g1Point.Y, deserialized.Y, "G1 point Y coordinate mismatch after deserialization")
}

func TestG2Point(t *testing.T) {
	x := [2]*big.Int{big.NewInt(1), big.NewInt(2)}
	y := [2]*big.Int{big.NewInt(3), big.NewInt(4)}

	// Test NewG2Point
	g2Point := bn254.NewG2Point(x, y)
	assert.NotNil(t, g2Point, "G2 point should not be nil")
	assert.Equal(t, x[1], g2Point.X.A0.BigInt(new(big.Int)), "G2 point X.A0 coordinate mismatch")
	assert.Equal(t, x[0], g2Point.X.A1.BigInt(new(big.Int)), "G2 point X.A1 coordinate mismatch")
	assert.Equal(t, y[1], g2Point.Y.A0.BigInt(new(big.Int)), "G2 point Y.A0 coordinate mismatch")
	assert.Equal(t, y[0], g2Point.Y.A1.BigInt(new(big.Int)), "G2 point Y.A1 coordinate mismatch")

	// Test Add
	zeroPoint := bn254.NewZeroG2Point()
	g2Point.Add(zeroPoint)
	assert.Equal(t, x[1], g2Point.X.A0.BigInt(new(big.Int)), "G2 point X.A0 coordinate mismatch after addition")
	assert.Equal(t, x[0], g2Point.X.A1.BigInt(new(big.Int)), "G2 point X.A1 coordinate mismatch after addition")
	assert.Equal(t, y[1], g2Point.Y.A0.BigInt(new(big.Int)), "G2 point Y.A0 coordinate mismatch after addition")
	assert.Equal(t, y[0], g2Point.Y.A1.BigInt(new(big.Int)), "G2 point Y.A1 coordinate mismatch after addition")

	// Test Sub
	g2Point.Sub(zeroPoint)
	assert.Equal(t, x[1], g2Point.X.A0.BigInt(new(big.Int)), "G2 point X.A0 coordinate mismatch after subtraction")
	assert.Equal(t, x[0], g2Point.X.A1.BigInt(new(big.Int)), "G2 point X.A1 coordinate mismatch after subtraction")
	assert.Equal(t, y[1], g2Point.Y.A0.BigInt(new(big.Int)), "G2 point Y.A0 coordinate mismatch after subtraction")
	assert.Equal(t, y[0], g2Point.Y.A1.BigInt(new(big.Int)), "G2 point Y.A1 coordinate mismatch after subtraction")

	// Test Serialize and Deserialize
	serialized := g2Point.Serialize()
	deserialized := bn254.DeserializeG2(serialized)
	assert.Equal(t, g2Point.X, deserialized.X, "G2 point X coordinate mismatch after deserialization")
	assert.Equal(t, g2Point.Y, deserialized.Y, "G2 point Y coordinate mismatch after deserialization")
}

func TestHashToG2(t *testing.T) {
	for i := 0; i < 10; i++ {
		var digest [32]byte
		rand.Read(digest[:])
		hashedPoint := bn254.HashToG2(digest)

		assert.NotNil(t, hashedPoint, "Hashed G2 point should not be nil")
		assert.True(t, hashedPoint.IsOnCurve(), "Hashed G2 point should be on 𝔽p² curve")
		assert.True(t, hashedPoint.IsInSubGroup(), "Hashed G2 point should be in the correct subgroup G2")
	}
}

func TestHashToG1(t *testing.T) {
	for i := 0; i < 10; i++ {
		var digest [32]byte
		rand.Read(digest[:])
		hashedPoint := bn254.HashToG1(digest)

		assert.NotNil(t, hashedPoint, "Hashed G1 point should not be nil")
		assert.True(t, hashedPoint.IsOnCurve(), "Hashed G1 point should be on 𝔽p curve")
	}
}

func TestG1G2Multiplication(t *testing.T) {
	// Test G1 multiplication by generator
	privKey := new(fr.Element).SetUint64(1234)
	g1Mul := bn254.MulByGeneratorG1(privKey)
	assert.NotNil(t, g1Mul, "G1 multiplication result should not be nil")
	expectedX1, _ := new(fp.Element).SetString("5240721337203810155063577104887775964429040310352786870634285698927658009894")
	expectedY1, _ := new(fp.Element).SetString("8895618777946819312582035270689922760507554319433213576472857911545059134563")
	assert.Equal(t, *expectedX1, g1Mul.X, "G1 X mismatch")
	assert.Equal(t, *expectedY1, g1Mul.Y, "G1 Y mismatch")

	// Test G2 multiplication by generator
	g2Mul := bn254.MulByGeneratorG2(privKey)
	assert.NotNil(t, g2Mul, "G2 multiplication result should not be nil")
	expectedXA0, _ := new(fp.Element).SetString("11092999225633600247987762624347164490249105779527674870198751554395839697955")
	expectedXA1, _ := new(fp.Element).SetString("20581924060851364827089112084266116502083385887431055789664064343317555539927")
	assert.Equal(t, *expectedXA0, g2Mul.X.A0, "G2 X.A0 mismatch")
	assert.Equal(t, *expectedXA1, g2Mul.X.A1, "G2 X.A1 mismatch")
	expectedYA0, _ := new(fp.Element).SetString("745165978236660430257472951680766366710880332450354657023528496579863883583")
	expectedYA1, _ := new(fp.Element).SetString("18968046579869378264211454225940124683613790899808490038348890384662573052093")
	assert.Equal(t, *expectedYA0, g2Mul.Y.A0, "G2 X.A0 mismatch")
	assert.Equal(t, *expectedYA1, g2Mul.Y.A1, "G2 X.A1 mismatch")
}
