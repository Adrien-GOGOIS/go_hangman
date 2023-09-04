package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadUserInput() (input string, err error) {
	valid := false
	for !valid {
		fmt.Print("Entrez une lettre : ")
		input, err = reader.ReadString('\n') // Lit jusqu'à l'appuie de la touche "entrée" : \n
		
		if err != nil {
			return input, err
		}

		input = strings.TrimSpace(input)
		if len(input) != 1 {
			fmt.Printf("Invalid input : letter = %v, len = %v\n", input, len(input))
			continue
		}
		valid = true
	}
	
	return input, err
}