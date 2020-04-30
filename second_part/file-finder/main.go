package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := os.Args[1:]
	if len(path) == 0 {
		fmt.Println("please provide the path")
	}
	files, err := ioutil.ReadDir(path[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}
	names := make([]byte, 0, total)
	fmt.Printf("\n total require size is %d ", total)
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
			//	fmt.Println(name)
		}

	}
	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%s", names)
	moji := 128572
	//moji := '😼'
	fmt.Printf("%c", moji)
}
