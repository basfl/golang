package main

import "fmt"

type printer interface {
	print()
}

//type list []*movie
type list []printer

func (ms list) print() {
	for _, it := range ms {
		it.print()
	}
}

func (ms list) discount(ratio float64) {
	type discounter interface {
		discount(ration float64)
	}
	for _, it := range ms {
		//_, isMovie := it.(*movie)
		_, isMovie := it.(discounter)
		if isMovie {
			fmt.Println("\nis movie")
		}
	}
	fmt.Println("*** list discount")
}
