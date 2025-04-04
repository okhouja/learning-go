package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find account balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		panic("Can't Continue, Sorry")
	}

	fmt.Println("Welcome to Go Bank!")

	// for i := 0; i < 200; i++	
	for {
	
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)
	
		// wantCheckBalance := choice == 1
		// if statment :
	
	// 		if choice == 1{
	// 		fmt.Println("Your balance is", accountBalance)
	// 	} else if choice == 2 {
	// 		fmt.Println("You deposit:")
	// 		var depositAmount float64
	// 		fmt.Scan(&depositAmount)
	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0.")
	// 			continue
	// 		}
	// 		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
	// 		fmt.Println("Balance Updated! Your new balance is:", accountBalance)
	// 	} else if choice == 3 {
	// 		fmt.Println("Your Withdraw amount:")
	// 		var withdrawAmount float64
	// 		fmt.Scan(&withdrawAmount)
	// 		if withdrawAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0.")
	// 			continue
	// 		}
	// 		if withdrawAmount > accountBalance {
	// 			fmt.Println("Invalid amount. You can't withdraw more than you have.")
	// 			continue
	// 		}
	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("Balance Updated! Your new balance is:", accountBalance)
	// 	} else {
	// 		fmt.Println("Goodbye!")
	// 		// return
	// 		break
	// 	}
	// }

	// switch statement :

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Println("You deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance Updated! Your new balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("Your Withdraw amount:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! Your new balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
			//break
		}

	}

}