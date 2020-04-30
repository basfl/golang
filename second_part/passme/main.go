package main

import (
	"fmt"
	"os"

	"./message"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(message.UserMessage1)
	} else if os.Args[1] != message.UserName {
		fmt.Println(message.UserNameErrMesg)
	} else if os.Args[2] != message.PassWord {
		fmt.Println(message.PassWordErrMesg)
	} else {
		fmt.Println(message.UserMessage2)
	}
}
