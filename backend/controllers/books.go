package controllers

import (
	"fmt"
	"net/http"
)

func RegisterBook(w http.ResponseWriter, r *http.Request) {
	// checks if user is admin
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return
	}
	status, _, err := check_admin(cookie)
	if err != nil {
		w.WriteHeader(status)
		fmt.Fprintln(w, "Erro ao validar usuário")
		return
	}

	if status != http.StatusOK {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Você não possui acesso a esta parte do sistema")
		return
	}

	err = register_book(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao criar livro")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Livro cadastrado")
}
