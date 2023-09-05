package web3_demo

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/sha3"
)

var client *ethclient.Client
var tokenInstance *Token
var signer types.Signer

var web3Api Web3Api

func TestCreateWeb3Account(t *testing.T) {
	createWeb3AccountRequest := CreateWeb3AccountRequest{
		AccountName: "first-web3-wallet3",
		HiddenOnUI:  false,
	}

	var createAccountResponse CreateWeb3AccountResponse

	if err := web3Api.CreateWeb3Account(createWeb3AccountRequest, &createAccountResponse); err != nil {
		panic(fmt.Errorf("failed to create web3 wallet account, %w", err))
	}

	log.Infof("web3 wallet account created, accountKey: %s, evm address: %s",
		createAccountResponse.AccountKey, createAccountResponse.AddressList[0].Address)
}

/*
1. Modify demo/api_demo/web3/config.yaml.example according to the comments
2. Execute Command: cp config.yaml.example config.yaml
3. Execute Command: go test -run TestETHSign
*/
func TestETHSign(t *testing.T) {
	// Create contract function data
	functionData := createTransferData(1, viper.GetString("toAddress"))

	// Create Raw Transaction
	rawTransaction := createTransaction(viper.GetString("evmAddress"), 0, viper.GetString("contractAddress"), functionData)
	txString, _ := json.Marshal(rawTransaction)
	log.Infof("Raw transaction data: %s", txString)

	// Encode the transaction and compute the hash value
	hash := signer.Hash(rawTransaction).Hex()

	// Sign transaction with Safeheron
	customerRefId := uuid.New().String()
	txKey := createWeb3EthSign(viper.GetString("accountKey"), customerRefId,
		rawTransaction.ChainId().String(), []string{hash})

	log.Infof("Web3 sign task has been created with Safeheron API, txKey: %s, customerRefId: %s", txKey, customerRefId)
	log.Info("You can approve the sign task with Safeheron mobile app or API Co-Signer according to your policy config")

	// Query
	signedTransaction := queryWeb3Sig(customerRefId)
	log.Infof("Got sig data: %s", signedTransaction.MessageHash.SigList[0].Sig)

	// Add 0x prefix to sig
	sigByte, _ := hexutil.Decode("0x" + signedTransaction.MessageHash.SigList[0].Sig)

	// Encode tx with sig
	signedTx, err := rawTransaction.WithSignature(signer, sigByte)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Transaction successful with hash: %s", signedTx.Hash().Hex())
}

/*
1. Modify demo/api_demo/web3/config.yaml.example according to the comments
2. Execute Command: cp config.yaml.example config.yaml
3. Execute Command: go test -run TestPersonalSign
*/
func TestPersonalSign(t *testing.T) {
	chainId, _ := client.NetworkID(context.Background())
	// Sign transaction with Safeheron
	customerRefId := uuid.New().String()
	txKey := createWeb3PersonalSign(viper.GetString("accountKey"), customerRefId,
		chainId.String(), "demo text")

	log.Infof("Web3 sign task has been created with Safeheron API, txKey: %s, customerRefId: %s", txKey, customerRefId)
	log.Info("You can approve the sign task with Safeheron mobile app or API Co-Signer according to your policy config")

	// Query
	signedTransaction := queryWeb3Sig(customerRefId)
	log.Infof("Got sig data: %s", signedTransaction.Message.Sig.Sig)
}

/*
1. Modify demo/api_demo/web3/config.yaml.example according to the comments
2. Execute Command: cp config.yaml.example config.yaml
3. Execute Command: go test -run TestEthSignTypedData
*/
func TestEthSignTypedData(t *testing.T) {
	chainId, _ := client.NetworkID(context.Background())
	// Sign transaction with Safeheron
	customerRefId := uuid.New().String()
	txKey := createWeb3EthSignTypedData(viper.GetString("accountKey"), customerRefId,
		chainId.String(), "{\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Person\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"wallet\",\"type\":\"address\"}],\"Mail\":[{\"name\":\"from\",\"type\":\"Person\"},{\"name\":\"to\",\"type\":\"Person\"},{\"name\":\"contents\",\"type\":\"string\"}]},\"primaryType\":\"Mail\",\"domain\":{\"name\":\"Ether Mail\",\"version\":\"1\",\"chainId\":1,\"verifyingContract\":\"0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC\"},\"message\":{\"from\":{\"name\":\"Cow\",\"wallet\":\"0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826\"},\"to\":{\"name\":\"Bob\",\"wallet\":\"0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB\"},\"contents\":\"Hello, Bob!\"}}", "ETH_SIGNTYPEDDATA_V4")

	log.Infof("Web3 sign task has been created with Safeheron API, txKey: %s, customerRefId: %s", txKey, customerRefId)
	log.Info("You can approve the sign task with Safeheron mobile app or API Co-Signer according to your policy config")

	// Query
	signedTransaction := queryWeb3Sig(customerRefId)
	log.Infof("Got sig data: %s", signedTransaction.Message.Sig.Sig)
}

/*
1. Modify demo/api_demo/web3/config.yaml.example according to the comments
2. Execute Command: cp config.yaml.example config.yaml
3. Execute Command: go test -run TestSignTransaction
*/
func TestSignTransaction(t *testing.T) {
	// Create contract function data
	functionData := createTransferData(1, viper.GetString("toAddress"))

	// Create Raw Transaction
	rawTransaction := createTransaction(viper.GetString("evmAddress"), 0, viper.GetString("contractAddress"), functionData)
	txString, _ := json.Marshal(rawTransaction)
	log.Infof("Raw transaction data: %s", txString)
	// Sign transaction with Safeheron
	customerRefId := uuid.New().String()
	txKey := createWeb3EthSignTransaction(viper.GetString("accountKey"), customerRefId,
		rawTransaction.To().String(), rawTransaction.Value().String(),
		rawTransaction.ChainId().String(), "", rawTransaction.Gas(),
		rawTransaction.GasTipCap().String(), rawTransaction.GasFeeCap().String(),
		rawTransaction.Nonce(), "0x"+hex.EncodeToString(rawTransaction.Data()))

	log.Infof("Web3 sign task has been created with Safeheron API, txKey: %s, customerRefId: %s", txKey, customerRefId)
	log.Info("You can approve the sign task with Safeheron mobile app or API Co-Signer according to your policy config")

	// Query
	signedTransaction := queryWeb3Sig(customerRefId)
	log.Infof("Got sig data: %s", signedTransaction.Transaction.Sig.Sig)

	// Add 0x prefix to sig
	sigByte, _ := hexutil.Decode("0x" + signedTransaction.Transaction.Sig.Sig[:])

	// Encode tx with sig
	signedTx, err := rawTransaction.WithSignature(signer, sigByte)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Transaction successful with hash: %s", signedTx.Hash().Hex())
}

func createTransferData(amount float64, to string) []byte {
	decimals, _ := tokenInstance.Decimals(&bind.CallOpts{})
	tokenAmount := new(big.Int)
	big.NewFloat(0).Mul(big.NewFloat(amount), big.NewFloat(math.Pow(10, float64(decimals)))).Int(tokenAmount)

	var data []byte
	data = append(data, createERC20FunctionMethodId("transfer(address,uint256)")...)
	data = append(data, common.LeftPadBytes(common.HexToAddress(to).Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(tokenAmount.Bytes(), 32)...)

	return data
}

func createERC20FunctionMethodId(method string) []byte {
	fnSignature := []byte(method)
	fnHash := sha3.NewLegacyKeccak256()
	fnHash.Write(fnSignature)
	return fnHash.Sum(nil)[:4]
}

func createTransaction(from string, value float64, to string, data []byte) *types.Transaction {
	fromAddress := common.HexToAddress(from)
	toAddress := common.HexToAddress(to)
	chainId, _ := client.NetworkID(context.Background())
	// Get data from block chain: nonce, gasPrice
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	// gasPrice, _ := client.SuggestGasPrice(context.Background())

	// Estimate gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Value: big.NewInt(0),
		Data:  data,
	})

	if err != nil {
		panic(err)
	}

	transferAmount := ethUnit.NewEther(big.NewFloat(value)).Wei()

	// Estimate maxFeePerGas, we assume maxPriorityFeePerGas's value is 2(gwei).
	// The baseFeePerGas is recommended to be 2 times the latest block's baseFeePerGas value.
	// maxFeePerGas must not less than baseFeePerGas + maxPriorityFeePerGas
	maxPriorityFeePerGas := ethUnit.NewGWei(big.NewFloat(2)).Wei()
	lastBlockHeader, _ := client.HeaderByNumber(context.Background(), nil)
	baseFee := lastBlockHeader.BaseFee
	suggestBaseFee := big.NewInt(0).Mul(baseFee, big.NewInt(2))
	maxFeePerGas := big.NewInt(0).Add(suggestBaseFee, maxPriorityFeePerGas)

	// Create raw transaction
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		To:        &toAddress,
		Value:     transferAmount,
		Data:      data,
		Gas:       gasLimit,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
	})

	jsonTx, _ := json.Marshal(tx)
	log.Infof("tx json data: %s", string(jsonTx))

	return tx
}

/*
1. Modify demo/api_demo/web3/config.yaml.example according to the comments
2. Execute Command: cp config.yaml.example config.yaml
3. Execute Command: go test -run TestLegacyTxSignTransaction
*/
func TestLegacyTxSignTransaction(t *testing.T) {
	// Create Raw Transaction
	rawTransaction := createLegacyTx(viper.GetString("evmAddress"), 0, viper.GetString("contractAddress"), viper.GetString("contractFunctionData"))
	txString, _ := json.Marshal(rawTransaction)
	log.Infof("Raw transaction data: %s", txString)

	// get chainId
	chainId, _ := client.NetworkID(context.Background())
	// Sign transaction with Safeheron
	customerRefId := uuid.New().String()
	txKey := createWeb3EthSignTransaction(viper.GetString("accountKey"), customerRefId,
		rawTransaction.To().String(), rawTransaction.Value().String(),
		chainId.String(), rawTransaction.GasPrice().String(), rawTransaction.Gas(),
		"", "",
		rawTransaction.Nonce(), string(rawTransaction.Data()))

	log.Infof("Web3 sign task has been created with Safeheron API, txKey: %s, customerRefId: %s", txKey, customerRefId)
	log.Info("You can approve the sign task with Safeheron mobile app or API Co-Signer according to your policy config")

	// Query
	signedTransaction := queryWeb3Sig(customerRefId)
	//log.Infof("Got signed transaction data: %s", signedTransaction)

	// Broadcast
	txHash := broadcast(signedTransaction.Transaction.SignedTransaction, rawTransaction.Type())
	log.Infof("Broadcast success, transaction hash: %s", txHash)

}

func createLegacyTx(from string, value float64, to string, data string) *types.Transaction {
	fromAddress := common.HexToAddress(from)
	toAddress := common.HexToAddress(to)

	// Get data from block chain: nonce, gasPrice
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, _ := client.SuggestGasPrice(context.Background())

	// Decode data
	callData, _ := hex.DecodeString(data[2:])

	// Estimate gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Value: big.NewInt(0),
		Data:  callData,
	})

	if err != nil {
		panic(err)
	}

	// Create raw transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    big.NewInt(0),
		Data:     []byte(data),
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})

	return tx
}

func createWeb3EthSignTransaction(accountKey string, customerRefId string, to string,
	value string, chainId string, gasPrice string, gasLimit uint64, maxPriorityFeePerGas string,
	maxFeePerGas string, nonce uint64, data string) string {
	ethSignTransactionRequest := EthSignTransactionRequest{
		AccountKey:    accountKey,
		CustomerRefId: customerRefId,
		Transaction: struct {
			To                   string `json:"to"`
			Value                string `json:"value"`
			ChainId              string `json:"chainId"`
			GasPrice             string `json:"gasPrice,omitempty"`
			GasLimit             string `json:"gasLimit"`
			MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
			MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
			Nonce                string `json:"nonce"`
			Data                 string `json:"data,omitempty"`
		}{
			To:                   to,
			Value:                value,
			ChainId:              chainId,
			GasLimit:             strconv.FormatUint(gasLimit, 10),
			GasPrice:             gasPrice,
			MaxPriorityFeePerGas: maxPriorityFeePerGas,
			MaxFeePerGas:         maxFeePerGas,
			Nonce:                strconv.FormatUint(nonce, 10),
			Data:                 data,
		},
	}
	var txKeyResult TxKeyResult
	if err := web3Api.EthSignTransaction(ethSignTransactionRequest, &txKeyResult); err != nil {
		panic(fmt.Errorf("failed to send create web3 eth_signTransaction, %w", err))
	}

	return txKeyResult.TxKey
}

func createWeb3EthSign(accountKey string, customerRefId string, chainId string, hashList []string) string {

	ethSignRequest := EthSignRequest{
		AccountKey:    accountKey,
		CustomerRefId: customerRefId,
		MessageHash: struct {
			ChainId string   `json:"chainId"`
			Hash    []string `json:"hash"`
		}{
			ChainId: chainId,
			Hash:    hashList,
		},
	}
	var txKeyResult TxKeyResult
	if err := web3Api.EthSign(ethSignRequest, &txKeyResult); err != nil {
		panic(fmt.Errorf("failed to send create web3 eth_sign, %w", err))
	}

	return txKeyResult.TxKey
}

func createWeb3PersonalSign(accountKey string, customerRefId string, chainId string, data string) string {

	personalSignRequest := PersonalSignRequest{
		AccountKey:    accountKey,
		CustomerRefId: customerRefId,
		Message: struct {
			ChainId string `json:"chainId"`
			Data    string `json:"data"`
		}{
			ChainId: chainId,
			Data:    data,
		},
	}
	var txKeyResult TxKeyResult
	if err := web3Api.PersonalSign(personalSignRequest, &txKeyResult); err != nil {
		panic(fmt.Errorf("failed to send create web3 personal_sign, %w", err))
	}

	return txKeyResult.TxKey
}

func createWeb3EthSignTypedData(accountKey string, customerRefId string, chainId string, data string, version string) string {

	ethSignTypedDataRequest := EthSignTypedDataRequest{
		AccountKey:    accountKey,
		CustomerRefId: customerRefId,
		Message: struct {
			ChainId string `json:"chainId"`
			Data    string `json:"data"`
			Version string `json:"version"`
		}{
			ChainId: chainId,
			Data:    data,
			Version: version,
		},
	}
	var txKeyResult TxKeyResult
	if err := web3Api.EthSignTypedData(ethSignTypedDataRequest, &txKeyResult); err != nil {
		panic(fmt.Errorf("failed to send create web3 eth_signTypedData, %w", err))
	}

	return txKeyResult.TxKey
}

func queryWeb3Sig(customerRefId string) Web3SignQueryResponse {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	web3SignQueryRequest := Web3SignQueryRequest{
		CustomerRefId: customerRefId,
	}

	var web3SignQueryResponse Web3SignQueryResponse

	for range ticker.C {
		if err := web3Api.QueryWeb3Sig(web3SignQueryRequest, &web3SignQueryResponse); err != nil {
			panic(fmt.Errorf("failed to query web3 eth_signTransaction result, %w", err))
		}

		log.Infof(`web3 sign task status: %s`, web3SignQueryResponse.TransactionStatus)

		if web3SignQueryResponse.TransactionStatus == "FAILED" || web3SignQueryResponse.TransactionStatus == "REJECTED" {
			panic(`web3 sign task was FAILED or REJECTED`)
		} else if web3SignQueryResponse.TransactionStatus == "SIGN_COMPLETED" {
			return web3SignQueryResponse
		} else {
			log.Infof(`will wait another 5 seconds`)
		}
	}

	panic("can't get web3 sign task result.")
}

func broadcast(signedTransaction string, txType uint8) string {
	signedTransactionBytes, err := hex.DecodeString(signedTransaction[2:])
	if err != nil {
		panic(fmt.Errorf("broadcast failed. %w", err))
	}
	tx := new(types.Transaction)

	if txType == types.LegacyTxType {
		rlp.DecodeBytes(signedTransactionBytes, &tx)
	} else {
		log.Infof("---------------------signedTransactionBytes data: %s", string(signedTransactionBytes))
		if err := tx.UnmarshalBinary(signedTransactionBytes); err != nil {
			panic(fmt.Errorf("decode signed transaction failed. %w", err))
		}
	}

	if err := client.SendTransaction(context.Background(), tx); err != nil {
		panic(fmt.Errorf("broadcast failed. %w", err))
	}

	return tx.Hash().Hex()
}

func setup() {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %w", err))
	}

	sc := safeheron.Client{Config: safeheron.ApiConfig{
		BaseUrl:               viper.GetString("baseUrl"),
		ApiKey:                viper.GetString("apiKey"),
		RsaPrivateKey:         viper.GetString("privateKeyPemFile"),
		SafeheronRsaPublicKey: viper.GetString("safeheronPublicKeyPemFile"),
	}}

	web3Api = Web3Api{Client: sc}
	client, _ = ethclient.Dial(viper.GetString("ethereumRpcApi"))
	chainId, _ := client.NetworkID(context.Background())
	signer = types.NewLondonSigner(chainId)
	tokenInstance, _ = NewToken(common.HexToAddress(viper.GetString("contractAddress")), client)
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
