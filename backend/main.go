package main

import (
	"biblioteca-a23/database"
	"biblioteca-a23/routes"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	//"biblioteca-a23/models"
	//"fmt"
)

func main() {
	// configura logger
	f, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	logger := slog.New(slog.NewJSONHandler(f, nil))
	slog.SetDefault(logger)

	// loads environment variables
	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// init database
	database.ConnectToDatabase()
	database.MakeMigration()

	// init router
	routes.SetupRoutes()
}
