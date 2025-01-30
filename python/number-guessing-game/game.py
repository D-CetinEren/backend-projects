def play_game(secret_number, attempts):
    for attempt in range(1, attempts + 1):
        try:
            guess = int(input(f"Enter your guess ({attempts - attempt + 1} attempts left): "))
        except ValueError:
            print("Invalid input. Please enter a number.")
            continue

        if guess == secret_number:
            print(f"🎉 Congratulations! You guessed the correct number in {attempt} attempts.")
            return
        elif guess < secret_number:
            print(f"Incorrect! The number is greater than {guess}.")
        else:
            print(f"Incorrect! The number is lesser than {guess}.")

    print(f"💀 Game over! The correct number was {secret_number}.")
