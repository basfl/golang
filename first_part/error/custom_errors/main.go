package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	r, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)

}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("can not process square root of negative")
	}
	return 42, nil
}
