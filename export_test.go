package typeGame

import "io"

func NewTestGame(i io.Reader, o io.Writer, q []string, t int) *Game {
	if len(q) == 0 {
		q = word
	}

	return &Game{input: i, output: o, quiz: q, score: 0, timeLimit: t}
}

func (g *Game) Score() int {
	return g.score
}
