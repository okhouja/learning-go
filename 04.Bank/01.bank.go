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
	}

	fmt.Println("You chose", choice)

}