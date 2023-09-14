package game

import "testing"

var word = []string{"t", "e", "s", "t"}
var foundLetters = []string{"t", "e", "_", "t"}
var goodGuess = "t"
var badGuess = "m"

func TestLetterInWord(t *testing.T) {
	hasLetter := letterInWord(goodGuess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s, got = %v", word, goodGuess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	hasLetter := letterInWord(badGuess, word)
	if letterInWord(badGuess, word) {
		t.Errorf("Word %s does not contain letter %s, got = %v", word, badGuess, hasLetter)
	}
}

func TestHasWon(t *testing.T) {
	if !hasWon(word, word) {
		t.Errorf("Player should win here. Word = %s, foundLetters = %s", word, word)
	}
}

func TestHasNotWonYet(t *testing.T) {
	if hasWon(word, foundLetters) {
		t.Errorf("Player should not win here. Word = %s, foundLetters = %s", word, foundLetters)
	}
}