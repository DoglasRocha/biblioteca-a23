package routes

import (
	"biblioteca-a23/controllers"

	"github.com/gorilla/mux"
)

func SetupGetRoutes(router *mux.Router) {
	router.HandleFunc("/api/livros/buscar", controllers.SearchBooksByName).Methods("GET")
	router.HandleFunc("/api/livros/buscar/{id}", controllers.SearchBookById).Methods("GET")
	router.HandleFunc("/api/admin/livros/buscar", controllers.SearchBooksByNameAdmin).Methods("GET")
	router.HandleFunc("/api/admin/livros/buscar/{id}", controllers.SearchBookByIdAdmin).Methods("GET")
	router.HandleFunc("/api/minhaconta", controllers.GetMyAccount).Methods("GET")
	router.HandleFunc("/api/admin/minhaconta", controllers.GetMyAccountAdmin).Methods("GET")
	router.HandleFunc("/api/admin/admins", controllers.ListAdmins).Methods("GET")
	router.HandleFunc("/api/admin/emprestimos/ativos", controllers.ActiveLoans).Methods("GET")
	router.HandleFunc("/api/admin/emprestimos/historico", controllers.HistoryOfLoans).Methods("GET")
	router.HandleFunc("/api/admin/emprestimos/solicitacoes", controllers.GetOpenRequests).Methods("GET")
	router.HandleFunc("/api/solicitacoes", controllers.GetReaderRequests).Methods("GET")
	router.HandleFunc("/api/emprestimos", controllers.GetUserLoans).Methods("GET")
	router.HandleFunc("/api/emprestimos/ativo", controllers.GetUserActiveLoan).Methods("GET")
	router.HandleFunc("/api/admin/leitor/{user_id}", controllers.GetReaderData).Methods("GET")
}
