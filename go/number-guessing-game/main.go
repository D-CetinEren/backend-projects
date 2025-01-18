package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the number guessing game")
	fmt.Println("I'm thinking of a number 1 to 100")
	fmt.Println("Your goal is to guess the number")
	fmt.Println("Let's start")

	secretNumber := rand.Intn(100) + 1
	fmt.Println("Secret number has been chosen!")

	var attempts int
	fmt.Println("Select difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		attempts = 10
	case 2:
		attempts = 5
	case 3:
		attempts = 3
	default:
		fmt.Println("Invalid choice. Defaulting to Medium difficulty.")
		attempts = 5
	}

	fmt.Printf("You have %d chances to guess the number.\n", attempts)

	for attempts > 0 {
		var guess int
		fmt.Printf("Enter your guess (%d attempts left): ", attempts)
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.", attempts)
			break
		} else if guess < secretNumber {
			fmt.Printf("Incorrect! The number is greater than %d\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is lesser than %d\n", guess)
		}

		attempts--
	}

	if attempts == 0 {
		fmt.Printf("Game over! The correct number was %d.\n", secretNumber)
	}

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
		main()
	} else {
		fmt.Println("Thanks for playing! Goodbye.")
	}

}
