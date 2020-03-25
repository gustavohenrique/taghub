package http

import (
	"github.com/labstack/echo/v4"

	"server/src/containers/service"
)

func VerifyAuth(services *service.ServiceContainer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
