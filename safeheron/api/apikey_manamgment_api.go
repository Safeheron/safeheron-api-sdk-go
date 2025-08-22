package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type ApiKeyManagement struct {
	Client safeheron.Client
}

func (e *ApiKeyManagement) DisableApikey(r *ResultResponse) error {
	return e.Client.SendRequest(nil, r, "/v1/apikey/disable")
}
