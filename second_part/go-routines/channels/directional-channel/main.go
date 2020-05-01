package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(2)
	//recieving from channel
	go func(ch <-chan int) {
		num := <-ch
		fmt.Println(num)
		wg.Done()

	}(ch)
	//writing to channel
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()

	}(ch)

	wg.Wait()

}
