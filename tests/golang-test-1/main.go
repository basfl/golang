package main

import "fmt"

func main() {
	a, b := 3, 8
	fmt.Println(add(a, b))
}

func add(a, b int) int {
	return a + b
}
