package main

import (
	"net/http"
	"os"

	"github.com/ahub-tech/hub-profile-api/db"
	"github.com/ahub-tech/hub-profile-api/profile"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route(router *echo.Echo, port string) {
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.POST("/register", func(c echo.Context) error {
		rProfile := profile.NewProfile(c.QueryParam("fullname"), c.QueryParam("age"), c.QueryParam("corp"), c.QueryParam("exp"), c.QueryParam("lkin"), c.QueryParam("tw"), c.QueryParam("fb"), c.QueryParam("ig"), c.QueryParam("aut"))

		db.AddProfile(rProfile)

		return c.String(http.StatusOK, "Registered sucesfully")
	})

	router.GET("/search/:name", func(c echo.Context) error {
		name := c.Param("name")

		return c.JSON(http.StatusOK, db.SearchProfile(name))
	})

	router.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, db.AllProfiles())
	})

	router.Start(":" + os.Getenv("PORT"))
}