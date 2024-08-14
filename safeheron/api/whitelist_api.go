package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type WhitelistApi struct {
	Client safeheron.Client
}

type CreateWhitelistRequest struct {
	WhitelistName string `json:"whitelistName,omitempty"`
	ChainType     string `json:"chainType,omitempty"`
	Address       string `json:"address,omitempty"`
}

type CreateWhitelistResponse struct {
	WhitelistKey string `json:"whitelistKey"`
}

func (e *MpcSignApi) CreateWhitelist(d CreateWhitelistRequest, r *CreateWhitelistResponse) error {
	return e.Client.SendRequest(d, r, "/v1/whitelist/create")
}

type OneWhitelistRequest struct {
	WhitelistKey string `json:"whitelistKey,omitempty"`
	Address      string `json:"address,omitempty"`
}

type WhitelistResponse struct {
	WhitelistKey    string `json:"whitelistKey,omitempty"`
	ChainType       string `json:"chainType,omitempty"`
	WhitelistName   string `json:"whitelistName,omitempty"`
	WhitelistStatus string `json:"whitelistStatus,omitempty"`
	CreateTime      int64  `json:"createTime,omitempty"`
	LastUpdateTime  int64  `json:"lastUpdateTime,omitempty"`
}

func (e *MpcSignApi) OneWhitelist(d OneWhitelistRequest, r *WhitelistResponse) error {
	return e.Client.SendRequest(d, r, "/v1/whitelist/one")
}

type ListWhitelistRequest struct {
	Direct          string `json:"direct,omitempty"`
	Limit           int32  `json:"limit,omitempty"`
	FromId          string `json:"fromId,omitempty"`
	ChainType       string `json:"chainType,omitempty"`
	WhitelistStatus string `json:"whitelistStatus,omitempty"`
	CreateTimeMin   int64  `json:"createTimeMin,omitempty"`
	CreateTimeMax   int64  `json:"createTimeMax,omitempty"`
}

func (e *MpcSignApi) ListWhitelist(d ListWhitelistRequest, r *[]WhitelistResponse) error {
	return e.Client.SendRequest(d, r, "/v1/whitelist/list")
}

type EditWhitelistRequest struct {
	WhitelistKey  string `json:"whitelistKey,omitempty"`
	WhitelistName string `json:"whitelistName,omitempty"`
	Address       string `json:"address,omitempty"`
	Force         bool   `json:"force,omitempty"`
}

func (e *MpcSignApi) EditWhitelist(d EditWhitelistRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/whitelist/edit")
}

type DeleteWhitelistRequest struct {
	WhitelistKey string `json:"whitelistKey,omitempty"`
}

func (e *MpcSignApi) DeleteWhitelist(d DeleteWhitelistRequest, r *ResultResponse) error {
	return e.Client.SendRequest(d, r, "/v1/whitelist/delete")
}
