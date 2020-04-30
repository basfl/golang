package main

import "fmt"

func main() {
	var m = make(map[string]int)
	m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
	m["d"] = 4
	fmt.Println("after add :", m)
	delete(m, "d")
	fmt.Println("after delete :", m)

	if _, ok := m["d"]; !ok {
		fmt.Println("did not find value")
	}
	m2 := make(map[string][]int)
	m2 = map[string][]int{
		"m1": {1, 2, 3},
		"m2": {34, 34},
	}
	fmt.Println(m2)

}
