package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func ConnectDB() error{
	var err error
	dbString := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=postgres port=5434")

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
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	con, _ := dbClient.DB()
	con.Close()
}