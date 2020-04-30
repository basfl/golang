package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v", os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(len(os.Args))

}
