package game

import "strings"

type Hangman struct {
	State string
	Letters []string
	FoundLetters []string
	AlreadyUsedLetters []string
	RemainingAttempts int
}

func New(turns int, word string) *Hangman {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for index := 0; index < len(letters); index++ {
		found[index] = " _ "
	}

	game := &Hangman {
		State: "",
		Letters: letters,
		FoundLetters: found,
		AlreadyUsedLetters: []string{},
		RemainingAttempts: turns,
	}

	return game
}