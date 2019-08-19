package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.bbc.com/news/uk-politics-49393556")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", body)
}
