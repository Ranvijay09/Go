package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "VoidPointer"

func GenerateJWTToken(email string, userId int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return jwtToken.SignedString([]byte(secretKey))
}
