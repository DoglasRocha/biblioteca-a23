package controllers

import (
	"fmt"

	"io"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func RegisterReader(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler requisição", http.StatusBadRequest)
		return
	}

	err = create_reader(body)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				http.Error(w, "Email já cadastrado", http.StatusNotAcceptable)
				return
			}
		}
		http.Error(w, "Erro ao criar usuário", http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Usuário leitor criado")
}

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler requisição", http.StatusBadRequest)
		return
	}

	err = create_admin(body)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				http.Error(w, "Email já cadastrado", http.StatusNotAcceptable)
				return
			}
		}
		fmt.Println(err)
		http.Error(w, "Erro ao criar admin", http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Usuário admin criado")
}
