package main

import (
	"fmt"
)

func pointerFunc(a *int) {
	fmt.Println("a inside the function is ", a)
	*a = *a + 1

}
func main() {
	a := 12
	fmt.Println("a=", a)
	pointerFunc(&a)
	var b *int = &a
	pointerFunc(b)
	fmt.Println("a=", a)
}
