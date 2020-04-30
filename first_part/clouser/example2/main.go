package main

import (
	"fmt"
)

var x int = 0

func increament() int {
	x++
	return x
}

func main() {
	fmt.Println("*********")
	fmt.Println("*", increament())
	fmt.Println("*", increament())
}
