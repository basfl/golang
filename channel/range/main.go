package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i + 2
		}
		close(ch)
	}(ch)
	for v := range ch {
		fmt.Println("value is ", v)
	}

}
