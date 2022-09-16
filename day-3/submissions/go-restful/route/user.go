package route

import (
	"go-restful/controller"

	"github.com/labstack/echo/v4"
)

func User(g, gAuth *echo.Group) {
	gAuth.GET("/users", controller.GetAllUser)
	g.POST("/users", controller.CreateUser)
	gAuth.GET("/users/:id", controller.GetUser)
	gAuth.PUT("/users/:id", controller.UpdateUser)
	gAuth.DELETE("/users/:id", controller.DeleteUser)
}
