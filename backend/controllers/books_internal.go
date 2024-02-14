package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"net/http"
	"net/url"
)

func register_book(request *http.Request) error {
	var book models.Book
	var copy models.Copy

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

	copy.BookID = book.ID
	err = database.DB.Create(&copy).Error
	if err != nil {
		return err
	}

	return nil
}

func search_book_by_name(query_params url.Values) (int, []models.Book) {
	var books []models.Book

	book_name := query_params.Get("name")

	if book_name == "" && len(query_params) != 0 {
		return http.StatusNotFound, books
	}
	database.DB.Where("name LIKE ?", book_name+"%").Limit(50).Find(&books)

	return http.StatusOK, books
}
