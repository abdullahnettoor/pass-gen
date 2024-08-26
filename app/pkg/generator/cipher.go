package encoder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("ERROR: ENCRYPT 1:", err.Error())
		return nil, err
	}

	cypherText := make([]byte, aes.BlockSize+len(text))
	iv := cypherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("ERROR: ENCRYPT 2:", err.Error())
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cypherText[aes.BlockSize:], text)
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

	// iv := text[:aes.BlockSize]
	// text = text[aes.BlockSize:]
	// cfb := cipher.NewCFBDecrypter(block, iv)
	// cfb.XORKeyStream(text, text)
	// data, err := base64.StdEncoding.DecodeString(string(text))
	// if err != nil {
	// 	fmt.Println("ERROR: DECRYPT 1:", err.Error())
	// 	return nil, err
	// }
	// return data, nil
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return text, nil
}
