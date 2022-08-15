package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func NewCBCDecrypter(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpadding(ciphertext), nil
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
