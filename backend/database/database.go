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
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_URL"),
		os.Getenv("MYSQL_DATABASE"),
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
	err := DB.First(user, id).Error

	user.Password = nil

	return err
}

func PopulateReader(reader *m.Reader, id uint) error {
	err := DB.First(reader, id).Error
	if err != nil {
		return err
	}

	return PopulateUser(&reader.User, reader.UserID)
}

func PopulateAdmin(admin *m.Admin, id uint) error {
	err := DB.First(admin, id).Error
	if err != nil {
		return err
	}

	return PopulateUser(&admin.User, admin.UserID)
}

func PopulateBook(book *m.Book, id uint) error {
	return DB.First(book, id).Error
}

func PopulateCopy(copy *m.Copy, id uint) error {
	err := DB.First(copy, id).Error
	if err != nil {
		return err
	}

	return PopulateBook(&copy.Book, copy.BookID)
}

func PopulateRequest(request *m.Request, id uint) error {
	err := DB.First(request, id).Error
	if err != nil {
		return err
	}

	err = PopulateBook(&request.Book, request.BookID)
	if err != nil {
		return err
	}

	return PopulateReader(&request.Reader, request.ReaderID)
}

func PopulateLoan(loan *m.Loan, id uint) error {
	err := DB.First(loan, id).Error
	if err != nil {
		return err
	}

	err = PopulateCopy(&loan.Copy, loan.CopyID)
	if err != nil {
		return err
	}

	return PopulateRequest(&loan.Request, loan.RequestID)
}
