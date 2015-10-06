package main

import (
	"fmt"
	"strings"
)

var target = []string{"a", "b", "e", "i", "l", "l", "e"}
var defaultState = []string{}
var alpha = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

type State struct {
	Deck     []string
	Failures int
	Letters  []string
}

type Brain struct {
	Map map[string]State
}

func NewBrain() Brain {
	b := Brain{
		Map: make(map[string]State, 0),
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

func main() {
	for i := 0; i < len(target); i++ {
		defaultState = append(defaultState, "_")
	}

	brain := NewBrain()
	fmt.Println("brain: %v", brain)

	state := NewState()
	for _, letter := range alpha {
		newState := state.ApplyLetter(letter)
		fmt.Println(letter, state, newState, newState.Hash())
	}
}
