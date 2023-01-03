package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	cfg "github.com/deall-users/config"
)

// Decrypt Data
func Decrypt(plainText string) (string, error) {
	var err error
	if len(plainText) <= 0 {
		return plainText, errors.New("empty text")
	}
	text, err := base64.StdEncoding.DecodeString(plainText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(cfg.KEY_ENCRYPTION))
	if err != nil {
		return "", err
	}

	if len(text) < aes.BlockSize {
		return "", errors.New("cipher text not match")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text), nil
}

// Encrypt Data
func Encrypt(plainText string) (string, error) {
	var err error
	if len(plainText) <= 0 {
		return plainText, errors.New("empty text")
	}
	text := []byte(plainText)
	block, err := aes.NewCipher([]byte(cfg.KEY_ENCRYPTION))
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], text)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
