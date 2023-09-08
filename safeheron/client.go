package safeheron

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/utils"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	Config ApiConfig
}

type SafeheronResponse struct {
	Code       int64  `form:"code" json:"code"`
	Message    string `form:"message" json:"message"`
	Sig        string `form:"sig" json:"sig"`
	Key        string `form:"key" json:"key"`
	BizContent string `form:"bizContent" json:"bizContent"`
	Timestamp  string `form:"timestamp" json:"timestamp"`
}

func (c Client) SendRequest(request any, response any, path string) error {
	respContent, err := c.execute(request, path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respContent, &response)
	return err
}

func (c Client) execute(request any, endpoint string) ([]byte, error) {
	// Use AES to encrypt request data
	aesKey := make([]byte, 32)
	rand.Read(aesKey)
	aesIv := make([]byte, 16)
	rand.Read(aesIv)
	// Create params map
	params := map[string]string{
		"apiKey":    c.Config.ApiKey,
		"timestamp": strconv.FormatInt(time.Now().UnixMicro(), 10),
	}
	if request != nil {
		payLoad, _ := json.Marshal(request)
		data := string(payLoad)
		log.Infof("send request data: %s", data)
		encryptBizContent, err := utils.EncryContentWithAES(data, aesKey, aesIv)
		if err != nil {
			return nil, err
		}
		params["bizContent"] = encryptBizContent
	}

	// Use Safeheron RSA public key to encrypt request's aesKey and aesIv
	encryptedKeyAndIv, err := utils.EncryptWithRSA(append(aesKey, aesIv...), c.Config.SafeheronRsaPublicKey)
	if err != nil {
		return nil, err
	}
	params["key"] = encryptedKeyAndIv

	// Sign the request data with your RSA private key
	signature, err := utils.SignParamsWithRSA(serializeParams(params), c.Config.RsaPrivateKey)
	if err != nil {
		return nil, err
	}
	params["sig"] = signature

	// Send post
	safeheronResponse, _ := c.Post(params, endpoint)

	// Decode json data into SafeheronResponse struct
	var responseStruct SafeheronResponse
	json.Unmarshal(safeheronResponse, &responseStruct)
	if responseStruct.Code != 200 {
		log.Warnf("request failed: %d, message: %s", responseStruct.Code, responseStruct.Message)
		return nil, fmt.Errorf("request failed, code: %d, message: %s", responseStruct.Code, responseStruct.Message)
	}

	responseStringMap := map[string]string{
		"code":       strconv.FormatInt(responseStruct.Code, 10),
		"message":    responseStruct.Message,
		"key":        responseStruct.Key,
		"timestamp":  responseStruct.Timestamp,
		"bizContent": responseStruct.BizContent,
	}

	// Verify sign
	verifyRet := utils.VerifySignWithRSA(serializeParams(responseStringMap), responseStruct.Sig, c.Config.SafeheronRsaPublicKey)
	if !verifyRet {
		return nil, errors.New("response signature verification failed")
	}

	// Use your RSA private key to decrypt response's aesKey and aesIv
	plaintext, _ := utils.DecryptWithRSA(responseStruct.Key, c.Config.RsaPrivateKey)
	resAesKey := plaintext[:32]
	resAesIv := plaintext[32:]
	// Use AES to decrypt bizContent
	ciphertext, _ := base64.StdEncoding.DecodeString(responseStruct.BizContent)
	respContent, _ := utils.NewCBCDecrypter(resAesKey, resAesIv, ciphertext)
	return respContent, nil
}

func (c Client) Post(params map[string]string, path string) ([]byte, error) {
	jsonValue, _ := json.Marshal(params)

	resp, err := http.Post(fmt.Sprintf("%s%s", c.Config.BaseUrl, path), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("request safeheron api error, api: %s", path)
		return nil, err
	}

	return body, nil
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
