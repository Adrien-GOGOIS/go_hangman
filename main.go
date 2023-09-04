package main

import (
	"exemple/hangman/game"
	"fmt"
	"os"
)

func main() {
	hangman := game.New(8, "Golang")
	fmt.Println(hangman)

	input, err := game.ReadUserInput()
	if err != nil {
		fmt.Printf("Something wrong happened : %v", err)
		os.Exit(1)
	}

	fmt.Print(input)
}