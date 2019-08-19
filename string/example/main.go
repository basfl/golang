// Package util is for main functions
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "string1"
	b := "string1"
	fmt.Println(compareStrings(a, b))
	st1 := "today is a good day"
	search := "day"
	fmt.Println(checkContain(st1, search))
	s := []string{"a", "b", "c"}
	joinResult := joinString(s, ",")
	fmt.Println(joinResult)
	splitresult := stringSplit(joinResult, "s")
	fmt.Println(splitresult)
	fmt.Println(findPosition("casabalanca", "la"))
}

//compareStrings recieves two strings and return int
func compareStrings(a string, b string) int {
	return strings.Compare(a, b)
}
func checkContain(s string, search string) bool {
	return strings.Contains(s, search)
}
func joinString(s []string, sep string) string {
	return strings.Join(s, sep)
}
func stringSplit(s string, sep string) []string {
	return strings.Split(s, sep)
}
func findPosition(s string, sub string) int {
	return strings.Index(s, sub)
}
