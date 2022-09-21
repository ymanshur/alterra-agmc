package route

import (
	cons "day-4/go-restful/constant"
	c "day-4/go-restful/controller"
	db "day-4/go-restful/database"
	mid "day-4/go-restful/middleware"
	m "day-4/go-restful/model"

	"github.com/labstack/echo/v4"
)

func RouteUser(g *echo.Group) {
	// Init
	database := db.New(db.Config{
		User:     cons.Getenv("DB_USER"),
		Password: cons.Getenv("DB_PASS"),
		Host:     cons.Getenv("DB_HOST"),
		Port:     cons.Getenv("DB_PORT"),
		Name:     cons.Getenv("DB_NAME"),
	})
	db.Load(database, &m.User{})
	controller := c.NewUserController(database)

	// Routes
	g.GET("/users", controller.GetAllUser, mid.IsLoggedIn)
	g.POST("/users", controller.CreateUser)
	g.GET("/users/:id", controller.GetUser, mid.IsLoggedIn, mid.IsAuthorized)
	g.PUT("/users/:id", controller.UpdateUser, mid.IsLoggedIn, mid.IsAuthorized)
	g.DELETE("/users/:id", controller.DeleteUser, mid.IsLoggedIn, mid.IsAuthorized)
	// Login route - Obtain jwt auth token
	g.POST("/login", controller.LoginUser)
}
