package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func SignerFromPem(pemBytes []byte, password []byte) (ssh.Signer, error) {

	// read pem block
	err := errors.New("Pem decode failed, no key found")
	pemBlock, _ := pem.Decode(pemBytes)
	if pemBlock == nil {
		return nil, err
	}

	// handle encrypted key
	// get RSA, EC or DSA key
	key, err := parsePemBlock(pemBlock)
	if err != nil {
		return nil, err
	}

	// generate signer instance from key
	signer, err := ssh.NewSignerFromKey(key)
	if err != nil {
		return nil, fmt.Errorf("Creating signer from encrypted key failed %v", err)
	}

	return signer, nil

}

// parsePemBlock parse pemblock to PrivateKey
func parsePemBlock(block *pem.Block) (interface{}, error) {
	switch block.Type {
	case "RSA PRIVATE KEY":
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("Parsing PKCS private key failed %v", err)
		} else {
			return key, nil
		}

	case "EC PRIVATE KEY":
		key, err := x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("Parsing EC private key failed %v", err)
		} else {
			return key, nil
		}

	case "DSA PRIVATE KEY":
		key, err := ssh.ParseDSAPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("Parsing DSA private key failed %v", err)
		} else {
			return key, nil
		}
	default:
		return nil, fmt.Errorf("Parsing private key failed, unsupported key type %q", block.Type)
	}
}

// EncryptRSA -
func EncryptRSA(datastring string, privkey string) string {
	pemRes, _ := pem.Decode([]byte(privkey))
	privateKeyI, _ := x509.ParsePKCS1PrivateKey(pemRes.Bytes)
	h := sha256.New()
	h.Write([]byte(datastring))
	sum := h.Sum(nil)

	sig, _ := rsa.SignPKCS1v15(rand.Reader, privateKeyI, crypto.SHA256, sum)

	return base64.StdEncoding.EncodeToString(sig)
}
