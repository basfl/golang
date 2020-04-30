package main

import "fmt"

func main() {
	fmt.Println(greet("ba", "sfl"))
}

func greet(f, l string) (r string) {

	r = "hello " + f + " " + l
	return

}
