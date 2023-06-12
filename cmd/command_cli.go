package cmd

import (
	"errors"
	"github.com/urfave/cli/v2"
	"github.com/yonisaka/xcrypto/pkg/xcrypto"
	"log"
)

// newGenerateKeyPair is a method
func (cmd *Command) newGenerateKeyPair() *cli.Command {
	return &cli.Command{
		Name:  "key:generate",
		Usage: "A command to run generate key pair",
		Action: func(c *cli.Context) error {
			privateKey, publicKey, err := xcrypto.GenerateKeyPair()
			if err != nil {
				return err
			}

			err = xcrypto.SavePrivateKeyToFile(privateKey, cmd.RSA.PrivateKeyPath)
			if err != nil {
				return err
			}

			err = xcrypto.SavePublicKeyToFile(publicKey, cmd.RSA.PublicKeyPath)
			if err != nil {
				return err
			}

			log.Print("generate key pair success")
			return nil
		},
	}
}

func (cmd *Command) newRSAEncryptor() *cli.Command {
	return &cli.Command{
		Name:  "rsa:encrypt",
		Usage: "A command to run rsa encryptor",
		Action: func(c *cli.Context) error {
			privateKey, err := xcrypto.LoadPrivateKeyFromFile(cmd.RSA.PrivateKeyPath)
			if err != nil {
				return err
			}

			publicKey, err := xcrypto.LoadPublicKeyFromFile(cmd.RSA.PublicKeyPath)
			if err != nil {
				return err
			}

			args := c.Args().First()
			if args == "" {
				return errors.New("please provide data to encrypt, e.g: go run main.go rsa:encrypt \"DEV0202\"")
			}

			data := []byte(args)

			crypto := xcrypto.NewRSAEncryptor(privateKey, publicKey)
			encryptedData, err := crypto.Encrypt(data)
			if err != nil {
				return err
			}

			log.Printf("Encrypted data: %x", encryptedData)

			decryptedData, err := crypto.Decrypt(encryptedData)
			if err != nil {
				return err
			}

			log.Printf("Decrypted data: %s", decryptedData)

			return nil
		},
	}
}

func (cmd *Command) newAESEncryptor() *cli.Command {
	return &cli.Command{
		Name:  "aes:encrypt",
		Usage: "A command to run aes encryptor",
		Action: func(c *cli.Context) error {
			args := c.Args().First()
			if args == "" {
				return errors.New("please provide data to encrypt, e.g: go run main.go aes:encrypt \"DEV0202\"")
			}

			data := []byte(args)
			key := []byte("1234567890123456")

			crypto := xcrypto.NewAESEncryptor(key)
			encryptedData, err := crypto.Encrypt(data)
			if err != nil {
				return err
			}

			log.Printf("Encrypted data: %x", encryptedData)

			decryptedData, err := crypto.Decrypt(encryptedData)
			if err != nil {
				return err
			}

			log.Printf("Decrypted data: %s", decryptedData)

			return nil
		},
	}
}

func (cmd *Command) newHashing() *cli.Command {
	return &cli.Command{
		Name:  "hash",
		Usage: "A command to run hashing",
		Action: func(c *cli.Context) error {
			args := c.Args().First()
			if args == "" {
				return errors.New("please provide data to encrypt, e.g: go run main.go hash \"DEV0202\"")
			}

			log.Printf("Plain Text: %s", args)

			hashedText, salt, err := xcrypto.Hash(args)
			if err != nil {
				return err
			}

			log.Printf("Hashed Text: %s", hashedText)

			match, err := xcrypto.CompareHash(hashedText, salt, args)
			if err != nil {
				return err
			}

			log.Printf("Hash Match: %t", match)

			return nil
		},
	}
}
