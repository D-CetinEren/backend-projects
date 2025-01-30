def select_difficulty():
    print("Select difficulty level:")
    print("1. Easy (10 chances)")
    print("2. Medium (5 chances)")
    print("3. Hard (3 chances)")

    while True:
        try:
            choice = int(input("Enter your choice: "))
            if choice == 1:
                return 10
            elif choice == 2:
                return 5
            elif choice == 3:
                return 3
            else:
                print("Invalid choice. Please enter 1, 2, or 3.")
        except ValueError:
            print("Invalid input. Please enter a number.")
