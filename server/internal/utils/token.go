package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {
	tokenGenerated := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":     id,
		"expiration": time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := tokenGenerated.SignedString([]byte(os.Getenv(("JWTSECRET"))))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSECRET")), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
