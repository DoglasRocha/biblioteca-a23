package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"log/slog"
)

func create_reader(request_body []byte) error {
	var reader models.Reader
	var user models.User

	// creates the user
	user, err := create_user(request_body)
	if err != nil {
		return err
	}

	// unpacks the json body to the reader struct
	err = json.Unmarshal(request_body, &reader)
	if err != nil {
		slog.Warn(
			"Erro ao preencher struct reader",
			"err", err,
		)
		database.DB.Delete(&user)
		return err
	}

	reader.User = user
	reader.UserID = user.ID

	// validates the reader fields
	err = reader.Validate()
	if err != nil {
		slog.Warn(
			"Erro ao validar struct reader",
			"err", err,
		)
		database.DB.Delete(&user)
		return err
	}

	// creates the reader in DB
	err = database.DB.Create(&reader).Error
	if err != nil {
		slog.Warn(
			"Erro ao criar struct reader no banco de dados",
			"err", err,
		)
		database.DB.Delete(&user)
		return err
	}

	return nil
}

func create_admin(request_body []byte) error {
	var admin models.Admin
	var user models.User

	// creates the user
	user, err := create_user(request_body)
	if err != nil {
		return err
	}

	admin.User = user
	admin.UserID = user.ID
	admin.IsCleared = false

	// validates the admin struct
	err = admin.Validate()
	if err != nil {
		slog.Warn(
			"Erro ao validar struct admin",
			"err", err,
		)
		return err
	}

	// creates the admin in DB
	err = database.DB.Create(&admin).Error
	if err != nil {
		slog.Warn(
			"Erro ao criar struct admin no banco de dados",
			"err", err,
		)
		database.DB.Delete(&user)
		return err
	}

	return nil
}
