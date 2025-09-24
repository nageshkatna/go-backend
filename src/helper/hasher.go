package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(stringToHash string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(stringToHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckHashPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	fmt.Println("Password match error: ", err)
	return err == nil
}	