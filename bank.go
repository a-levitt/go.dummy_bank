package main

import "fmt"

var accountBalance float64 = 1000

func main() {

	fmt.Println("Welcom to Go Bank!")
	fmt.Println("Choose an option:")
	fmt.Println("1. Check balance")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Exit")

	var choice uint8
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	userChoice(choice)

}

func userChoice(choice uint8) {
	for choice != 4 {
		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("Withdraw operations")
		} else if choice == 3 {
			fmt.Println("Deposit operations")
		} else {
			fmt.Println("Invalid choice")
		}
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	}
	fmt.Println("Exit...")
}
