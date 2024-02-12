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
	router.HandleFunc("/api/cadastro", controllers.RegisterReader).Methods("POST")
	router.HandleFunc("/api/admin/cadastro", controllers.RegisterAdmin).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/admin/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/check", controllers.CheckReader).Methods("POST")
	router.HandleFunc("/api/admin/check", controllers.CheckAdmin).Methods("POST")

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
