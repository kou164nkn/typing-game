package main

import (
	"os"

	typeGame "github.com/kou164nkn/typing-game"
)

func main() {
	g := typeGame.New(os.Stdin, os.Stdout)

	g.Do()
}
