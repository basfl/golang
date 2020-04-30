package main

import (
	"fmt"
)

func main() {
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)
	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)
	plante, distance := "venus", 265.5
	fmt.Printf("%v is %v away!,think about is %[2]v km away ,%[1]v omg", plante, distance)

}
