package route

import (
	"day-4/go-restful/controller"
	"day-4/go-restful/middleware"

	"github.com/labstack/echo/v4"
)

func Book(g *echo.Group) {
	handler := new(controller.BookHandler)

	g.GET("/books", handler.GetAllBook)
	g.POST("/books", handler.CreateBook, middleware.IsLoggedIn)
	g.GET("/books/:id", handler.GetBook)
	g.PUT("/books/:id", handler.UpdateBook, middleware.IsLoggedIn)
	g.DELETE("/books/:id", handler.DeleteBook, middleware.IsLoggedIn)
}
