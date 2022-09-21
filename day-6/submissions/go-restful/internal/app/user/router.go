package user

import (
	"day-6/go-restful/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (c *controller) Route(g *echo.Group) {
	// Routes
	g.GET("", c.GetAll, middleware.IsLoggedIn)
	g.GET("/:id", c.Get, middleware.IsLoggedIn, middleware.IsAuthorized)
	g.PUT("/:id", c.Update, middleware.IsLoggedIn, middleware.IsAuthorized)
	g.DELETE("/:id", c.Delete, middleware.IsLoggedIn, middleware.IsAuthorized)
}
