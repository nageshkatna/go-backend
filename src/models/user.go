package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"size:256;not null;"`
	LastName string `gorm:"size:256;not null;"`
	Email string `gorm:"size:256;not null;unique"`
	Password string `gorm:"size:16;not null;"`
	UserRoles *[]Role
}

type Role struct {
	gorm.Model
	Name string `gorm:"size:256;not null;"`
	Permissions string `gorm:"size:500;not null"`
}