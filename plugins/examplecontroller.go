package plugins

import (
	"GoWork/models"
)

type ExampleController struct{}

func (c *ExampleController) HandleRequest(req models.Request) models.Response {
	// Perform work based on the request
	result := []string{"example result"}
	return models.Response{
		RequestId:  req.Id,
		InternalId: req.InternalId,
		ExternalId: req.ExternalId,
		Params:     req.Params,
		Idsid:      req.Idsid,
		Domain:     req.Domain,
		Command:    req.Command,
		Result:     result,
	}
}
