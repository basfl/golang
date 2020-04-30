package util

import (
	"fmt"
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{12, 12}, 24},
		test{[]int{3, 5, 2}, 10},
		test{[]int{4, 4, 10}, 18},
	}
	for _, v := range tests {
		result := MySum(v.data...)
		if result != v.answer {
			t.Error("expected", v.answer, " we got ", result)
		}
	}

}
func ExampleMySum() {
	fmt.Println(MySum(2, 3))
	//Output:
	//5
}
