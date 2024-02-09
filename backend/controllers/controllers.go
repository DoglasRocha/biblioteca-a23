package controllers

import (
	//"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func CreateReader(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var reader models.Reader
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = user.Validate(validate)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	err = json.Unmarshal(body, &reader)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Println(reader)

	fmt.Fprintln(w, reader)

}
