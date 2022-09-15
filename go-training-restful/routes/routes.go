package routes

import (
	"go-training-restful/constants"
	"go-training-restful/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(
		// middleware.HTTPSRedirect(),
		middleware.RemoveTrailingSlash(),
	)

	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.CreateUserController)
	// e.GET("/users/:id", controllers.GetUserController)
	// Obtain jwt auth token
	e.POST("/login", controllers.LoginUserController)

	// Implement middelware with group routing
	eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(m.BasicAuthDB))
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	eAuth.GET("/users/:id", controllers.GetUserController)

	return e
}
