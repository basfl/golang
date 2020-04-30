package main

import (
	"fmt"
	"os"
	"strconv"

	"./messages"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(messages.UserMsg)
		return
	}
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	meter := feet * 0.3048
	fmt.Printf("%v feet is %.2f meter", feet, meter)

}
