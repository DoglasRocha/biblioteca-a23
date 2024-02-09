package main

import (
	"biblioteca-a23/database"
	"biblioteca-a23/routes"
	//"biblioteca-a23/models"
	//"fmt"
)

func main() {

	// init database
	database.ConnectToDatabase()
	database.MakeMigration()

	// init router
	routes.SetupRoutes()
}
