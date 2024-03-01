package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetMyAccount(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	err = database.DB.Where("user_id = ?", id).First(&reader).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao localizar usuário leitor na base de dados")
		return
	}

	err = database.DB.First(&reader.User, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao localizar usuário leitor na base de dados")
		return
	}

	// hide password from response
	reader.User.Password = nil

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&reader)
}

func GetMyAccountAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	err = database.DB.Where("user_id = ?", id).First(&admin).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao localizar usuário leitor na base de dados")
		return
	}

	err = database.DB.First(&admin.User, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao localizar usuário leitor na base de dados")
		return
	}

	// hide password from response
	admin.User.Password = nil

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&admin)
}

func UpdateMyAccount(w http.ResponseWriter, r *http.Request) {
	var updated_reader_data models.Reader
	var current_reader_data models.Reader
	var new_password newPassword

	if status := is_reader_authenticated(w, r); status != http.StatusOK {
		return
	}

	id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	updated_reader_data, new_password, err = read_reader_request_body(w, r)
	if err != nil {
		return
	}

	current_reader_data, err = find_reader_by_id(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar usuário")
		return
	}

	// checks if password is correct
	if err = bcrypt.CompareHashAndPassword(
		[]byte(*current_reader_data.User.Password),
		[]byte(*updated_reader_data.User.Password),
	); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Senha incorreta!")
		return
	}

	// password comes unhashed from request
	updated_reader_data.User.Password = current_reader_data.User.Password

	// needs to do it otherwise it is going to throw an error
	updated_reader_data.UserID = uint(id)
	updated_reader_data.User.ID = uint(id)

	// updates password
	if new_password.NewPassword != "" {
		if err = update_password(w, new_password, updated_reader_data.User); err != nil {
			return
		}
	}

	if err = updated_reader_data.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao validar dados de usuário")
		return
	}

	// needs to do it otherwise it is going to throw an error
	updated_reader_data.User.CreatedAt = current_reader_data.User.CreatedAt
	updated_reader_data.CreatedAt = current_reader_data.CreatedAt

	if err = update_reader_in_db(updated_reader_data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao salvar usuário no banco de dados")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Usuário atualizado com sucesso")
}

func UpdateMyAccountAdmin(w http.ResponseWriter, r *http.Request) {
	var updated_admin_data models.Admin
	var current_admin_data models.Admin
	var new_password newPassword

	if status := is_admin_authenticated(w, r); status != http.StatusOK {
		return
	}

	id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	updated_admin_data, new_password, err = read_admin_request_body(w, r)
	if err != nil {
		return
	}

	current_admin_data, err = find_admin_by_id(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar usuário")
		return
	}

	// checks if password is correct
	if err = bcrypt.CompareHashAndPassword(
		[]byte(*current_admin_data.User.Password),
		[]byte(*updated_admin_data.User.Password),
	); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Senha incorreta!")
		return
	}

	// password comes unhashed from request
	updated_admin_data.User.Password = current_admin_data.User.Password

	// needs to do it otherwise it is going to throw an error
	updated_admin_data.UserID = uint(id)
	updated_admin_data.User.ID = uint(id)

	// updates password
	if new_password.NewPassword != "" {
		if err = update_password(w, new_password, updated_admin_data.User); err != nil {
			return
		}
	}

	if err = updated_admin_data.Validate(); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao validar dados de usuário")
		return
	}

	// needs to do it otherwise it is going to throw an error
	updated_admin_data.User.CreatedAt = current_admin_data.User.CreatedAt
	updated_admin_data.CreatedAt = current_admin_data.CreatedAt

	if err = update_admin_in_db(updated_admin_data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao salvar usuário no banco de dados")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Usuário atualizado com sucesso")
}
