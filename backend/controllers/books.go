package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterBook(w http.ResponseWriter, r *http.Request) {
	// checks if user is admin
	status := is_admin_autenticated(w, r)
	if status != http.StatusOK {
		return
	}

	err := register_book(r)
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

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	query_params := r.URL.Query()
	status, books = search_book_by_name(query_params)

	json.NewEncoder(w).Encode(books)
}

func SearchBookById(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	id := mux.Vars(r)["id"]

	database.DB.First(&book, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func SearchBooksByNameAdmin(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	status := is_admin_autenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	query_params := r.URL.Query()
	status, books = search_book_by_name(query_params)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(books)
}

func SearchBookByIdAdmin(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	status := is_admin_autenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	id := mux.Vars(r)["id"]

	database.DB.First(&book, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
