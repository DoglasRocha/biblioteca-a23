package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListAdmins(w http.ResponseWriter, r *http.Request) {
	var admins []models.Admin

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	err := database.DB.Find(&admins).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao buscar admins na base de dados")
		return
	}

	for i := range admins {
		// find user
		err = database.DB.Find(&admins[i].User, admins[i].UserID).Error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Erro ao buscar admins na base de dados")
			return
		}

		// hide password
		admins[i].User.Password = nil
	}

	json.NewEncoder(w).Encode(admins)
}
