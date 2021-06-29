package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func initial(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(r)
		for {
			scanner.Scan()
			ch <- scanner.Text()
		}
	}()

	return ch
}

func Do(input io.Reader) chan string {
	ich := initial(input)

	msg := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// var score int

	go func() {
		defer cancel()
		defer close(msg)

		msg <- "Game start!"
		for {
			select {
			case <-ctx.Done():
				goto ENDGAME
			case <-ich:
				fmt.Println("foo")
			}
		}
	ENDGAME:

		msg <- ("Finished!")
	}()

	return msg
}

func main() {
	msg := Do(os.Stdin)

	for {
		message, ok := <-msg
		if ok {
			fmt.Println(message)
		} else {
			break
		}
	}
}
