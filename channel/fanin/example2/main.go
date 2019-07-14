package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)
	go send(even, odd)
	receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println("fmt", v)
	}
}
func send(even, odd chan<- int) {
	for i := 0; i < 1; i++ {
		if i%2 == 0 {
			fmt.Println("even")
			even <- i
		} else {
			fmt.Println("odd")
			odd <- i
		}
	}
	close(even)
	close(odd)
}
func receive(even, odd <-chan int, fanin chan<- int) {
	fmt.Println("receive")
	for {
		select {
		case v := <-even:
			fanin <- v
		case v := <-odd:
			fanin <- v
		}
	}
	//	close(fanin)
}
