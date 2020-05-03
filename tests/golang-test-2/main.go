package main

import "fmt"

type myStruc struct {
	a []int
}

func (ms myStruc) myAdd(is ...int) int {
	sum := 0
	for _, v := range ms.a {
		sum += v + is[0]
	}
	return sum
}
func main() {
	ms := myStruc{
		[]int{1, 4, 5},
	}
	fmt.Println(ms.myAdd(2))
}
