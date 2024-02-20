package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"main/services/server"
	"main/services/storage"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbManagementType := os.Getenv("DB_MANAGEMENT_TYPE")

	if dbManagementType == "gorm" {
		log.Println("sql gorm initialized")
		//storage.DBGormInit()
	}
	if dbManagementType == "sql" {
		log.Println("sql initialized")
		storage.InitDB()
	}

	server.Init()
}
