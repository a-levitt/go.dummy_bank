package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "user_balance.txt"

func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceSting := string(data)
	balance, _ := strconv.ParseFloat(balanceSting, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceString), 0644)
}

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
	accountBalance := readBalanceFromFile()
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
				writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)

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
