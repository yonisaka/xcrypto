package xcrypto

// CryptoStrategy is an interface for encryption and decryption
type CryptoStrategy interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
