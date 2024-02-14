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
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}
	status, message, err := check_admin(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
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

	// check if user is authorized
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}

	status, message, err := check_reader(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
		return
	}

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

func SearchBookById(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// check if user is authorized
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}

	status, message, err := check_reader(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
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

	// check if user is authorized
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}

	status, message, err := check_admin(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
		return
	}

	// get parameters
	query_params := r.URL.Query()
	/*book_name := query_params.Get("name")

	if book_name == "" && len(query_params) != 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}
	database.DB.Where("name LIKE ?", book_name+"%").Limit(50).Find(&books)*/
	status, books = search_book_by_name(query_params)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(books)
}

func SearchBookByIdAdmin(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// check if user is authorized
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}

	status, message, err := check_admin(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
		return
	}

	// get parameters
	id := mux.Vars(r)["id"]

	database.DB.First(&book, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
