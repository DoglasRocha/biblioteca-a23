package controllers

import (
	"biblioteca-a23/database"
	"biblioteca-a23/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type login_struct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func create_user(request_body []byte) (models.User, error) {
	var user models.User

	// unpacks the json from body to user struct
	err := json.Unmarshal(request_body, &user)
	if err != nil {
		return models.User{}, err
	}

	// validates the user fields
	err = user.Validate()
	if err != nil {
		return models.User{}, err
	}

	// creates the password hash
	var password_hash []byte
	password_hash, err = bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.MinCost)
	if err != nil {
		return models.User{}, err
	}
	*user.Password = string(password_hash)

	// creates the user in DB
	err = database.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

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

func user_exists_and_passwords_match(request_body []byte) (models.User, error) {
	var login login_struct
	var user models.User

	// unpacks the json body to the login_struct
	err := json.Unmarshal(request_body, &login)
	if err != nil {
		return models.User{}, err
	}

	// checks if email exists in DB
	err = database.DB.Where("email = ?", login.Email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	// checks if passwords match
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(login.Password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func sign_jwt(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	return tokenString, err
}

func create_cookie(key string, value string) http.Cookie {
	cookie := http.Cookie{
		Name:     key,
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: false,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	}

	return cookie
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
