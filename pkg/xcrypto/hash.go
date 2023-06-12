package xcrypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func Hash(plainText string) (string, string, error) {
	// Generate a random salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}

	textBytes := []byte(plainText)

	textWithSalt := append(salt, textBytes...)

	hash := sha256.New()

	hash.Write(textWithSalt)

	hashedTextBytes := hash.Sum(nil)

	saltString := hex.EncodeToString(salt)
	hashedText := hex.EncodeToString(hashedTextBytes)

	return hashedText, saltString, nil
}

// CompareHash compares a hashed text with a plain text
func CompareHash(hashedText, salt, plainText string) (bool, error) {
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return false, err
	}

	textBytes := []byte(plainText)

	textWithSalt := append(saltBytes, textBytes...)

	hash := sha256.New()

	hash.Write(textWithSalt)

	computedHashedTextBytes := hash.Sum(nil)

	computedHashedText := hex.EncodeToString(computedHashedTextBytes)

	// Compare the computed hashed text with the stored hashed text
	return hashedText == computedHashedText, nil
}
