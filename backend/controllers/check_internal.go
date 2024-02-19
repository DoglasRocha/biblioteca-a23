package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"net/http"
)

func check_reader(cookie *http.Cookie) (int, string, error) {
	id, err := parse_cookie(cookie)

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
	id, err := parse_cookie(cookie)

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
