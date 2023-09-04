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
	}
}