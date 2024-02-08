package database

import (
	m "biblioteca-a23/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	err := godotenv.Load("database/.env")

	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/biblioteca?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func MakeMigration(db *gorm.DB) {
	db.AutoMigrate(
		&m.User{},
		&m.Reader{},
		&m.Admin{},
		&m.Book{},
		&m.Copy{},
		&m.Loan{},
	)
}
