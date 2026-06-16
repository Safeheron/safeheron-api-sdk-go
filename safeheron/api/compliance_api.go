package api

import (
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
)

type ComplianceApi struct {
	Client safeheron.Client
}

// KYT Report

type KytReportRequest struct {
	TxKey         string `json:"txKey,omitempty"`
	CustomerRefId string `json:"customerRefId,omitempty"`
}

type KytReportResponse struct {
	TxKey                      string      `json:"txKey"`
	CustomerRefId              string      `json:"customerRefId"`
	AmlScreeningTriggeredState string      `json:"amlScreeningTriggeredState"`
	AmlList                    []AmlReport `json:"amlList"`
}

type AmlReport struct {
	Provider       string `json:"provider"`
	Timestamp      string `json:"timestamp"`
	Status         string `json:"status"`
	RiskLevel      string `json:"riskLevel"`
	LastUpdateTime string `json:"lastUpdateTime"`
	Payload        any    `json:"payload"`
}

func (e *ComplianceApi) KytReport(d KytReportRequest, r *KytReportResponse) error {
	return e.Client.SendRequest(d, r, "/v1/compliance/kyt/report")
}

// KYA Screening Create

type KyaScreeningRequest struct {
	Address   string   `json:"address"`
	ChainType string   `json:"chainType"`
	Network   string   `json:"network,omitempty"`
	Providers []string `json:"providers"`
}

type KyaScreeningOrder struct {
	ScreenOrderId string `json:"screenOrderId"`
	Provider      string `json:"provider"`
}

type KyaScreeningCreateResponse struct {
	ScreenId   string              `json:"screenId"`
	Address    string              `json:"address"`
	ChainType  string              `json:"chainType"`
	Network    string              `json:"network"`
	Orders     []KyaScreeningOrder `json:"orders"`
	CreateTime int64               `json:"createTime"`
}

func (e *ComplianceApi) KyaScreeningCreate(d KyaScreeningRequest, r *KyaScreeningCreateResponse) error {
	return e.Client.SendRequest(d, r, "/v1/compliance/kya/screening/create")
}

// KYA Screening Summary

type KyaScreeningOneRequest struct {
	ScreenId string `json:"screenId"`
}

type KyaScreeningOrderSummary struct {
	ScreenOrderId string `json:"screenOrderId"`
	Provider      string `json:"provider"`
	Status        string `json:"status"`
	RiskLevel     string `json:"riskLevel"`
	CompletedAt   int64  `json:"completedAt"`
}

type KyaScreeningOneResponse struct {
	ScreenId   string                     `json:"screenId"`
	Address    string                     `json:"address"`
	ChainType  string                     `json:"chainType"`
	Network    string                     `json:"network"`
	Status     string                     `json:"status"`
	CreateTime int64                      `json:"createTime"`
	Orders     []KyaScreeningOrderSummary `json:"orders"`
}

func (e *ComplianceApi) KyaScreeningOne(d KyaScreeningOneRequest, r *KyaScreeningOneResponse) error {
	return e.Client.SendRequest(d, r, "/v1/compliance/kya/screening/one")
}

// KYA Screening Order Details

type KyaScreeningOrderOneRequest struct {
	ScreenOrderId string `json:"screenOrderId"`
}

type KyaScreeningOrderOneResponse struct {
	ScreenOrderId string `json:"screenOrderId"`
	ScreenId      string `json:"screenId"`
	Provider      string `json:"provider"`
	Address       string `json:"address"`
	AddressType   string `json:"addressType"`
	ChainType     string `json:"chainType"`
	Network       string `json:"network"`
	Status        string `json:"status"`
	RiskLevel     string `json:"riskLevel"`
	CompletedAt   int64  `json:"completedAt"`
	CreateTime    int64  `json:"createTime"`
	Payload       any    `json:"payload"`
}

func (e *ComplianceApi) KyaScreeningOrderOne(d KyaScreeningOrderOneRequest, r *KyaScreeningOrderOneResponse) error {
	return e.Client.SendRequest(d, r, "/v1/compliance/kya/screening/order/one")
}

// KYA Supported Networks & Providers

type KyaSupportedNetworksResponse struct {
	Network   string   `json:"network"`
	ChainType string   `json:"chainType"`
	Providers []string `json:"providers"`
}

type KyaSupportedNetworksResponseList []KyaSupportedNetworksResponse

func (e *ComplianceApi) KyaSupportedNetworks(r *KyaSupportedNetworksResponseList) error {
	return e.Client.SendRequest(nil, r, "/v1/compliance/kya/supportedNetworks")
}
