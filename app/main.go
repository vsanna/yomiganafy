package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vsanna/yomiganafy/handlers"
	"net/http"
)

// GAEに上げる場合main関数ではなくinitを使うらしい。
// なのでlocalでテストする場合は常にdev_appserver.pyを経由するしか無い
func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/", handlers.Root())
	e.GET("/yomiganafy", handlers.Yomiganafy()) // TODO: developmentでだけ有効化したい
	e.POST("/yomiganafy", handlers.Yomiganafy())

	// GAEのため、/ で受け取った処理をinit関数でeにわたす
	// e.Startができない...><
	http.Handle("/", e)
}

