package main

import (
	"fmt"
)

func main() {
	welcome.welcome()

	var secret int
	secret = secretnumber.secretNumber()

	var attempts int
	attempts = difficulty.selectingDifficulty()

	game.game(secret, attempts)

	var playAgain string

	fmt.Println("Do you want to play again? (yes/no)")
	fmt.Scan(&playAgain)

	for {
		fmt.Println("Do you want to play again? (yes/no)")
		fmt.Scan(&playAgain)

		if playAgain == "yes" || playAgain == "no" {
			break
		}

		fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
	}

	if playAgain == "yes" {
		game.game()
	} else {
		fmt.Println("Thanks for playing! Goodbye.")
	}

}
