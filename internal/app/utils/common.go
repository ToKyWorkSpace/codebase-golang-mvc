package utils

import (
	"codebase-golang/internal/app/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var key = []byte(config.GetEnv("SECRET_KEY"))

func Encrypt(text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	result := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", result)
}

func Decrypt(text string) string {
	enc, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphered := enc[:nonceSize], enc[nonceSize:]

	result, err := aesGCM.Open(nil, nonce, ciphered, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", result)
}
