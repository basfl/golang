package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.yahoo.com/",
		"https://twitter.com/",
	}
	wg.Add(len(urls))
	for _, url := range urls {
		//wg.Add(1)
		go fetch(url)
	}
	fmt.Println("NOT WAITING!!!")
	wg.Wait()
	fmt.Println("*****")
}
func fetch(url string) {
	_, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("requesting url= ", url)
	wg.Done()
	//fmt.Printf("%s", res)
}
