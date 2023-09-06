package game

import "fmt"

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
			|    /|\  You lose
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
	fmt.Print("Mot : ")
	drawLetters(game.FoundLetters)

	fmt.Print("Lettres utilisées : ")
	drawLetters(game.AlreadyUsedLetters)

	switch game.State {
	case "goodGuess":
		fmt.Print("Bon choix !")
	case "alreadyGuessed":
		fmt.Printf("Lettre %v déjà utilisée !", input)
	case "badGuess":
		fmt.Printf("Mauvais choix, %v n'est pas dans ce mot !", input)
	case "lost":
		fmt.Print("Perdu !")
		drawLetters(game.Letters)
	case "won":
		fmt.Print("Tu as choisi...judicieusement !")
		drawLetters(game.Letters)
	}
}

func drawLetters(letters []string) {
	for _, char := range letters {
		fmt.Printf("%v ", char)
	}
	fmt.Println()
}