package controllers

import (
	"fmt"
	"net/http"
)

func CheckReader(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		http.Error(w, "Erro ao ler cookie", http.StatusUnauthorized)
		return
	}

	status, message, error := check_reader(cookie)
	fmt.Println(status, message, error)
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Cookies())
	fmt.Println(r.Header.Get("Authorization"))
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		http.Error(w, "Erro ao ler cookie", http.StatusUnauthorized)
		return
	}

	status, message, error := check_reader(cookie)
	fmt.Println(status, message, error)
}
