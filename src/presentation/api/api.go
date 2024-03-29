package api

import (
	"github.com/labstack/echo/v4"
	apiMiddleware "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/middleware"
)

// @title			Clean DDD TAGHS PoC Contacts
// @version			0.0.1
// @description		Clean Architecture & DDD with Go, Tailwind, Alpine.js, HTMX, and SQLite: A Proof of Concept

// @contact.name	Northon Torga
// @contact.url		https://ntorga.com/
// @contact.email	northontorga+github@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/v1
func ApiInit() {
	e := echo.New()

	basePath := "/v1"
	baseRoute := e.Group(basePath)

	e.Pre(apiMiddleware.TrailingSlash(basePath))
	e.Use(apiMiddleware.PanicHandler)
	e.Use(apiMiddleware.SetDefaultHeaders)

	router := NewRouter(baseRoute)
	router.RegisterRoutes()

	e.Start(":8080")
}
