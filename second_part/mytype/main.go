package main

import (
	"fmt"

	"./weight"
)

//Gram in main
type Gram float64
type ounce float64

const num = 100

func main() {
	var g Gram = 1000
	var o ounce
	o = ounce(g * 0.035274)
	fmt.Printf("%v grame = %.2f ounce", g, o)
	fmt.Println("\n ***", weight.Gram(1))
	fmt.Printf("the type weight Gram is %T", weight.Gram(1))

}
