package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type AccountApi struct {
	Client safeheron.Client
}

type ListAccountRequest struct {
	PageNumber    int    `json:"pageNumber,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	HiddenOnUI    *bool  `json:"hiddenOnUI,omitempty"`
	NamePrefix    string `json:"namePrefix,omitempty"`
	NameSuffix    string `json:"nameSuffix,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

type AccountResponse struct {
	AccountKey    string `json:"accountKey"`
	CustomerRefId string `json:"customerRefId"`
	AccountName   string `json:"accountName"`
	AccountIndex  int32  `json:"accountIndex"`
	AccountType   string `json:"accountType"`
	AccountTag    string `json:"accountTag"`
	HiddenOnUI    bool   `json:"hiddenOnUI"`
	UsdBalance    string `json:"usdBalance"`
	PubKeys       []struct {
		SignAlg string `json:"signAlg"`
		PubKey  string `json:"pubKey"`
	} `json:"pubKeys"`
}

type ListAccountResponse struct {
	PageNumber    int32             `json:"pageNumber"`
	PageSize      int32             `json:"pageSize"`
	TotalElements int64             `json:"totalElements"`
	Content       []AccountResponse `json:"content"`
}

func (e *AccountApi) ListAccounts(d ListAccountRequest, r *ListAccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/list")
}

type OneAccountRequest struct {
	AccountKey    string `json:"accountKey,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

func (e *AccountApi) OneAccounts(d OneAccountRequest, r *AccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/one")
}

type CreateAccountRequest struct {
	AccountName   string   `json:"accountName,omitempty"`
	CustomerRefId string   `json:"customerRefId,omitempty"`
	HiddenOnUI    bool     `json:"hiddenOnUI,omitempty"`
	AccountTag    string   `json:"accountTag,omitempty"`
	CoinKeyList   []string `json:"coinKeyList,omitempty"`
}

type CreateAccountResponse struct {
	AccountKey string `json:"accountKey"`
	PubKeys    []struct {
		SignAlg string `json:"signAlg"`
		PubKey  string `json:"pubKey"`
	} `json:"pubKeys"`
	CoinAddressList []struct {
		CoinKey          string `json:"coinKey"`
		AddressGroupKey  string `json:"addressGroupKey"`
		AddressGroupName string `json:"addressGroupName"`
		AddressList      []struct {
			Address     string `json:"address"`
			AddressType string `json:"addressType"`
			DerivePath  string `json:"derivePath"`
		} `json:"addressList"`
	} `json:"coinAddressList"`
}

func (e *AccountApi) CreateAccount(d CreateAccountRequest, r *CreateAccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/create")
}

type BatchCreateAccountRequest struct {
	AccountName string `json:"accountName,omitempty"`
	HiddenOnUI  bool   `json:"hiddenOnUI,omitempty"`
	Count       int32  `json:"count"`
	AccountTag  string `json:"accountTag,omitempty"`
}

type BatchCreateAccountResponse struct {
	AccountKeyList []string `json:"accountKeyList"`
}

func (e *AccountApi) BatchCreateAccount(d BatchCreateAccountRequest, r *BatchCreateAccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/batch/create")
}

func (e *AccountApi) BatchCreateAccountV2(d BatchCreateAccountRequest, r *[]CreateAccountResponse) error {
	return e.Client.SendRequest(d, r, "/v2/account/batch/create")
}

type UpdateAccountShowStateRequest struct {
	AccountKey string `json:"accountKey"`
	HiddenOnUI bool   `json:"hiddenOnUI"`
}

type ResultResponse struct {
	Result bool `json:"result"`
}

func (e *AccountApi) UpdateAccountShowState(d UpdateAccountShowStateRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/update/show/state")
}

type BatchUpdateAccountTagRequest struct {
	AccountKeyList []string `json:"accountKeyList"`
	AccountTag     string   `json:"accountTag"`
}

func (e *AccountApi) batchUpdateAccountTag(d BatchUpdateAccountTagRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/batch/update/tag")
}

type AddCoinRequest struct {
	CoinKey    string `json:"coinKey,omitempty"`
	AccountKey string `json:"accountKey,omitempty"`
}

type AddCoinResponse []struct {
	Address     string `json:"address"`
	AddressType string `json:"addressType"`
	DerivePath  string `json:"derivePath"`
}

func (e *AccountApi) AddCoin(d AddCoinRequest, r *AddCoinResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/create")
}

type AddCoinV2Request struct {
	CoinKeyList []string `json:"coinKeyList,omitempty"`
	AccountKey  string   `json:"accountKey,omitempty"`
}

type AddCoinV2Response struct {
	AccountKey      string `json:"accountKey"`
	CoinAddressList []struct {
		CoinKey          string `json:"coinKey"`
		AddressGroupKey  string `json:"addressGroupKey"`
		AddressGroupName string `json:"addressGroupName"`
		AddressList      []struct {
			Address     string `json:"address"`
			AddressType string `json:"addressType"`
			DerivePath  string `json:"derivePath"`
		} `json:"addressList"`
	} `json:"coinAddressList"`
}

func (e *AccountApi) AddCoinV2(d AddCoinV2Request, r *AddCoinV2Response) error {
	return e.Client.SendRequest(d, r, "/v2/account/coin/create")
}

type BatchCreateAccountCoinRequest struct {
	CoinKey          string   `json:"coinKey"`
	AccountKeyList   []string `json:"accountKeyList"`
	AddressGroupName string   `json:"addressGroupName,omitempty"`
}

type BatchCreateAccountCoinResponse []struct {
	AccountKey       string `json:"accountKey"`
	AddressGroupKey  string `json:"addressGroupKey"`
	AddressGroupName string `json:"addressGroupName"`
	AddressList      []struct {
		Address     string `json:"address"`
		AddressType string `json:"addressType"`
		DerivePath  string `json:"derivePath"`
	} `json:"addressList"`
}

func (e *AccountApi) BatchCreateAccountCoin(d BatchCreateAccountCoinRequest, r *BatchCreateAccountCoinResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/batch/coin/create")
}

type ListAccountCoinRequest struct {
	AccountKey string `json:"accountKey"`
}

type AccountCoinResponse []struct {
	CoinKey           string `json:"coinKey"`
	CoinFullName      string `json:"coinFullName"`
	CoinName          string `json:"coinName"`
	CoinDecimal       int32  `json:"coinDecimal"`
	TxRefUrl          string `json:"txRefUrl"`
	AddressRefUrl     string `json:"addressRefUrl"`
	LogoUrl           string `json:"logoUrl"`
	Symbol            string `json:"symbol"`
	IsMultipleAddress string `json:"isMultipleAddress"`
	FeeCoinKey        string `json:"feeCoinKey"`
	FeeUnit           string `json:"feeUnit"`
	FeeDecimal        int32  `json:"feeDecimal"`
	ShowCoinDecimal   int32  `json:"showCoinDecimal"`
	Balance           string `json:"balance"`
	UsdBalance        string `json:"usdBalance"`
	AddressList       []struct {
		Address        string `json:"address"`
		AddressType    string `json:"addressType"`
		DerivePath     string `json:"derivePath"`
		AddressBalance string `json:"addressBalance"`
	} `json:"addressList"`
}

func (e *AccountApi) ListAccountCoin(d ListAccountCoinRequest, r *AccountCoinResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/list")
}

type ListAccountCoinAddressRequest struct {
	PageNumber    int    `json:"pageNumber,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	CoinKey       string `json:"coinKey"`
	AccountKey    string `json:"accountKey"`
	CustomerRefId string `json:"customerRefId"`
}

type AccountCoinAddressResponse struct {
	PageNumber    int32 `json:"pageNumber"`
	PageSize      int32 `json:"pageSize"`
	TotalElements int64 `json:"totalElements"`
	Content       []struct {
		AddressGroupKey  string `json:"addressGroupKey"`
		AddressGroupName string `json:"addressGroupName"`
		CustomerRefId    string `json:"customerRefId"`
		AddressList      []struct {
			Address        string `json:"address"`
			AddressType    string `json:"addressType"`
			DerivePath     string `json:"derivePath"`
			AddressBalance string `json:"addressBalance"`
		} `json:"addressList"`
	} `json:"content"`
}

func (e *AccountApi) ListAccountCoinAddress(d ListAccountCoinAddressRequest, r *AccountCoinAddressResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/address/list")
}

type InfoAccountCoinAddressRequest struct {
	CoinKey string `json:"coinKey"`
	Address string `json:"address"`
}

type InfoAccountCoinAddressResponse struct {
	Address        string `json:"address"`
	AddressType    string `json:"addressType"`
	DerivePath     string `json:"derivePath"`
	AddressBalance string `json:"addressBalance"`
	AccountKey     string `json:"accountKey"`
}

func (e *AccountApi) InfoAccountCoinAddress(d InfoAccountCoinAddressRequest, r *InfoAccountCoinAddressResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/address/info")
}

type RenameAccountCoinAddressRequest struct {
	AddressGroupKey  string `json:"addressGroupKey"`
	AddressGroupName string `json:"addressGroupName"`
}

func (e *AccountApi) RenameAccountCoinAddress(d RenameAccountCoinAddressRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/address/name")
}

type CreateAccountCoinAddressRequest struct {
	CoinKey          string `json:"coinKey"`
	AccountKey       string `json:"accountKey"`
	AddressGroupName string `json:"addressGroupName"`
	CustomerRefId    string `json:"customerRefId"`
}

type CreateAccountCoinAddressResponse struct {
	Address     string `json:"address"`
	AddressType string `json:"addressType"`
	DerivePath  string `json:"derivePath"`
}

func (e *AccountApi) CreateAccountCoinAddress(d CreateAccountCoinAddressRequest, r *CreateAccountCoinAddressResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/address/create")
}

type CreateAccountCoinAddressV2Response struct {
	AddressGroupKey  string `json:"addressGroupKey"`
	AddressGroupName string `json:"addressGroupName"`
	AddressList      []struct {
		Address        string `json:"address"`
		AddressType    string `json:"addressType"`
		DerivePath     string `json:"derivePath"`
		AddressBalance string `json:"addressBalance"`
	} `json:"addressList"`
}

func (e *AccountApi) CreateAccountCoinAddressV2(d CreateAccountCoinAddressRequest, r *CreateAccountCoinAddressV2Response) error {
	return e.Client.SendRequest(d, r, "/v2/account/coin/address/create")
}

type BatchCreateAccountCoinUTXORequest struct {
	CoinKey          string `json:"coinKey"`
	AccountKey       string `json:"accountKey"`
	Count            int32  `json:"count"`
	AddressGroupName string `json:"addressGroupName"`
}

type BatchCreateAccountCoinUTXOResponse []struct {
	AccountKey       string `json:"accountKey"`
	AddressGroupKey  string `json:"addressGroupKey"`
	AddressGroupName string `json:"addressGroupName"`
	AddressList      []struct {
		Address     string `json:"address"`
		AddressType string `json:"addressType"`
		DerivePath  string `json:"derivePath"`
	} `json:"addressList"`
}

func (e *AccountApi) BatchCreateAccountCoinUTXO(d BatchCreateAccountCoinUTXORequest, r *BatchCreateAccountCoinUTXOResponse) error {
	return e.Client.SendRequest(d, r, "/v1/account/coin/utxo/batch/create")
}
