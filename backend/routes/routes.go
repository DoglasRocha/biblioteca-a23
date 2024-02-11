package routes

import (
	"biblioteca-a23/controllers"
	"log"
	"net/http"

	//"biblioteca-a23/models"
	//"fmt"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cadastro", controllers.CreateReader).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
