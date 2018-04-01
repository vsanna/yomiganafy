package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)

func Root() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Root")
	}
}
