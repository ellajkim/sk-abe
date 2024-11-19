package abe

import (
	"math/big"
	"testing"
	"bytes"
)

func TestABEScheme(t *testing.T) {
	// Test parameters
	modulus := big.NewInt(1000000007) // Example prime modulus
	length := 4                        // Vector length
	msg := []byte("Hello, ABE!")      // Test message
	
	// Test vector x (attribute vector)
	x := make([]*big.Int, length)
	for i := 0; i < length; i++ {
		x[i] = big.NewInt(int64(i + 1))
	}

	// Test function f (policy vector)
	f := make([]*big.Int, length)
	for i := 0; i < length; i++ {
		f[i] = big.NewInt(int64(i + 1))
	}

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
	ct, err := ABEEnc(msk, x, msg)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decryption
	decrypted, err := ABEDec(skf, x, ct)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Verify decryption
	if !bytes.Equal(msg, decrypted) {
		t.Errorf("Decryption mismatch: got %v, want %v", decrypted, msg)
	}
}

func TestABESchemeInvalidInput(t *testing.T) {
	// Test parameters
	modulus := big.NewInt(1000000007)
	length := 4
	msg := []byte("Hello, ABE!")

	// Test with invalid length
	_, err := ABESetup(modulus, -1)
	if err == nil {
		t.Error("Setup should fail with negative length")
	}

	// Setup with valid parameters for subsequent tests
	msk, _ := ABESetup(modulus, length)

	// Test KeyGen with invalid function vector length
	invalidF := make([]*big.Int, length+1) // Wrong length
	_, err = ABEKeyGen(msk, invalidF)
	if err == nil {
		t.Error("KeyGen should fail with invalid function vector length")
	}

	// Test Encryption with invalid attribute vector length
	invalidX := make([]*big.Int, length+1) // Wrong length
	_, err = ABEEnc(msk, invalidX, msg)
	if err == nil {
		t.Error("Encryption should fail with invalid attribute vector length")
	}
} 