package main

import (
	"biblioteca-a23/database"
	"biblioteca-a23/routes"

	"github.com/joho/godotenv"
	//"biblioteca-a23/models"
	//"fmt"
)

func main() {
	// loads environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// init database
	database.ConnectToDatabase()
	database.MakeMigration()

	// init router
	routes.SetupRoutes()
}
