package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func get_id_from_cookie(cookie *http.Cookie) (int, error) {
	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil {
		return -1, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := claims["id"]
		return int(id.(float64)), nil
	}
	return -1, err
}

func check_reader(cookie *http.Cookie) (int, string, error) {
	id, err := get_id_from_cookie(cookie)

	if err != nil {
		return http.StatusUnauthorized, "Erro ao processar token", err
	}

	err = database.DB.Where("user_id = ?", id).First(&models.Reader{}).Error
	// user is reader
	if err == nil {
		return http.StatusOK, "OK", nil
	}

	err = database.DB.Where("user_id = ?", id).First(&models.Admin{}).Error
	// user is admin
	if err == nil {
		return http.StatusForbidden, "Usuário é administrador", nil
	}

	// is not user
	return http.StatusUnauthorized, "Favor logar-se", err
}

func check_admin(cookie *http.Cookie) (int, string, error) {
	id, err := get_id_from_cookie(cookie)

	if err != nil {
		return http.StatusUnauthorized, "Erro ao processar token", err
	}

	err = database.DB.Where("user_id = ?", id).First(&models.Admin{}).Error
	// user is admin
	if err == nil {
		return http.StatusOK, "OK", nil
	}

	err = database.DB.Where("user_id = ?", id).First(&models.Reader{}).Error
	// user is reader
	if err == nil {
		return http.StatusForbidden, "Usuário é leitor", nil
	}

	// is not user
	return http.StatusUnauthorized, "Favor logar-se", err
}
