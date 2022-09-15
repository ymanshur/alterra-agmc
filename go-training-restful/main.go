package main

import (
	"go-training-restful/config"
	"go-training-restful/middlewares"
	"go-training-restful/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	// implement middleware logger
	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
