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

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	updated_reader_data, new_password, err = read_request_body(w, r)
	if err != nil {
		return
	}

	err = database.DB.Where("user_id = ?", id).First(&current_reader_data).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar usuário")
		return
	}

	err = database.DB.Where("id = ?", id).First(&current_reader_data.User).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar usuário")
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(*current_reader_data.User.Password),
		[]byte(*updated_reader_data.User.Password),
	)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Senha incorreta!")
		return
	}

	err = new_password.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "A senha deve possuir mais de 8 caracteres")
		return
	}

	// needs to do it otherwise it is going to throw an error
	updated_reader_data.UserID = uint(id)
	updated_reader_data.User.ID = uint(id)

	err = updated_reader_data.Validate()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao validar dados de usuário")
		return
	}

	err = updated_reader_data.User.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao validar dados de usuário")
		return
	}

	if new_password.NewPassword != "" {
		password_hash, err := bcrypt.GenerateFromPassword([]byte(new_password.NewPassword), bcrypt.MinCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro ao hashear a nova senha")
		}

		*updated_reader_data.User.Password = string(password_hash)
	}

	// needs to do it otherwise it is going to throw an error
	updated_reader_data.User.CreatedAt = current_reader_data.User.CreatedAt

	err = database.DB.Save(&updated_reader_data.User).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao salvar usuário no banco de dados")
		return
	}

	err = database.DB.Save(&updated_reader_data).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao salvar usuário no banco de dados")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Usuário atualizado com sucesso")
}

func UpdateMyAccountAdmin(w http.ResponseWriter, r *http.Request) {

}
