package database

import (
	"errors"
	"go-restful/config"
	"go-restful/lib/auth"
	"go-restful/model"
)

func Login(user *model.User) (interface{}, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("required both email and password")
	}

	if err := config.DB.Where(model.User{Email: user.Email, Password: user.Password}).Take(&user).Error; err != nil {
		return nil, errors.New("these credentials do not match our records")
	}

	token, err := auth.CreateJWt(user)
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}
