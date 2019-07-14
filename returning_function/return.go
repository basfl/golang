package main

import (
	"fmt"
)

func main() {
	b := bar()

	fmt.Println("return is ", b(45))
}
func bar() func(n int) int {
	return func(n int) int {
		return n
	}
}
