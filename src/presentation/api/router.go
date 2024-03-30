package api

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	apiController "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/controller"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "embed"

	_ "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/api/docs"
)

type Router struct {
	baseRoute       *echo.Group
	persistentDbSvc *db.PersistentDatabaseService
}

func NewRouter(
	baseRoute *echo.Group,
	persistentDbSvc *db.PersistentDatabaseService,
) *Router {
	return &Router{
		baseRoute:       baseRoute,
		persistentDbSvc: persistentDbSvc,
	}
}

func (router *Router) swaggerRoute() {
	swaggerGroup := router.baseRoute.Group("/swagger")
	swaggerGroup.GET("/*", echoSwagger.WrapHandler)
}

func (router *Router) contactRoutes() {
	accountGroup := router.baseRoute.Group("/v1/contact")
	contactController := apiController.NewContactController(router.persistentDbSvc)

	accountGroup.GET("/", contactController.Read)
	accountGroup.POST("/", contactController.Create)
	accountGroup.PUT("/", contactController.Update)
}

func (router *Router) RegisterRoutes() {
	router.swaggerRoute()
	router.contactRoutes()
}
