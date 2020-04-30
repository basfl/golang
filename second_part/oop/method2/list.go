package main

type list []*movie

func (ms list) print() {
	for _, it := range ms {
		it.print()
	}
}
