package middleware

import (
	inutil "day-6/go-restful/internal/pkg/util"
	"day-6/go-restful/pkg/constant"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &inutil.JwtClaims{},
		SigningKey: []byte(constant.Env.Get("JWT_SECRET", "")),
	})
)

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, _ := strconv.Atoi(ctx.Param("id"))
		if userId != 0 {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*inutil.JwtClaims)
			if claims.UserId != uint(userId) {
				return ctx.JSON(http.StatusUnauthorized, echo.Map{
					"message": "unauthorized",
				})
			}
		}

		return next(ctx)
	}
}
