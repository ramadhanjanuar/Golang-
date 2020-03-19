// passwords.go
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "qwerty"
	salt := "rXI$L@bdND"
	hash, _ := HashPassword(salt + password) // ignore error for the sake of simplicity
	hashQwerty := "$2y$08$mHMknlurxKSvJ.Eh63kOD.ySX4XzfRUGU8tc2uQQWjwpPEkTRYam6"

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	//oldHash := "$2a$14$YvnBHanEsmraYA7naQmCS.KMN06WKckZZNsvVRwil/fiOxIoH9Khi"
	match := CheckPasswordHash(salt+password, hashQwerty)
	fmt.Println("Match:   ", match)
}
