package web3_demo

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var client *ethclient.Client

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
3. Execute Command: go test -run TestEthSignTransaction
*/
func TestEthSignTransaction(t *testing.T) {
	// Create Raw Transaction
	rawTransaction := createTransaction(viper.GetString("evmAddress"), 0, viper.GetString("contractAddress"), viper.GetString("contractFunctionData"))

	// Sign transaction with Safeheron
	customerRefId := uuid.New().String()
	txKey := createWeb3EthSignTransaction(viper.GetString("accountKey"), customerRefId,
		rawTransaction.To().String(), rawTransaction.Value().String(),
		rawTransaction.ChainId().String(), rawTransaction.Gas(),
		rawTransaction.GasTipCap().String(), rawTransaction.GasFeeCap().String(),
		rawTransaction.Nonce(), string(rawTransaction.Data()))
	log.Infof("web3 eth_signTransaction has been created, txKey: %s", txKey)

	// Query
	signedTransaction := queryWeb3Sig(customerRefId)
	log.Infof("got signed transaction data: %s", signedTransaction)

	// Broadcast
	txHash := broadcast(signedTransaction, rawTransaction.Type())
	log.Infof("You can view the transaction at: https://goerli.etherscan.io/tx/%s", txHash)

}

func createTransaction(from string, value float64, to string, data string) *types.Transaction {
	fromAddress := common.HexToAddress(from)
	toAddress := common.HexToAddress(to)

	// Get data from block chain: nonce, chainId
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	chainId, _ := client.NetworkID(context.Background())

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
		Value:     big.NewInt(0),
		Data:      []byte(data),
		Gas:       gasLimit,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
	})

	jsonTx, _ := json.Marshal(tx)
	log.Infof("tx json data: %s", string(jsonTx))

	return tx
}

func createWeb3EthSignTransaction(accountKey string, customerRefId string, to string,
	value string, chainId string, gasLimit uint64, maxPriorityFeePerGas string,
	maxFeePerGas string, nonce uint64, data string) string {
	ethSignTransactionRequest := EthSignTransactionRequest{
		AccountKey:    accountKey,
		CustomerRefId: customerRefId,
		Transaction: struct {
			To                   string `json:"to,omitempty"`
			Value                string `json:"value,omitempty"`
			ChainId              string `json:"chainId,omitempty"`
			GasPrice             string `json:"gasPrice,omitempty"`
			GasLimit             string `json:"gasLimit,omitempty"`
			MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
			MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
			Nonce                string `json:"nonce,omitempty"`
			Data                 string `json:"data,omitempty"`
		}{
			To:                   to,
			Value:                value,
			ChainId:              chainId,
			GasLimit:             strconv.FormatUint(gasLimit, 10),
			MaxPriorityFeePerGas: maxPriorityFeePerGas,
			MaxFeePerGas:         maxFeePerGas,
			Nonce:                strconv.FormatUint(nonce, 10),
			Data:                 data,
		},
	}

	var ethSignTransactionResponse EthSignTransactionResponse
	if err := web3Api.EthSignTransaction(ethSignTransactionRequest, &ethSignTransactionResponse); err != nil {
		panic(fmt.Errorf("failed to send create web3 eth_signTransaction, %w", err))
	}

	return ethSignTransactionResponse.TxKey
}

func queryWeb3Sig(customerRefId string) string {
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

		log.Infof(`web3 eth_signTransaction status: %s, sub status: %s`, web3SignQueryResponse.TransactionStatus, web3SignQueryResponse.TransactionSubStatus)

		if web3SignQueryResponse.TransactionStatus == "FAILED" || web3SignQueryResponse.TransactionStatus == "REJECTED" {
			panic(`web3 eth_signTransaction was FAILED or REJECTED`)
		}

		if web3SignQueryResponse.TransactionStatus == "COMPLETED" {
			log.Infof("result :%v", web3SignQueryResponse)
			return web3SignQueryResponse.Transaction.SignedTransaction
		}
	}

	panic("can't get web3 eth_signTransaction sign result.")
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
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
