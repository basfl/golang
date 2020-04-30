package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title     string
	price     float64
	published interface{}
}

func (b *book) print() {
	p := formated(b.published)
	fmt.Printf("\n the book title is %s and the price is %.2f and published at %s ", b.title, b.price, p)
}

// func (b *book) discount(ration float64) {
// 	b.price = b.price * ration

// }

func formated(v interface{}) string {
	var t int
	//fmt.Printf("tpe %v", )
	switch value := v.(type) {
	case int:
		t = value
	case string:
		t, _ = strconv.Atoi(value)
	default:
		return "unknown"
	}
	// if v == nil {

	// 	return "unknown"
	// }

	// if value, ok := v.(int); ok {
	// 	t = value
	// }
	// if value, ok := v.(string); ok {
	// 	t, _ = strconv.Atoi(value)
	// }
	u := time.Unix(int64(t), 0)
	return u.String()
}
