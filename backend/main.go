package main

import (
	"fmt"

	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func main() {
	// teste validator

	user := models.User{
		Name:     "Doglas",
		Surname:  "Rocha",
		Email:    "teste@teste.com",
		Password: "12345678",
	}

	//fmt.Println(user)
	fmt.Println(user.Validate(validate))

	// teste hash
	var teste string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	hash, err := bcrypt.GenerateFromPassword([]byte(teste), bcrypt.MinCost)

	if err != nil {
		panic("aaaaa")
	}

	fmt.Print(string(hash))

	// teste database
	db, err := database.ConnectToDatabase()

	fmt.Println(db.First(&user, 1))
}
