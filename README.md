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

* Define request parameter data object
    ```go
    type CreateAccountRequest struct {
        AccountName string `json:"accountName,omitempty"`
        HiddenOnUI  bool   `json:"hiddenOnUI,omitempty"`
    }
    ```

* Define response data object
    ```go
    type CreateAccountResponse struct {
        AccountKey string `json:"accountKey"`
        PubKeys    []struct {
            SignAlg string `json:"signAlg"`
            PubKey  string `json:"pubKey"`
        } `json:"pubKeys"`
    }
    ```

* Define the interface to use
    ```go
    type AccountApi struct {
        Client safeheron.Client
    }

    const createpath = "/v1/account/create"

    func (e *AccountApi) CreateAccount(d CreateAccountRequest, r *CreateAccountResponse) error {
        return e.Client.SendRequest(d, r, createpath)
    }

    // Other interfaces
    ```
* Construct `safeheron.ApiConfig` 
    ```go
    sc := safeheron.Client{Config: safeheron.ApiConfig{
            BaseUrl:               "https://api.safeheron.vip",
            ApiKey:                "d1ad6*****a572e7",
            RsaPrivateKey:         "pems/my_private.pem",
            SafeheronRsaPublicKey: "pems/safeheron_public.pem",
    }}
    ```
* Call `CreateAccount` api with `sc`
    ```go
    accountApi = AccountApi{Client: sc}

    req := CreateAccountRequest{
            HiddenOnUI: true,
    }

    var res CreateAccountResponse
    err := accountApi.CreateAccount(req, &res)
    if err != nil {

    }

    // Your code to process response
    ...
    ...
    ```
