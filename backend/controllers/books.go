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
		fmt.Println(w, "Erro ao validar usuário")
	}

	if status != http.StatusOK {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println(w, "Você não possui acesso a esta parte do sistema")
		return
	}

	err = register_book(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao criar livro")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Livro cadastrado")
}
