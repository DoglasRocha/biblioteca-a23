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
	router.HandleFunc("/api/admin/livros/cadastrar", controllers.RegisterBook).Methods("POST")
	router.HandleFunc("/api/livros/buscar", controllers.SearchBooksByName).Methods("GET")
	router.HandleFunc("/api/livros/buscar/{id}", controllers.SearchBookById).Methods("GET")
	router.HandleFunc("/api/admin/livros/buscar", controllers.SearchBooksByNameAdmin).Methods("GET")
	router.HandleFunc("/api/admin/livros/buscar/{id}", controllers.SearchBookByIdAdmin).Methods("GET")
	router.HandleFunc("/api/admin/livros/editar/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/minhaconta", controllers.GetMyAccount).Methods("GET")
	router.HandleFunc("/api/admin/minhaconta", controllers.GetMyAccountAdmin).Methods("GET")
	router.HandleFunc("/api/minhaconta", controllers.UpdateMyAccount).Methods("PUT")
	router.HandleFunc("/api/admin/minhaconta", controllers.UpdateMyAccountAdmin).Methods("PUT")

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
		":8080",
		handlers.CORS(headersOk, originsOk, methodsOk, credentialsOk)(router),
	))
}
