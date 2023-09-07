package main

import (
	"encoding/json"
	"errors"
	"exemple/hangman/game"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	game.Welcome()
	word, err := getARandomWorld()
	if err != nil {
		fmt.Println("error picking a word :", err)
		os.Exit(1)
	}
	hangman := game.New(8, word)
	input := ""
	for {
		game.DrawGame(hangman, input)

		if (hangman.State == "won" || hangman.State == "lost") {
			os.Exit(0)
		}

		input, err := game.ReadUserInput()
		if err != nil {
			fmt.Printf("Something wrong happened : %v", err)
			os.Exit(1)
		}

		hangman.MakeAGuess(input)
	}	
}

func getARandomWorld() (string, error) {
	data, err := os.ReadFile("./dataset/words.json")
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", errors.New("Empty content")
	}

	var words []string
	err = json.Unmarshal(data, &words)
    if err != nil {
        fmt.Println("error:", err)
    }

	maxIndex := len(words) - 1
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(maxIndex)
	return string(words[randomIndex]), nil
}