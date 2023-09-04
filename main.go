package main

import (
	"exemple/hangman/game"
	"fmt"
)

func main() {
	game := game.New(8, "Golang")
	fmt.Println(game)
}