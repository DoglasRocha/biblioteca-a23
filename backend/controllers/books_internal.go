package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"net/http"
)

func register_book(request *http.Request) error {
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		return err
	}

	err = book.Validate()
	if err != nil {
		return err
	}

	err = database.DB.Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}
