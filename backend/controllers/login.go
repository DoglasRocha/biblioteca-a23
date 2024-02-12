package controllers

import (
	"biblioteca-a23/models"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler requisição", http.StatusBadRequest)
		return
	}

	// checks if user exists and password match
	user, err = user_exists_and_passwords_match(body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Email não encontrado", http.StatusNotAcceptable)
			return
		}

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			http.Error(w, "Senha incorreta", http.StatusNotAcceptable)
			return
		}

		http.Error(
			w,
			"Erro ao processar solicitação. Tente novamente mais tarde.",
			http.StatusBadRequest,
		)
		return
	}

	// generates jwt
	token, err := sign_jwt(user)
	if err != nil {
		fmt.Println(err)
		http.Error(
			w,
			"Erro ao assinar token de acesso. Tente novamente mais tarde.",
			http.StatusInternalServerError,
		)
		return
	}

	// creates cookie
	cookie := create_cookie("accessToken", token)

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged-in")
}
