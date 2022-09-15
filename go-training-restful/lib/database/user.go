package database

import (
	"go-training-restful/config"
	"go-training-restful/models"

	"github.com/labstack/echo/v4"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(c echo.Context) (interface{}, error) {
	userId := c.Param("id")
	var user models.User
	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
