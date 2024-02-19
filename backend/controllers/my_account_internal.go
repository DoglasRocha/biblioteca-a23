package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type newPassword struct {
	NewPassword string `json:"new_password" validate:"required,gte=8"`
}

func (new_password *newPassword) Validate() error {
	return models.Validator.Struct(new_password)
}

func read_reader_request_body(w http.ResponseWriter, r *http.Request) (models.Reader, newPassword, error) {
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

func update_password(w http.ResponseWriter, r *http.Request, new_password newPassword, user models.User) error {
	err := new_password.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "A senha deve possuir mais de 8 caracteres")
		return err
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(new_password.NewPassword), bcrypt.MinCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao hashear a nova senha")
		return err
	}

	*user.Password = string(password_hash)
	return nil
}

func update_reader_in_db(reader models.Reader) error {
	err := database.DB.Save(&reader.User).Error
	if err != nil {
		return err
	}

	err = database.DB.Save(&reader).Error
	if err != nil {
		return err
	}

	return nil
}

func read_admin_request_body(w http.ResponseWriter, r *http.Request) (models.Admin, newPassword, error) {
	var admin models.Admin
	var new_password newPassword

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao ler requisição")
		return models.Admin{}, newPassword{}, err
	}

	err = json.Unmarshal(body, &admin.User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao ler corpo da requisição")
		return models.Admin{}, newPassword{}, err
	}

	err = json.Unmarshal(body, &admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao ler corpo da requisição")
		return models.Admin{}, newPassword{}, err
	}

	err = json.Unmarshal(body, &new_password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao ler corpo da requisição")
		return models.Admin{}, newPassword{}, err
	}

	return admin, new_password, nil
}

func find_admin_by_id(id int) (models.Admin, error) {
	var admin models.Admin

	err := database.DB.Where("user_id = ?", id).First(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}

	err = database.DB.Where("id = ?", id).First(&admin.User).Error
	if err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func update_admin_in_db(admin models.Admin) error {
	err := database.DB.Save(&admin.User).Error
	if err != nil {
		return err
	}

	err = database.DB.Save(&admin).Error
	if err != nil {
		return err
	}

	return nil
}
