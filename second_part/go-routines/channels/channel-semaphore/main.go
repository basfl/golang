package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	cnt  int
	mutx sync.Mutex
)

func main() {

	c := make(chan string)
	done := make(chan bool)

	go increament(c, done, "foo")
	go increament(c, done, "bar")
	go func() {
		//	wg.Wait()
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Printf("\n %s ", n)
	}

}

func increament(c chan string, done chan bool, name string) {
	for i := 0; i < 5; i++ {
		mutx.Lock()
		cnt++
		c <- "name:" + name + strconv.Itoa(cnt)
		mutx.Unlock()
	}
	done <- true

	//	wg.Done()
}
