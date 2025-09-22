package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(stringToHash string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(stringToHash), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(hashedPassword)
}

func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}	