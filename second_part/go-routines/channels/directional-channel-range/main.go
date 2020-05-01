package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int, 50)
	wg.Add(2)
	go func(c <-chan int) {
		for ele := range c {
			fmt.Println(ele)
		}
		wg.Done()

	}(ch)
	go func(c chan<- int) {
		c <- 4
		c <- 12
		c <- 14
		close(c)
		wg.Done()
	}(ch)

	wg.Wait()

}
