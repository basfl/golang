package main

import (
	"fmt"
)

func main() {
	fmt.Println("***")
	x := 0
	increament := func() int {
		x++
		return x
	}
	fmt.Println(increament())
	fmt.Println(increament())
}
