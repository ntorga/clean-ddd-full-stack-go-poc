package apiHelper

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/liaison"
)

type formattedResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func ResponseWrapper(
	c echo.Context,
	liaisonOutput liaison.LiaisonOutput,
) error {
	responseStatus := http.StatusOK
	switch liaisonOutput.Status {
	case liaison.Created:
		responseStatus = http.StatusCreated
	case liaison.MultiStatus:
		responseStatus = http.StatusMultiStatus
	case liaison.UserError:
		responseStatus = http.StatusBadRequest
	case liaison.InfraError:
		responseStatus = http.StatusInternalServerError
	}

	formattedResponse := formattedResponse{
		Status: responseStatus,
		Body:   liaisonOutput.Body,
	}
	return c.JSON(responseStatus, formattedResponse)
}
