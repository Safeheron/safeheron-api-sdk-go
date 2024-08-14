package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type WebhookApi struct {
	Client safeheron.Client
}

type ResendWebhookRequest struct {
	Category string `json:"category,omitempty"`
	TxKey    string `json:"txKey,omitempty"`
}

func (e *MpcSignApi) ResendWebhook(d ResendWebhookRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/webhook/resend")
}

type ResendFailedRequest struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
}

type MessagesCountResponse struct {
	MessagesCount int32 `json:"messagesCount"`
}

func (e *MpcSignApi) ResendFailed(d ResendFailedRequest, r *MessagesCountResponse) error {
	return e.Client.SendRequest(d, r, "/v1/webhook/resend/failed")
}
