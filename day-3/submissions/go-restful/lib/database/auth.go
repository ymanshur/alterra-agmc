package database

import (
	"errors"
	"fmt"
	"go-restful/config"
	"go-restful/constant"
	"go-restful/model"
	"time"

	"github.com/golang-jwt/jwt"
)

func Login(user *model.User) (interface{}, error) {
	fmt.Println(user)
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("required both email and password")
	}

	if err := config.DB.Where(model.User{Email: user.Email, Password: user.Password}).Take(&user).Error; err != nil {
		return nil, errors.New("these credentials do not match our records")
	}

	// Create token
	claims := jwt.MapClaims{
		"authorized": true,
		"Email":      user.Email,
		"exp":        time.Now().Add(time.Hour + 1).Unix(), // Token expires after 1 hour
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(constant.SECRET_JWT))
	if err != nil {
		return nil, err
	}

	user.Token = token
	return user, nil
}
