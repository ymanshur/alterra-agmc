package route

import (
	c "day-4/go-restful/controller"
	mid "day-4/go-restful/middleware"
	"day-4/go-restful/model"

	"github.com/labstack/echo/v4"
)

func RouteBook(g *echo.Group) {
	controller := c.NewBookController(map[int]*model.Book{})

	g.GET("/books", controller.GetAllBook)
	g.POST("/books", controller.CreateBook, mid.IsLoggedIn)
	g.GET("/books/:id", controller.GetBook)
	g.PUT("/books/:id", controller.UpdateBook, mid.IsLoggedIn)
	g.DELETE("/books/:id", controller.DeleteBook, mid.IsLoggedIn)
}
