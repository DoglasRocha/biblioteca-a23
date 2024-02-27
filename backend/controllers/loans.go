package controllers

import (
	//"biblioteca-a23/models"

	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func GetUserActiveLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	user_id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	loan, err = get_active_user_loan(user_id, w)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&loan)
}

func RenewLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get loan id
	loan_id := mux.Vars(r)["loan_id"]

	// get loan in DB
	err := database.DB.First(&loan, loan_id).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Não existe empréstimo com esse identificador!")
		return
	}

	// renew loan
	loan.HasRenewed = true
	loan.ReturnDate = loan.ReturnDate.Add(7 * 24 * time.Hour)

	err = database.DB.Save(&loan).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao renovar empréstimo")
		return
	}

	err = database.PopulateLoan(&loan, loan.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar dados do empréstimo")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&loan)
}

func ReturnLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get loan
	loan_id := mux.Vars(r)["loan_id"]
	err := database.DB.First(&loan, loan_id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao buscar empréstimo")
		return
	}

	err = database.PopulateLoan(&loan, loan.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao buscar dados do empréstimo")
		return
	}

	// mark as returned
	loan.HasReturned = true
	loan.Copy.IsBorrowed = false

	// save in db
	err = database.DB.Save(&loan.Copy).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao salvar devolucação no banco de dados")
		return
	}

	err = database.DB.Save(&loan).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao salvar devolucação no banco de dados")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Empréstimo devolvido com sucesso")
}
