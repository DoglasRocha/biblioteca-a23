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
	//var active_loans []models.Loan

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	active_loans, err := get_history_of_loans(w)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(active_loans)
}
