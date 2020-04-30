package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cnt  int
	wg   sync.WaitGroup
	mutx sync.Mutex
)

func main() {
	wg.Add(2)
	go increament("foo")
	go increament("bar")
	wg.Wait()
}

func increament(s string) {
	for i := 0; i < 10; i++ {

		fmt.Printf("\n%s : cnt %d ", s, cnt)
		mutx.Lock()
		cnt++
		time.Sleep(time.Duration(3 * time.Millisecond))
		mutx.Unlock()
	}
	wg.Done()
}
