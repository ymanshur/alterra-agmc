package controllers

import (
	"go-training-restful/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all users",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	user, err := database.CreateUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success create new user",
		"data":    user,
	})
}

func GetUserController(c echo.Context) error {
	user, err := database.GetUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get a user",
		"data":    user,
	})
}
