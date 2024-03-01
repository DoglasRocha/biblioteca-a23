package routes

import (
	"biblioteca-a23/controllers"

	"github.com/gorilla/mux"
)

func SetupPostRoutes(router *mux.Router) {
	router.HandleFunc("/api/cadastro", controllers.RegisterReader).Methods("POST")
	router.HandleFunc("/api/admin/cadastro", controllers.RegisterAdmin).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/admin/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/check", controllers.CheckReader).Methods("POST")
	router.HandleFunc("/api/admin/check", controllers.CheckAdmin).Methods("POST")
	router.HandleFunc("/api/admin/livros/cadastrar", controllers.RegisterBook).Methods("POST")
	router.HandleFunc("/api/emprestar/{book_id}", controllers.CreateRequest).Methods("POST")
	router.HandleFunc("/api/admin/emprestimos/aprovar/{request_id}", controllers.ApproveRequest).Methods("POST")
}
