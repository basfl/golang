package main

import (
	"fmt"
)

func main() {
	num := 12
	var pn = &num
	fmt.Println("num=", num)
	change(&num)
	fmt.Println("num=", num)
	change(pn)
	fmt.Println("num=", num)

}
func change(a *int) {
	*a++
}
