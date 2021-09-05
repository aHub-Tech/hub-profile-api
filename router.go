package main

import (
	"net/http"
	"os"

	"github.com/ahub-tech/hub-profile-api/db"
	"github.com/ahub-tech/hub-profile-api/profile"
	"github.com/henriquetied472/logplus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route(router *echo.Echo, port string) {
	var profiles []profile.Profile

	profiles = append(profiles, profile.Profile{
		FullName: "Teste",
		Age: "0",
		Corporation: "Hub",
		Experience: "Estagiario",
		LinkedIn: "",
		Twitter: "",
		Facebook: "",
		Instagram: "",
	})

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.POST("/register", func(c echo.Context) error {
		rProfile := profile.NewProfile(c.QueryParam("fullname"), c.QueryParam("age"), c.QueryParam("corp"), c.QueryParam("exp"), c.QueryParam("lkin"), c.QueryParam("tw"), c.QueryParam("fb"), c.QueryParam("ig"), c.QueryParam("aut"))
		profiles = append(profiles, rProfile)

		db.AddProfile(rProfile)

		logplus.Debugf("%v", profiles)

		c.Response().Status = http.StatusOK
		return nil
	})

	router.GET("/search/:name", func(c echo.Context) error {
		name := c.Param("name")
		var err error
		db.SearchProfile(name)

		for _, v := range profiles {
			if v.FullName != name {
				continue
			}

			err = c.JSON(http.StatusOK, v)
		}

		return err
	})

	router.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, profiles)
	})

	router.Start(":" + os.Getenv("PORT"))
}