package abe

import (
	"bytes"
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
