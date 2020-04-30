package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// in := bufio.NewScanner(os.Stdin)
	// words := make(map[string]bool)
	// in.Split(bufio.ScanWords)

	// if _, err := os.Stat("./temp.txt"); err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("file does not exist")
	// 	} else {
	// 		fmt.Println("file does do exist")
	// 	}
	// }

	// for in.Scan() {
	// 	word := strings.ToLower(in.Text())
	// 	fmt.Printf("* %s", word)
	// 	if len(word) > 2 {
	// 		words[word] = true
	// 	}

	// }
	// query := "sun"
	// if words[query] {
	// 	fmt.Printf("the input contains %s", query)
	// } else {
	// 	fmt.Printf("sorry coldnot find %s", query)
	// }
	// fmt.Println("\n-----------------------")
	content, err := ioutil.ReadFile("/data/temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", content)

}
