package main

import (
	"biblioteca-a23/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func testes() {
	// teste validator

	user := models.User{
		Name:     "Doglas",
		Surname:  "Rocha",
		Email:    "teste@teste.com",
		Password: "12345678",
	}

	//fmt.Println(user)
	fmt.Println(user.Validate())

	// teste hash
	var teste string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	hash, err := bcrypt.GenerateFromPassword([]byte(teste), bcrypt.MinCost)

	if err != nil {
		panic("aaaaa")
	}

	fmt.Print(string(hash))
}
