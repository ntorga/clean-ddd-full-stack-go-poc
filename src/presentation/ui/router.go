package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/ui/pages"
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
	contactGroup := router.baseRoute.Group("/contact")

	contactPresenter := pages.NewContactPresenter(router.persistentDbSvc)
	contactGroup.GET("/", contactPresenter.Handler)
}

func (router *Router) RegisterRoutes() {
	router.contactRoutes()
}
