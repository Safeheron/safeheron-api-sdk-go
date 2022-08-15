package api_demo

import (
	"os"
	"testing"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"

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

func TestCreateAccount(t *testing.T) {
	req := CreateAccountRequest{
		HiddenOnUI: true,
	}

	var res CreateAccountResponse
	err := accountApi.CreateAccount(req, &res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	log.Infof("created, account key: %s", res.AccountKey)

}

func setup() {
	sc := safeheron.Client{Config: safeheron.ApiConfig{
		BaseUrl:               "https://api.safeheron.vip",
		ApiKey:                "d1ad6a******1ba572e7",
		RsaPrivateKey:         "pems/my_private.pem",
		SafeheronRsaPublicKey: "pems/safeheron_public.pem",
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
