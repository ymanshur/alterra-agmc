package middleware

import (
	cons "day-4/go-restful/constant"
	"day-4/go-restful/lib/auth"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     &auth.JwtClaims{},
	SigningKey: []byte(cons.Getenv("SECRET_JWT")),
})

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, _ := strconv.Atoi(ctx.Param("id"))
		if userId != 0 {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*auth.JwtClaims)
			if claims.UserId != uint(userId) {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{
					"message": "unauthorized",
				})
			}
		}

		return next(ctx)
	}
}
