package routes

import (
	"biblioteca-a23/controllers"

	"github.com/gorilla/mux"
)

func SetupPatchRoutes(router *mux.Router) {
	router.HandleFunc("/api/admin/autorizar/{admin_id}", controllers.AuthorizeAdmin).Methods("PATCH")
	router.HandleFunc("/api/admin/revogar/{admin_id}", controllers.RevokeAdmin).Methods("PATCH")
	router.HandleFunc("/api/renovar/{loan_id}", controllers.RenewLoan).Methods("PATCH")
	router.HandleFunc("/api/admin/emprestimos/devolver/{loan_id}", controllers.ReturnLoan).Methods("PATCH")
}
