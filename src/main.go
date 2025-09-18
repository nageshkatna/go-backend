package main

import (
	"fmt"
	"go-backend/database"
	"log"
)

func main() {
	fmt.Println("Connecting to DB...")
	err := database.ConnectDB()

	if (err != nil){
		log.Println("❌ Failed to connect to db %v", err)
	}
}