package routes

import (
	"log"
	"net/http"
	"os"

	//"biblioteca-a23/models"
	//"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
	router := mux.NewRouter()

	SetupPostRoutes(router)
	SetupDeleteRoutes(router)
	SetupGetRoutes(router)
	SetupPutRoutes(router)
	SetupPatchRoutes(router)

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
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"})
	credentialsOk := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(
		":"+os.Getenv("PORT"),
		handlers.CORS(headersOk, originsOk, methodsOk, credentialsOk)(router),
	))
}
