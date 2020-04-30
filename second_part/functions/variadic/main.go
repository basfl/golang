package main

import "fmt"

func main() {
	fmt.Println("\n average is ", average([]int{2, 4, 5, 7}, func(n int) int {
		return n * 2
	}))
	defer goodby()
	greeting := func() {
		fmt.Println("greeting")
	}
	greeting()
	res := func(a int) int {
		fmt.Println("running", a, " times")
		return a + 3
	}(2)
	fmt.Println(res)
	fmt.Println("-----------------")
	vari(1, 2, 3, 4, 5, 6)
}
func average(vs []int, callback func(int) int) (av float64) {
	fmt.Printf("\n type is %T", vs)
	fmt.Printf("\n length is %d ", len(vs))
	var sum int
	for _, v := range vs {
		sum += callback(v)

	}
	av = float64(sum) / float64(len(vs))
	return

}
func vari(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func goodby() {
	fmt.Println("bye!!")
}
