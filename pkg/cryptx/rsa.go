package cryptx

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func (c *cryptox) DecryptRSAWithBase64(base64PrivateKey string, base64Ciphertext string) (string, error) {
	// Decode the Base64 private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(base64PrivateKey)
	if err != nil {
		return "", ErrDecodeBase64
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return "", ErrParsePEMBlock
	}

	var privateKey *rsa.PrivateKey

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// If PKCS#1 parsing fails, try PKCS#8
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			err = fmt.Errorf("%w: %w", ErrParsePrivateKey, err)
			return "", fmt.Errorf("failed to parse private key: %w", err)
		}
		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return "", ErrUnsupportedRSAKeyFormat
		}
	}

	// Decode the Base64 ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		return "", ErrDecodeBase64CipherText
	}

	// Decrypt the ciphertext using the private key
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
