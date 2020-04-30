package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("\n start main :", time.Now().String())
	ctf1 := timefunc("*func1->")
	ctf2 := timefunc("*func2->")
	fmt.Println("\n", <-ctf1)
	fmt.Println("\n", <-ctf2)
}

func timefunc(s string) chan string {
	c := make(chan string)
	go func() {
		t := s + ": " + time.Now().String()
		c <- t
		close(c)
	}()

	return c
}
