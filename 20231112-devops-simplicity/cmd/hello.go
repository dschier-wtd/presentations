package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCurrentTime() string {
	currentTime := time.Now().Format("02.01.2006 - 15:04:05")
	return currentTime
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.GET("/time", func(c echo.Context) error {

		now := getCurrentTime()
		return c.JSON(http.StatusOK, now)

	})

	e.Logger.Fatal(e.Start(":3000"))
}
