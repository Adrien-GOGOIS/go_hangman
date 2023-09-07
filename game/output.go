package game

import (
	"fmt"

	"github.com/fatih/color"
)

func Welcome() {
	fmt.Println(`
	{}    {}    {}{}     {}    {}    {}}}}}    {}      {}    {}{}     {}    {}
	{}    {}   {}  {}    {}}}  {}   {}    {}   {}}}  {{{}   {}  {}    {}}}  {}
	{}{{}}{}  {}{{}}{}   {} {} {}   {}         {} {{}} {}  {}{{}}{}   {} {} {}
	{}    {}  {}    {}   {}  {{{}   {}  {{{{   {}  {}  {}  {}    {}   {}  {{{}
	{}    {}  {}    {}   {}    {}    {}}}}}    {}      {}  {}    {}   {}    {}   
	`)
}

func DrawGame(game *Hangman, input string) {
	drawTurns(game.RemainingAttempts)
	drawState(game, input)
}

func drawTurns(turns int) {
	var draw string
	switch turns {
		case 0:
			draw = `
			----------
			|     |
			|     @
			|    /|\  YOU LOSE !
			|     |
			|    / \
			|
			|
			______________
			`
		case 1:
			draw = `
			----------
			|     |
			|     @
			|     |
			|     |
			|    / \
			|
			|
			______________
			`
		case 2:
			draw = `
			----------
			|     |
			|     @
			|     |  
			|     |
			|    
			|
			|
			______________
			`
		case 3:
			draw = `
			----------
			|     |
			|     @
			|     
			|     
			|    
			|
			|
			______________
			`
		case 4:
			draw = `
			----------
			|     
			|    
			|      
			|     
			|    
			|
			|
			______________
			`
		case 5:
			draw = `
			
			|     
			|     
			|     
			|     
			|    
			|
			|
			______________
			`
		case 6:
			draw = `
			
			  
			|    
			|
			|
			______________
			`
		case 7:
			draw = `
			
			
			______________
			`
		case 8:
			draw = `
			
			`
	}
	fmt.Println(draw)
}

func drawState(game *Hangman, input string)  {
	color.Blue("Mot : ")
	fmt.Print(" \n")
	drawLetters(game.FoundLetters)
	fmt.Print(" \n")
	color.Yellow("Lettres utilisées : ")
	fmt.Print(" \n")
	drawLetters(game.AlreadyUsedLetters)
	fmt.Print(" \n")

	switch game.State {
	case "goodGuess":
		color.Green("Bon choix !\n")
		fmt.Print(" \n")
	case "alreadyGuessed":
		color.Yellow("Lettre déjà utilisée !\n")
	case "badGuess":
		color.Red("Mauvais choix !\n")
		fmt.Printf("Il te reste %v essais\n", game.RemainingAttempts)
		fmt.Print(" \n")
	case "lost":
		fmt.Print("Perdu !\n")
		drawLetters(game.Letters)
	case "won":
		drawWin()
	}
}

func drawLetters(letters []string) {
	for _, char := range letters {
		fmt.Printf("%v ", char)
	}
	fmt.Println()
}

func drawWin() {
	fmt.Printf(`
	{}             {}   {}   {}      {}   {}
	 {}           {}    {}   {} {}   {}   {}
	  {}   {}    {}     {}   {}  {}  {}   {}
	   {} {} {} {}      {}   {}   {} {}
	    {}    {}        {}   {}      {}   {}
	`)
}