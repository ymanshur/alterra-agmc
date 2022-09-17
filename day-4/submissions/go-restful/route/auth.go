package route

import (
	"day-4/go-restful/controller"

	"github.com/labstack/echo/v4"
)

func Auth(g *echo.Group) {
	// Obtain jwt auth token
	g.POST("/login", controller.Login)
}
