package utils

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
	"os"
	"strings"
)

const ECB_OAEP = "ECB_OAEP"

func SignParamsWithRSA(data string, privateKeyPathOrStr string) (string, error) {
	// Sign data with your RSA private key
	var privateKey *rsa.PrivateKey
	var err error
	if strings.HasSuffix(privateKeyPathOrStr, ".pem") {
		privateKey, err = loadPrivateKeyFromPath(privateKeyPathOrStr)
	} else {
		privateKey, err = ParsePrivateKey(privateKeyPathOrStr)
	}
	if err != nil {
		return "", err
	}
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	// Encode to base64 format
	b64sig := base64.StdEncoding.EncodeToString(signature)
	return b64sig, err
}

func SignParamsWithRSAPSS(data string, privateKeyPathOrStr string) (string, error) {
	// Sign data with your RSA private key
	var privateKey *rsa.PrivateKey
	var err error
	if strings.HasSuffix(privateKeyPathOrStr, ".pem") {
		privateKey, err = loadPrivateKeyFromPath(privateKeyPathOrStr)
	} else {
		privateKey, err = ParsePrivateKey(privateKeyPathOrStr)
	}
	if err != nil {
		return "", err
	}
	hashed := sha256.Sum256([]byte(data))

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash, Hash: crypto.SHA256})

	if err != nil {
		return "", err
	}
	// Encode to base64 format
	b64sig := base64.StdEncoding.EncodeToString(signature)
	return b64sig, err
}

func DecryptWithRSA(base64Data string, privateKeyPathOrStr string) ([]byte, error) {
	var privateKey *rsa.PrivateKey
	var err error
	if strings.HasSuffix(privateKeyPathOrStr, ".pem") {
		privateKey, err = loadPrivateKeyFromPath(privateKeyPathOrStr)
	} else {
		privateKey, err = ParsePrivateKey(privateKeyPathOrStr)
	}
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func DecryptWithOAEP(base64Data string, privateKeyPathOrStr string) ([]byte, error) {
	var privateKey *rsa.PrivateKey
	var err error
	if strings.HasSuffix(privateKeyPathOrStr, ".pem") {
		privateKey, err = loadPrivateKeyFromPath(privateKeyPathOrStr)
	} else {
		privateKey, err = ParsePrivateKey(privateKeyPathOrStr)
	}
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func ParsePublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	var derByte []byte
	if strings.HasPrefix(pubKeyStr, "-----BEGIN PUBLIC KEY-----") {
		var block, _ = pem.Decode([]byte(pubKeyStr))
		if block == nil {
			return nil, errors.New("failed to decode PEM block")
		}
		derByte = block.Bytes
	} else {
		var err error
		derByte, err = base64.StdEncoding.DecodeString(pubKeyStr)
		if err != nil {
			return nil, err
		}
	}
	pub, err := x509.ParsePKIXPublicKey(derByte)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func ParsePrivateKey(privKeyStr string) (*rsa.PrivateKey, error) {
	var derByte []byte
	if strings.HasPrefix(privKeyStr, "-----BEGIN PRIVATE KEY-----") {
		var block, _ = pem.Decode([]byte(privKeyStr))
		if block == nil {
			return nil, errors.New("failed to decode PEM block")
		}
		derByte = block.Bytes
	} else {
		var err error
		derByte, err = base64.StdEncoding.DecodeString(privKeyStr)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS8PrivateKey(derByte)
	if err != nil {
		return x509.ParsePKCS1PrivateKey(derByte)
	}
	return key.(*rsa.PrivateKey), nil
}

func EncryptWithRSA(data []byte, publicKeyPathOrStr string) (string, error) {
	var pubKey *rsa.PublicKey
	var err error
	if strings.HasSuffix(publicKeyPathOrStr, ".pem") {
		pubKey, err = loadPublicKeyFromPath(publicKeyPathOrStr)
	} else {
		pubKey, err = ParsePublicKey(publicKeyPathOrStr)
	}
	if err != nil {
		return "", err
	}
	signPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		return "", err
	}
	// Base64 encode
	ciphertext := base64.StdEncoding.EncodeToString(signPKCS1v15)
	return ciphertext, nil
}

func EncryptWithOAEP(data []byte, publicKeyPathOrStr string) (string, error) {
	var pubKey *rsa.PublicKey
	var err error
	if strings.HasSuffix(publicKeyPathOrStr, ".pem") {
		pubKey, err = loadPublicKeyFromPath(publicKeyPathOrStr)
	} else {
		pubKey, err = ParsePublicKey(publicKeyPathOrStr)
	}
	if err != nil {
		return "", err
	}
	signPKOAEP, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, data, nil)
	if err != nil {
		return "", err
	}
	// Base64 encode
	ciphertext := base64.StdEncoding.EncodeToString(signPKOAEP)
	return ciphertext, nil
}

func VerifySignWithRSA(data string, base64Sign string, raspublicKeyPathOrStr string) bool {
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false
	}
	var publicKey *rsa.PublicKey
	if strings.HasSuffix(raspublicKeyPathOrStr, ".pem") {
		publicKey, err = loadPublicKeyFromPath(raspublicKeyPathOrStr)
	} else {
		publicKey, err = ParsePublicKey(raspublicKeyPathOrStr)
	}
	if err != nil {
		return false
	}
	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], sign)
	return err == nil
}

func VerifySignWithRSAPSS(data string, base64Sign string, raspublicKeyPathOrStr string) bool {
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false
	}

	var publicKey *rsa.PublicKey
	if strings.HasSuffix(raspublicKeyPathOrStr, ".pem") {
		publicKey, err = loadPublicKeyFromPath(raspublicKeyPathOrStr)
	} else {
		publicKey, err = ParsePublicKey(raspublicKeyPathOrStr)
	}
	if err != nil {
		return false
	}
	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], sign, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash, Hash: crypto.SHA256})
	return err == nil
}

func loadPublicKeyFromPath(path string) (*rsa.PublicKey, error) {
	var err error
	readFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(readFile)
	if pemBlock == nil {
		return nil, fmt.Errorf("Could not read public key from[%s]. Please make sure the file in pem format, with headers and footers.(e.g. '-----BEGIN PUBLIC KEY-----' and '-----END PUBLIC KEY-----')", path)
	}
	var pkixPublicKey interface{}
	if pemBlock.Type == "RSA PUBLIC KEY" {
		// -----BEGIN RSA PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	} else if pemBlock.Type == "PUBLIC KEY" {
		// -----BEGIN PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes)
	}
	if err != nil {
		return nil, err
	}
	publicKey := pkixPublicKey.(*rsa.PublicKey)
	return publicKey, nil
}

func loadPrivateKeyFromPath(path string) (*rsa.PrivateKey, error) {
	context, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(context)
	if pemBlock == nil {
		return nil, fmt.Errorf("Could not read private key from[%s]. Please make sure the file in pem format, with headers and footers.(e.g. '-----BEGIN PRIVATE KEY-----' and '-----END PRIVATE KEY-----')", path)
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	return privateKey.(*rsa.PrivateKey), err
}
