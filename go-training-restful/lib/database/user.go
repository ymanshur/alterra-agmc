package database

import (
	"errors"
	"go-training-restful/config"
	"go-training-restful/middlewares"
	"go-training-restful/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(userId uint) (interface{}, error) {
	var user models.User
	if err := config.DB.Where(models.User{ID: userId}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(user *models.User) (interface{}, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("required both email and password")
	}

	if err := config.DB.Where(models.User{Email: user.Email, Password: user.Password}).Take(&user).Error; err != nil {
		return nil, errors.New("these credentials do not match our records")
	}

	token, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}
