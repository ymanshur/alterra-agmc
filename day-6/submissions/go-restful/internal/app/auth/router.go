package auth

import (
	"github.com/labstack/echo/v4"
)

func (c *controller) Route(g *echo.Group) {
	// Routes
	g.POST("/signup", c.SignUp)
	g.POST("/signin", c.SignIn)
}
