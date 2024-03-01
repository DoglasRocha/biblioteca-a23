package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func ListAdmins(w http.ResponseWriter, r *http.Request) {
	var admins []models.Admin

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	err := database.DB.Find(&admins).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar todos os admins",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar admins na base de dados")
		return
	}

	for i := range admins {
		// find user
		err = database.DB.Find(&admins[i].User, admins[i].UserID).Error
		if err != nil {
			slog.Warn(
				"Erro ao popular admin",
				"err", err,
				"id", admins[i].UserID,
			)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro ao buscar admins na base de dados")
			return
		}

		// hide password
		admins[i].User.Password = nil
	}

	json.NewEncoder(w).Encode(admins)
}

func AuthorizeAdmin(w http.ResponseWriter, r *http.Request) {
	var admin_to_authorize models.Admin
	var current_admin models.Admin

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get current admin
	current_admin_id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}

	err = database.DB.Where("user_id = ?", current_admin_id).Find(&current_admin).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar admin corrente na base de dados",
			"err", err,
			"id", current_admin_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar admin corrente na base de dados")
		return
	}

	// get admin to authorize
	admin_to_authorize_id := mux.Vars(r)["admin_id"]
	err = database.DB.Where("user_id = ?", admin_to_authorize_id).Find(&admin_to_authorize).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar admin para autorizar na base de dados",
			"err", err,
			"id", admin_to_authorize_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar admin para autorizar na base de dados")
		return
	}

	admin_to_authorize.IsCleared = true
	err = database.DB.Save(&admin_to_authorize).Error
	if err != nil {
		slog.Warn(
			"Erro ao atualizar adm na base de dados",
			"err", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao autorizar administrador")
		return
	}

	slog.Info(
		"Admin autorizado",
		"id", admin_to_authorize.UserID,
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Admin autorizado com sucesso")
}

func RevokeAdmin(w http.ResponseWriter, r *http.Request) {
	var admin_to_revoke models.Admin
	var current_admin models.Admin

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get current admin
	current_admin_id, err := get_id_from_request_cookie(w, r)
	if err != nil {
		return
	}
	err = database.DB.Where("user_id = ?", current_admin_id).Find(&current_admin).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar admin corrente na base de dados",
			"err", err,
			"id", current_admin_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar admin corrente na base de dados")
		return
	}

	// get admin to revoke
	admin_to_revoke_id := mux.Vars(r)["admin_id"]
	err = database.DB.Where("user_id = ?", admin_to_revoke_id).Find(&admin_to_revoke).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar admin para revogar na base de dados",
			"err", err,
			"id", admin_to_revoke_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao encontrar admin para revogar na base de dados")
		return
	}

	if admin_to_revoke.UserID == current_admin.UserID {
		slog.Warn(
			"kkkkkkk admin tentando se revogar",
			"id", admin_to_revoke_id,
		)
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, "Um admin não pode se revogar")
		return
	}

	admin_to_revoke.IsCleared = false
	err = database.DB.Save(&admin_to_revoke).Error
	if err != nil {
		slog.Warn(
			"Erro ao atualizar admin revogado na base de dados",
			"err", err,
			"id", admin_to_revoke_id,
		)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao autorizar administrador")
		return
	}

	slog.Info(
		"Admin revogado",
		"id", admin_to_revoke_id,
	)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Admin revogado com sucesso")
}
