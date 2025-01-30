import welcome
import secretnumber
import difficulty
import game

def main():
    welcome.welcome()

    while True:
        secret = secretnumber.generate_secret_number()
        attempts = difficulty.select_difficulty()
        game.play_game(secret, attempts)

        while True:
            play_again = input("Do you want to play again? (yes/no): ").strip().lower()
            if play_again in ["yes", "no"]:
                break
            print("Invalid input. Please enter 'yes' or 'no'.")

        if play_again == "no":
            print("Thanks for playing! Goodbye. 👋")
            break

if __name__ == "__main__":
    main()
