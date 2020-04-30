package main

import "fmt"

func main() {
	moji := 128572
	//moji := '😼'

	fmt.Printf("character %c , unicode %U ", moji, moji)
	values := []rune{'♛', '♠', '♧', '♬', 128572, 'a'}
	for i, value := range values {
		fmt.Printf("\n chracter %c ,unicode %#U ,string %s , int %d  ,position %d", value, value, string(value), value, i)
	}
	unicode := '\u00AE'
	//r := 0174
	fmt.Printf("\n%q and converted from ascicode %-12x ", unicode, unicode)
	fmt.Println("\n-----------------------------------------------")
	valueStr := []string{"♛", "♠", "♧", "♬", "a"}
	for _, value := range valueStr {
		fmt.Printf("\n value in string %s value in rune is %q len is %d", value, value, len(value))
	}

}
