package seeds

import (
	"go-backend/helper"
	"go-backend/models"
	"log"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{Name: "admin", Permissions: pq.StringArray{"create", "read", "update", "delete"}},
		{Name: "user", Permissions: pq.StringArray{"read"}},
		{Name: "manager", Permissions: pq.StringArray{"create", "read", "update"}},
	}

	for _, role := range roles {
		err := db.FirstOrCreate(&role, models.Role{Name: role.Name}).Error

		if err != nil {
			log.Printf("❌ Failed to seed Roles table %v\n", err)
			return err
		}
	}
	log.Println("✅ Roles table seeded successfully!")
	return nil
}

func SeedUsers(db *gorm.DB) error {
	hashedPassword, err := helper.HashPassword("password123")

	if err != nil {
		log.Printf("❌ Failed to hash password %v\n", err)
		return err
	}

	users := []models.User{
		{Id: uuid.New(), FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com", Password: hashedPassword},
		{Id: uuid.New(), FirstName: "Jane", LastName: "Smith", Email: "jane.smith@gmail.com", Password: hashedPassword},
		{Id: uuid.New(), FirstName: "Alice", LastName: "Johnson", Email: "alice.johnson@gmail.com", Password: hashedPassword},	
	}

	for _, user := range users {
		err := db.FirstOrCreate(&user, models.User{Email:user.Email}).Error
		if err != nil {
			log.Printf("❌ Failed to seed Users table %v\n", err)
			return err
		}
	}

	log.Println("✅ Users table seeded successfully!")
	return nil
}

func SeedUserRoles(db *gorm.DB) error{
	var users []models.User	
	err := db.Find(&users).Error
	if err != nil {
		log.Printf("❌ Failed to fetch Users table data %v\n", err)
		return err
	}

	var roles []models.Role
	errs := db.Find(&roles).Error

	if errs != nil {
		log.Printf("❌ Failed to fetch Roles table data %v\n", errs)
		return errs
	}
	
	userroles := []models.UserRole{
		{UserId: users[0].Id, RoleId: uint(roles[0].Id)}, // John Doe as admin
		{UserId: users[1].Id, RoleId: uint(roles[1].Id)}, // Jane Smith as user
		{UserId: users[2].Id, RoleId: uint(roles[2].Id)}, // Alice Johnson as manager
	}

	for _, userrole := range userroles {
		err := db.FirstOrCreate(&userrole, models.UserRole{UserId: userrole.UserId, RoleId: userrole.RoleId}).Error
		if(err != nil){
			log.Printf("❌ Failed to seed UserRoles table %v\n", err)
			return err
		}
	}

	log.Println("✅ UserRoles seed is done successfully!")
	return nil
}