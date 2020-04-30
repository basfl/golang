package main

import (
	"fmt"
)

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

type allbooks [4]string

func main() {

	books := allbooks{
		"book0",
		"book1",
		"book2",
		"books0" + "@second addition",
	}
	var (
		//	books     [yearly]string
		wbooks    [winter]string
		sbooks    [summer]string
		published [len(books)]bool
	)
	// books[0] = "book0"
	// books[1] = "book1"
	// books[2] = "book2"
	// books[3] += books[0] + "@second addition"
	fmt.Printf("books: %T ", books)
	fmt.Printf("\nbooks %#v ", books)
	wbooks[0] = books[0]
	fmt.Printf("\nwinter books %#v ", wbooks)
	// sbooks[0] = books[1]
	// sbooks[1] = books[2]
	// sbooks[2] = books[3]
	for i := range sbooks {
		sbooks[i] = books[i+1]
	}
	fmt.Printf("\nsummer books %#v ", sbooks)
	published[0] = true
	published[len(books)-1] = true
	for i, ok := range published {
		if ok {
			fmt.Printf("\n published books : %s", books[i])
		}
	}
	fmt.Println("\n --------------------------")
	multi := [...][3]int{
		{1, 2, 3},
		{6, 7, 8},
	}

	fmt.Printf("\nmulti = %#v ", len(multi))

}
