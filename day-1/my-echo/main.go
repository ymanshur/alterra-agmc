package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func main() {
	// create nerw echo intance
	e := echo.New()
	// Route / to handler function
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.GET("/", HelloController)
	e.GET("/users/:id", GetUserController)
	e.GET("/users", GetAllUserController)
	// e.GET("/users", UserSearchController)
	e.POST("/users", CreateUser)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

// handler - Simple handler to make sure environment in setup
func HelloController(c echo.Context) error {
	// return the string "Hello, World!" as the response body
	// with an http.StatusOK (200) status
	return c.String(http.StatusOK, "Hello, World!")
}

// e.GET("/users/:id", GetUserController)
func GetUserController(c echo.Context) error {
	// User ID from path `users/:id`
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{id, "Ismail", "ismail@alrerra.id"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

// e.GET("/users", GetAllUserController)
func GetAllUserController(c echo.Context) error {
	user := User{1, "Ismail", "ismail@alrerra.id"}
	return c.JSON(http.StatusOK, []interface{}{user})
}

func UserSearchController(c echo.Context) error {
	// get data from query param
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"adi", "aan", "asif"},
	})
}

func CreateUser(c echo.Context) error {
	// get data from value
	// name := c.FormValue("name")
	// email := c.FormValue("email")

	// var user User
	// user.Id = 1
	// user.Name = name
	// user.Email = email

	// get data from raw body
	// user := new(User)
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}
