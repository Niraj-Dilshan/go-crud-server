package config

import (
	"crud-server/utils"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	portStr := os.Getenv("DATABASE_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Error converting port to integer: %v\n", err)
		// Set a default port or handle the error as appropriate
		port = 5432 // Default PostgreSQL port
	}

	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		dbName   = os.Getenv("DATABASE_NAME")
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
