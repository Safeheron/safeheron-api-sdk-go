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

type RetrieveMpcSignRequest struct {
	CustomerRefId string `json:"customerRefId,omitempty"`
	TxKey         string `json:"txKey,omitempty"`
}

type RetrieveMpcSignResponse struct {
	TxKey                string `json:"txKey,omitempty"`
	TransactionStatus    string `json:"transactionStatus,omitempty"`
	TransactionSubStatus string `json:"transactionSubStatus,omitempty"`
	CreateTime           int    `json:"createTime,omitempty"`
	SourceAccountKey     string `json:"sourceAccountKey,omitempty"`
	CustomerRefId        string `json:"customerRefId,omitempty"`
	Hashs                []struct {
		Hash string `json:"hash,omitempty"`
		Sig  string `json:"sig,omitempty"`
		Note string `json:"note,omitempty"`
	} `json:"hashs,omitempty"`
}

func (e *MpcSignApi) CreateMpcSign(d CreateMpcSignRequest, r *CreateMpcSignResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/mpcsign/create")
}

func (e *MpcSignApi) RetrieveSig(d RetrieveMpcSignRequest, r *RetrieveMpcSignResponse) error {
	return e.Client.SendRequest(d, r, "/v1/transactions/mpcsign/one")
}
