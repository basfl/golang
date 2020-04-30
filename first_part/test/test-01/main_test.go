package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	result := mySum(2, 3)
	if result != 5 {
		t.Error("expected 5 got ", result)
	}

}
