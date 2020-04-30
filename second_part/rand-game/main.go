package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurn = 5

func main() {
	fmt.Println(os.Args[0])
	if len(os.Args) < 2 {
		fmt.Println("please provide yout guess number!!")
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if guess < 0 {
		fmt.Println("please pick positive number")
		return
	}
	rand.Seed(time.Now().UnixNano())
	for turn := 0; turn < maxTurn; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println("you win!!")
			return
		}
	}
	fmt.Println("you lost")
}
