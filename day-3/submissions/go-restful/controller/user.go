package controller

import (
	"errors"
	"go-restful/lib/database"
	"go-restful/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)

	createdUser, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success create new user",
		"data":    createdUser,
	})
}

func GetUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUser(uint(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get a user",
		"data":    user,
	})
}

func UpdateUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	// Return http.StatusNotFound if user does not exist
	_, err := database.GetUser(uint(userId))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	user := new(model.User)
	c.Bind(&user)

	updatedUser, err := database.UpdateUser(uint(userId), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update a user",
		"data":    updatedUser,
	})
}

func DeleteUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := database.DeleteUser(uint(userId)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete a user",
	})
}

func GetAllUser(c echo.Context) error {
	users, err := database.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all users",
		"data":    users,
	})
}
