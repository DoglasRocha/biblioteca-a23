package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	var request models.Request
	var book models.Book

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	user_id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	// get parameters
	book_id := mux.Vars(r)["book_id"]

	// checks if book exists
	err = database.DB.Where("id = ?", book_id).First(&book).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Livro não encontrado")
		return
	}

	// create request
	request.BookID = int(book.ID)
	request.ReaderID = user_id
	request.IsAccepted = false

	// validate request
	err = request.Validate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao validar solicitação")
		return
	}

	// create request in db
	err = database.DB.Create(&request).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao criar solicitação")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Solicitação criada com sucesso.")
}