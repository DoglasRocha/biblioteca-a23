package routes

import (
	"biblioteca-a23/controllers"
	"log"
	"net/http"

	//"biblioteca-a23/models"
	//"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cadastro", controllers.CreateReader).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"Access-Control-Allow-Origin",
		"Accept",
		"Accept-Language",
		"Content-Language",
		"Origin",
	})

	originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS"})
	credentialsOk := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(
		":8080",
		handlers.CORS(headersOk, originsOk, methodsOk, credentialsOk)(router),
	))
}
