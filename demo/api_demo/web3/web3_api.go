package web3_demo

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
	AccountKey string `json:"accountKey,omitempty"`
	PubKeyList []struct {
		SignAlg string `json:"signAlg,omitempty"`
		PubKey  string `json:"pubKey,omitempty"`
	} `json:"pubKeyList"`
	AddressList []struct {
		BlockchainType string `json:"blockchainType,omitempty"`
		Address        string `json:"address,omitempty"`
	} `json:"addressList"`
}

func (e *Web3Api) CreateWeb3Account(d CreateWeb3AccountRequest, r *CreateWeb3AccountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/account/create")
}

type EthSignTransactionRequest struct {
	AccountKey    string `json:"accountKey,omitempty"`
	CustomerRefId string `json:"customerRefId"`
	Note          string `json:"note,omitempty"`
	Transaction   struct {
		To                   string `json:"to,omitempty"`
		Value                string `json:"value,omitempty"`
		ChainId              string `json:"chainId,omitempty"`
		GasPrice             string `json:"gasPrice,omitempty"`
		GasLimit             string `json:"gasLimit,omitempty"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
		MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
		Nonce                string `json:"nonce,omitempty"`
		Data                 string `json:"data,omitempty"`
	} `json:"transaction"`
}

type EthSignTransactionResponse struct {
	TxKey string `json:"txKey"`
}

func (e *Web3Api) EthSignTransaction(d EthSignTransactionRequest, r *EthSignTransactionResponse) error {
	return e.Client.SendRequest(d, r, "/v1/web3/sign/ethSignTransaction")
}

type Web3SignQueryRequest struct {
	TxKey         string `json:"txKey"`
	CustomerRefId string `json:"customerRefId"`
}

type Web3SignQueryResponse struct {
	TxKey                string `json:"txKey"`
	AccountKey           string `json:"accountKey,omitempty"`
	CustomerRefId        string `json:"customerRefId"`
	SourceAddress        string `json:"sourceAddress,omitempty"`
	TransactionStatus    string `json:"transactionStatus,omitempty"`
	TransactionSubStatus string `json:"transactionSubStatus,omitempty"`
	CreatedByUserKey     string `json:"createdByUserKey,omitempty"`
	CreatedByUserName    string `json:"createdByUserName,omitempty"`
	CreateTime           int64  `json:"createTime,omitempty"`
	AuditUserKey         string `json:"auditUserKey,omitempty"`
	AuditUserName        string `json:"auditUserName,omitempty"`
	Note                 string `json:"note,omitempty"`
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
		GasLimit             int64  `json:"gasLimit,omitempty"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
		MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
		Nonce                int64  `json:"nonce,omitempty"`
		Data                 string `json:"data,omitempty"`
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
