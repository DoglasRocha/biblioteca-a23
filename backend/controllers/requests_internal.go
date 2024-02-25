package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"fmt"
	"net/http"
	"time"
)

func get_open_requests() ([]models.Request, error) {
	var requests []models.Request

	err := database.DB.Where("is_accepted = ?", false).Find(&requests).Error
	if err != nil {
		return []models.Request{}, err
	}

	for i := range requests {
		err = database.PopulateRequest(&requests[i], requests[i].ID)
		if err != nil {
			return []models.Request{}, err
		}
	}

	return requests, nil
}

func create_loan_in_db(request models.Request, w http.ResponseWriter, r *http.Request) error {
	var loan models.Loan
	var copy models.Copy
	// checks if there is available copies
	var available_copies int64
	query := database.DB.Model(&models.Copy{}).Where(
		"book_id = ? AND is_borrowed = ?", request.BookID, false,
	)
	err := query.Count(&available_copies).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao contar cópias disponíveis do livro")
		return err
	}

	if available_copies == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, "Não há cópias disponíveis do livro")
		return fmt.Errorf("")
	}

	// gets copy
	query.First(&copy)

	// creates loan
	loan.Copy = copy
	loan.CopyID = copy.ID
	loan.Request = request
	loan.RequestID = request.ID
	loan.HasRenewed = false
	loan.HasReturned = false
	loan.StartDate = get_next_saturday()
	loan.ReturnDate = loan.StartDate.Add(7 * 24 * time.Hour)

	// validates loan
	err = loan.Validate()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao validar empréstimo")
		return err
	}

	err = database.DB.Create(&loan).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro criar empréstimo no banco de dados")
		return err
	}

	// makes copy unavailable
	copy.IsBorrowed = true
	database.DB.Save(&copy)

	return nil
}

func get_next_saturday() time.Time {
	var next_saturday time.Time

	today := time.Now()
	weekday := today.Weekday()

	if weekday == time.Saturday {
		next_saturday = today.Add(7 * 24 * time.Hour)
	} else {
		days_to_saturday := time.Duration(6 - weekday)
		next_saturday = today.Add(days_to_saturday * 24 * time.Hour)
	}

	return next_saturday
}
