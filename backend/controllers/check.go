package controllers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func CheckReader(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		slog.Warn(
			"Erro ao acessar cookie",
			"err", err,
		)
		http.Error(w, "Erro ao ler cookie", http.StatusUnauthorized)
		return
	}

	status, message, _ := check_reader(cookie)

	w.WriteHeader(status)
	fmt.Fprintln(w, message)
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		slog.Warn(
			"Erro ao acessar cookie",
			"err", err,
		)
		http.Error(w, "Erro ao ler cookie", http.StatusUnauthorized)
		return
	}

	status, message, _ := check_admin(cookie)

	w.WriteHeader(status)
	fmt.Fprintln(w, message)
}
