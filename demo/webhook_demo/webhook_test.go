package webhook_demo

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var webhookConverter WebhookConverter

func setup() {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %w", err))
	}

	webhookConverter = WebhookConverter{Config: WebHookConfig{
		SafeheronWebHookRsaPublicKey: viper.GetString("safeheronWebHookRsaPublicKey"),
		WebHookRsaPrivateKey:         viper.GetString("webHookRsaPrivateKey"),
	}}

}

func teardown() {
}

func TestConvert(t *testing.T) {
	//The webHook received by the controller
	var webHook WebHook
	webHookBizContent, _ := webhookConverter.Convert(webHook)
	//According to different types of WebHook, the customer handles the corresponding type of business logic.
	log.Infof("webHookBizContent: %s", webHookBizContent)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
