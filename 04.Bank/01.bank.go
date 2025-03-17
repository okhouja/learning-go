package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Enter your choice: ")
	fmt.Scan(&choice)

	// wantCheckBalance := choice == 1

	if choice == 1{
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Println("You deposit:")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
		fmt.Println("Balance Updated! Your new balance is:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Your Withdraw amount:")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Balance Updated! Your new balance is:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}