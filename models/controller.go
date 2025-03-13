package models

type Controller interface {
	HandleRequest(req Request) Response
}
