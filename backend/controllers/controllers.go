package controllers

import (
	//"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"

	//"fmt"
	"io"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func CreateReader(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler requisição", http.StatusBadRequest)
		return
	}

	reader, err = create_reader(body)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			http.Error(w, "Email já cadastrado", http.StatusBadRequest)
			return
		}
		http.Error(w, "Erro ao criar usuário", http.StatusBadRequest)
		return
	}

	// hides password from response
	reader.User.Password = nil
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reader)
}

func Login(w http.ResponseWriter, r *http.Request) {

}
