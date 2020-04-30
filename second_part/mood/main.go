package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},

		{"sad 😞", "bad 👎", "terrible 😩"},
	}
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("please enter the name and mood")
		return
	}
	name, mood := arg[0], arg[1]

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(moods[0]))
	fmt.Println(name, mood, moods, n)
	var mi int
	if mood != "positive" {

		mi = 1
	}
	fmt.Printf("%s feels %s\n", name, moods[mi][n])

}
