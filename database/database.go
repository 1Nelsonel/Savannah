package database

import (
	"fmt"
	"os"

	"github.com/1Nelsonel/Savannah/models"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil{
		danger := color.New(color.FgHiRed)
	 	danger.Printf("Error loading .env file: %s", err)
	}
	
	// Access environment variables
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSL_MODE"),
        os.Getenv("DB_TIMEZONE"),
    )
	// var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		danger := color.New(color.FgHiRed)
		danger.Println("Error connecting to the database.....")
	}
	green := color.New(color.FgGreen)
	green.Println("Database connected successfully.......")

	// Auto Migration
	if err := db.AutoMigrate(
		&models.Customer{},
		&models.Order{}); err != nil {
		danger := color.New(color.FgHiRed)
		danger.Println("Auto migration failed:: ", err)
	} else {
		cyan := color.New(color.FgHiCyan)
		cyan.Println("Auto migration successful......")
	}

	DBConn = db
}
