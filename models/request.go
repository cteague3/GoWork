package models

type Request struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	InternalId  string `json:"internalId"`
	ExternalId  string `json:"externalId"`
	Params      string `json:"params"`
	Idsid       string `json:"idsid"`
	Domain      string `json:"domain"`
	Command     string `json:"command"`
	PayloadType string `json:"payloadType"`
	Payload     []byte `json:"payload"`
}
