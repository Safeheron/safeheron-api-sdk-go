package webhook

import (
	"encoding/base64"
	"errors"
	"sort"
	"strings"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/utils"
)

type WebhookConverter struct {
	Config WebHookConfig
}

type WebHookConfig struct {
	SafeheronWebHookRsaPublicKey string `comment:"safeheronWebHookRsaPublicKey"`
	WebHookRsaPrivateKey         string `comment:"webHookRsaPrivateKey"`
}

type WebHook struct {
	Timestamp  string `json:"timestamp"`
	Sig        string `json:"sig"`
	Key        string `json:"key"`
	BizContent string `json:"bizContent"`
	RsaType    string `json:"rsaType"`
	AesType    string `json:"aesType"`
}

type WebHookResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *WebhookConverter) Convert(d WebHook) (string, error) {
	responseStringMap := map[string]string{
		"key":        d.Key,
		"timestamp":  d.Timestamp,
		"bizContent": d.BizContent,
	}
	// Verify sign
	verifyRet := utils.VerifySignWithRSA(serializeParams(responseStringMap), d.Sig, c.Config.SafeheronWebHookRsaPublicKey)
	if !verifyRet {
		return "", errors.New("response signature verification failed")
	}
	// Use your RSA private key to decrypt response's aesKey and aesIv
	var plaintext []byte
	if d.RsaType == utils.ECB_OAEP {
		plaintext, _ = utils.DecryptWithOAEP(d.Key, c.Config.WebHookRsaPrivateKey)
	} else {
		plaintext, _ = utils.DecryptWithRSA(d.Key, c.Config.WebHookRsaPrivateKey)
	}

	resAesKey := plaintext[:32]
	resAesIv := plaintext[32:]
	// Use AES to decrypt bizContent
	ciphertext, _ := base64.StdEncoding.DecodeString(d.BizContent)
	var webHookContent []byte
	if d.AesType == utils.GCM {
		webHookContent, _ = utils.NewGCMDecrypter(resAesKey, resAesIv, ciphertext)
	} else {
		webHookContent, _ = utils.NewCBCDecrypter(resAesKey, resAesIv, ciphertext)
	}
	return string(webHookContent), nil
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
