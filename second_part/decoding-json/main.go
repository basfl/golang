package main

import (
	"encoding/json"
	"fmt"
)

//Permission type
type Permission map[string]bool
type user struct {
	Name string
	Pass string
	Permission
}

func main() {
	users := []user{
		{"babak", "123", nil},
		{"amir", "****", Permission{"write": true}},
	}
	// s1 := []string{}
	// s1 = []string{"a", "b"}
	// fmt.Println(s1)
	// a := make(map[string]bool)
	// a = map[string]bool{
	// 	"a": true,
	// }
	// a := Permission{"a": false}
	//fmt.Printf("\n%v", a)
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(out))

}
