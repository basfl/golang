package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}
	type book struct {
		text text
		isbn string
	}

	moby := book{text: text{title: "moby dick", words: 2000}, isbn: "10000"}
	fmt.Printf("\n%v", moby)
}
