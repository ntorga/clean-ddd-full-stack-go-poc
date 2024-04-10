package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
)

func UiInit(e *echo.Echo, persistentDbSvc *db.PersistentDatabaseService) {
	basePath := ""
	baseRoute := e.Group(basePath)

	router := NewRouter(baseRoute, persistentDbSvc)
	router.RegisterRoutes()
}
