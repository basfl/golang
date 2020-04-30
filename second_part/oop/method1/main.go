package main

import "fmt"

func main() {
	var (
		book  = book{"book1", 20}
		movie = movie{"movie1", 10}
	)

	book.print()
	fmt.Println("\n book discounted 20%")
	book.discount(.2)
	book.print()
	fmt.Println("\n-------------------")
	movie.print()
	fmt.Println("\n book discounted 30%")
	movie.discount(.3)
	movie.print()

}
