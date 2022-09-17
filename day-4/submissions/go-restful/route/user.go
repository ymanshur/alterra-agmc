package route

import (
	cg "day-4/go-restful/config"
	c "day-4/go-restful/controller"
	m "day-4/go-restful/middleware"

	"github.com/labstack/echo/v4"
)

func RouteUser(g *echo.Group) {
	controller := c.NewUserController(cg.DB)

	g.GET("/users", controller.GetAllUser, m.IsLoggedIn)
	g.POST("/users", controller.CreateUser)
	g.GET("/users/:id", controller.GetUser, m.IsLoggedIn, m.IsAuthorized)
	g.PUT("/users/:id", controller.UpdateUser, m.IsLoggedIn, m.IsAuthorized)
	g.DELETE("/users/:id", controller.DeleteUser, m.IsLoggedIn, m.IsAuthorized)

	// Obtain jwt auth token
	g.POST("/login", controller.LoginUser)
}
