package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type login_struct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func create_cookie(key string, value string) http.Cookie {
	cookie := http.Cookie{
		Name:     key,
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: false,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	}

	return cookie
}

func sign_jwt(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	return tokenString, err
}

func user_exists_and_passwords_match(request_body []byte) (models.User, error) {
	var login login_struct
	var user models.User

	// unpacks the json body to the login_struct
	err := json.Unmarshal(request_body, &login)
	if err != nil {
		return models.User{}, err
	}

	// checks if email exists in DB
	err = database.DB.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	// checks if passwords match
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(login.Password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func create_user(request_body []byte) (models.User, error) {
	var user models.User

	// unpacks the json from body to user struct
	err := json.Unmarshal(request_body, &user)
	if err != nil {
		return models.User{}, err
	}

	// validates the user fields
	err = user.Validate()
	if err != nil {
		return models.User{}, err
	}

	// creates the password hash
	var password_hash []byte
	password_hash, err = bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.MinCost)
	if err != nil {
		return models.User{}, err
	}
	*user.Password = string(password_hash)

	// creates the user in DB
	err = database.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func is_reader_authenticated(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return http.StatusInternalServerError
	}
	status, message, _ := check_reader(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
		return status
	}

	return status
}

func is_admin_autenticated(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Erro ao acessar cookie")
		return http.StatusInternalServerError
	}
	status, message, _ := check_admin(cookie)
	if status != http.StatusOK {
		w.WriteHeader(status)
		fmt.Fprintln(w, message)
		return status
	}

	return status
}
