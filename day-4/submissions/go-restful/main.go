package main

import (
	"day-4/go-restful/config"
	"day-4/go-restful/route"
	"fmt"
	"log"

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
	// Init database
	config.InitDB()
	// Init echo
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	// Default group
	g := e.Group("")

	// Routes
	route.Auth(g)
	route.Book(g)
	route.User(g)

	fmt.Println(g)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
