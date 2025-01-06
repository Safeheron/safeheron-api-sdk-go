package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type TransactionApi struct {
	Client safeheron.Client
}

type TransactionsRequest struct {
	TxKey                      string               `json:"txKey"`
	TxHash                     string               `json:"txHash"`
	CoinKey                    string               `json:"coinKey"`
	TxAmount                   string               `json:"txAmount"`
	SourceAccountKey           string               `json:"sourceAccountKey"`
	SourceAccountType          string               `json:"sourceAccountType"`
	SourceAddress              string               `json:"sourceAddress"`
	SourceAddressList          []SourceAddress      `json:"sourceAddressList"`
	DestinationAccountKey      string               `json:"destinationAccountKey"`
	DestinationAccountType     string               `json:"destinationAccountType"`
	DestinationAddress         string               `json:"destinationAddress"`
	Memo                       string               `json:"memo"`
	DestinationAddressList     []DestinationAddress `json:"destinationAddressList"`
	DestinationTag             string               `json:"destinationTag"`
	TransactionType            string               `json:"transactionType"`
	TransactionStatus          string               `json:"transactionStatus"`
	TransactionSubStatus       string               `json:"transactionSubStatus"`
	CreateTime                 int64                `json:"createTime"`
	Note                       string               `json:"note"`
	AuditUserKey               string               `json:"auditUserKey"`
	CreatedByUserKey           string               `json:"createdByUserKey"`
	TxFee                      string               `json:"txFee"`
	FeeCoinKey                 string               `json:"feeCoinKey"`
	ReplaceTxHash              string               `json:"replaceTxHash"`
	CustomerRefId              string               `json:"customerRefId"`
	CustomerExt1               string               `json:"customerExt1"`
	CustomerExt2               string               `json:"customerExt2"`
	AmlLock                    string               `json:"amlLock"`
	BlockHeight                int64                `json:"blockHeight"`
	CompletedTime              int64                `json:"completedTime"`
	RealDestinationAccountType string               `json:"realDestinationAccountType"`
	TransactionSubStatusDesc   string               `json:"transactionSubStatusDesc"`
	TxAmountToUsd              string               `json:"txAmountToUsd"`
	SourceAccountName          string               `json:"sourceAccountName"`
	SourceAccountTypeName      string               `json:"sourceAccountTypeName"`
	DestinationAccountName     string               `json:"destinationAccountName"`
	DestinationAccountTypeName string               `json:"destinationAccountTypeName"`
	AuditUserName              string               `json:"auditUserName"`
	CreatedByUserName          string               `json:"createdByUserName"`
	TransactionDirection       string               `json:"transactionDirection"`
}

type ListTransactionsV1Request struct {
	PageNumber                 int    `json:"pageNumber,omitempty"`
	PageSize                   int    `json:"pageSize,omitempty"`
	SourceAccountKey           string `json:"sourceAccountKey,omitempty"`
	SourceAccountType          string `json:"sourceAccountType,omitempty"`
	DestinationAccountKey      string `json:"destinationAccountKey,omitempty"`
	DestinationAccountType     string `json:"destinationAccountType,omitempty"`
	CreateTimeMin              int64  `json:"createTimeMin,omitempty"`
	CreateTimeMax              int64  `json:"createTimeMax,omitempty"`
	TxAmountMin                string `json:"txAmountMin,omitempty"`
	TxAmountMax                string `json:"txAmountMax,omitempty"`
	CoinKey                    string `json:"coinKey,omitempty"`
	FeeCoinKey                 string `json:"feeCoinKey,omitempty"`
	TransactionStatus          string `json:"transactionStatus,omitempty"`
	TransactionSubStatus       string `json:"transactionSubStatus,omitempty"`
	CompletedTimeMin           int64  `json:"completedTimeMin,omitempty"`
	CompletedTimeMax           int64  `json:"completedTimeMax,omitempty"`
	CustomerRefId              string `json:"customerRefId,omitempty"`
	RealDestinationAccountType string `json:"realDestinationAccountType,omitempty"`
	HideSmallAmountUsd         string `json:"hideSmallAmountUsd,omitempty"`
	TransactionDirection       string `json:"transactionDirection,omitempty"`
}

type TransactionsResponseV1 struct {
	PageNumber    int32                 `json:"pageNumber"`
	PageSize      int32                 `json:"pageSize"`
	TotalElements int64                 `json:"totalElements"`
	Content       []TransactionsRequest `json:"content"`
}

func (e *TransactionApi) ListTransactionsV1(d ListTransactionsV1Request, r *TransactionsResponseV1) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/list")
}

type ListTransactionsV2Request struct {
	Direct                     string `json:"direct,omitempty"`
	Limit                      int32  `json:"limit,omitempty"`
	FromId                     string `json:"fromId,omitempty"`
	SourceAccountKey           string `json:"sourceAccountKey,omitempty"`
	SourceAccountType          string `json:"sourceAccountType,omitempty"`
	DestinationAccountKey      string `json:"destinationAccountKey,omitempty"`
	DestinationAccountType     string `json:"destinationAccountType,omitempty"`
	CreateTimeMin              int64  `json:"createTimeMin,omitempty"`
	CreateTimeMax              int64  `json:"createTimeMax,omitempty"`
	TxAmountMin                string `json:"txAmountMin,omitempty"`
	TxAmountMax                string `json:"txAmountMax,omitempty"`
	CoinKey                    string `json:"coinKey,omitempty"`
	FeeCoinKey                 string `json:"feeCoinKey,omitempty"`
	TransactionStatus          string `json:"transactionStatus,omitempty"`
	TransactionSubStatus       string `json:"transactionSubStatus,omitempty"`
	CompletedTimeMin           int64  `json:"completedTimeMin,omitempty"`
	CompletedTimeMax           int64  `json:"completedTimeMax,omitempty"`
	CustomerRefId              string `json:"customerRefId,omitempty"`
	RealDestinationAccountType string `json:"realDestinationAccountType,omitempty"`
	HideSmallAmountUsd         string `json:"hideSmallAmountUsd,omitempty"`
	TransactionDirection       string `json:"transactionDirection,omitempty"`
}

type TransactionsResponseV2 []TransactionsRequest

func (e *TransactionApi) ListTransactionsV2(d ListTransactionsV2Request, r *TransactionsResponseV2) error {
	return e.Client.SendRequest(d, r, "/v2/transactions/list")
}

type CreateTransactionsRequest struct {
	CustomerRefId          string     `json:"customerRefId"`
	CustomerExt1           string     `json:"customerExt1,omitempty"`
	CustomerExt2           string     `json:"customerExt2,omitempty"`
	Note                   string     `json:"note,omitempty"`
	CoinKey                string     `json:"coinKey"`
	TxFeeLevel             string     `json:"txFeeLevel,omitempty"`
	FeeRateDto             FeeRateDto `json:"feeRateDto,omitempty"`
	MaxTxFeeRate           string     `json:"maxTxFeeRate,omitempty"`
	TxAmount               string     `json:"txAmount"`
	TreatAsGrossAmount     bool       `json:"treatAsGrossAmount,omitempty"`
	SourceAccountKey       string     `json:"sourceAccountKey"`
	SourceAccountType      string     `json:"sourceAccountType"`
	DestinationAccountKey  string     `json:"destinationAccountKey,omitempty"`
	DestinationAccountType string     `json:"destinationAccountType"`
	DestinationAddress     string     `json:"destinationAddress,omitempty"`
	Memo                   string     `json:"memo,omitempty"`
	DestinationTag         string     `json:"destinationTag,omitempty"`
	IsRbf                  *bool      `json:"isRbf,omitempty"`
	FailOnContract         *bool      `json:"failOnContract,omitempty"`
	Nonce                  int64      `json:"nonce,omitempty"`
	SequenceNumber         int64      `json:"sequenceNumber,omitempty"`
	BalanceVerifyType      string     `json:"balanceVerifyType,omitempty"`
}

type FeeRateDto struct {
	FeeRate        string `json:"feeRate,omitempty"`
	GasLimit       string `json:"gasLimit,omitempty"`
	MaxPriorityFee string `json:"maxPriorityFee,omitempty"`
	MaxFee         string `json:"maxFee,omitempty"`
	GasPremium     string `json:"gasPremium,omitempty"`
	GasFeeCap      string `json:"gasFeeCap,omitempty"`
	GasBudget      string `json:"gasBudget,omitempty"`
	GasUnitPrice   string `json:"gasUnitPrice,omitempty"`
	MaxGasAmount   string `json:"maxGasAmount,omitempty"`
}

type TxKeyResult struct {
	TxKey string `json:"txKey"`
}

func (e *TransactionApi) CreateTransactions(d CreateTransactionsRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v2/transactions/create")
}

type CreateTransactionsUTXOMultiDestRequest struct {
	CustomerRefId          string               `json:"customerRefId"`
	CustomerExt1           string               `json:"customerExt1,omitempty"`
	CustomerExt2           string               `json:"customerExt2,omitempty"`
	Note                   string               `json:"note,omitempty"`
	CoinKey                string               `json:"coinKey"`
	TxFeeLevel             string               `json:"txFeeLevel,omitempty"`
	FeeRateDto             FeeRateDto           `json:"feeRateDto,omitempty"`
	MaxTxFeeRate           string               `json:"maxTxFeeRate,omitempty"`
	SourceAccountKey       string               `json:"sourceAccountKey"`
	SourceAccountType      string               `json:"sourceAccountType"`
	DestinationAddressList []DestinationAddress `json:"destinationAddressList,omitempty"`
	DestinationTag         string               `json:"destinationTag,omitempty"`
	IsRbf                  bool                 `json:"isRbf,omitempty"`
}

type SourceAddress struct {
	Address         string `json:"address"`
	AddressGroupKey string `json:"addressGroupKey"`
}

type DestinationAddress struct {
	Address         string `json:"address"`
	Memo            string `json:"memo"`
	Amount          string `json:"amount"`
	AddressGroupKey string `json:"addressGroupKey"`
}

func (e *TransactionApi) CreateTransactionsUTXOMultiDest(d CreateTransactionsUTXOMultiDestRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/utxo/multidest/create")
}

type RecreateTransactionRequest struct {
	TxKey      string     `json:"txKey"`
	TxHash     string     `json:"txHash"`
	CoinKey    string     `json:"coinKey"`
	TxFeeLevel string     `json:"txFeeLevel,omitempty"`
	FeeRateDto FeeRateDto `json:"feeRateDto,omitempty"`
}

func (e *TransactionApi) RecreateTransactions(d RecreateTransactionRequest, r *TxKeyResult) error {
	return e.Client.SendRequest(d, r, "/v2/transactions/recreate")
}

type OneTransactionsRequest struct {
	TxKey         string `json:"txKey,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

type OneTransactionsResponse struct {
	TxKey                      string                `json:"txKey"`
	TxHash                     string                `json:"txHash"`
	CoinKey                    string                `json:"coinKey"`
	TxAmount                   string                `json:"txAmount"`
	SourceAccountKey           string                `json:"sourceAccountKey"`
	SourceAccountType          string                `json:"sourceAccountType"`
	SourceAddress              string                `json:"sourceAddress"`
	SourceAddressList          []SourceAddress       `json:"sourceAddressList"`
	DestinationAccountKey      string                `json:"destinationAccountKey"`
	DestinationAccountType     string                `json:"destinationAccountType"`
	DestinationAddress         string                `json:"destinationAddress"`
	Memo                       string                `json:"memo"`
	DestinationAddressList     []DestinationAddress  `json:"destinationAddressList"`
	DestinationTag             string                `json:"destinationTag"`
	TransactionType            string                `json:"transactionType"`
	TransactionStatus          string                `json:"transactionStatus"`
	TransactionSubStatus       string                `json:"transactionSubStatus"`
	CreateTime                 int64                 `json:"createTime"`
	Note                       string                `json:"note"`
	AuditUserKey               string                `json:"auditUserKey"`
	CreatedByUserKey           string                `json:"createdByUserKey"`
	TxFee                      string                `json:"txFee"`
	FeeCoinKey                 string                `json:"feeCoinKey"`
	ReplaceTxHash              string                `json:"replaceTxHash"`
	CustomerRefId              string                `json:"customerRefId"`
	CustomerExt1               string                `json:"customerExt1"`
	CustomerExt2               string                `json:"customerExt2"`
	AmlLock                    string                `json:"amlLock"`
	BlockHeight                int64                 `json:"blockHeight"`
	CompletedTime              int64                 `json:"completedTime"`
	RealDestinationAccountType string                `json:"realDestinationAccountType"`
	TransactionSubStatusDesc   string                `json:"transactionSubStatusDesc"`
	TxAmountToUsd              string                `json:"txAmountToUsd"`
	SourceAccountName          string                `json:"sourceAccountName"`
	SourceAccountTypeName      string                `json:"sourceAccountTypeName"`
	DestinationAccountName     string                `json:"destinationAccountName"`
	DestinationAccountTypeName string                `json:"destinationAccountTypeName"`
	AuditUserName              string                `json:"auditUserName"`
	CreatedByUserName          string                `json:"createdByUserName"`
	SpeedUpHistory             []TransactionsRequest `json:"speedUpHistory"`
	TransactionDirection       string                `json:"transactionDirection"`
}

func (e *TransactionApi) OneTransactions(d OneTransactionsRequest, r *OneTransactionsResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/one")
}

type TransactionsFeeRateRequest struct {
	CoinKey                string               `json:"coinKey"`
	TxHash                 string               `json:"txHash,omitempty"`
	SourceAccountKey       string               `json:"sourceAccountKey,omitempty"`
	SourceAddress          string               `json:"sourceAddress,omitempty"`
	DestinationAddress     string               `json:"destinationAddress"`
	DestinationAddressList []DestinationAddress `json:"destinationAddressList"`
	Value                  string               `json:"value,omitempty"`
}

type FeeRate struct {
	FeeRate        string `json:"feeRate"`
	Fee            string `json:"fee"`
	GasLimit       string `json:"gasLimit"`
	BaseFee        string `json:"baseFee"`
	MaxPriorityFee string `json:"maxPriorityFee"`
	MaxFee         string `json:"maxFee"`
	BytesSize      string `json:"bytesSize"`
	GasPremium     string `json:"gasPremium "`
	GasFeeCap      string `json:"gasFeeCap"`
	GasBudget      string `json:"gasBudget"`
	GasUnitPrice   string `json:"gasUnitPrice"`
	MaxGasAmount   string `json:"maxGasAmount"`
}

type TransactionsFeeRateResponse struct {
	FeeUnit       string  `json:"feeUnit"`
	MinFeeRate    FeeRate `json:"minFeeRate"`
	LowFeeRate    FeeRate `json:"lowFeeRate"`
	MiddleFeeRate FeeRate `json:"middleFeeRate"`
	HighFeeRate   FeeRate `json:"highFeeRate"`
}

func (e *TransactionApi) TransactionFeeRate(d TransactionsFeeRateRequest, r *TransactionsFeeRateResponse) error {
	return e.Client.SendRequest(d, r, "/v2/transactions/getFeeRate")
}

type CancelTransactionRequest struct {
	TxKey  string `json:"txKey"`
	TxType string `json:"txType,omitempty"`
}

func (e *TransactionApi) CancelTransactions(d CancelTransactionRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/cancel")
}

type CollectionTransactionsUTXORequest struct {
	CustomerRefId          string `json:"customerRefId"`
	CustomerExt1           string `json:"customerExt1,omitempty"`
	CustomerExt2           string `json:"customerExt2,omitempty"`
	Note                   string `json:"note,omitempty"`
	CoinKey                string `json:"coinKey"`
	TxFeeRate              string `json:"txFeeRate,omitempty"`
	TxFeeLevel             string `json:"txFeeLevel,omitempty"`
	MaxTxFeeRate           string `json:"maxTxFeeRate,omitempty"`
	MinCollectionAmount    string `json:"minCollectionAmount,omitempty"`
	SourceAccountKey       string `json:"sourceAccountKey"`
	SourceAccountType      string `json:"sourceAccountType"`
	DestinationAccountKey  string `json:"destinationAccountKey"`
	DestinationAccountType string `json:"destinationAccountType"`
	DestinationAddress     string `json:"destinationAddress,omitempty"`
	DestinationTag         string `json:"destinationTag,omitempty"`
}

type CollectionTransactionsUTXOResponse struct {
	TxKey            string `json:"txKey"`
	CollectionAmount string `json:"collectionAmount"`
	CollectionNum    int32  `json:"collectionNum"`
}

func (e *TransactionApi) CollectionTransactionsUTXO(d CollectionTransactionsUTXORequest, r *CollectionTransactionsUTXOResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/utxo/collection")
}
