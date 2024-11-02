package main

import "fmt"

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

	if choice == 1 {
		fmt.Println("Balance operations")
	} else if choice == 2 {
		fmt.Println("Withdraw operations")
	} else if choice == 3 {
		fmt.Println("Deposit operations")
	} else if choice == 4 {
		fmt.Println("Choose an option:")
	} else {
		fmt.Println("Invalid choice")
	}

}
