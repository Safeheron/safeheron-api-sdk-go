package safeheron

type ApiConfig struct {
	BaseUrl               string `comment:"Safeheron Request Base URL"`
	ApiKey                string `comment:"api key, you can get from safeheron web console"`
	RsaPrivateKey         string `comment:"Your RSA private key"`
	SafeheronRsaPublicKey string `comment:"Api key's platform public key, you can get from safeheron web console"`
	RequestTimeout        int64  `comment:"RequestTimeout (Millisecond)"`
}
