package game

import "fmt"

func Game(secretNumber, attempts int) {
	remainder := attempts
	for attempts > 0 {
		var guess int
		fmt.Printf("%d", secretNumber)
		fmt.Printf("Enter your guess (%d attempts left): ", attempts)

		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			fmt.Scanln()
			continue
		}

		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", remainder-attempts+1)
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
