package main

import "fmt"

type home struct {
	address string
	rooms   int
}

func main() {
	var counter int
	p := &counter
	fmt.Printf("type is %T value %v *p is %v", p, p, *p)
	h := home{address: "brg", rooms: 2}
	fmt.Printf("\n home address %v has %v rooms", h.address, h.rooms)
	update(&h)
	fmt.Printf("\n after update home address %v has %v rooms", h.address, h.rooms)

}
func update(h *home) {
	h.rooms++
}
