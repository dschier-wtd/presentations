package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func helloHandler(c echo.Context) error {
	ip := c.RealIP()
	host := c.Request().Host

	html := `<div align="center"><h1>Hello, DecompileD!</h1><p>You are visiting from ` + ip + `.<br>My address is ` + host + `.</p></div>`

	return c.HTML(http.StatusOK, html)
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloHandler)
	e.GET("/health", healthHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
