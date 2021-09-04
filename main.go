package main

import (
	"github.com/henriquetied472/logplus"
	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	logplus.Info("Running")
	Route(router, ":3000")
}
