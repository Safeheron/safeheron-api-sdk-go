package mpc_demo

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type MpcSignApi struct {
	Client safeheron.Client
}

type CreateMpcSignRequest struct {
	CustomerRefId    string `json:"customerRefId,omitempty"`
	SourceAccountKey string `json:"sourceAccountKey,omitempty"`
	SignAlg          string `json:"signAlg,omitempty"`
	Hashs            []struct {
		Hash string `json:"hash,omitempty"`
		Note string `json:"note,omitempty"`
	} `json:"hashs,omitempty"`
}

type CreateMpcSignResponse struct {
	TxKey string `json:"txKey"`
}

func (e *MpcSignApi) CreateMpcSign(d CreateMpcSignRequest, r *CreateMpcSignResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/mpcsign/create")
}

type OneMPCSignTransactionsRequest struct {
	CustomerRefId string `json:"customerRefId,omitempty"`
	TxKey         string `json:"txKey,omitempty"`
}

type MPCSignTransactionsResponse struct {
	TxKey                string `json:"txKey,omitempty"`
	TransactionStatus    string `json:"transactionStatus,omitempty"`
	TransactionSubStatus string `json:"transactionSubStatus,omitempty"`
	CreateTime           int    `json:"createTime,omitempty"`
	SourceAccountKey     string `json:"sourceAccountKey,omitempty"`
	AuditUserKey         string `json:"auditUserKey,omitempty"`
	CreatedByUserKey     string `json:"createdByUserKey,omitempty"`
	CustomerRefId        string `json:"customerRefId,omitempty"`
	CustomerExt1         string `json:"customerExt1,omitempty"`
	CustomerExt2         string `json:"customerExt2,omitempty"`
	SignAlg              string `json:"signAlg,omitempty"`
	AuditUserName        string `json:"auditUserName,omitempty"`
	CreatedByUserName    string `json:"createdByUserName,omitempty"`
	Hashs                []struct {
		Data string `json:"data,omitempty"`
		Sig  string `json:"sig,omitempty"`
		Note string `json:"note,omitempty"`
	} `json:"hashs,omitempty"`
}

func (e *MpcSignApi) OneMPCSignTransactions(d OneMPCSignTransactionsRequest, r *MPCSignTransactionsResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/mpcsign/one")
}

type ListMPCSignTransactionsRequest struct {
	Direct        string `json:"direct,omitempty"`
	Limit         int32  `json:"limit,omitempty"`
	FromId        string `json:"fromId,omitempty"`
	CreateTimeMin int64  `json:"createTimeMin,omitempty"`
	CreateTimeMax int64  `json:"createTimeMax,omitempty"`
}

func (e *MpcSignApi) ListMPCSignTransactions(d ListMPCSignTransactionsRequest, r *[]MPCSignTransactionsResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/mpcsign/list")
}
