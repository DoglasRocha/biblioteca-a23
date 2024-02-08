package main

import (
	"biblioteca-a23/database"
	//"biblioteca-a23/models"
	//"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func main() {

	// teste database
	db, err := database.ConnectToDatabase()

	if err != nil {
		panic("erro na conexao")
	}

	database.MakeMigration(db)

	// db.Create(&models.User{Name: "Doglas", Surname: "Rocha", Email: "doglas", Password: "rocha"})
	/*var user models.User
	db.First(&user)
	fmt.Println(user)

	db.Delete(&user)*/
}
