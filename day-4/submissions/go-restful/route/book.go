package route

import (
	"day-4/go-restful/controller"
	m "day-4/go-restful/middleware"

	"github.com/labstack/echo/v4"
)

func Book(g *echo.Group) {
	c := new(controller.BookController)

	g.GET("/books", c.GetAllBook)
	g.POST("/books", c.CreateBook, m.IsLoggedIn)
	g.GET("/books/:id", c.GetBook)
	g.PUT("/books/:id", c.UpdateBook, m.IsLoggedIn)
	g.DELETE("/books/:id", c.DeleteBook, m.IsLoggedIn)
}
