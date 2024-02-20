package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {

	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionData := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	db, err = sql.Open(
		"postgres", connectionData)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {

		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}
