package main

import (
	"go-restful/config"
	"go-restful/route"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting .env file, %v", err)
	}
}

func main() {
	// Initial config
	config.InitDB()

	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	g := e.Group("")
	gAuth := e.Group("")
	gAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	// Routes
	route.Auth(g)
	route.Book(g, gAuth)
	route.User(g, gAuth)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
