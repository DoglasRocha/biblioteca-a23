package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type newPassword struct {
	NewPassword string `json:"new_password" validate:"required,gte=8"`
}

func (new_password *newPassword) Validate() error {
	return models.Validator.Struct(new_password)
}

func read_request_body(w http.ResponseWriter, r *http.Request) (models.Reader, newPassword, error) {
	var reader models.Reader
	var new_password newPassword

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao ler requisição")
		return models.Reader{}, newPassword{}, err
	}

	err = json.Unmarshal(body, &reader.User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao ler corpo da requisição")
		return models.Reader{}, newPassword{}, err
	}

	err = json.Unmarshal(body, &reader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao ler corpo da requisição")
		return models.Reader{}, newPassword{}, err
	}

	err = json.Unmarshal(body, &new_password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao ler corpo da requisição")
		return models.Reader{}, newPassword{}, err
	}

	return reader, new_password, nil
}

func find_reader_by_id(id int) (models.Reader, error) {
	var reader models.Reader

	err := database.DB.Where("user_id = ?", id).First(&reader).Error
	if err != nil {
		return models.Reader{}, err
	}

	err = database.DB.Where("id = ?", id).First(&reader.User).Error
	if err != nil {
		return models.Reader{}, err
	}

	return reader, nil
}
