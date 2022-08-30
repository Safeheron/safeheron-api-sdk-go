package mpcdemo

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math"
	"math/big"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

// Replace with your infura api key
const ropsten = "https://ropsten.infura.io/v3/802******bc2fcb"

// Replace with your private key
const PRIVATE_KEY = "6b493b4*******ad01d357"
const READ_ONLY_FROM_ADDRESS = "0x0000000000000000000000000000000000000000"
const ERC20_CONTRACT_ADDRESS = "0x7fab42998149d35C03376b09D042220b6c7c778B"

var client *ethclient.Client
var privateKey *ecdsa.PrivateKey
var fromAddress common.Address
var eip1559Signer types.Signer
var chainID *big.Int
var tokenInstance *Token

func init() {
	client, _ = ethclient.Dial(ropsten)
	privateKey, _ = crypto.HexToECDSA(PRIVATE_KEY)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	chainId, _ := client.NetworkID(context.Background())
	eip1559Signer = types.NewLondonSigner(chainId)
	tokenInstance, _ = NewToken(common.HexToAddress(ERC20_CONTRACT_ADDRESS), client)
}

func SendEther(to string, value float64) {
	log.Infof("Attempting to send transaction from %s to %s", fromAddress, to)

	// Create transaction
	tx := createTransaction(value, to, nil)

	// Encode the transaction and compute the hash value
	hash := eip1559Signer.Hash(tx).Hex()

	// Sign with private key
	sig := getLocalSigWithPrivateKey(hash)

	// Or Sign with safeheron mpc
	// sig := getMpcSigFromSafeheron(hash[2:])
	log.Infof("sig: %s", sig)

	// Add 0x prefix to sig
	sigByte, _ := hexutil.Decode("0x" + sig[:])

	// Encode tx with sig
	signedTx, err := tx.WithSignature(eip1559Signer, sigByte)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction
	client.SendTransaction(context.Background(), signedTx)
	log.Infof("transaction hash: %s", signedTx.Hash().Hex())

}

func SendERC20Token(to string, value float64) {
	log.Infof("Attempting to send erc20 token transaction from %s to %s, contract address: %s",
		fromAddress, to, ERC20_CONTRACT_ADDRESS)

	// Create contract function data
	functionData := createTransferData(value, to)

	// Create transaction
	tx := createTransaction(0, ERC20_CONTRACT_ADDRESS, functionData)

	// Encode the transaction and compute the hash value
	hash := eip1559Signer.Hash(tx).Hex()

	// Sign with private key
	sig := getLocalSigWithPrivateKey(hash)

	// Or Sign with safeheron mpc
	// sig := getMpcSigFromSafeheron(hash[2:])
	log.Infof("sig: %s", sig)

	// Add 0x prefix to sig
	sigByte, _ := hexutil.Decode("0x" + sig[:])

	// Encode tx with sig
	signedTx, err := tx.WithSignature(eip1559Signer, sigByte)
	if err != nil {
		log.Fatal(err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("transaction hash: %s", signedTx.Hash().Hex())
}

func createTransaction(value float64, to string, data []byte) *types.Transaction {
	toAddress := common.HexToAddress(to)

	// Get data from block chain: nonce, gasPrice
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	// gasPrice, _ := client.SuggestGasPrice(context.Background())

	// Estimate gas
	gasLimit, _ := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Data: data,
	})

	transferAmount := ethUnit.NewEther(big.NewFloat(value)).Wei()

	// Estimate maxFeePerGas, we assume maxPriorityFeePerGas's value is 2(gwei).
	// The baseFeePerGas is recommended to be 1.25 times the latest block's baseFeePerGas value.
	// maxFeePerGas must not less than baseFeePerGas + maxPriorityFeePerGas
	maxPriorityFeePerGas := ethUnit.NewGWei(big.NewFloat(2)).Wei()
	lastBlockHeader, _ := client.HeaderByNumber(context.Background(), nil)
	baseFeeFloat := big.NewFloat(0).SetInt(lastBlockHeader.BaseFee)
	suggestBaseFee := big.NewFloat(0).Mul(baseFeeFloat, big.NewFloat(1.25))
	maxFeePerGas, _ := big.NewFloat(0).Add(suggestBaseFee, big.NewFloat(0).SetInt(maxPriorityFeePerGas)).Int(nil)

	// Create raw transaction
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &toAddress,
		Value:     transferAmount,
		Data:      data,
		Gas:       gasLimit * 2,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
	})

	jsonTx, _ := json.Marshal(tx)
	log.Infof("tx json data: %s", string(jsonTx))

	return tx
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

func getLocalSigWithPrivateKey(hash string) string {
	// Sign the transaction
	unsignedHash, _ := hexutil.Decode(hash)
	signature, _ := crypto.Sign(unsignedHash, privateKey)
	// Return the signature without 0x prefix, the purpose of doing
	// this is simply to be consistent with safeheron mpc and explain it better
	return hexutil.Encode(signature)[2:]
}

func getMpcSigFromSafeheron(hash string) string {
	// Get sig's value from safeheron with hash(32-byte hex string without '0x' prefix).
	// 1. Request "v1/transactions/mpcsign/create" api to sign then waiting for completed.
	// 2. Get sig from webhook or "v1/transactions/mpcsign/one" api

	// This is an example.
	// The value of sig consists of 32 bytes r + 32 bytes s + 1 byte v
	sig := "f7fdd1c4d0e2d90a0c2dc*********4566d9578daa0e532d571b"
	return sig
}
