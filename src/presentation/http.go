package presentation

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api"
)

func HttpServerInit() {
	e := echo.New()

	api.ApiInit(e)

	e.Start(":8080")
}
