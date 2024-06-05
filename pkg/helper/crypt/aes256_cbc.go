package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/mergermarket/go-pkcs7"
)

func Decrypt(encrypted string, passphrase string) (string, error) {
	key := []byte(CreateHash(passphrase))

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}
	iv := bytes.NewBuffer(make([]byte, 16))

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv.Bytes())
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext, _ = pkcs7.Unpad(ciphertext, aes.BlockSize)

	return string(ciphertext), nil
}

func Encrypt(unencrypted, passphrase string) (string, error) {
	key := []byte(CreateHash(passphrase))
	plainText := []byte(unencrypted)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)

	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := bytes.NewBuffer(make([]byte, 16))

	mode := cipher.NewCBCEncrypter(block, iv.Bytes())
	mode.CryptBlocks(plainText, plainText)

	return base64.StdEncoding.EncodeToString(plainText), nil
}
