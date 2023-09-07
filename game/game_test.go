package game

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"t", "e", "s", "t"}
	goodGuess := "t"
	hasLetter := letterInWord(goodGuess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s, got = %v", word, goodGuess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"t", "e", "s", "t"}
	badGuess := "a"
	hasLetter := letterInWord(badGuess, word)
	if letterInWord(badGuess, word) {
		t.Errorf("Word %s does not contain letter %s, got = %v", word, badGuess, hasLetter)
	}
}