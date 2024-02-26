package controllers

import (
	//"biblioteca-a23/models"

	"biblioteca-a23/models"
	"encoding/json"
	"net/http"
)

func ActiveLoans(w http.ResponseWriter, r *http.Request) {
	var active_loans []models.Loan

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	active_loans, err := get_active_loans(w)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(active_loans)
}

func HistoryOfLoans(w http.ResponseWriter, r *http.Request) {
	var history_of_loans []models.Loan

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	history_of_loans, err := get_history_of_loans(w)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history_of_loans)
}

func GetUserLoans(w http.ResponseWriter, r *http.Request) {
	var loans []models.Loan

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	user_id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	loans, err = get_user_loans(user_id, w)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&loans)
}
