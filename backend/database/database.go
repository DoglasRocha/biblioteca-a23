package database

import (
	m "biblioteca-a23/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	err := godotenv.Load("database/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/biblioteca?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func MakeMigration() {
	DB.AutoMigrate(
		&m.User{},
		&m.Reader{},
		&m.Admin{},
		&m.Book{},
		&m.Copy{},
		&m.Loan{},
	)
}
