package route

import (
	"go-restful/controller"
	"go-restful/middleware"

	"github.com/labstack/echo/v4"
)

func User(g *echo.Group) {
	g.GET("/users", controller.GetAllUser, middleware.IsLoggedIn)
	g.POST("/users", controller.CreateUser)
	g.GET("/users/:id", controller.GetUser, middleware.IsLoggedIn, middleware.IsAuthorized)
	g.PUT("/users/:id", controller.UpdateUser, middleware.IsLoggedIn, middleware.IsAuthorized)
	g.DELETE("/users/:id", controller.DeleteUser, middleware.IsLoggedIn, middleware.IsAuthorized)
}
