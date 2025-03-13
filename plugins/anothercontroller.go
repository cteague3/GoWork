package plugins

import (
	"GoWork/models"
)

type AnotherController struct{}

func (c *AnotherController) HandleRequest(req models.Request) models.Response {
	// Perform work based on the request
	result := []string{"another result"}
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
