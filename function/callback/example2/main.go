package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("**", filter(nums, callback))

}
func callback(a int) bool {
	return a > 2
}
func filter(a []int, callback func(a int) bool) []int {
	xs := []int{}
	for _, v := range a {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	return xs
}
