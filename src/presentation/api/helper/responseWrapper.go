package apiHelper

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/service"
)

type formattedResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func ResponseWrapper(
	c echo.Context,
	serviceOutput service.ServiceOutput,
) error {
	responseStatus := http.StatusOK
	switch serviceOutput.Status {
	case service.Created:
		responseStatus = http.StatusCreated
	case service.MultiStatus:
		responseStatus = http.StatusMultiStatus
	case service.UserError:
		responseStatus = http.StatusBadRequest
	case service.InfraError:
		responseStatus = http.StatusInternalServerError
	}

	formattedResponse := formattedResponse{
		Status: responseStatus,
		Body:   serviceOutput.Body,
	}
	return c.JSON(responseStatus, formattedResponse)
}
