package middleware

import (
	"go-restful/lib/auth"
	"log"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
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

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("id"))
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.JwtClaims)
		if claims.UserId != uint(userId) {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
