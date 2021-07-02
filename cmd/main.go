package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/og-zero/go-api-echo/internal/pkg/configs"
)

var config = configs.Load()

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", home)
	e.GET("/ping", ping)
	e.GET("/pong", pong)
	e.GET("/ip", ip)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func home(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/ip")
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "ping")
}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func ip(c echo.Context) error {
	clientIP := c.RealIP()
	return c.String(http.StatusOK, clientIP)
}
