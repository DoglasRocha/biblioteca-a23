package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"log/slog"
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
		slog.Info(
			"kkkkkkkkkkk admin tentando entrar na parte de usuário",
			"user_id", id,
		)
		return http.StatusForbidden, "Usuário é administrador", nil
	}

	slog.Warn(
		"Não tem usuário algum com esse id",
		"user_id", id,
	)
	// is not user
	return http.StatusUnauthorized, "Favor logar-se", err
}

func check_admin(cookie *http.Cookie) (int, string, error) {
	var admin models.Admin
	id, err := parse_cookie(cookie)

	if err != nil {
		return http.StatusUnauthorized, "Erro ao processar token", err
	}

	err = database.DB.Where("user_id = ?", id).First(&admin).Error
	// user is admin
	if err == nil {
		if admin.IsCleared {
			return http.StatusOK, "OK", nil
		}
		slog.Warn(
			"Admin não autorizado tentando entrar",
			"user_id", id,
		)
		return http.StatusNotAcceptable, "Admin não autorizado", nil
	}

	err = database.DB.Where("user_id = ?", id).First(&models.Reader{}).Error
	// user is reader
	if err == nil {
		slog.Warn(
			"Leitor tentando entrar no layer de admin",
			"user_id", id,
		)
		return http.StatusForbidden, "Usuário é leitor", nil
	}

	slog.Warn(
		"Não tem usuário algum com esse id",
		"user_id", id,
	)
	// is not user
	return http.StatusUnauthorized, "Favor logar-se", err
}
