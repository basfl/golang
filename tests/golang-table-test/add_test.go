package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{3, 3}, 6},
		test{[]int{5, 5}, 10},
		test{[]int{9, 9}, 18},
		test{[]int{2, 0}, 1},
	}
	for _, v := range tests {
		res := add(v.data[0], v.data[1])
		if res != v.answer {
			t.Errorf("myAdd(%d,%d) = %d; expect %d", v.data[0], v.data[1], res, v.answer)
		}

	}

}
