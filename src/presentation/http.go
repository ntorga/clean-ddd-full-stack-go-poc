package presentation

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/api"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/ui"
)

func HttpServerInit(persistentDbSvc *db.PersistentDatabaseService) {
	e := echo.New()

	api.ApiInit(e, persistentDbSvc)
	ui.UiInit(e, persistentDbSvc)

	e.HideBanner = true

	serverBanner := `
░█▀▀░█░░░█▀▀░█▀█░█▀█░░░█▀▄░█▀▄░█▀▄              
░█░░░█░░░█▀▀░█▀█░█░█░░░█░█░█░█░█░█        v0.0.1
░▀▀▀░▀▀▀░▀▀▀░▀░▀░▀░▀░░░▀▀░░▀▀░░▀▀░              
░█▀▀░█░█░█░░░█░░░░░█▀▀░▀█▀░█▀█░█▀▀░█░█░░░█▀▀░█▀█
░█▀▀░█░█░█░░░█░░░░░▀▀█░░█░░█▀█░█░░░█▀▄░░░█░█░█░█
░▀░░░▀▀▀░▀▀▀░▀▀▀░░░▀▀▀░░▀░░▀░▀░▀▀▀░▀░▀░░░▀▀▀░▀▀▀
`

	fmt.Println(serverBanner)
	e.Start(":8080")
}
