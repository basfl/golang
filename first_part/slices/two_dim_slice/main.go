package main

import (
	"fmt"
)

func main() {
	td := make([][]string, 0, 5)
	fmt.Println(td)
	r1 := []string{"r1", "ford"}
	td = append(td, r1)
	r2 := []string{"honda", "kia"}
	td = append(td, r2)
	for _, v := range td {
		fmt.Println("v=", v)
		for _, v1 := range v {
			fmt.Println("v1=", v1)
		}
	}
}
