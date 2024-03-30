package apiMiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
)

func SetDatabaseServices(
	persistentDbSvc *db.PersistentDatabaseService,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("persistentDbSvc", persistentDbSvc)
			return next(c)
		}
	}
}
