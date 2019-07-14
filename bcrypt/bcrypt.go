package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `password123`
	fmt.Println("Hello, playground")
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	err = bcrypt.CompareHashAndPassword(bs, []byte(pass))
	if err != nil {
		fmt.Println("you failed")
	}
	fmt.Println("you logged in")
}
