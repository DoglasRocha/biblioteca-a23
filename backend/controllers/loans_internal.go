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
