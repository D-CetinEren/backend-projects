import random

def main():
    lower_bound = 1
    upper_bound = 100
    secret_number = random.randint(lower_bound, upper_bound)
    attempts = 0

    print(f"Guess the number between {lower_bound} and {upper_bound}.")

    while True:
        try:
            guess = int(input("Enter your guess: "))
            attempts += 1

            if guess < lower_bound or guess > upper_bound:
                print(f"Please enter a number between {lower_bound} and {upper_bound}.")
                continue

            if guess < secret_number:
                print("Too low!")
            elif guess > secret_number:
                print("Too high!")
            else:
                print(f"Congratulations! You've guessed the number in {attempts} attempts.")
                break
        except ValueError:
            print("Invalid input. Please enter a valid integer.")

if __name__ == "__main__":
    main()
