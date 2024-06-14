package cosigner

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

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
	RsaType    string `json:"rsaType"`
	AesType    string `json:"aesType"`
}

func (c *CoSignerConverter) RequestConvert(d CoSignerCallBack) (string, error) {
	responseStringMap := map[string]string{
		"key":        d.Key,
		"timestamp":  d.Timestamp,
		"bizContent": d.BizContent,
	}
	// Verify sign
	verifyRet := utils.VerifySignWithRSA(serializeParams(responseStringMap), d.Sig, c.Config.ApiPubKey)
	if !verifyRet {
		return "", errors.New("CoSignerCallBack signature verification failed")
	}
	// Use your RSA private key to decrypt response's aesKey and aesIv
	var plaintext []byte
	if d.RsaType == utils.ECB_OAEP {
		plaintext, _ = utils.DecryptWithOAEP(d.Key, c.Config.BizPrivKey)
	} else {
		plaintext, _ = utils.DecryptWithRSA(d.Key, c.Config.BizPrivKey)
	}
	resAesKey := plaintext[:32]
	resAesIv := plaintext[32:]
	// Use AES to decrypt bizContent
	ciphertext, _ := base64.StdEncoding.DecodeString(d.BizContent)
	var callBackContent []byte
	if d.AesType == utils.GCM {
		callBackContent, _ = utils.NewGCMDecrypter(resAesKey, resAesIv, ciphertext)
	} else {
		callBackContent, _ = utils.NewCBCDecrypter(resAesKey, resAesIv, ciphertext)
	}
	return string(callBackContent), nil
}

type CoSignerResponse struct {
	Approve bool   `json:"approve"`
	TxKey   string `json:"txKey"`
}

// It has been Deprecated,Please use convertCoSignerResponseWithNewCryptoType
func (c *CoSignerConverter) ResponseConverter(d any) (map[string]string, error) {
	// Use AES to encrypt request data
	aesKey := make([]byte, 32)
	rand.Read(aesKey)
	aesIv := make([]byte, 16)
	rand.Read(aesIv)
	// Create params map
	params := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().UnixMilli(), 10),
		"code":      "200",
		"message":   "SUCCESS",
	}
	if d != nil {
		payLoad, _ := json.Marshal(d)
		data := string(payLoad)
		encryptBizContent, err := utils.EncryContentWithAES(data, aesKey, aesIv)
		if err != nil {
			return nil, err
		}
		params["bizContent"] = encryptBizContent
	}

	// Use Safeheron RSA public key to encrypt request's aesKey and aesIv
	encryptedKeyAndIv, err := utils.EncryptWithRSA(append(aesKey, aesIv...), c.Config.ApiPubKey)
	if err != nil {
		return nil, err
	}
	params["key"] = encryptedKeyAndIv

	// Sign the request data with your RSA private key
	signature, err := utils.SignParamsWithRSA(serializeParams(params), c.Config.BizPrivKey)
	if err != nil {
		return nil, err
	}
	params["sig"] = signature
	return params, nil
}

func (c *CoSignerConverter) ResponseConverterWithNewCryptoType(d any) (map[string]string, error) {
	// Use AES to encrypt request data
	aesKey := make([]byte, 32)
	rand.Read(aesKey)
	aesIv := make([]byte, 16)
	rand.Read(aesIv)
	// Create params map
	params := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().UnixMilli(), 10),
		"code":      "200",
		"message":   "SUCCESS",
	}
	if d != nil {
		payLoad, _ := json.Marshal(d)
		data := string(payLoad)
		encryptBizContent, err := utils.EncryContentWithAESGCM(data, aesKey, aesIv)
		if err != nil {
			return nil, err
		}
		params["bizContent"] = encryptBizContent
	}

	// Use Safeheron RSA public key to encrypt request's aesKey and aesIv
	encryptedKeyAndIv, err := utils.EncryptWithOAEP(append(aesKey, aesIv...), c.Config.ApiPubKey)
	if err != nil {
		return nil, err
	}
	params["key"] = encryptedKeyAndIv

	// Sign the request data with your RSA private key
	signature, err := utils.SignParamsWithRSA(serializeParams(params), c.Config.BizPrivKey)
	if err != nil {
		return nil, err
	}
	params["sig"] = signature
	params["rsaType"] = utils.ECB_OAEP
	params["aesType"] = utils.GCM
	return params, nil
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
