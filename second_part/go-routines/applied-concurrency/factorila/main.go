package main

import (
	"fmt"
)

var factm map[string]float64

func main() {

	c4 := decreament(4)
	fmt.Printf("factorila of 4 is %f", <-c4)
	c8 := decreament(8)
	fmt.Printf("\nfactorila of 8 is %f", <-c8)
	c10 := decreament(10)
	fmt.Printf("\nfactorila of 10 is %f", <-c10)
	c30 := decreament(12)
	fmt.Printf("\nfactorila of 12 is %f", <-c30)
	c400 := decreament(14)
	fmt.Printf("\nfactorila of 14 is %f", <-c400)

}

func decreament(n int) chan float64 {
	c := make(chan float64)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- float64(total)
		close(c)

	}()
	return c

}
