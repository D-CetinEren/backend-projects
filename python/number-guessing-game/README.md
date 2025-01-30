# Number Guessing Game (Python)

This is a simple command-line number guessing game implemented in Python. The program randomly selects a number within a specified range, and the player must guess the number with the help of hints.

## Features

- The game selects a random number between **1 and 100**.
- The player inputs a guess, and the program provides feedback:
  - **"Too low!"** if the guess is lower than the secret number.
  - **"Too high!"** if the guess is higher than the secret number.
  - **"Congratulations!"** when the correct number is guessed.
- Tracks the number of attempts taken.
- Ensures valid input (only integers within the allowed range).
- Handles invalid inputs gracefully.

## Folder Structure
```plaintext
📂 number_guessing_game/
│
├── difficulty.py
├── game.py
├── secretnumber.py
├── welcome.py
├── main.py
└── README.md
```plaintext