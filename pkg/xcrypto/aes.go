package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type aesEncryptor struct {
	key []byte
}

func NewAESEncryptor(key []byte) CryptoStrategy {
	return &aesEncryptor{
		key: key,
	}
}

func (e *aesEncryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	// Create a new Galois Counter Mode (GCM) cipher using the AES block.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Generate a random nonce.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the data using the GCM cipher and the nonce.
	ciphertext := gcm.Seal(nil, nonce, data, nil)

	// Append the nonce to the ciphertext.
	ciphertext = append(nonce, ciphertext...)

	return ciphertext, nil
}

func (e *aesEncryptor) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	// Create a new Galois Counter Mode (GCM) cipher using the AES block.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Extract the nonce from the ciphertext.
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("invalid ciphertext")
	}
	nonce := data[:nonceSize]
	data = data[nonceSize:]

	// Decrypt the ciphertext using the GCM cipher and the nonce.
	plaintext, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
