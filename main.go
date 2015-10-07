package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var target = []string{}
var defaultState = []string{}
var alpha = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

type State struct {
	Deck     []string
	Failures int
	Letters  []string
	Counter  int
}

type Brain struct {
	Map           map[string]*State
	FailuresLimit int
}

func NewBrain() Brain {
	b := Brain{
		Map: make(map[string]*State, 0),
	}
	return b
}

func NewState() State {
	s := State{
		Deck:     make([]string, len(defaultState)),
		Failures: 0,
		Letters:  make([]string, 0),
	}
	copy(s.Deck, defaultState)
	return s
}

func (b *Brain) AddState(state *State) bool {
	hash := state.Hash()

	if found, ok := b.Map[hash]; ok {
		found.Counter++
		return false
	}
	state.Counter = 1
	b.Map[state.Hash()] = state
	return true
}

func (s *State) ApplyLetter(letter string) *State {
	newState := NewState()

	copy(newState.Deck, s.Deck)
	newState.Failures = s.Failures

	found := false
	for i := 0; i < len(target); i++ {
		if target[i] == letter {
			found = true
			newState.Deck[i] = letter
		}
	}
	if !found {
		newState.Failures++
	}

	return &newState
}

func (s *State) Hash() string {
	return fmt.Sprintf("%s:%d", strings.Join(s.Deck, ""), s.Failures)
}

func tryAlphabet(brain Brain, state *State) {
	for _, letter := range alpha {
		newState := state.ApplyLetter(letter)
		isNew := brain.AddState(newState)
		fmt.Println(letter, state, newState, newState.Hash())

		if isNew && newState.Failures < brain.FailuresLimit {
			tryAlphabet(brain, newState)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "pendu-dont-vous-etes-le-heros"
	app.Usage = "un pendu pour les héros"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "word",
			Value: "abeille",
			Usage: "Word",
		},
		cli.IntFlag{
			Name:  "failures-limit",
			Value: 6,
			Usage: "Failures limit",
		},
	}

	app.Action = func(c *cli.Context) {
		Pendu(c.String("word"), c.Int("failures-limit"))
	}

	app.Run(os.Args)
}

func Pendu(word string, failuresLimit int) {
	brain := NewBrain()
	brain.FailuresLimit = failuresLimit

	for _, letter := range word {
		target = append(target, string(letter))
	}

	for i := 0; i < len(target); i++ {
		defaultState = append(defaultState, "_")
	}

	state := NewState()
	brain.AddState(&state)

	tryAlphabet(brain, &state)

	for _, value := range brain.Map {
		fmt.Println(value.Hash(), value.Counter)
	}
	fmt.Println(len(brain.Map))
}
