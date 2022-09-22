package main

import (
	"day-6/go-restful/internal/factory"
	"day-6/go-restful/internal/http"
	"day-6/go-restful/pkg/constant"
	"day-6/go-restful/pkg/util"

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
	e.Use(middleware.Recover())

	// Register validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	f := factory.NewFactory()
	http.New(e.Group("/api"), f)

	// Start server
	port := constant.Env.Get("SERVER_PORT", "")
	e.Logger.Fatal(e.Start(":" + port))
}
