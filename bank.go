package main

import (
	"fmt"
	"go.dummy_bank/fileops"
)

const accountBalanceFile = "user_balance.txt"

func main() {

	presentOptions()

	var choice uint8
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	userChoice(choice)

}

func userChoice(choice uint8) {
	var accountBalance, err = fileops.ReadFloatFromFile(accountBalanceFile, 1000)
	if err != nil {
		fmt.Println("Error.", err, "\n-----------")
	}
	for choice != 4 {
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)

			//break
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}

			//break
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

			//break
		default:
			fmt.Println("Invalid choice")

			//break
			return

		}

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	}
	fmt.Println("Exit...")
}
