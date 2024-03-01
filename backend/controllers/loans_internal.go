package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"fmt"
	"log/slog"
	"net/http"

	"gorm.io/gorm"
)

func get_active_loans(w http.ResponseWriter) ([]models.Loan, error) {
	var active_loans []models.Loan

	// get loans
	err := database.DB.Where("has_returned = ?", false).Find(&active_loans).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar empréstimos ativos",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro buscar empréstimos ativos na base de dados")
		return []models.Loan{}, err
	}

	for i := range active_loans {
		err = database.PopulateLoan(&active_loans[i], active_loans[i].ID)
		if err != nil {
			slog.Warn(
				"Erro ao popular empréstimo",
				"err", err,
				"loan_id", active_loans[i].ID,
			)
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
		slog.Warn(
			"Erro ao buscar empréstimos inativos",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro buscar empréstimos inativos na base de dados")
		return []models.Loan{}, err
	}

	for i := range history_of_loans {
		err = database.PopulateLoan(&history_of_loans[i], history_of_loans[i].ID)
		if err != nil {
			slog.Warn(
				"Erro ao popular empréstimo",
				"err", err,
				"loan_id", history_of_loans[i].ID,
			)
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
		slog.Warn(
			"Erro ao buscar solicitacoes aceitas do usuario",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar solicitações do usuario")
		return true
	}

	// gets active loans from user
	err = database.DB.Model(&models.Loan{}).
		Where("has_returned = ? AND request_id IN ?", false, requests_from_user).
		Count(&active_loans).Error

	if err != nil {
		slog.Warn(
			"Erro ao buscar emprestimos ativos do usuario",
			"err", err,
			"user_id", user_id,
		)
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

func get_user_loans(user_id int, w http.ResponseWriter) ([]models.Loan, error) {
	var loans []models.Loan

	var requests []uint
	// request ids from user
	err := database.DB.Model(&models.Request{}).
		Select("id").
		Where(
			"reader_id = ? AND is_accepted = ?", user_id, true,
		).
		Scan(&requests).Error

	if err != nil {
		slog.Warn(
			"Erro ao buscar solicitacoes aceitas do usuario",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar solicitações do usuario")
		return []models.Loan{}, err
	}

	// get all loans from user
	err = database.DB.
		Where("request_id IN ?", requests).
		Find(&loans).Error

	if err != nil {
		slog.Warn(
			"Erro ao buscar emprestimos do usuario",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, "Erro ao buscar empréstimos ativos do usuario")
		return []models.Loan{}, err
	}

	// populate loans
	for i := range loans {
		err = database.PopulateLoan(&loans[i], loans[i].ID)
		if err != nil {
			slog.Warn(
				"Erro ao popular emprestimo",
				"err", err,
				"loan_id", loans[i].ID,
			)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro buscar empréstimo ", loans[i].ID, " na base de dados")
			return []models.Loan{}, err
		}
	}

	return loans, nil
}

func get_active_user_loan(user_id int, w http.ResponseWriter) (models.Loan, error) {
	var loan models.Loan

	var request models.Request
	// get last accepted loan, because there is only one active loan per user
	err := database.DB.Model(&models.Request{}).
		Select("id").
		Where(
			"reader_id = ? AND is_accepted = ?", user_id, true,
		).
		Last(&request).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusOK)
			return models.Loan{}, err
		}
		slog.Warn(
			"Erro ao buscar solicitacoes do usuario",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar solicitações do usuario")
		return models.Loan{}, err
	}

	// get active loan from user
	err = database.DB.
		Where("request_id = ? AND has_returned = ?", request.ID, false).
		First(&loan).Error

	if err != nil {
		slog.Warn(
			"Erro ao buscar emprestimos ativos do usuario",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, "Erro ao buscar empréstimo ativo do usuario")
		return models.Loan{}, err
	}

	// populate loan
	err = database.PopulateLoan(&loan, loan.ID)
	if err != nil {
		slog.Warn(
			"Erro ao popular emprestimo",
			"err", err,
			"loan_id", loan.ID,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro buscar empréstimo ", loan.ID, " na base de dados")
		return models.Loan{}, err
	}

	return loan, nil
}
