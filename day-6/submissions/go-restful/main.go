package main

import (
	lib "day-6/go-restful/lib"
	r "day-6/go-restful/route"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init echo
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	// e.Use(middleware.Recover())

	// Register validator
	e.Validator = &lib.CustomValidator{Validator: validator.New()}

	// Default routes group
	g := e.Group("")
	// Routes
	r.RouteBook(g)
	r.RouteUser(g)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
