package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterBook(w http.ResponseWriter, r *http.Request) {
	// checks if user is admin
	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	err := register_book(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao criar livro")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Livro cadastrado")
}

func SearchBooksByName(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	query_params := r.URL.Query()
	status, books = search_available_books_by_name(query_params)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(books)
}

func SearchBookById(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	id := mux.Vars(r)["id"]

	// search in db
	database.DB.First(&book, id)

	// set number of available copies
	var available_copies int64
	database.DB.Model(&models.Copy{}).
		Where("book_id = ? AND is_borrowed = ?", book.ID, false).
		Count(&available_copies)
	book.CopiesCount = uint(available_copies)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func SearchBooksByNameAdmin(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	query_params := r.URL.Query()
	status, books = search_book_by_name(query_params)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(books)
}

func SearchBookByIdAdmin(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	id := mux.Vars(r)["id"]

	database.DB.First(&book, id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book_fields_to_update, not_updated_book models.Book

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// json from request
	err := json.NewDecoder(r.Body).Decode(&book_fields_to_update)
	if err != nil {
		slog.Warn(
			"Erro ao decodificar corpo da request",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao processar requisição")
		return
	}

	// get parameters
	id := mux.Vars(r)["id"]
	err = database.DB.First(&not_updated_book, id).Error
	if err != nil {
		slog.Warn(
			"Erro ao buscar livro a ser atualizado",
			"err", err,
			"id", id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Não há livro com esse identificador")
		return
	}
	book_id := int(not_updated_book.ID)

	// logic to create or delete copies
	copies_diff := int(book_fields_to_update.CopiesCount) - int(not_updated_book.CopiesCount)
	if copies_diff > 0 {
		err = create_copies(copies_diff, book_id)
		if err != nil {
			slog.Warn(
				"Erro ao criar cópias do livro",
				"err", err,
				"book_id", book_id,
			)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro ao criar cópias do livro")
			return
		}
	} else if copies_diff < 0 {
		err = delete_copies(-copies_diff, book_id)
		if err != nil {
			slog.Warn(
				"Erro ao deletar cópias do livro",
				"err", err,
				"book_id", book_id,
			)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro ao deletar cópias do livro")
			return
		}
	}

	// makes sure the id update is the id from the url
	book_fields_to_update.ID = not_updated_book.ID

	// updates the count of copies
	var copies_count int64
	database.DB.Model(&models.Copy{}).Where("book_id = ?", book_id).Count(&copies_count)
	book_fields_to_update.CopiesCount = uint(copies_count)

	// updates the book
	err = database.DB.Save(&book_fields_to_update).Error
	if err != nil {
		slog.Warn(
			"Erro ao atualizar livro na base de dados",
			"err", err,
			"book_id", book_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao atualizar livro")
		return
	}

	slog.Info(
		"Livro atualizado",
		"book_id", book_id,
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Livro atualizado com sucesso")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get parameters
	book_id := mux.Vars(r)["book_id"]

	// check if book exists
	err := database.DB.First(&book, book_id).Error
	if err != nil {
		slog.Warn(
			"Não há livro com esse id",
			"err", err,
			"book_id", book_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Não há livro com esse identificador")
		return
	}

	// check if book has borrowed copies
	var borrowed_copies int64
	err = database.DB.
		Where("book_id = ? AND is_borrowed = ?", book_id, true).
		Find(&[]models.Copy{}).
		Count(&borrowed_copies).Error
	if err != nil || borrowed_copies != 0 {
		slog.Warn(
			"Não é possivel deletar o livro, pois há cópias emprestadas",
			"err", err,
			"book_id", book_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Não é possivel deletar o livro, pois há cópias emprestadas")
		return
	}

	// delete copies
	err = database.DB.Where("book_id = ?", book_id).Delete(&[]models.Copy{}).Error
	if err != nil {
		slog.Warn(
			"Erro ao deletar cópias do livro",
			"err", err,
			"book_id", book_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao deletar cópias do livro")
		return
	}

	// delete book
	err = database.DB.Delete(&book).Error
	if err != nil {
		slog.Warn(
			"Erro ao deletar livro",
			"err", err,
			"book_id", book_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao deletar livro")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Livro deletado")
}
