package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type ComplianceApi struct {
	Client safeheron.Client
}

type KytReportRequest struct {
	TxKey         string `json:"txKey,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

type KytRepostResponse struct {
	TxKey         string          `json:"txKey"`
	CustomerRefId string          `json:"customerRefId"`
	AmlList       []AmlAndPayload `json:"amlList"`
}

type AmlAndPayload struct {
	Provider       string `json:"provider"`
	Timestamp      string `json:"timestamp"`
	Status         string `json:"status"`
	RiskLevel      string `json:"riskLevel"`
	LastUpdateTime string `json:"lastUpdateTime"`
	Payload        any    `json:"payload"`
}

func (e *ComplianceApi) KytReport(d KytReportRequest, r *KytRepostResponse) error {
	return e.Client.SendRequest(d, r, "/v1/compliance/kyt/report")
}
