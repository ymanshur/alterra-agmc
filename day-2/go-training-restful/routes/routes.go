package routes

import (
	"go-training-restful/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserControllers)

	return e
}
