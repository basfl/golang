package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
)

func main() {
	fmt.Println("kb", kb)
	fmt.Println("mb", mb)
}
