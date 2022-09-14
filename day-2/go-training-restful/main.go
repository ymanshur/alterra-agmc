package main

import (
	"fmt"
	"go-training-restful/config"
	"go-training-restful/routes"
)

func main() {
	config.InitDB()

	e := routes.New()

	fmt.Println("Starting web server at http://localhost:8080")

	e.Logger.Fatal(e.Start(":8000"))
}
