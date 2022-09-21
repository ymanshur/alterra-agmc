package util

import (
	"day-6/go-restful/internal/model"
	"day-6/go-restful/pkg/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func CreateJwt(user *model.User) (string, error) {
	// Set claims
	// claims := jwt.MapClaims{
	// 	"email": user.ID,
	// 	"exp":   time.Now().Add(time.Hour + 1).Unix(), // Token expires after 1 hour
	// }
	claims := &JwtClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires after 1 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.Env.Get("SECRET_JWT", "")))
}
