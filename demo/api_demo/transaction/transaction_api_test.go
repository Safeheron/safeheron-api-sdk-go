package transaction_api_demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var transactionApi TransactionApi

func TestSendTransaction(t *testing.T) {
	createTransactionRequest := CreateTransactionRequest{
		SourceAccountKey:       viper.GetString("accountKey"),
		SourceAccountType:      "VAULT_ACCOUNT",
		DestinationAccountType: "ONE_TIME_ADDRESS",
		DestinationAddress:     viper.GetString("destinationAddress"),
		CoinKey:                "ETH_GOERLI",
		TxAmount:               "0.001",
		TxFeeLevel:             "MIDDLE",
		CustomerRefId:          uuid.New().String(),
	}

	var createTransactionResponse CreateTransactionResponse
	if err := transactionApi.SendTransaction(createTransactionRequest, &createTransactionResponse); err != nil {
		panic(fmt.Errorf("failed to send transaction, %w", err))
	}

	log.Infof("transaction has been created, txKey: %s", createTransactionResponse.TxKey)

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

	transactionApi = TransactionApi{Client: sc}
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
