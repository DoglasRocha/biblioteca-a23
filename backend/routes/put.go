package routes

import (
	"biblioteca-a23/controllers"

	"github.com/gorilla/mux"
)

func SetupPutRoutes(router *mux.Router) {
	router.HandleFunc("/api/admin/livros/editar/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/minhaconta", controllers.UpdateMyAccount).Methods("PUT")
	router.HandleFunc("/api/admin/minhaconta", controllers.UpdateMyAccountAdmin).Methods("PUT")
}
