package main

import (
	"fmt"

	r "./removeslice"
)

func main() {
	a := []int{
		1, 2, 3, 4, 5, 6,
	}
	fmt.Println("before: ", a)
	after := r.Remove(a, 2)
	fmt.Println("after: ", after)

}
