package middleware

import (
	"go-restful/lib/auth"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn echo.MiddlewareFunc

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting .env file, %v", err)
	}

	IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.JwtClaims{},
		SigningKey: []byte(os.Getenv("SECRET_JWT")),
	})
}
