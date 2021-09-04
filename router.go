package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route(router *echo.Echo, port string) {
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/", func(c echo.Context) error {
		return c.File("views/index.html")
	})

	router.Start(port)
}