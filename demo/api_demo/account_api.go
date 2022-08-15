package api_demo

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type AccountApi struct {
	Client safeheron.Client
}

const listpath = "/v1/account/list"

type ListAccountRequest struct {
	PageNumber int    `json:"pageNumber,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	HiddenOnUI bool   `json:"hiddenOnUI,omitempty"`
	NamePrefix string `json:"namePrefix,omitempty"`
	NameSuffix string `json:"nameSuffix,omitempty"`
}

type ListAccountResponse struct {
	PageNumber    int32 `json:"pageNumber"`
	PageSize      int32 `json:"pageSize"`
	TotalElements int64 `json:"totalElements"`
	Content       []struct {
		AccountKey        string `json:"accountKey"`
		AccountName       string `json:"accountName"`
		AccountIndex      int32  `json:"accountIndex"`
		AccountType       string `json:"accountType"`
		HiddenOnUI        bool   `json:"hiddenOnUI"`
		UsdBalance        string `json:"usdBalance"`
		FrozenUsdBalance  string `json:"frozenUsdBalance"`
		AmlLockUsdBalance string `json:"amlLockUsdBalance"`
		PubKeys           []struct {
			SignAlg string `json:"signAlg"`
			PubKey  string `json:"pubKey"`
		} `json:"pubKeys"`
	} `json:"content"`
}

func (e *AccountApi) ListAccounts(d ListAccountRequest, r *ListAccountResponse) error {
	return e.Client.SendRequest(d, r, listpath)
}

const createpath = "/v1/account/create"

type CreateAccountRequest struct {
	AccountName string `json:"accountName,omitempty"`
	HiddenOnUI  bool   `json:"hiddenOnUI,omitempty"`
}

type CreateAccountResponse struct {
	AccountKey string `json:"accountKey"`
	PubKeys    []struct {
		SignAlg string `json:"signAlg"`
		PubKey  string `json:"pubKey"`
	} `json:"pubKeys"`
}

func (e *AccountApi) CreateAccount(d CreateAccountRequest, r *CreateAccountResponse) error {
	return e.Client.SendRequest(d, r, createpath)
}
