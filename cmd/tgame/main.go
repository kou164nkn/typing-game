package main

import (
	"fmt"
	"os"

	typeGame "github.com/kou164nkn/typing-game"
)

func main() {
	msg := typeGame.Do(os.Stdin)

	for {
		message, ok := <-msg
		if ok {
			fmt.Println(message)
		} else {
			break
		}
	}
}
