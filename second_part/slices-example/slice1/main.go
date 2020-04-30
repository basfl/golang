package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	type myStrArr [5]string
	var temp myStrArr
	fmt.Printf("%#v", temp)
	rand.Seed(time.Now().UnixNano())
	var unique []int
parent:
	for found := 0; found < 10; {
		n := rand.Intn(10) + 1
		fmt.Printf("%d", n)
		for _, v := range unique {
			if v == n {
				continue parent
			}

		}
		unique = append(unique, n)
		found++
	}
	fmt.Printf("%#v", unique)
	final := []int{0, 0, 0}
	final = append(final, unique...)
	fmt.Println("\n _____________________")
	fmt.Printf("\n%#v", final)
	args := []int{4, 12, 3}
	first := args[0:1]
	second := args[1:3]
	fmt.Printf("\n%#v", first)
	fmt.Printf("\n%#v", second)
	args[0] = 123
	fmt.Printf("\n%#v", first)
	fmt.Printf("\n%#v", second)
	ages1 := []int{1, 2, 3}
	fmt.Printf("\n%#v", ages1)
	ages1 = append(ages1, 12)
	cap()

}
