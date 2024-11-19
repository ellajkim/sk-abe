package ske

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "errors"
    "io"
)

// Encrypt encrypts the plaintext using the key 'k'.
func Encrypt(k []byte, plaintext []byte) ([]byte, error) {
    // Derive a 32-byte key using SHA-256
    key := sha256.Sum256(k)

    // Create a new AES cipher using the derived key
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return nil, err
    }

    // Use Galois/Counter Mode (GCM) for AES encryption
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    // Generate a random nonce of the appropriate size
    nonce := make([]byte, aesGCM.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    // Encrypt the plaintext and prepend the nonce to the ciphertext
    ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}

// Decrypt decrypts the ciphertext using the key 'k'.
func Decrypt(k []byte, ciphertext []byte) ([]byte, error) {
    // Derive the same 32-byte key using SHA-256
    key := sha256.Sum256(k)

    // Create a new AES cipher using the derived key
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return nil, err
    }

    // Use Galois/Counter Mode (GCM) for AES decryption
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonceSize := aesGCM.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }

    // Extract the nonce from the beginning of the ciphertext
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    // Decrypt the ciphertext
    plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}