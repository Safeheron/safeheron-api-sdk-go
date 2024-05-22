package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

const GCM = "GCM_NOPADDING"

func NewCBCDecrypter(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpadding(ciphertext), nil
}

func NewGCMDecrypter(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return nil, err
	}

	decrypted, err := aesGCM.Open(nil, iv, ciphertext, nil)

	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func EncryContentWithAES(data string, aesKey []byte, aesIv []byte) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	plaintext := []byte(data)
	plaintext = padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	if _, err := io.ReadFull(rand.Reader, aesIv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, aesIv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext[aes.BlockSize:]), nil
}

func EncryContentWithAESGCM(data string, aesKey []byte, aesIv []byte) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	plaintext := []byte(data)
	aesGCM, err := cipher.NewGCMWithNonceSize(block, len(aesIv))
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, aesIv, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}
