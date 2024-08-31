/*
Package encoder provides encryption and decryption utilities using AES-CFB.
It handles secure password storage and retrieval for the password manager.
*/
package encoder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// Encrypt takes a key and plaintext, returns the encrypted ciphertext
// Uses AES-CFB mode with a random IV for encryption
func Encrypt(key, text []byte) ([]byte, error) {
	// Create new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR: ENCRYPT 1:", err.Error())
		return nil, err
	}

	// Create ciphertext buffer with space for IV
	cypherText := make([]byte, aes.BlockSize+len(text))
	iv := cypherText[:aes.BlockSize]

	// Generate random IV
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("ERROR: ENCRYPT 2:", err.Error())
		return nil, err
	}

	// Encrypt using CFB mode
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cypherText[aes.BlockSize:], text)
	return cypherText, nil
}

// Decrypt takes a key and ciphertext, returns the decrypted plaintext
// Uses AES-CFB mode and extracts IV from ciphertext
func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	// Extract IV and decrypt
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return text, nil
}
