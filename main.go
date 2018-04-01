package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/", handlers.Root())
	//e.GET("/yomiganafy", handlers.Yomiganafy())
	e.POST("/yomiganafy", handlers.Yomiganafy())

	e.Start(":1234")
}

