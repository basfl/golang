package main

import (
	"fmt"
	"path"
	"time"
)

func main() {
	var (
		name     string
		prevName string
		now      time.Time
	)
	name = "babak"
	prevName = name
	name = "amir"
	fmt.Println("prevname name is ", prevName)
	fmt.Println("current name is ", name)
	fmt.Println("current time ", now)
	now = time.Now()
	fmt.Println("current time ", now)
	_, file := path.Split("src/css/main.css")
	fmt.Println("file is ", file)
}
