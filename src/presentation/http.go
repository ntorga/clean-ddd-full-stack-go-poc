package presentation

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api"
)

func HttpServerInit(persistentDbSvc *db.PersistentDatabaseService) {
	e := echo.New()

	api.ApiInit(e, persistentDbSvc)

	e.Start(":8080")
}
