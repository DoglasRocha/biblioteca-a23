package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
)

func create_reader(request_body []byte) (models.Reader, error) {
	var reader models.Reader
	var user models.User

	// creates the user
	user, err := create_user(request_body)
	if err != nil {
		return models.Reader{}, err
	}

	// unpacks the json body to the reader struct
	err = json.Unmarshal(request_body, &reader)
	if err != nil {
		database.DB.Delete(&user)
		return models.Reader{}, err
	}

	reader.User = user
	reader.UserID = user.ID

	// validates the reader fields
	err = reader.Validate()
	if err != nil {
		database.DB.Delete(&user)
		return models.Reader{}, err
	}

	// creates the reader in DB
	err = database.DB.Create(&reader).Error
	if err != nil {
		database.DB.Delete(&user)
		return models.Reader{}, err
	}

	// hides password from response
	reader.User.Password = nil
	return reader, nil
}

func create_admin(request_body []byte) (models.Admin, error) {
	var admin models.Admin
	var user models.User

	// creates the user
	user, err := create_user(request_body)
	if err != nil {
		return admin, err
	}

	admin.User = user
	admin.UserID = user.ID
	admin.IsCleared = false

	// validates the admin struct
	err = admin.Validate()
	if err != nil {
		return models.Admin{}, err
	}

	// creates the admin in DB
	err = database.DB.Create(&admin).Error
	if err != nil {
		database.DB.Delete(&user)
		return models.Admin{}, err
	}

	// hides password from response
	admin.User.Password = nil
	return admin, nil
}
