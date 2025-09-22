package database

import (
	"log"

	"go-backend/database/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func ConnectDB() error{
	var err error
	dbString := "host=localhost user=postgres password=postgres dbname=postgres port=5434"

	dbClient, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})
	
	if(err != nil) {
		return err
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if(err != nil){
		return err
	}

	sqlDB.SetConnMaxIdleTime(15)
	sqlDB.SetConnMaxLifetime(5)
	sqlDB.SetMaxOpenConns(100)

	
	log.Println("✅ Successfully connected to DB")

	SeedDB()
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	con, _ := dbClient.DB()
	con.Close()
}

func SeedDB() {
	seeds.SeedRoles(dbClient)
	seeds.SeedUsers(dbClient)
	seeds.SeedUserRoles(dbClient)
}