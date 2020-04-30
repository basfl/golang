package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	name := "babak"
	myAge, yourAge := 30, 35
	var avarge float64
	avarge = float64(myAge+yourAge) / 2
	fmt.Printf("avrage age is %.2f", avarge)
	//////////////////////////////////
	celsuis, _ := strconv.ParseFloat(os.Args[1], 64)
	//cel := float64(celsuisStr)
	fahrenhit := (9*celsuis + 160) / 5
	fmt.Printf("\nthe %v celsuis =%.2f fahrenhit\n", celsuis, fahrenhit)
	fmt.Println(utf8.RuneCountInString(name))

}
