package xcrypto

import (
	"crypto/rand"
	"crypto/rsa"
)

// rsaEncryptor is a struct
type rsaEncryptor struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewRSAEncryptor is a function
func NewRSAEncryptor(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) CryptoStrategy {
	return &rsaEncryptor{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// Encrypt is a function
func (e *rsaEncryptor) Encrypt(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, e.publicKey, data)
}

// Decrypt is a function
func (e *rsaEncryptor) Decrypt(data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, e.privateKey, data)
}
