package main

import (
	"fmt"
)

func main() {
	a := 12
	var b *int = &a
	fmt.Println("a=", a)
	fmt.Println("&a=", &a)
	fmt.Println("b=", b)
	fmt.Println("&b=", &b)
	fmt.Println("*b=", *b)

}
