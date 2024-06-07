package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ValidateHmac(inputMac, message, secretKey string) bool {
	expectedMac := CreateHmac(message, secretKey)
	return hmac.Equal([]byte(inputMac), []byte(expectedMac))
}

func CreateHmac(message, secretKey string) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
