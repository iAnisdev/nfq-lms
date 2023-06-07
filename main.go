package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}
}

func main() {
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

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	fmt.Println("database connection successful!")
	fmt.Printf(db.Name())
	fmt.Println()
	if err != nil {
		log.Fatal("Error connection to database")
	}
	//gin router setup
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "App is running",
			"connStr": connStr,
		})
	})

	router.Run()
}
