package controllers

import (
	//"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	//"fmt"
	"io"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
			http.Error(w, "Email já cadastrado", http.StatusNotAcceptable)
			return
		}
		http.Error(w, "Erro ao criar usuário", http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reader)
}

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
	cookie := http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged-in")
}
