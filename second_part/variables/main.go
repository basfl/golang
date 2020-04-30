package main

import "fmt"

var safe = true

func main() {
	var (
		speed  int
		speed2 int
		heat   float64
		off    bool
		brand  string
	)

	fmt.Println(
		speed,
		heat,
		off,
		brand,
	)
	fmt.Printf("%q\n", brand)
	_ = speed2
	fmt.Println("safe is ", safe)
	safe1, speed3 := true, 12
	fmt.Println("safe1 and spped3 are ", safe1, speed3)
}
