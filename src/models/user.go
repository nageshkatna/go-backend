package models

import (
	"time"

	"github.com/lib/pq"
)

type BaseModel struct {
	Id 	uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	FirstName string `gorm:"size:256;not null;"`
	LastName string `gorm:"size:256;not null;"`
	Email string `gorm:"size:256;not null;unique"`
	Password string `gorm:"size:16;not null;"`
	UserRoles *[]UserRole
}

type Role struct {
	BaseModel
	Name string `gorm:"size:256;not null;"`
	Permissions pq.StringArray `gorm:"size:500;not null;type:permission_enum[];"`
	UserRoles *[]UserRole
}

type UserRole struct {
	BaseModel
	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	UserId uint `gorm:"not null;"`
	RoleId uint `gorm:"not null;"`
}