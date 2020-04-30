package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\n foo #%d", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\n bar #%d", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
