package game

import "fmt"

func game(secretNumber int, attempts int) {
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
}
