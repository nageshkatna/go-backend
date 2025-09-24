package main

import (
	"fmt"
	"go-backend/api"
	"go-backend/database"
	"log"
)

func main() {
	fmt.Println("Connecting to DB...")
	err := database.ConnectDB()

	if (err != nil){
		log.Printf("❌ Failed to connect to db %v\n", err)
	}
	api.InitServer()
}