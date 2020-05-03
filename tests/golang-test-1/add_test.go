package main

import "testing"

func TestAdd(t *testing.T) {
	res := add(3, 3)
	if res != 6 {
		t.Errorf("myAdd(2,3) = %d; expect 6", res)
	}
}
