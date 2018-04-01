package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)

func Root() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<a href='https://github.com/vsanna/yomiganafy'>how to use(github)</a>")
	}
}
