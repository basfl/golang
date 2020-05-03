package main

import "testing"

func TestMyAdd(t *testing.T) {
	ms := myStruc{
		[]int{1, 4, 5},
	}
	result := ms.myAdd(2)
	if result != 16 {
		t.Errorf("ms.myAdd(2) = %d; expect 16", result)
	}
}
