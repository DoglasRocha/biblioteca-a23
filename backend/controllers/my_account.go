package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMyAccount(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader

	status := is_reader_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// getting id from cookie
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao ler cookie")
		return
	}

	id, err := get_id_from_cookie(cookie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao ler identificador")
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

}
