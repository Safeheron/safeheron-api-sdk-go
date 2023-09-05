package cosigner_demo

import (
	"encoding/base64"
	"errors"
	"sort"
	"strings"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/utils"
)

type CoSignerConverter struct {
	Config CoSignerConfig
}

type CoSignerConfig struct {
	ApiPubKey  string `comment:"apiPubKey"`
	BizPrivKey string `comment:"bizPrivKey"`
}

type CoSignerCallBack struct {
	Timestamp  string `json:"timestamp"`
	Sig        string `json:"sig"`
	Key        string `json:"key"`
	BizContent string `json:"bizContent"`
}

func (c *CoSignerConverter) Convert(d CoSignerCallBack) (string, error) {
	responseStringMap := map[string]string{
		"key":        d.Key,
		"timestamp":  d.Timestamp,
		"bizContent": d.BizContent,
	}
	// Verify sign
	verifyRet := utils.VerifySignWithRSA(serializeParams(responseStringMap), d.Sig, c.Config.ApiPubKey)
	if !verifyRet {
		return "", errors.New("response signature verification failed")
	}
	// Use your RSA private key to decrypt response's aesKey and aesIv
	plaintext, _ := utils.DecryptWithRSA(d.Key, c.Config.BizPrivKey)
	resAesKey := plaintext[:32]
	resAesIv := plaintext[32:]
	// Use AES to decrypt bizContent
	ciphertext, _ := base64.StdEncoding.DecodeString(d.BizContent)
	respContent, _ := utils.NewCBCDecrypter(resAesKey, resAesIv, ciphertext)
	return string(respContent), nil
}

func serializeParams(params map[string]string) string {
	// Sort by key and serialize all request param into apiKey=...&bizContent=... format
	var data []string
	for k, v := range params {
		data = append(data, strings.Join([]string{k, v}, "="))
	}
	sort.Strings(data)
	return strings.Join(data, "&")
}
