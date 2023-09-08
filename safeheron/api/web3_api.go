package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type Web3Api struct {
	Client safeheron.Client
}

type CreateWeb3AccountRequest struct {
	AccountName string `json:"accountName,omitempty"`
	HiddenOnUI  bool   `json:"hiddenOnUI,omitempty"`
}

type CreateWeb3AccountResponse struct {
	AccountKey  string `json:"accountKey"`
	AccountName string `json:"accountName"`
	HiddenOnUI  bool   `json:"hiddenOnUI"`
	PubKeyList  []struct {
		SignAlg string `json:"signAlg"`
		PubKey  string `json:"pubKey"`
	} `json:"pubKeyList"`
	AddressList []struct {
		BlockchainType string `json:"blockchainType"`
		Address        string `json:"address"`
	} `json:"addressList"`
}

func (e *Web3Api) CreateWeb3Account(d CreateWeb3AccountRequest, r *CreateWeb3AccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/account/create")
}

type BatchCreateWeb3AccountRequest struct {
	AccountName string `json:"accountName,omitempty"`
	Count       int32  `json:"count"`
}

type BatchCreateWeb3AccountResponse []struct {
	AccountKey string `json:"accountKey"`
	PubKeyList []struct {
		SignAlg string `json:"signAlg"`
		PubKey  string `json:"pubKey"`
	} `json:"pubKeyList"`
	AddressList []struct {
		BlockchainType string `json:"blockchainType"`
		Address        string `json:"address"`
	} `json:"addressList"`
}

func (e *Web3Api) BatchCreateWeb3Account(d BatchCreateWeb3AccountRequest, r *BatchCreateWeb3AccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/batch/account/create")
}

type ListWeb3AccountRequest struct {
	Direct     string `json:"direct,omitempty"`
	Limit      int32  `json:"limit,omitempty"`
	FromId     string `json:"fromId,omitempty"`
	NamePrefix string `json:"namePrefix,omitempty"`
}

func (e *Web3Api) ListWeb3Accounts(d ListWeb3AccountRequest, r *[]CreateWeb3AccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/account/list")
}

type EthSignRequest struct {
	AccountKey    string `json:"accountKey"`
	CustomerRefId string `json:"customerRefId"`
	Note          string `json:"note,omitempty"`
	CustomerExt1  string `json:"customerExt1,omitempty"`
	CustomerExt2  string `json:"customerExt2,omitempty"`
	MessageHash   struct {
		ChainId int64    `json:"chainId"`
		Hash    []string `json:"hash"`
	} `json:"messageHash"`
}

func (e *Web3Api) EthSign(d EthSignRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/ethSign")
}

type PersonalSignRequest struct {
	AccountKey    string `json:"accountKey"`
	CustomerRefId string `json:"customerRefId"`
	Note          string `json:"note,omitempty"`
	CustomerExt1  string `json:"customerExt1,omitempty"`
	CustomerExt2  string `json:"customerExt2,omitempty"`
	Message       struct {
		ChainId int64  `json:"chainId"`
		Data    string `json:"data"`
	} `json:"message"`
}

func (e *Web3Api) PersonalSign(d PersonalSignRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/personalSign")
}

type EthSignTypedDataRequest struct {
	AccountKey    string `json:"accountKey"`
	CustomerRefId string `json:"customerRefId"`
	Note          string `json:"note,omitempty"`
	CustomerExt1  string `json:"customerExt1,omitempty"`
	CustomerExt2  string `json:"customerExt2,omitempty"`
	Message       struct {
		ChainId int64  `json:"chainId"`
		Data    string `json:"data"`
		Version string `json:"version"`
	} `json:"message"`
}

func (e *Web3Api) EthSignTypedData(d EthSignTypedDataRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/ethSignTypedData")
}

type EthSignTransactionRequest struct {
	AccountKey    string `json:"accountKey"`
	CustomerRefId string `json:"customerRefId"`
	Note          string `json:"note,omitempty"`
	CustomerExt1  string `json:"customerExt1,omitempty"`
	CustomerExt2  string `json:"customerExt2,omitempty"`
	Transaction   struct {
		To                   string `json:"to"`
		Value                string `json:"value"`
		ChainId              int64  `json:"chainId"`
		GasPrice             string `json:"gasPrice,omitempty"`
		GasLimit             int32  `json:"gasLimit"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
		MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
		Nonce                int64  `json:"nonce"`
		Data                 string `json:"data,omitempty"`
	} `json:"transaction"`
}

func (e *Web3Api) EthSignTransaction(d EthSignTransactionRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/ethSignTransaction")
}

type CancelWeb3SignRequest struct {
	TxKey string `json:"txKey"`
}

func (e *Web3Api) CancelWeb3Sign(d CancelWeb3SignRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/cancel")
}

type Web3SignQueryRequest struct {
	TxKey         string `json:"txKey,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

type Web3SignQueryResponse struct {
	TxKey                string `json:"txKey"`
	AccountKey           string `json:"accountKey,omitempty"`
	SourceAddress        string `json:"sourceAddress,omitempty"`
	TransactionStatus    string `json:"transactionStatus,omitempty"`
	TransactionSubStatus string `json:"transactionSubStatus,omitempty"`
	CreatedByUserKey     string `json:"createdByUserKey,omitempty"`
	CreatedByUserName    string `json:"createdByUserName,omitempty"`
	CreateTime           int64  `json:"createTime,omitempty"`
	AuditUserKey         string `json:"auditUserKey,omitempty"`
	AuditUserName        string `json:"auditUserName,omitempty"`
	CustomerRefId        string `json:"customerRefId"`
	Note                 string `json:"note,omitempty"`
	CustomerExt1         string `json:"customerExt1,omitempty"`
	CustomerExt2         string `json:"customerExt2,omitempty"`
	Balance              string `json:"balance,omitempty"`
	TokenBalance         string `json:"tokenBalance,omitempty"`
	Symbol               string `json:"symbol,omitempty"`
	TokenSymbol          string `json:"tokenSymbol,omitempty"`
	SubjectType          string `json:"subjectType,omitempty"`
	Transaction          struct {
		To                   string `json:"to,omitempty"`
		Value                string `json:"value,omitempty"`
		ChainId              int64  `json:"chainId,omitempty"`
		GasPrice             string `json:"gasPrice,omitempty"`
		GasLimit             int32  `json:"gasLimit,omitempty"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
		MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
		Nonce                int64  `json:"nonce,omitempty"`
		Data                 string `json:"data,omitempty"`
		TxHash               string `json:"txHash,omitempty"`
		SignedTransaction    string `json:"signedTransaction,omitempty"`
		Sig                  struct {
			Hash string `json:"hash,omitempty"`
			Sig  string `json:"sig,omitempty"`
		} `json:"sig,omitempty"`
	} `json:"transaction"`
	Message struct {
		ChainId int64  `json:"hash,omitempty"`
		Data    string `json:"data,omitempty"`
		Sig     struct {
			Hash string `json:"hash,omitempty"`
			Sig  string `json:"sig,omitempty"`
		} `json:"sig,omitempty"`
	} `json:"message,omitempty"`
	MessageHash struct {
		ChainId int64 `json:"hash,omitempty"`
		SigList []struct {
			Hash string `json:"hash,omitempty"`
			Sig  string `json:"sig,omitempty"`
		} `json:"sigList,omitempty"`
	} `json:"messageHash,omitempty"`
}

func (e *Web3Api) QueryWeb3Sig(d Web3SignQueryRequest, r *Web3SignQueryResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/one")
}

type ListWeb3SignRequest struct {
	Direct            string   `json:"direct,omitempty"`
	Limit             int32    `json:"limit,omitempty"`
	FromId            string   `json:"fromId,omitempty"`
	SubjectType       string   `json:"subjectType,omitempty"`
	TransactionStatus []string `json:"transactionStatus,omitempty"`
	AccountKey        string   `json:"accountKey,omitempty"`
	CreateTimeMin     int64    `json:"createTimeMin,omitempty"`
	CreateTimeMax     int64    `json:"createTimeMax,omitempty"`
}

func (e *Web3Api) ListWeb3Sign(d ListWeb3SignRequest, r *[]Web3SignQueryResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/list")
}
