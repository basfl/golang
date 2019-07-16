package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(10)
		for e := 0; e < 10; e++ {

			go func() {
				for i := 0; i < 10; i++ {
					c <- i
				}
				wg.Done()
			}()

		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
