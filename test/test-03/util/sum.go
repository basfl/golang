// Package util is for utility functions
package util

//MySum recieves unlimited int numbers and return the sum
func MySum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
