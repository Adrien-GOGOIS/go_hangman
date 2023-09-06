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

func (game *Hangman) MakeAGuess(input string) {
	input = strings.ToUpper(input)
	if letterInWord(input, game.AlreadyUsedLetters) {
		game.State = "alreadyGuessed"
	} else if letterInWord(input, game.Letters) {
		game.State = "goodGuess"
		game.revealLetter(input)
		if hasWon(game.Letters, game.FoundLetters) {
			game.State = "won"
		}
	}
}

func letterInWord(input string, letters []string) bool {
	for _, letter := range letters {
		if letter == input {
			return true
		}
	}
	return false
}

func (game *Hangman) revealLetter(input string) {
	game.AlreadyUsedLetters = append(game.AlreadyUsedLetters, input)
	for index, letter := range game.Letters {
		if letter == input {
			game.FoundLetters[index] = input
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for index := range letters {
		if letters[index] != foundLetters[index] {
			return false
		}
	}
	return true
}