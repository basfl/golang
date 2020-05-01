package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

var ch = make(chan string, 50)

func main() {

	wg.Add(2)
	func1()
	func2()

	// func() {
	wg.Wait()
	// }()

}

func func1() {

	go func(c <-chan string) {
		for ele := range c {
			fmt.Println(strings.ToUpper(ele))
		}
		wg.Done()

	}(ch)

}

func func2() {

	go func(c chan<- string) {
		file, err := os.Open("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}

		close(c)
		wg.Done()
	}(ch)

}
