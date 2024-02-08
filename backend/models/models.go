package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,gte=3,lte=50" gorm:"type:varchar(50)"`
	Surname  string `json:"surname" validate:"required,gte=3,lte=100" gorm:"type:varchar(100)"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique"`
	Password string `json:"password" validate:"required,lte=8" gorm:"type:varchar(255)"`
}

type Reader struct {
	gorm.Model
	UserID      int       `json:"user_id" validate:"required"`
	User        User      ``
	Birthday    time.Time `json:"birthday" validate:"required,datetime" gorm:"type:date"`
	Address     string    `json:"address" validate:"required" gorm:"type:text"`
	PhoneNumber string    `json:"phone_number" validate:"required" gorm:"type:varchar(20)"`
}

type Admin struct {
	gorm.Model
	UserID    int `json:"user_id"`
	User      User
	IsCleared bool `json:"is_creared" validate:"required" gorm:"type:boolean"`
}

type Book struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,lte=3" gorm:"type:varchar(100)"`
	ISBN        *string `json:"isbn" gorm:"type:varchar(20)"`
	Description string  `json:"description" gorm:"type:text"`
	Gender      string  `json:"gender" gorm:"type:varchar(30)"`
	Copies      int     `json:"copies" gorm:"default:1"`
}

func (user *User) Validate(validator *validator.Validate) error {
	return validator.Struct(user)
}
