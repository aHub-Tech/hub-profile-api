package main

import (
	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	Route(router, ":3000")
}
