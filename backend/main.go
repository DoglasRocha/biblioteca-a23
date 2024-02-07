package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var teste string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	hash, err := bcrypt.GenerateFromPassword([]byte(teste), bcrypt.MinCost)

	if err != nil {
		panic("aaaaa")
	}

	fmt.Print(string(hash))
}
