package cosigner_demo

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var coSignerConverter CoSignerConverter

func setup() {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %w", err))
	}

	coSignerConverter = CoSignerConverter{Config: CoSignerConfig{
		ApiPubKey:  viper.GetString("apiPubKey"),
		BizPrivKey: viper.GetString("bizPrivKey"),
	}}

}

func teardown() {
}

func TestConvert(t *testing.T) {
	//The CoSignerCallBack received by the controller
	var coSignerCallBack CoSignerCallBack
	coSignerBizContent, _ := coSignerConverter.Convert(coSignerCallBack)
	//According to different types of CoSignerCallBack, the customer handles the corresponding type of business logic.
	log.Infof("coSignerBizContent: %s", coSignerBizContent)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
