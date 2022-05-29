package main

import (
	"fmt"
	"strings"
)

func main() {

	//declare a word ansign some value at same time
	word := "elephant"
	//give some hint about the word
	hint := "biggest animal"
	//give chances to user to guess the letter
	chances := 10
	//declare emty slice for placeholder
	guesses := []string{}

	//append uderscores into empty slice
	for i := 0; i < len(word); i++ {
		guesses = append(guesses, "_")
	}

	//this is infinite loop
	for {
		//add condition show his loss
		userGuess := strings.Join(guesses, "")
		if chances == 0 && userGuess != word {
			fmt.Println("Failed to guess please try again")
			break
		}

		//Show the user that he wins the game and end the infinit loop
		if chances > 0 && userGuess == word {
			fmt.Println("Congratulations")
			break
		}

		//print the hint to guess the word
		fmt.Println(hint)
		//print the right guess made by user
		fmt.Print(guesses)
		//print message
		fmt.Println("Enter a your guess")
		//declare and initilize a one veriable to store user input
		guess := ""
		//store user input
		fmt.Scanln(&guess)

		//add remaining code here

	}

}
