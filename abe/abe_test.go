package abe

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"
)

func TestABEScheme(t *testing.T) {
	// Test parameters
	modulus := big.NewInt(1009) // Small prime for testing
	length := 5
	message := []byte("Hello, ABE!")

	// Test vectors
	x := make([]*big.Int, length)
	f := make([]*big.Int, length)

	// Set up vectors where <f,x> = 0 mod modulus
	// Example: f = [1,1,1,1,1], x = [1,2,3,4,-10]
	// Inner product: 1+2+3+4-10 = 0
	for i := 0; i < length-1; i++ {
		x[i] = big.NewInt(int64(i + 1))
		f[i] = big.NewInt(1)
	}
	// Make the last element balance the sum
	sum := big.NewInt(0)
	for i := 0; i < length-1; i++ {
		tmp := new(big.Int).Mul(x[i], f[i])
		sum.Add(sum, tmp)
	}
	x[length-1] = big.NewInt(-10) // or any value that makes sum zero
	f[length-1] = big.NewInt(1)

	// Add debug prints
	t.Logf("x values: %v", x)
	t.Logf("f values: %v", f)

	// Verify inner product is 0 mod modulus
	innerProduct := big.NewInt(0)
	for i := 0; i < length; i++ {
		tmp := new(big.Int).Mul(x[i], f[i])
		innerProduct.Add(innerProduct, tmp)
	}
	innerProduct.Mod(innerProduct, modulus)
	t.Logf("Inner product mod %v: %v", modulus, innerProduct)

	// Setup
	msk, err := ABESetup(modulus, length)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Key Generation
	skf, err := ABEKeyGen(msk, f)
	if err != nil {
		t.Fatalf("KeyGen failed: %v", err)
	}

	// Encryption
	ct, err := ABEEnc(msk, x, message)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decryption
	decrypted, err := ABEDec(skf, x, ct)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Check decryption result
	if !bytes.Equal(decrypted, message) {
		t.Fatalf("Decryption result does not match the original message")
	}
}

func BenchmarkABEScheme(b *testing.B) {
	// Test sizes
	sizes := []int{5, 10, 20, 50, 100, 1000, 10000, 100000}

	for _, length := range sizes {
		b.Run(fmt.Sprintf("Vector_Size_%d", length), func(b *testing.B) {
			// Setup test parameters
			modulus, _ := big.NewInt(0).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF61", 16)
			message := []byte("Hello, ABE!")

			// Initialize vectors
			x := make([]*big.Int, length)
			f := make([]*big.Int, length)

			// Fill vectors (similar pattern as test, but simplified)
			for i := 0; i < length; i++ {
				x[i] = big.NewInt(int64(i + 1))
				f[i] = big.NewInt(1)
			}
			// Adjust last element to ensure inner product is 0
			x[length-1] = big.NewInt(int64(-((length * (length - 1)) / 2)))

			// Setup phase (done outside the benchmark loop)
			msk, _ := ABESetup(modulus, length)
			skf, _ := ABEKeyGen(msk, f)
			ct, _ := ABEEnc(msk, x, message)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Benchmark the decryption operation
				_, _ = ABEDec(skf, x, ct)
			}
		})
	}
}
