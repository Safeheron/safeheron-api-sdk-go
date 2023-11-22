# Go SDK for Safeheron API

![GitHub last commit](https://img.shields.io/github/last-commit/Safeheron/safeheron-api-sdk-go)
![GitHub top language](https://img.shields.io/github/languages/top/Safeheron/safeheron-api-sdk-go?color=red)

# API Documentation
- [Official documentation](https://docs.safeheron.com/api/index.html)

# Installation

safeheron-api-sdk-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/Safeheron/safeheron-api-sdk-go
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/Safeheron/safeheron-api-sdk-go/safeheron"
```

and run `go get` without parameters.

# Usage
```go
import "github.com/Safeheron/safeheron-api-sdk-go/safeheron"
```

> Take [`/v1/account/create`](https://docs.safeheron.com/api/index.html) as an example to explain, the complete code can be found in `demo/api_demo` directory

* Construct `safeheron.ApiConfig` 
    ```go
    // You can get `ApiKey` and `SafeheronRsaPublicKey` from Safeheron Web Console: https://www.safeheron.com/console.
    sc := safeheron.Client{Config: safeheron.ApiConfig{
            BaseUrl:               "https://api.safeheron.vip",
            ApiKey:                "d1ad6*****a572e7",
            RsaPrivateKey:         "pems/my_private.pem",
            SafeheronRsaPublicKey: "pems/safeheron_public.pem",
    }}
    ```
* Call `CreateAccount` api with `sc`
    ```go

    accountApi := api.AccountApi{Client: sc}

    req := api.CreateAccountRequest{
		AccountName: "first-wallet-account",
		HiddenOnUI:  true,
	}

    var res api.CreateAccountResponse
    err := accountApi.CreateAccount(req, &res)
    if err != nil {
        // Your code to process err
    }

    // Your code to process response
    ...
    ...
    ```

# Test

## Test Create Wallet Account
* Before run the test code, modify `demo/api_demo/account/config.yaml.example` according to the comments
    ```yaml
    # Your API key, you can get it from Safeheron Web Console: https://www.safeheron.com/console.
    apiKey: 080db****6e60
    # path to your private key file, pem encoded
    privateKeyPemFile: /path/to/your/privatekey.pem
    # path to Safeheron API public key file, pem encoded, you can get it from Safeheron Web Console: https://www.safeheron.com/console.
    # You can convert the Safeheron API Public Key obtained from Safeheron Web Console to PEM format with openssl: 
    # echo MIICIjANBgkqhk*****UbNkcCAwEAAQ== | base64 -d | openssl rsa -pubin -inform DER -outform PEM -out /path/to/safeheron/api/publickey.pem
    safeheronPublicKeyPemFile: /path/to/safeheron/api/publickey.pem
    # Safeheron API base url
    baseUrl: https://api.safeheron.vip
    ```
* Run the test
    ```bash
    $ cd demo/api_demo/account
    $ cp config.yaml.example config.yaml
    $ go test -run TestCreateAccountAndAddCoin
    ```

## Test Send A Transaction
* Before run the test code, modify `demo/api_demo/transaction/config.yaml.example` according to the comments
    ```yaml
    # Your API key, you can get it from Safeheron Web Console: https://www.safeheron.com/console.
    apiKey: 080db****6e60
    # path to your private key file, pem encoded
    privateKeyPemFile: /path/to/your/privatekey.pem
    # path to Safeheron API public key file, pem encoded, you can get it from Safeheron Web Console: https://www.safeheron.com/console.
    # You can convert the Safeheron API Public Key obtained from Safeheron Web Console to PEM format with openssl: 
    # echo MIICIjANBgkqhk*****UbNkcCAwEAAQ== | base64 -d | openssl rsa -pubin -inform DER -outform PEM -out /path/to/safeheron/api/publickey.pem
    safeheronPublicKeyPemFile: /path/to/safeheron/api/publickey.pem
    # Safeheron API base url
    baseUrl: https://api.safeheron.vip
    # Wallet Account key
    accountKey: account****5ecad40
    # To address
    destinationAddress: 0x9437A****0BF95f5
    ```
* Run the test
    ```bash
    $ cd demo/api_demo/transaction
    $ cp config.yaml.example config.yaml
    $ go test -run TestSendTransaction
    ```

## Test MPC Sign
* Before run the test code, modify `demo/mpc_demo/config.yaml.example` according to the comments
    ```yaml
    # Your API key, you can get it from Safeheron Web Console: https://www.safeheron.com/console.
    apiKey: 080db****6e60
    # path to your private key file, pem encoded
    privateKeyPemFile: /path/to/your/privatekey.pem
    # path to Safeheron API public key file, pem encoded, you can get it from Safeheron Web Console: https://www.safeheron.com/console.
    # You can convert the Safeheron API Public Key obtained from Safeheron Web Console to PEM format with openssl: 
    # echo MIICIjANBgkqhk*****UbNkcCAwEAAQ== | base64 -d | openssl rsa -pubin -inform DER -outform PEM -out /path/to/safeheron/api/publickey.pem
    safeheronPublicKeyPemFile: /path/to/safeheron/api/publickey.pem
    # Safeheron API base url
    baseUrl: https://api.safeheron.vip
    # Wallet Account key
    accountKey: account****5ecad40
    # Goerli testnet token address in wallet account
    accountTokenAddress: 0x970****4ffD59
    # erc20 token contract address
    erc20ContractAddress: 0x078****Eaa37F
    # address to receive token
    toAddress: 0x53B****321789
    # Ethereum RPC API
    ethereumRpcApi: https://goerli.infura.io/v3/802******bc2fcb
    ```

* Run the test
    ```bash
    $ cd demo/mpc_demo
    $ cp config.yaml.example config.yaml
    $ go test -run TestMpcSgin
    ```
