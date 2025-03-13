package models

type Response struct {
	RequestId  string      `json:"requestId"`
	InternalId string      `json:"internalId"`
	ExternalId string      `json:"externalId"`
	Params     string      `json:"params"`
	Idsid      string      `json:"idsid"`
	Domain     string      `json:"domain"`
	Command    string      `json:"command"`
	Result     interface{} `json:"result"`
}
