package main

import "fmt"

type movie struct {
	title string
	price float64
}

func (m movie) print() {
	fmt.Printf("\n the movie title is %s and price is %.2f ", m.title, m.price)
}

func (m *movie) discount(ration float64) {
	m.price *= ration
}
