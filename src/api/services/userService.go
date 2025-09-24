package services

import (
	"go-backend/api/dto"
	"go-backend/database"
	"go-backend/helper"
	"go-backend/models"
	"log"
)

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AuthenticateUser(req *dto.LoginRequest) (string, error) {
	db := database.GetDB()

	email, password := req.Email, req.Password

	var users []models.User
	err := db.Where(&models.User{Email: email}).First(&users).Error
	if(err != nil) {
		log.Printf("Couldn't find the user: %v", err)
		return "", err
	}

	if !helper.CheckHashPassword(users[0].Password, password) {
		log.Printf("Password doesn't match")
		return "", err
	}

	return users[0].Id.String(), nil
}