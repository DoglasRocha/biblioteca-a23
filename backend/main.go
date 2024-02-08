package main

import (
	"biblioteca-a23/controllers"
	"biblioteca-a23/database"
	"log"
	"net/http"

	//"biblioteca-a23/models"
	//"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func main() {

	// teste database
	database.ConnectToDatabase()
	database.MakeMigration()

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.CreateReader)
	// db.Create(&models.User{Name: "Doglas", Surname: "Rocha", Email: "doglas", Password: "rocha"})
	/*var user models.User
	db.First(&user)
	fmt.Println(user)

	db.Delete(&user)*/
	log.Fatal(http.ListenAndServe(":8080", router))
}
