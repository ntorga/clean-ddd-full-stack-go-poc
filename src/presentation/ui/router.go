package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/infra/db"
	"github.com/ntorga/clean-ddd-full-stack-go-poc/src/presentation/ui/presenter"
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

func (router *Router) contactRoutes() {
	// Ideally, this feature would be implemented in a separate group, such as:
	// contactGroup := router.baseRoute.Group("/contact")

	contactPresenter := presenter.NewContactPresenter(router.persistentDbSvc)
	router.baseRoute.GET("/", contactPresenter.Handler)
}

func (router *Router) RegisterRoutes() {
	router.contactRoutes()
}
