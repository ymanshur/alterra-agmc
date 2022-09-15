package middlewares

import (
	"go-training-restful/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour + 1).Unix(), // Token expires after 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
