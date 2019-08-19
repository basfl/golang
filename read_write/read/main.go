package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", content)
}
