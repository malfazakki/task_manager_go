package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtsecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(UserID uint) (string, error) {

	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtsecret)

}
