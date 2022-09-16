package controller

import (
	"go-restful/lib/database"
	"go-restful/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	user := new(model.User)
	c.Bind(&user)

	loggedInUser, err := database.Login(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success loggin",
		"data":    loggedInUser,
	})
}
