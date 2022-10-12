package transaction_api_demo

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type TransactionApi struct {
	Client safeheron.Client
}

type CreateTransactionRequest struct {
	CustomerRefId string `json:"customerRefId"`
	CoinKey       string `json:"coinKey,omitempty"`
	TxFeeLevel    string `json:"txFeeLevel,omitempty"`
	FeeRateDto    struct {
		FeeRate        string `json:"feeRate,omitempty"`
		GasLimit       string `json:"gasLimit,omitempty"`
		MaxPriorityFee string `json:"maxPriorityFee,omitempty"`
		MaxFee         string `json:"maxFee,omitempty"`
	} `json:"feeRateDto,omitempty"`
	TxAmount               string `json:"txAmount,omitempty"`
	SourceAccountKey       string `json:"sourceAccountKey,omitempty"`
	SourceAccountType      string `json:"sourceAccountType,omitempty"`
	DestinationAccountKey  string `json:"destinationAccountKey,omitempty"`
	DestinationAccountType string `json:"destinationAccountType,omitempty"`
	DestinationAddress     string `json:"destinationAddress,omitempty"`
}

type CreateTransactionResponse struct {
	TxKey string `json:"txKey"`
}

func (e *TransactionApi) SendTransaction(d CreateTransactionRequest, r *CreateTransactionResponse) error {
	return e.Client.SendRequest(d, r, "/v2/transactions/create")
}
