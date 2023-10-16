package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
)

func SignParamsWithRSA(data string, privateKeyPath string) (string, error) {
	// Sign data with your RSA private key
	privateKey, err := loadPrivateKeyFromPath(privateKeyPath)
	if err != nil {
		return "", err
	}

	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])

	// Encode to base64 format
	b64sig := base64.StdEncoding.EncodeToString(signature)
	return b64sig, err
}

func DecryptWithRSA(base64Data string, privateKeyPath string) ([]byte, error) {
	privateKey, err := loadPrivateKeyFromPath(privateKeyPath)
	if err != nil {
		return nil, err
	}

	data, _ := base64.StdEncoding.DecodeString(base64Data)
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func EncryptWithRSA(data []byte, publicKeyPath string) (string, error) {
	pubKey, err := loadPublicKeyFromPath(publicKeyPath)
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

func VerifySignWithRSA(data string, base64Sign string, rasPublicKeyPath string) bool {
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false
	}

	publicKey, err := loadPublicKeyFromPath(rasPublicKeyPath)
	if err != nil {
		return false
	}

	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], sign)
	return err == nil
}

func loadPublicKeyFromPath(path string) (*rsa.PublicKey, error) {
	var err error
	readFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(readFile)
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
	context, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(context)
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	return privateKey.(*rsa.PrivateKey), err
}
