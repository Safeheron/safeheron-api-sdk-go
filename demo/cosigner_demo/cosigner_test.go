package cosigner_demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/cosigner"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var coSignerConverter cosigner.CoSignerConverter

func setup() {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %w", err))
	}

	coSignerConverter = cosigner.CoSignerConverter{Config: cosigner.CoSignerConfig{
		CoSignerPubKey:                    viper.GetString("coSignerPubKey"),
		ApprovalCallbackServicePrivateKey: viper.GetString("approvalCallbackServicePrivateKey"),
	}}

}

func teardown() {
}

func TestConvert(t *testing.T) {
	//The CoSignerCallBack received by the controller
	var coSignerCallBack cosigner.CoSignerCallBackV3
	coSignerBizContent, _ := coSignerConverter.RequestV3Convert(coSignerCallBack)
	//According to different types of CoSignerCallBack, the customer handles the corresponding type of business logic.
	log.Infof("coSignerBizContent: %s", coSignerBizContent)

	var coSignerResponse cosigner.CoSignerResponseV3
	//coSignerBizContent.ApprovalId
	coSignerResponse.ApprovalId = "approvalId"
	coSignerResponse.Action = "APPROVE"
	encryptResponse, _ := coSignerConverter.ResponseV3Converter(coSignerResponse)
	log.Infof("encryptResponse: %s", encryptResponse)
	//The customer returns encryptResponse after processing the business logic.
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
