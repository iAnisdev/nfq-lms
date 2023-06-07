package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDB() {
	//postgresql setup
	appEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env")
	}
	DB_HOST := appEnv["DB_HOST"]
	DB_USER := appEnv["DB_USER"]
	DB_PASSWORD := appEnv["DB_PASSWORD"]
	DB_NAME := appEnv["DB_NAME"]
	DB_PORT := appEnv["DB_PORT"]
	DB_SSLMODE := appEnv["DB_SSLMODE"]
	DB_TIMEZONE := appEnv["DB_TIMEZONE"]
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE, DB_TIMEZONE)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	fmt.Printf("database connection successful! connected to : %v", db.Name())
	fmt.Println()

	DB = DbInstance{
		Db: db,
	}
}
