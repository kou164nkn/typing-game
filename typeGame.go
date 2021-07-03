package typeGame

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var word = [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}

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

func question(l int) string {
	rand.Seed(time.Now().UnixNano())

	return word[rand.Intn(l)]
}

func Do(input io.Reader) chan string {
	ich := initial(input)

	msg := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	var score int
	len := len(word)

	go func() {
		defer cancel()
		defer close(msg)

		msg <- "Game start!"
		for {
			q := question(len)
			msg <- q
			select {
			case <-ctx.Done():
				goto ENDGAME
			case a := <-ich:
				if q == a {
					msg <- "Correct!"
					score++
				} else {
					msg <- "Wrong..."
				}
			}
		}
	ENDGAME:

		msg <- "Finished!"
		msg <- fmt.Sprintf("Your Score: %d", score)
	}()

	return msg
}
