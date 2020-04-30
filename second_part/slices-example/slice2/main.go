package main

import "fmt"

func main() {
	s := []string{
		"a", "b", "c", "d",
		"e", "f", "g", "h", "k",
		"n", "m", "w", "z",
	}
	// fmt.Printf("\n%#v", s[0:4])
	// fmt.Printf("\n%#v", s[4:8])
	// fmt.Printf("\n%#v", s[8:12])
	// fmt.Printf("\n%#v", s[12:])
	l := len(s)
	const pageSize = 4
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		//fmt.Printf("\nfrom %d to %d ", from, to)
		if to > l {
			to = l
		}
		fmt.Printf("\n%#v", s[from:to])

	}
}
