package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	args := os.Args
	for _, v := range args {
		fmt.Println(v)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
