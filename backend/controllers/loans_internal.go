package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"fmt"
	"net/http"
)

func get_active_loans(w http.ResponseWriter) ([]models.Loan, error) {
	var active_loans []models.Loan

	// get loans
	err := database.DB.Where("has_returned = ?", false).Find(&active_loans).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro buscar empréstimos ativos na base de dados")
		return []models.Loan{}, err
	}

	for i := range active_loans {
		err = database.PopulateLoan(&active_loans[i], active_loans[i].ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro buscar empréstimo ", active_loans[i].ID, " na base de dados")
			return []models.Loan{}, err
		}
	}

	return active_loans, nil
}

func get_history_of_loans(w http.ResponseWriter) ([]models.Loan, error) {
	var history_of_loans []models.Loan

	// get loans
	err := database.DB.Where("has_returned = ?", true).Find(&history_of_loans).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro buscar empréstimos ativos na base de dados")
		return []models.Loan{}, err
	}

	for i := range history_of_loans {
		err = database.PopulateLoan(&history_of_loans[i], history_of_loans[i].ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro buscar empréstimo ", history_of_loans[i].ID, " na base de dados")
			return []models.Loan{}, err
		}
	}

	return history_of_loans, nil
}

func has_active_loan(user_id int, w http.ResponseWriter) bool {
	var requests_from_user []uint
	var active_loans int64

	// request ids from user
	err := database.DB.Model(&models.Request{}).
		Select("id").
		Where(
			"reader_id = ? AND is_accepted = ?", user_id, true,
		).
		Scan(&requests_from_user).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar solicitações do usuario")
		return true
	}

	// gets active loans from user
	err = database.DB.Model(&models.Loan{}).
		Where("has_returned = ? AND request_id IN ?", false, requests_from_user).
		Count(&active_loans).Error

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, "Erro ao buscar empréstimos ativos do usuario")
		return true
	}

	if active_loans != 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, "Usuário já possui empréstimo ativo")
		return true
	}

	return false
}
