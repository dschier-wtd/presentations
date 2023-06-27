package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
