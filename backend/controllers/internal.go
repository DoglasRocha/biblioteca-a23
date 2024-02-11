package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func create_user(request_body []byte) (models.User, error) {
	var user models.User

	err := json.Unmarshal(request_body, &user)
	if err != nil {
		return user, err
	}

	err = user.Validate()
	if err != nil {
		return user, err
	}

	var password_hash []byte
	password_hash, err = bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	*user.Password = string(password_hash)

	err = database.DB.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func create_reader(request_body []byte) (models.Reader, error) {
	var reader models.Reader
	var user models.User

	user, err := create_user(request_body)
	if err != nil {
		fmt.Println(err)
		return reader, err
	}

	err = json.Unmarshal(request_body, &reader)
	if err != nil {
		database.DB.Delete(&user)
		return reader, err
	}

	reader.User = user
	reader.UserID = user.ID

	err = reader.Validate()
	if err != nil {
		database.DB.Delete(&user)
		return reader, err
	}

	err = database.DB.Create(&reader).Error
	if err != nil {
		database.DB.Delete(&user)
		return reader, err
	}
	return reader, nil
}
