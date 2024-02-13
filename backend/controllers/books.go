package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterBook(w http.ResponseWriter, r *http.Request) {
	// checks if user is admin
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}
	status, _, err := check_admin(cookie)
	if err != nil {
		w.WriteHeader(status)
		fmt.Fprintln(w, "Erro ao validar usuário")
		return
	}

	if status != http.StatusOK {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Você não possui acesso a esta parte do sistema")
		return
	}

	err = register_book(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao criar livro")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Livro cadastrado")
}

func SearchBooksByName(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	// get parameters
	query_params := r.URL.Query()
	book_name := query_params.Get("name")

	if book_name == "" && len(query_params) != 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}
	database.DB.Where("name LIKE ?", book_name+"%").Limit(50).Find(&books)

	json.NewEncoder(w).Encode(books)
}
