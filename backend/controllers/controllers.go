package controllers

import (
	//"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
)

func CreateReader(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reader, err = create_reader(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// hides password from response
	reader.User.Password = nil
	w.Header().Add("Set-Cookie", "boa tarde")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reader)
}
