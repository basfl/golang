package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := remove(xi, 2)
	fmt.Println(res)
}
func remove(a []int, pos int) []int {
	a = append(a[:pos], a[pos+1:]...)
	return a
}
