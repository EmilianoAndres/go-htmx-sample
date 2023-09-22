package main

import (
	"github.com/go-htmx-sample/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	// Create the router
	r := router.CreateRouter()
	r.Use(middleware.Logger())
	// Start the server
	r.Logger.Debug(r.Start(":8080"))
}
