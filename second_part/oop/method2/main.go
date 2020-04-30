package main

import "fmt"

func main() {
	var (
		m1 = movie{"m1", 10.5}
		m2 = movie{"m2", 20}
	)

	movies := []*movie{&m1, &m2}
	fmt.Printf("\n %v", movies)
	list := list(movies)
	list.print()

}
