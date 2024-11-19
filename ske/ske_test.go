package ske

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		key       []byte
		plaintext []byte
	}{
		{
			name:      "Basic encryption/decryption",
			key:       []byte("mysecretkey12345"),
			plaintext: []byte("Hello, World!"),
		},
		{
			name:      "Empty plaintext",
			key:       []byte("mysecretkey12345"),
			plaintext: []byte(""),
		},
		{
			name:      "Long plaintext",
			key:       []byte("mysecretkey12345"),
			plaintext: bytes.Repeat([]byte("Long message. "), 100),
		},
		{
			name:      "Binary data",
			key:       []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
			plaintext: []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := Encrypt(tt.key, tt.plaintext)
			if err != nil {
				t.Fatalf("Encrypt failed: %v", err)
			}

			// Verify ciphertext is different from plaintext
			if bytes.Equal(ciphertext, tt.plaintext) {
				t.Error("Ciphertext should be different from plaintext")
			}

			// Decrypt
			decrypted, err := Decrypt(tt.key, ciphertext)
			if err != nil {
				t.Fatalf("Decrypt failed: %v", err)
			}

			// Verify decrypted matches original
			if !bytes.Equal(decrypted, tt.plaintext) {
				t.Errorf("Decrypted text doesn't match original.\nGot: %v\nWant: %v",
					decrypted, tt.plaintext)
			}
		})
	}
}

func TestEncryptionErrors(t *testing.T) {
	
} 