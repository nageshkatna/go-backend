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

func (us *UserService) GetUserByEmail(req *dto.LoginRequest) ([]models.User, error) {
	db := database.GetDB()

	email, password := req.Email, req.Password

	var users []models.User
	err := db.Where(&models.User{Email: email}).First(&users).Error
	if(err != nil) {
		log.Printf("Couldn't find the user: %v", err)
		return users, err
	}

	if !helper.CheckHashPassword(users[0].Password, password) {
		log.Printf("Password doesn't match")
		return users, err
	}

	return users, nil
}

func (us *UserService) CreateUser(req *dto.RegisterUserRequest, role *[]models.UserRole) ( dto.LoginResponse, error) {
	db := database.GetDB()

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		log.Printf("❌ Failed to hash password %v\n", err)
		return dto.LoginResponse{}, err
	}
	
	newUser := models.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: hashedPassword,
		UserRoles: role,
	}

	rowsAffected := db.FirstOrCreate(&newUser, models.User{Email: newUser.Email}).RowsAffected
	if rowsAffected == 0 {
		c := helper.CustomErrors{
			Message: "User already exists",
			Field: req.Email,
		}
		return dto.LoginResponse{}, c.CreateUserError()
	} else {
		return dto.LoginResponse{
			UserId: newUser.Id.String(),
		}, nil
	}
}