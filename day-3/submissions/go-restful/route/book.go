package route

import (
	"go-restful/controller"

	"github.com/labstack/echo/v4"
)

func Book(g, gAuth *echo.Group) {
	g.GET("/books", controller.GetAllBook)
	gAuth.POST("/books", controller.CreateBook)
	g.GET("/books/:id", controller.GetBook)
	gAuth.PUT("/books/:id", controller.UpdateBook)
	gAuth.DELETE("/books/:id", controller.DeleteBook)
}
