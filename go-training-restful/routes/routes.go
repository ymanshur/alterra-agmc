package routes

import (
	c "go-training-restful/controllers"
	m "go-training-restful/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(
		// middleware.HTTPSRedirect(),
		middleware.RemoveTrailingSlash(),
	)

	e.POST("/users", c.CreateUserController)

	// implement middelware with group routing
	eAuth := e.Group("")
	eAuth.Use(middleware.BasicAuth(m.BasicAuthDB))

	eAuth.GET("/users", c.GetUsersController)
	eAuth.GET("/users/:id", c.GetUserController)

	return e
}
