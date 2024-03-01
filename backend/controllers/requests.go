package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"log/slog"
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

	// checks if user has active loan
	if has_active_loan(user_id, w) {
		return
	}

	// checks if book exists
	err = database.DB.Where("id = ?", book_id).First(&book).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Livro não encontrado")
		return
	}

	// create request
	request.BookID = book.ID
	request.ReaderID = uint(user_id)
	request.IsAccepted = false

	// validate request
	err = request.Validate()
	if err != nil {
		slog.Warn(
			"Erro ao validar solicitação de empréstimo",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao validar solicitação")
		return
	}

	// create request in db
	err = database.DB.Create(&request).Error
	if err != nil {
		slog.Warn(
			"Erro ao criar solicitação na base de dados",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao criar solicitação")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Solicitação criada com sucesso.")
}

func GetOpenRequests(w http.ResponseWriter, r *http.Request) {
	var requests []models.Request

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	requests, err := get_open_requests()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao encontrar solicitações em aberto")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&requests)
}

func ApproveRequest(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get request id from url
	request_id := mux.Vars(r)["request_id"]

	// checks if request exists and is not approved
	err := database.DB.
		Where("id = ? AND is_accepted = ?", request_id, false).
		First(&request).Error
	if err != nil {
		slog.Warn(
			"Erro ao buscar solicitação na base de dados",
			"err", err,
			"request_id", request_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar solicitacao de emprestimo")
		return
	}

	// checks if user has active loan
	if has_active_loan(int(request.ReaderID), w) {
		return
	}

	err = create_loan_in_db(request, w)
	if err != nil {
		return
	}

	request.IsAccepted = true
	err = database.DB.Save(&request).Error
	if err != nil {
		slog.Warn(
			"Erro ao atualizar solicitação na base de dados",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao aceitar solicitação")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Empréstimo aprovado com sucesso")
}

func DenyRequest(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get request id from url
	request_id := mux.Vars(r)["request_id"]

	// checks if request exists and is not approved
	err := database.DB.
		Where("id = ? AND is_accepted = ?", request_id, false).
		First(&request).Error
	if err != nil {
		slog.Warn(
			"Erro ao buscar solicitação na base de dados",
			"err", err,
			"request_id", request_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar solicitacao de emprestimo")
		return
	}

	// deletes request
	err = database.DB.Delete(&request).Error
	if err != nil {
		slog.Warn(
			"Erro ao deletar solicitação",
			"err", err,
			"request_id", request_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao deletar solicitação")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Solicitação rejeitada com sucesso")
}

func GetReaderRequests(w http.ResponseWriter, r *http.Request) {
	var requests []models.Request

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	reader_id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	err = database.DB.Where(
		"reader_id = ? AND is_accepted = ?",
		reader_id, false,
	).Find(&requests).Error
	if err != nil {
		slog.Warn(
			"Erro ao buscar solicitações do usuário",
			"err", err,
			"reader_id", reader_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar solicitações do usuário")
		return
	}

	for i := range requests {
		err = database.PopulateRequest(&requests[i], requests[i].ID)
		if err != nil {
			slog.Warn(
				"Erro ao popular solicitação",
				"err", err,
				"requests_id", requests[i].ID,
			)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro ao buscar solicitações do usuário")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&requests)
}
