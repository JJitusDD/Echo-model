package crypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
)

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func VerifyHash(value string, hash string) (bool, string) {
	bs := HashSha1(value)
	if bs == hash {
		return true, bs
	}
	return false, bs
}

func HashSha1(value string) string {
	hash := sha1.New()
	hash.Write([]byte(value))
	bs := hash.Sum(nil)

	return hex.EncodeToString(bs)
}

func HashSha256(value string) string {
	hash := sha256.New()
	hash.Write([]byte(value))
	bs := hash.Sum(nil)

	return hex.EncodeToString(bs)
}

func VerifyHash256(value string, hash string) (bool, string) {
	bs := HashSha256(value)
	if bs == hash {
		return true, bs
	}
	return false, bs
}

func HashhmacSha256(content, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(content))
	expectedMAC := mac.Sum(nil)
	return hex.EncodeToString(expectedMAC)
}

func HashHMACSH(input, key string) string {
	key_for_sign := []byte(key)
	h := hmac.New(sha1.New, key_for_sign)
	h.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func VerifyHashRequest(c map[string]interface{}, apiSecretKey string) (e error) {
	stringHash := ""
	hashValue := ""
	var keys []string
	for s := range c {
		if s != "hash" {
			keys = append(keys, s)
		}

	}
	sort.Strings(keys)
	for _, v := range keys {
		switch c[v].(type) {
		case float64:
			if c[v].(float64) > 1000 {
				convert := c[v].(float64)
				c[v] = int64(convert)
			}
		default:
			c[v] = c[v]
		}
		str := fmt.Sprint(c[v])

		if stringHash != "" {
			stringHash = stringHash + "|" + str
		} else {
			stringHash = stringHash + str
		}
	}

	if c["hash"] != nil {
		hashValue = fmt.Sprint(c["hash"])
	}

	sha1Verify, _ := VerifyHash(stringHash+"|"+apiSecretKey, hashValue)

	if !sha1Verify {
		return fmt.Errorf("VerifyHash line 66 error ")
	}
	return
}
