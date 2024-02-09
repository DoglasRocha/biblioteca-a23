package controllers

import (
	//"biblioteca-a23/database"
	"biblioteca-a23/models"
	"fmt"
	"io"
	"net/http"
)

func CreateReader(w http.ResponseWriter, r *http.Request) {
	var reader models.Reader
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err, http.StatusBadRequest)
		return
	}

	reader, err = create_reader(body)
	if err != nil {
		fmt.Fprintln(w, err, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, reader)
	// err = json.Unmarshal(body, &reader)
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// }

	// fmt.Println(reader)

	// fmt.Fprintln(w, reader)

}
