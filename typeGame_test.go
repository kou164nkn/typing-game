package typeGame_test

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	typeGame "github.com/kou164nkn/typing-game"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		inputMethod io.Reader
		expect      *typeGame.Game
	}{
		"cli_mode": {os.Stdin, typeGame.NewTestGame(os.Stdin, os.Stdout, nil, 15)},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := typeGame.New(tt.inputMethod, os.Stdout)

			if !reflect.DeepEqual(actual, tt.expect) {
				t.Errorf("New want %v but got %v", tt.expect, actual)
			}
		})
	}
}

func TestDo(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		testQuiz    []string
		testTime    int
		answer      string
		expectScore int
	}{
		"no_answer": {[]string{"sample"}, 1, "", 0},
		"2_correct": {[]string{"sample"}, 1, "sample\nwrong\nsample", 2},
	}

	for name, tt := range cases {
		g := typeGame.NewTestGame(strings.NewReader(tt.answer), io.Discard, tt.testQuiz, tt.testTime)

		t.Run(name, func(t *testing.T) {
			g.Do()

			if g.Score() != tt.expectScore {
				t.Errorf("Expect score is %d but %d", tt.expectScore, g.Score())
			}
		})
	}
}
