package services

import (
	"errors"
	"fmt"
	"go-backend/api/dto"
	"go-backend/database"
	"go-backend/helper"
	"go-backend/models"
	"log"
	"math"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUserByEmail(req *dto.LoginRequest) ([]models.User, error) {
	var db = database.GetDB()
	email, password := req.Email, req.Password

	var users []models.User
	err := db.Preload("UserRoles.Role").Where(&models.User{Email: email}).First(&users).Error

	if !helper.CheckHashPassword(users[0].Password, password) {
		log.Printf("Password doesn't match")
		return []models.User{}, fmt.Errorf("password doesn't match")
	}

	if err != nil {
		log.Printf("Couldn't find the user: %v", err)
		return []models.User{}, err
	}

	return users, nil
}

func (us *UserService) CreateUser(req *dto.RegisterUserRequest, role *[]models.UserRole) (dto.LoginResponse, error) {
	var db = database.GetDB()
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		log.Printf("❌ Failed to hash password %v\n", err)
		return dto.LoginResponse{}, err
	}

	newUser := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		UserRoles: role,
	}

	rowsAffected := db.FirstOrCreate(&newUser, models.User{Email: newUser.Email}).RowsAffected
	if rowsAffected == 0 {
		c := helper.CustomErrors{
			Message: "User already exists",
			Field:   req.Email,
		}
		return dto.LoginResponse{}, c.CreateUserError()
	} else {
		return dto.LoginResponse{
			UserId: newUser.Id.String(),
		}, nil
	}
}

func (*UserService) GetPaginatedUser(req *dto.Pagination) (dto.FetchUserRoleWithPaginatedResponse, error) {
	var db = database.GetDB()
	var users []models.User
	var total int64

	paginate := dto.PaginatedResponse{}
	response := dto.FetchUserRoleWithPaginatedResponse{}

	if err := db.Model(&models.User{}).Count(&total).Error; err == nil {
		paginate = dto.PaginatedResponse{
			TotalRecords: total,
			TotalPages:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
			Pagination: dto.Pagination{
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}
	}

	page := (req.Page - 1) * req.PageSize

	result := db.Preload("UserRoles.Role").Limit(req.PageSize).Offset(page).Find(&users)

	if result.RowsAffected == 0 {
		return response, fmt.Errorf("no more rows to render")
	}

	var usersWithRole []dto.FetchUserWithRole

	for _, user := range users {
		usersWithRole = append(usersWithRole, dto.FetchUserWithRole{
			UserBaseRequest: dto.UserBaseRequest{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			},
			RoleName: (*user.UserRoles)[0].Role.Name,
		})
	}

	response = dto.FetchUserRoleWithPaginatedResponse{
		Users:             usersWithRole,
		PaginatedResponse: paginate,
	}

	return response, nil
}

func (*UserService) UpdateUser(userId string, updateObj dto.UpdateUserObj) (string, error) {
	var db = database.GetDB()
	updates := map[string]interface{}{}

	if updateObj.FirstName != nil {
		if *updateObj.FirstName != "" {
			updates["first_name"] = *updateObj.FirstName
		}
	}
	if updateObj.LastName != nil {
		if *updateObj.LastName != "" {
			updates["last_name"] = *updateObj.LastName
		}
	}
	if updateObj.Email != nil {
		if *updateObj.Email != "" {
			updates["email"] = *updateObj.Email
		}
	}
	if updateObj.RoleId != nil {
		if *updateObj.RoleId != 0 {
			if err := db.Model(&models.UserRole{}).Where("user_id = ?", userId).Update("role_id", *updateObj.RoleId).Error; err != nil {
				log.Printf("Error while updating user role %v", err)
				return "", err
			}
		}
	}

	if len(updates) == 0 {
		return "", errors.New("no fields to update")
	}

	if err := db.Model(&models.User{}).
		Where("id = ?", userId).Updates(updates).Error; err != nil {
		log.Printf("Error while updating user %v", err)
		return "", err
	}

	return "updated", nil
}

func (*UserService) DeleteUser(userId string) error {
	db := database.GetDB()

	if err := db.Where("id = ?", userId).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
