package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// 加密函数
func Encrypt(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 解密函数
func Decrypt(key []byte, ciphertext string) (string, error) {
	ciphertextByte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := aesgcm.NonceSize()
	nonce, ciphertext := ciphertextByte[:nonceSize], string(ciphertextByte[nonceSize:])
	plaintext, err := aesgcm.Open(nil, nonce, []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
