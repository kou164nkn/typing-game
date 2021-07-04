package typeGame

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var word = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}
var defaultTime = 15

type Game struct {
	input     io.Reader
	output    io.Writer
	quiz      []string
	score     int
	timeLimit int
}

func New(r io.Reader, w io.Writer) *Game {
	g := Game{
		input:     r,
		output:    w,
		quiz:      word,
		timeLimit: defaultTime,
	}

	return &g
}

func (g *Game) Do() {
	ich := initial(g.input)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(g.timeLimit)*time.Second)
	defer cancel()

	g.score = 0
	len := len(g.quiz)
	rand.Seed(time.Now().UnixNano())

	fmt.Fprintln(g.output, "Game start!")
	for {
		q := g.quiz[rand.Intn(len)]
		fmt.Fprintln(g.output, q)
		select {
		case <-ctx.Done():
			goto ENDGAME
		case a := <-ich:
			if q == a {
				fmt.Fprintln(g.output, "Correct!")
				g.score++
			} else {
				fmt.Fprintln(g.output, "Wrong...")
			}
		}
	}
ENDGAME:

	fmt.Fprintln(g.output, "Finished!")
	fmt.Fprintf(g.output, "Your score: %d\n", g.score)
}

func initial(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}
