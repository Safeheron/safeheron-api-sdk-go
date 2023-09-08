package mpc_demo

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"
	"time"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/api"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/sha3"
)

var mpcSignApi api.MpcSignApi

const READ_ONLY_FROM_ADDRESS = "0x0000000000000000000000000000000000000000"

var client *ethclient.Client
var signer types.Signer
var chainID *big.Int
var tokenInstance *Token

func sendERC20Token(accountKey string, accountTokenAddress string, erc20ContractAddress string, toAddress string, value float64) {
	log.Infof("Attempting to send erc20 token transaction from %s to %s, contract address: %s",
		accountTokenAddress, toAddress, erc20ContractAddress)

	// Create contract function data
	functionData := createTransferData(value, toAddress)

	// Create transaction
	tx := createTransaction(common.HexToAddress(accountTokenAddress), 0, erc20ContractAddress, functionData)

	// Encode the transaction and compute the hash value
	hash := signer.Hash(tx).Hex()

	// Sign with safeheron mpc
	customerRefId := uuid.NewString()
	txKey := requestMpcSig(customerRefId, accountKey, hash)
	log.Infof("transaction created, txKey: %s", txKey)

	// Get sig
	mpcSig := retrieveSig(customerRefId)
	log.Infof("got mpc sign result, sig: %s", mpcSig)

	// Add 0x prefix to sig
	sigByte, _ := hexutil.Decode("0x" + mpcSig[:])

	// Encode tx with sig
	signedTx, err := tx.WithSignature(signer, sigByte)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Transaction successful with hash: %s", signedTx.Hash().Hex())
	log.Infof("You can view the transaction at: https://goerli.etherscan.io/tx/%s", signedTx.Hash().Hex())

}

func createTransaction(fromAddress common.Address, value float64, to string, data []byte) *types.Transaction {
	toAddress := common.HexToAddress(to)

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
		ChainID:   chainID,
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

func requestMpcSig(customerRefId string, accountKey string, dataList string) string {
	createMpcSignRequest := api.CreateMpcSignRequest{
		CustomerRefId:    customerRefId,
		SourceAccountKey: accountKey,
		SignAlg:          "Secp256k1",
		DataList: []struct {
			Data string `json:"data,omitempty"`
			Note string `json:"note,omitempty"`
		}{{Data: dataList[2:]}},
	}

	var createMpcSignResponse api.CreateMpcSignResponse
	if err := mpcSignApi.CreateMpcSign(createMpcSignRequest, &createMpcSignResponse); err != nil {
		panic(err)
	}

	return createMpcSignResponse.TxKey
}

func retrieveSig(customerRefId string) string {
	// As an example, we use the API to get the signature, and you can also use webhooks to get it
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	OneMPCSignTransactionsRequest := api.OneMPCSignTransactionsRequest{
		CustomerRefId: customerRefId,
	}
	var MPCSignTransactionsResponse api.MPCSignTransactionsResponse
	for range ticker.C {
		if err := mpcSignApi.OneMPCSignTransactions(OneMPCSignTransactionsRequest, &MPCSignTransactionsResponse); err != nil {
			panic(err)
		}

		log.Infof(`mpc sign transaction status: %s, sub status: %s`, MPCSignTransactionsResponse.TransactionStatus, MPCSignTransactionsResponse.TransactionSubStatus)

		if MPCSignTransactionsResponse.TransactionStatus == "FAILED" || MPCSignTransactionsResponse.TransactionStatus == "REJECTED" {
			panic(`mpc sign transaction was FAILED or REJECTED`)
		}

		if MPCSignTransactionsResponse.TransactionStatus == "COMPLETED" && MPCSignTransactionsResponse.TransactionSubStatus == "CONFIRMED" {
			return MPCSignTransactionsResponse.DataList[0].Sig
		}
	}

	panic("can't get sig.")
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

	mpcSignApi = api.MpcSignApi{Client: sc}

	client, _ = ethclient.Dial(viper.GetString("ethereumRpcApi"))
	chainId, _ := client.NetworkID(context.Background())
	signer = types.NewLondonSigner(chainId)
	tokenInstance, _ = NewToken(common.HexToAddress(viper.GetString("erc20ContractAddress")), client)
}

func teardown() {
}

func TestMpcSgin(t *testing.T) {
	sendERC20Token(viper.GetString("accountKey"), viper.GetString("accountTokenAddress"), viper.GetString("erc20ContractAddress"), viper.GetString("toAddress"), 1)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
