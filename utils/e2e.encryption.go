package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

// EncryptMessage encrypts a message using AES-GCM
func EncryptMessage(key []byte, plaintext string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Error creating new GCM: %v", err)
		return nil, err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Printf("Error reading random bytes: %v", err)
		return nil, err
	}

	ciphertext := aesGcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return ciphertext, nil
}

// DecryptMessage decrypts a message using AES-GCM
func DecryptMessage(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Error creating new GCM: %v", err)
		return nil, err
	}

	nonceSize := aesGcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Printf("Error: ciphertext too short")
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Printf("Error decrypting message: %v", err)
		return nil, err
	}

	return plaintext, nil
}
