package difficulty

import "fmt"

func SelectingDifficulty() (attempts int) {
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

	return attempts
}
