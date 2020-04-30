package main

import "fmt"

func main() {
	//label:
	for i := 0; i < 5; i++ {
		fmt.Println("inside I", i)
		for j := 0; j < 3; j++ {
			fmt.Println("inside J", j)
			break

		}
	}
}
