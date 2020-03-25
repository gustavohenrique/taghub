package http

import (
	"github.com/labstack/echo/v4"

	"server/src/containers/service"
)

func VerifyAuth(services *service.ServiceContainer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("ManagerID", "6bb8d106-dca1-494c-b488-239c2607d00f")
			return next(c)
		}
	}
}
