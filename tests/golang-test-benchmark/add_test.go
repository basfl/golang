package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(3, 3)
	if res != 6 {
		t.Errorf("myAdd(2,3) = %d; expect 6", res)
	}
}
func Exampleadd() {
	fmt.Println(add(3, 3))
	//Output:
	//6
}
func BenchMarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(3, 3)
	}
}
