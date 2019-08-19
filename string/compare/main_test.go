package main

import (
	"testing"
)

func TestCompareStrings(t *testing.T) {
	result := compareStrings("string1", "string1")
	if result != 0 {
		t.Error("expected 0 but got ", result)
	}

}

func TestCheckContain(t *testing.T) {
	result := checkContain("today is a good day", "goad")
	if result != true {
		t.Error("expected a true but got ", result)
	}
}
