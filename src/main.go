package main

import (
	"fmt"
	"go-backend/api"
	"go-backend/database"
	"log"
)

//	@title			Go Backend user api
//	@version		1.0
//	@description	This is a Go backend user API.

//	@host		localhost:5000
//	@BasePath	/api/v1/

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and then your JWT token.

func main() {
	fmt.Println("Connecting to DB...")
	err := database.ConnectDB()

	if err != nil {
		log.Printf("❌ Failed to connect to db %v\n", err)
	}
	api.InitServer()
}
