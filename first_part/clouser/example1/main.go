package main

import (
	"fmt"
)

func main() {
	x := 12
	fmt.Println("x is ", x)
	{
		y := 17
		fmt.Println("y is ", y)
		fmt.Println("x is ", x)
	}
}
