package route

import (
	"go-restful/controller"

	"github.com/labstack/echo/v4"
)

func User(g *echo.Group) {
	g.GET("/users", controller.GetAllUser)
	g.POST("/users", controller.CreateUser)
	g.GET("/users/:id", controller.GetUser)
	g.PUT("/users/:id", controller.UpdateUser)
	g.DELETE("/users/:id", controller.DeleteUser)
}
