package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const dammy = "secret"

func GenerateToken(email string, username int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": username,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(dammy))

}
