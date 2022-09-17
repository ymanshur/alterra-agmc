package route

import (
	"go-restful/controller"
	"go-restful/middleware"

	"github.com/labstack/echo/v4"
)

func Book(g *echo.Group) {
	g.GET("/books", controller.GetAllBook)
	g.POST("/books", controller.CreateBook, middleware.IsLoggedIn)
	g.GET("/books/:id", controller.GetBook)
	g.PUT("/books/:id", controller.UpdateBook, middleware.IsLoggedIn)
	g.DELETE("/books/:id", controller.DeleteBook, middleware.IsLoggedIn)
}
