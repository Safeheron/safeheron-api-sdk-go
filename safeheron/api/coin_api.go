package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type CoinApi struct {
	Client safeheron.Client
}

type CoinResponse []struct {
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
	CoinType          string `json:"coinType"`
	TokenIdentifier   string `json:"tokenIdentifier"`
	MinTransferAmount string `json:"minTransferAmount"`
	BlockChain        string `json:"blockChain"`
	Network           string `json:"network"`
	GasLimit          int32  `json:"gasLimit"`
	IsMemo            string `json:"isMemo"`
	IsUtxo            string `json:"isUtxo"`
	BlockchainType    string `json:"blockchainType"`
}

func (e *CoinApi) ListCoin(r *CoinResponse) error {
	return e.Client.SendRequest(nil, r, "/v1/coin/list")
}

type CoinMaintainResponse []struct {
	CoinKey   string `json:"coinKey"`
	Maintain  bool   `json:"maintain"`
	Title     string `json:"title"`
	Content   int32  `json:"content"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (e *CoinApi) ListCoinMaintain(r *CoinMaintainResponse) error {
	return e.Client.SendRequest(nil, r, "/v1/coin/maintain/list")
}

type CheckCoinAddressRequest struct {
	CoinKey           string `json:"coinKey"`
	Address           string `json:"address"`
	CheckContract     bool   `json:"checkContract,omitempty"`
	CheckAml          bool   `json:"checkAml,omitempty"`
	CheckAddressValid bool   `json:"checkAddressValid,omitempty"`
}

type CheckCoinAddressResponse struct {
	Contract     bool `json:"contract"`
	AmlValid     bool `json:"amlValid"`
	AddressValid bool `json:"addressValid"`
}

func (e *CoinApi) CheckCoinAddress(d CheckCoinAddressRequest, r *CheckCoinAddressResponse) error {
	return e.Client.SendRequest(d, r, "/v1/coin/address/check")
}

type CoinBalanceSnapshotRequest struct {
	Gmt8Date string `json:"gmt8Date"`
}

type CoinBalanceSnapshotResponse struct {
	CoinKey     string `json:"coinKey"`
	CoinBalance string `json:"coinBalance"`
}

func (e *CoinApi) CoinBalanceSnapshot(d CoinBalanceSnapshotRequest, r *CoinBalanceSnapshotResponse) error {
	return e.Client.SendRequest(d, r, "/v1/coin/balance/snapshot")
}

type CoinBlockHeightRequest struct {
	CoinKey string `json:"coinKey"`
}

type CoinBlockHeightResponse struct {
	CoinKey          string `json:"coinKey"`
	LocalBlockHeight string `json:"localBlockHeight"`
}

func (e *CoinApi) CoinBlockHeight(d CoinBlockHeightRequest, r *CoinBlockHeightResponse) error {
	return e.Client.SendRequest(d, r, "/v1/coin/block/height")
}
