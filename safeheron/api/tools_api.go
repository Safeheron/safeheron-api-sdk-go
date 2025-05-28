package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type ToolsApi struct {
	Client safeheron.Client
}

type AmlCheckerRequestRequest struct {
	Network string `json:"network,omitempty"`
	Address string `json:"address,omitempty"`
}

type AmlCheckerRequestResponse struct {
	RequestId string `json:"requestId"`
}

func (e *ToolsApi) CmlCheckerRequest(d CreateWhitelistRequest, r *CreateWhitelistResponse) error {
	return e.Client.SendRequest(d, r, "/v1/tools/aml-checker/request")
}

type AmlCheckerRetrievesRequest struct {
	RequestId string `json:"requestId,omitempty"`
}

type AmlCheckerRetrievesResponse struct {
	RequestId          string    `json:"requestId"`
	CreateTime         string    `json:"createTime"`
	Network            string    `json:"network"`
	Address            string    `json:"address"`
	IsMaliciousAddress bool      `json:"isMaliciousAddress"`
	MistTrack          MistTrack `json:"mistTrack"`
}

type MistTrack struct {
	Status         string       `json:"status"`
	EvaluationTime string       `json:"evaluationTime"`
	Score          string       `json:"score"`
	RiskLevel      string       `json:"riskLevel"`
	DetailList     []string     `json:"detailList"`
	RiskDetail     []RiskDetail `json:"riskDetail"`
}

type RiskDetail struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Address string `json:"address"`
	Volume  string `json:"volume"`
	Percent bool   `json:"percent"`
}

func (e *ToolsApi) CreateFromTransactionWhitelist(d CreateFromTransactionWhitelistRequest, r *CreateWhitelistResponse) error {
	return e.Client.SendRequest(d, r, "/v1/tools/aml-checker/retrieves")
}
