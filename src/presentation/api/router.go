package api

import (
	"github.com/labstack/echo/v4"
	apiController "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/controller"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "embed"

	_ "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/docs"
)

type Router struct {
	baseRoute *echo.Group
}

func NewRouter(baseRoute *echo.Group) *Router {
	return &Router{
		baseRoute: baseRoute,
	}
}

func (router *Router) swaggerRoute() {
	swaggerGroup := router.baseRoute.Group("/swagger")
	swaggerGroup.GET("/*", echoSwagger.WrapHandler)
}

func (router *Router) contactRoutes() {
	accountGroup := router.baseRoute.Group("/contact")
	contactController := apiController.NewContactController()

	accountGroup.GET("/", contactController.GetContacts)
}

func (router *Router) RegisterRoutes() {
	router.swaggerRoute()
	router.contactRoutes()
}
