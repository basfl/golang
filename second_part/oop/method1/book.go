package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b *book) print() {
	fmt.Printf("\n the book title is %s and the price is %.2f ", b.title, b.price)
}

func (b *book) discount(ration float64) {
	b.price = b.price * ration

}
