package main

import (
	"fmt"
)

func main() {
	var name, food, sport string
	fmt.Print("Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&food)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&sport)
	if err != nil {
		panic(err)
	}
	fmt.Println(name, food, sport)

}
