# Number Guessing Game

Welcome to the Number Guessing Game! 🎉 This is a simple CLI-based game where you try to guess a number chosen randomly by the computer. The game is built using Go.

## Features

- The computer randomly selects a number between 1 and 100.
- Players can choose a difficulty level:
  - **Easy**: 10 chances
  - **Medium**: 5 chances
  - **Hard**: 3 chancesS
- Hints are provided after each incorrect guess:
  - Whether the secret number is greater or lesser than your guess.
- Option to play multiple rounds.
- Validates user input to ensure smooth gameplay.

## How to Play

1. Clone the repository or download the project files.
2. Make sure you have Go installed (version 1.23 or higher).
3. Run the game:
   ```bash
   go run main.go

## Directory Structure
```plaintext
📂 number-guessing-game/
│
├── internal/
│   ├── difficulty/
│   │   └── difficulty.go
│   ├── game/
│   │   └── game.go
│   ├── secretnumber/
│   │   └── secretnumber.go
│   └── welcome/
│       └── selcome.go
│
├── go.mod
├── main.go
└── README.md
```
