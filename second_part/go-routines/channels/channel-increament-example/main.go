package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	cnt  int
	wg   sync.WaitGroup
	mutx sync.Mutex
)

func main() {

	c := make(chan string)
	wg.Add(2)
	go increament(c, "foo")
	go increament(c, "bar")
	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Printf("\n %s ", n)
	}

}

func increament(c chan string, name string) {
	for i := 0; i < 5; i++ {
		mutx.Lock()
		cnt++
		c <- "name:" + name + strconv.Itoa(cnt)
		mutx.Unlock()
	}
	wg.Done()
}
