package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"server/src/domain"
)

func BindAndValidate(c echo.Context, item interface{}) error {
	if err := c.Bind(item); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Err: err.Error()})
		return err
	}
	if err := c.Validate(item); err != nil {
		msg := fmt.Sprintf("%s: %s", "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, domain.Response{Err: msg})
		return err
	}

	return nil
}
