package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error in hashing passsword")
	}
	return string(hash)
}

func ComparePassword(hashPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
