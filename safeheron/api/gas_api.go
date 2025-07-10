package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type GasApi struct {
	Client safeheron.Client
}

type GasStatusResponse struct {
	GasBalance []GasBalance `json:"gasBalance"`
}

type GasBalance struct {
	Symbol string `json:"symbol"`
	Amount string `json:"amount"`
}

func (e *GasApi) GasStatus(r *GasStatusResponse) error {
	return e.Client.SendRequest(nil, r, "/v1/gas/status")
}

type GasTransactionsGetByTxKeyRequest struct {
	TxKey string `json:"txKey,omitempty"`
}

type GasTransactionsGetByTxKeyResponse struct {
	TxKey       string   `json:"txKey"`
	Symbol      string   `json:"symbol"`
	TotalAmount string   `json:"totalAmount"`
	DetailList  []Detail `json:"detailList"`
}
type Detail struct {
	GasServiceTxKey string `json:"gasServiceTxKey"`
	Symbol          string `json:"symbol"`
	Amount          string `json:"amount"`
	Balance         string `json:"balance"`
	Status          string `json:"status"`
	ResourceType    string `json:"resourceType"`
	Timestamp       string `json:"timestamp"`
}

func (e *GasApi) GasTransactionsGetByTxKey(d GasTransactionsGetByTxKeyRequest, r *GasTransactionsGetByTxKeyResponse) error {
	return e.Client.SendRequest(d, r, "/v1/gas/transactions/getByTxKey")
}
