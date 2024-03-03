package routes

import (
	"biblioteca-a23/controllers"

	"github.com/gorilla/mux"
)

func SetupDeleteRoutes(router *mux.Router) {
	router.HandleFunc("/api/admin/emprestimos/rejeitar/{request_id}", controllers.DenyRequest).Methods("DELETE")
	router.HandleFunc("/api/admin/livros/deletar/{book_id}", controllers.DeleteBook).Methods("DELETE")

}
