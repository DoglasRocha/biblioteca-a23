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

func GetReaderData(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader

	status := is_admin_authenticated(w, r)
	if status != http.StatusOK {
		return
	}

	// get reader data
	user_id := mux.Vars(r)["user_id"]

	err := database.DB.First(&reader, user_id).Error
	if err != nil {
		slog.Warn(
			"Erro ao encontrar struct reader",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Usuário não encontrado")
		return
	}

	err = database.PopulateReader(&reader, reader.UserID)
	if err != nil {
		slog.Warn(
			"Erro ao popular struct reader",
			"err", err,
			"user_id", user_id,
		)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Erro ao buscar dados de usuário")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&reader)
}
