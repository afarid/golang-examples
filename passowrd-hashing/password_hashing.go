package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func main() {
	password := os.Getenv("PASSWORD")
	hash, err := HashPassword(password)
	if err != nil {
		panic(err)
	}

	fmt.Println(CheckPasswordHash(password, hash))

}
