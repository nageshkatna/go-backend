package services

import (
	"fmt"
	"go-backend/api/dto"
	"go-backend/database"
	"go-backend/helper"
	"go-backend/models"
	"log"
	"math"
)

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUserByEmail(req *dto.LoginRequest) ([]models.User, error) {
	db := database.GetDB()

	email, password := req.Email, req.Password

	var users []models.User
	err := db.Preload("UserRoles.Role").Where(&models.User{Email: email}).First(&users).Error
	if(err != nil) {
		log.Printf("Couldn't find the user: %v", err)
		return []models.User{}, err
	}

	if !helper.CheckHashPassword(users[0].Password, password) {
		log.Printf("Password doesn't match")
		return []models.User{}, err
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

func (*UserService) GetPaginatedUser(req *dto.Pagination) (dto.FetchUserRoleWithPaginatedResponse, error) {
	db := database.GetDB()

	var users []models.User
	var total int64

	paginate := dto.PaginatedResponse{}
	response := dto.FetchUserRoleWithPaginatedResponse{}

	if err := db.Model(&models.User{}).Count(&total).Error; err == nil {
		paginate = dto.PaginatedResponse{
			TotalRecords: total,
			TotalPages: uint(math.Ceil(float64(total)/float64(req.PageSize))),
			Pagination: dto.Pagination{
				Page: req.Page,
				PageSize: req.PageSize,
			},
		}
	}

	page := (req.Page - 1) * req.PageSize

	result := db.Preload("UserRoles.Role").Limit(req.PageSize).Offset(page).Find(&users)

	if(result.RowsAffected == 0) {
		return response, fmt.Errorf("no more rows to render")
	}
	
	var usersWithRole []dto.FetchUserWithRole

	for _, user := range users {
		usersWithRole = append(usersWithRole, dto.FetchUserWithRole{
			UserBaseRequest: dto.UserBaseRequest{
				FirstName: user.FirstName,
				LastName: user.LastName,
				Email: user.Email,
			},
			RoleName: (*user.UserRoles)[0].Role.Name,
		})
	}

	response = dto.FetchUserRoleWithPaginatedResponse{
		Users: usersWithRole,
		PaginatedResponse: paginate,
	}
	
	return  response, nil
}