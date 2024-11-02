package main

import "fmt"

var accountBalance float64 = 1000

func main() {

	fmt.Println("Welcom to Go Bank!")
	fmt.Println("Choose an option:")
	fmt.Println("1. Check balance")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Exit\n")

	var choice uint8
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	userChoice(choice)

}

func userChoice(choice uint8) {
	for choice != 4 {
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)

			break
		case 2:
			fmt.Println("Withdraw operations")
			fmt.Print("You withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Wrong amount")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Not enough balance! Available balance is", accountBalance)
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Success... New amount:", accountBalance)
			}

			break
		case 3:
			fmt.Println("Deposit operations")
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Wrong amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Success... New amount:", accountBalance)

			break
		default:
			fmt.Println("Invalid choice")

			break
		}

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	}
	fmt.Println("Exit...")
}
