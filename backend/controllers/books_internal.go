package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
)

func register_book(request *http.Request) error {
	var book models.Book
	var copy models.Copy

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		slog.Warn(
			"Erro ao decodificar corpo da request",
			"err", err,
		)
		return err
	}

	err = book.Validate()
	if err != nil {
		slog.Warn(
			"Erro ao validar livro",
			"err", err,
		)
		return err
	}

	err = database.DB.Create(&book).Error
	if err != nil {
		slog.Warn(
			"Erro ao criar livro na base de dados",
			"err", err,
		)
		return err
	}

	copy.BookID = book.ID
	err = database.DB.Create(&copy).Error
	if err != nil {
		slog.Warn(
			"Erro ao criar cópia do livro na base de dados",
			"err", err,
			"id", book.ID,
		)
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

func search_available_books_by_name(query_params url.Values) (int, []models.Book) {
	var books []models.Book
	var available_copies_ids []uint

	// available books
	err := database.DB.Model(&models.Copy{}).
		Distinct("book_id").
		Select("book_id").
		Where("is_borrowed = ?", false).
		Scan(&available_copies_ids).Error
	if err != nil {
		slog.Warn(
			"Erro ao pesquisar livros disponíveis",
			"err", err,
		)
		return http.StatusNotFound, books
	}

	book_name := query_params.Get("name")

	if book_name == "" && len(query_params) != 0 {
		return http.StatusNotFound, books
	}

	err = database.DB.
		Where("id IN ? AND name LIKE ?", available_copies_ids, book_name+"%").
		Limit(50).
		Find(&books).Error
	if err != nil {
		slog.Warn(
			"Erro ao pesquisar livros disponíveis com nome",
			"err", err,
			"book_name", book_name,
		)
		return http.StatusNotFound, books
	}

	return http.StatusOK, books
}
