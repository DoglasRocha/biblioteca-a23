package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var Validator *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" validate:"required,gte=3,lte=50" gorm:"type:varchar(50)"`
	Surname  string `json:"surname" validate:"required,gte=3,lte=100" gorm:"type:varchar(100)"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique"`
	Password string `json:"password" validate:"required,gte=8" gorm:"type:varchar(255)"`
}

func (user *User) Validate() error {
	return Validator.Struct(user)
}

type Reader struct {
	gorm.Model
	ID          uint/*`json:"id" */ `gorm:"primary_key"`
	UserID      uint/*`json:"user_id" */ `validate:"required"`
	User        User
	Birthday    time.Time `json:"birthday" validate:"required" gorm:"type:date"`
	Address     string    `json:"address" validate:"required" gorm:"type:text"`
	PhoneNumber string    `json:"phone_number" validate:"required" gorm:"type:varchar(20)"`
	Loans       []Loan
}

func (reader *Reader) Validate() error {
	return Validator.Struct(reader)
}

type Admin struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primary_key"`
	UserID    uint `json:"user_id"`
	User      User
	IsCleared bool `json:"is_creared" validate:"required" gorm:"type:boolean"`
}

func (admin *Admin) Validate() error {
	return Validator.Struct(admin)
}

type Book struct {
	gorm.Model
	ID          uint    `json:"id" gorm:"primary_key"`
	Name        string  `json:"name" validate:"required,lte=3" gorm:"type:varchar(100)"`
	ISBN        *string `json:"isbn" gorm:"type:varchar(20)"`
	Description string  `json:"description" gorm:"type:text"`
	Gender      string  `json:"gender" gorm:"type:varchar(30)"`
	CopiesCount uint    `json:"copies" gorm:"default:1"`
	Copies      []Copy
}

func (book *Book) Validate() error {
	return Validator.Struct(book)
}

type Copy struct {
	gorm.Model
	ID         uint `json:"id" gorm:"primary_key"`
	BookID     uint
	Loans      []Loan
	IsBorrowed bool `json:"is_borrowed" gorm:"default:false"`
}

type Loan struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primary_key"`
	CopyID       int       `json:"copy_id" validate:"required"`
	ReaderID     int       `json:"reader_id" validate:"required"`
	StartDate    time.Time `json:"start_date" gorm:"type:date"`
	ReturnDate   time.Time `json:"return_date" gorm:"type:date"`
	HasRenewed   bool      `json:"has_renewed" gorm:"default:false"`
	HasReturned  bool      `json:"has_returned" gorm:"default:false"`
	IsAuthorized bool      `json:"is_authorized" gorm:"default:false"`
}

func (loan *Loan) Validate() error {
	return Validator.Struct(loan)
}
