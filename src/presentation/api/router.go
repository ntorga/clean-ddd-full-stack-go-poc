package api

import (
	_ "embed"

	"github.com/labstack/echo/v4"
	_ "github.com/ntorga/clean-ddd-go-taghs-poc-contacts/src/presentation/api/docs"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) RegisterRoutes(baseRoute *echo.Group) {
}
