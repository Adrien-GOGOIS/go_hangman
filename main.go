package main

import (
	"exemple/hangman/game"
	"fmt"
	"os"
)

func main() {
	game.Welcome()
	hangman := game.New(8, "Golang")
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