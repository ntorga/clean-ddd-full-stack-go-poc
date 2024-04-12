package apiMiddleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TrailingSlash(apiBasePath string) echo.MiddlewareFunc {
	trailingSlashSkipRegex := regexp.MustCompile(
		`^` + apiBasePath + `/swagger`,
	)

	return middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusTemporaryRedirect,
		Skipper: func(c echo.Context) bool {
			currentPath := c.Request().URL.Path
			isNotApi := !strings.HasPrefix(currentPath, apiBasePath)
			if isNotApi {
				return true
			}

			return trailingSlashSkipRegex.MatchString(currentPath)
		},
	})
}
