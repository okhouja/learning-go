package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	// Get user input
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	
	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&tax_rate)

	// Calculate values after getting input
	ebt := revenue - expenses // ebt = earning before tax
	profit := ebt * (1 - tax_rate / 100)
	ratio := ebt / profit
	
	// Output results
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}