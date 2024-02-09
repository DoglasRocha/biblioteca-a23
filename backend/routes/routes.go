package routes

import (
	"biblioteca-a23/controllers"
	"log"
	"net/http"

	//"biblioteca-a23/models"
	//"fmt"

	"github.com/gorilla/mux"
)

func SetupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cadastro", controllers.CreateReader)

	log.Fatal(http.ListenAndServe(":8080", router))
}
