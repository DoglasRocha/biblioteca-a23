package database

import (
	m "biblioteca-a23/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {

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
		&m.Request{},
		&m.Copy{},
		&m.Loan{},
	)
}

// populating models
func PopulateUser(user *m.User, id uint) error {
	return DB.First(user, id).Error
}

func PopulateReader(reader *m.Reader, id uint) error {
	return PopulateUser(&reader.User, reader.UserID)
}

func PopulateAdmin(admin *m.Admin, id uint) error {
	return PopulateUser(&admin.User, admin.UserID)
}

func PopulateBook(book *m.Book, id uint) error {
	return DB.First(book, id).Error
}

func PopulateCopy(copy *m.Copy, id uint) error {
	return PopulateBook(&copy.Book, copy.BookID)
}

func PopulateRequest(request *m.Request, id uint) error {
	err := PopulateBook(&request.Book, request.BookID)
	if err != nil {
		return err
	}

	return PopulateReader(&request.Reader, request.ReaderID)
}

func PopulateLoan(loan *m.Loan, id uint) error {
	err := PopulateCopy(&loan.Copy, loan.CopyID)
	if err != nil {
		return err
	}

	return PopulateRequest(&loan.Request, loan.RequestID)
}
