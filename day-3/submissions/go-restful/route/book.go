package route

import (
	"go-restful/controller"

	"github.com/labstack/echo/v4"
)

func Book(g *echo.Group) {
	g.GET("/books", controller.GetAllBook)
	g.POST("/books", controller.CreateBook)
	g.GET("/books/:id", controller.GetBook)
	g.PUT("/books/:id", controller.UpdateBook)
	g.DELETE("/books/:id", controller.DeleteBook)
}
