package main

import (
	"fmt"
)

func main() {
	mp := map[string]int{"a": 1, "b": 2, "c": 5}
	// m = make(map[string]int)
	fmt.Println(mp)
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
