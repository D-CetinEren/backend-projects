package main

import (
	"fmt"
	"strings"

	"github.com/D-CetinEren/backend-projects/go/number-guessing-game/internal/difficulty"
	"github.com/D-CetinEren/backend-projects/go/number-guessing-game/internal/game"
	"github.com/D-CetinEren/backend-projects/go/number-guessing-game/internal/secretnumber"
	"github.com/D-CetinEren/backend-projects/go/number-guessing-game/internal/welcome"
)

func main() {
	welcome.Welcome()

	for {
		secret := secretnumber.SecretNumber()
		attempts := difficulty.SelectingDifficulty()

		game.Game(secret, attempts)

		var playAgain string
		for {
			fmt.Println("Do you want to play again? (yes/no)")
			_, err := fmt.Scanf("%s", &playAgain)
			if err != nil {
				fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
				fmt.Scanln()
				continue
			}
			//fmt.Scan(&playAgain)

			playAgain = strings.ToLower(playAgain)
			if playAgain == "yes" || playAgain == "no" {
				break
			}

			fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
		}

		if playAgain == "no" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}

}
