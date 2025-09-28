package services

import (
	"go-backend/database"
	"go-backend/models"
	"log"
)

type NewRoleService struct {}


func (*NewRoleService) GetRoleById(roleId uint) ([]models.Role, error) {
	db := database.GetDB()

	var role []models.Role

	err := db.First(&role, roleId).Error
	if err != nil {
		log.Printf("Couldn't find the role: %v", err)
		return []models.Role{}, err
	}

	return role, nil
}