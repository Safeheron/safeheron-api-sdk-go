package webhook_demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/webhook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var webhookConverter webhook.WebhookConverter

func setup() {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %w", err))
	}

	webhookConverter = webhook.WebhookConverter{Config: webhook.WebHookConfig{
		SafeheronWebHookRsaPublicKey: viper.GetString("safeheronWebHookRsaPublicKey"),
		WebHookRsaPrivateKey:         viper.GetString("webHookRsaPrivateKey"),
	}}

}

func teardown() {
}

func TestConvert(t *testing.T) {
	//The webHook received by the controller
	var webHook webhook.WebHook
	webHookBizContent, _ := webhookConverter.Convert(webHook)
	//According to different types of WebHook, the customer handles the corresponding type of business logic.
	log.Infof("webHookBizContent: %s", webHookBizContent)

	var webHookResponse webhook.WebHookResponse
	webHookResponse.Code = "200"
	webHookResponse.Message = "SUCCESS"
	//The customer returns WebHookResponse after processing the business logic.
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
