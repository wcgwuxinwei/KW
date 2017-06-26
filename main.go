package main

import (
	"net/http"

	"github.com/wcgwuxinwei/KW/conf"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	conf.Init()

	log.Warnf("Start Server")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// health_check
	e.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.Start(":8080")
}
