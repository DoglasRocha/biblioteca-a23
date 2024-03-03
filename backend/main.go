package main

import (
	"biblioteca-a23/database"
	"biblioteca-a23/routes"
	"log/slog"

	"github.com/joho/godotenv"
	//"biblioteca-a23/models"
	//"fmt"
)

func main() {

	// loads environment variables
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file")
	}

	// init database
	database.ConnectToDatabase()
	database.MakeMigration()

	// init router
	routes.SetupRoutes()
}
