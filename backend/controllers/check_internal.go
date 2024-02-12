package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func get_id_from_cookie(cookie *http.Cookie) (int, error) {
	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return os.Getenv("SIGNING_KEY"), nil
	})

	if err != nil {
		return -1, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["id"])
		return 0, nil
	}
	return -1, err
}

func check_reader(cookie *http.Cookie) (int, string, error) {
	id, err := get_id_from_cookie(cookie)

	if err != nil {
		return http.StatusUnauthorized, "Falha ao validar token. Por favor logar-se novamente", err
	}

	fmt.Println(id)
	return -1, "", err
}
