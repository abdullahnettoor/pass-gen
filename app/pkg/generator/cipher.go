package generator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func Erncypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR: ENCRYPT 1:", err.Error())
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)
	cypherText := make([]byte, aes.BlockSize+len(b))
	iv := cypherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("ERROR: ENCRYPT 2:", err.Error())
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cypherText[aes.BlockSize:], []byte(b))
	return cypherText, nil
}

func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
