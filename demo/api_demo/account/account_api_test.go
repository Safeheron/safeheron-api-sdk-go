package account_api_demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

var accountApi AccountApi

func TestListAccounts(t *testing.T) {
	req := ListAccountRequest{
		PageNumber: 1,
		PageSize:   10,
	}

	var res ListAccountResponse
	err := accountApi.ListAccounts(req, &res)
	assert.Nil(t, err)
	assert.Greater(t, len(res.Content), 0)

	for _, account := range res.Content {
		log.Infof("accountKey: %s, accountName: %s", account.AccountKey, account.AccountName)
	}

}

func TestCreateAccountAndAddCoin(t *testing.T) {
	createAccountRequest := CreateAccountRequest{
		AccountName: "first-wallet-account",
		HiddenOnUI:  true,
	}

	var createAccountResponse CreateAccountResponse
	if err := accountApi.CreateAccount(createAccountRequest, &createAccountResponse); err != nil {
		panic(fmt.Errorf("failed to create wallet account, %w", err))
	}

	log.Infof("wallet account created, account key: %s", createAccountResponse.AccountKey)

	addCoinRequest := AddCoinRequest{
		AccountKey: createAccountResponse.AccountKey,
		CoinKey:    "ETH_GOERLI",
	}

	var addCoinResponse AddCoinResponse

	if err := accountApi.AddCoin(addCoinRequest, &addCoinResponse); err != nil {
		panic(fmt.Errorf("failed to add coin in wallet, %w", err))
	}

	log.Infof("Token[ETH_GOERLI] address: %s", addCoinResponse[0].Address)

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

	accountApi = AccountApi{Client: sc}
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
