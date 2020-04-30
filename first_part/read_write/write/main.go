package main

import (
	"io/ioutil"
	"log"
)

func main() {
	mesg := "this is my mesg"
	b := []byte(mesg)
	err := ioutil.WriteFile("C:\\Users\\jupit\\Documents\\programming\\GO\\my_files\\read_write\\write\\write_temp.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
