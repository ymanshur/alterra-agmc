package main

import (
	"fmt"
	"go-training-restful/config"
	m "go-training-restful/middlewares"
	"go-training-restful/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	// implement middleware logger
	m.LogMiddlewares(e)

	fmt.Println("Starting web server at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8000"))
}
