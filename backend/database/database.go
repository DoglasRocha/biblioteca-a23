package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectToDatabase() (*gorm.DB, error) {
	err := godotenv.Load("database/.env")

	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/biblioteca",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
